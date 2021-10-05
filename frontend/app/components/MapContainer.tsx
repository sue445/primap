import { GoogleApiWrapper, InfoWindow, Map, Marker } from "google-maps-react";
import React from "react";
import { GeoFireClient } from "geofirex";
import * as Sentry from "@sentry/react";
import { ShopEntity } from "./ShopEntity";
import { correctLongitude, formatAddress, getShopMarkerIconUrl } from "./Util";
import SeriesCheckbox from "./SeriesCheckbox";
import SearchConditionRadio from "./SearchConditionRadio";

type Props = {
  latitude: number;
  longitude: number;
  zoom: number;
  geo: GeoFireClient;
};

const emptyShop = { series: new Set([]) } as ShopEntity;

const shopLimit = 2000;

const defaultSeries = ["primagi", "prichan", "pripara"];

type SearchCondition = "or" | "and";

export class MapContainer extends React.Component<Props, {}> {
  state = {
    activeMarker: {} as google.maps.Marker,
    selectedShop: emptyShop,
    showingInfoWindow: false,
    shops: [] as Array<ShopEntity>,
    latitude: this.props.latitude,
    longitude: this.props.longitude,
    series: new Set(defaultSeries),
    seriesArray: defaultSeries,
    searchCondition: "or" as SearchCondition,
  };

  shopCache = {};

  onMapReady = (mapProps, map: google.maps.Map) => {
    if (!navigator.geolocation) {
      return;
    }

    navigator.geolocation.getCurrentPosition((pos) => {
      const data = pos.coords;
      map.setCenter(new google.maps.LatLng(data.latitude, data.longitude));
      this.setState({
        latitude: data.latitude,
        longitude: data.longitude,
      });
    });
  };

  loadShops = (map: google.maps.Map) => {
    const bounds = map.getBounds();
    if (!bounds) {
      if (process.env.NODE_ENV != "production") {
        console.log("[INFO] bounds is undefined");
      }
      Sentry.captureMessage("bounds is undefined", Sentry.Severity.Info);
      return;
    }

    const geo = this.props.geo;

    const centerLatitude = map.getCenter().lat();
    const centerLongitude = correctLongitude(map.getCenter().lng());
    const center = geo.point(centerLatitude, centerLongitude);

    const distance = geo.distance(
      geo.point(bounds.getSouthWest().lat(), bounds.getSouthWest().lng()),
      geo.point(bounds.getNorthEast().lat(), bounds.getNorthEast().lng())
    );

    const firestoreRef = geo.app
      .firestore()
      .collection("Shops")
      .where("deleted", "==", false)
      .limit(shopLimit);
    const query = geo
      .query(firestoreRef)
      .within(center, distance / 2, "geography");

    query.subscribe((hits) => {
      const shops = [] as Array<ShopEntity>;
      hits.forEach((data) => {
        const shop = new ShopEntity(data);
        shops.push(shop);
        this.shopCache[shop.name] = shop;
      });
      this.setState({ shops: shops });
    });
  };

  onMapRefresh = (mapProps, map: google.maps.Map) => {
    this.loadShops(map);
  };

  onMarkerClick = (props, marker: google.maps.Marker) =>
    this.setState({
      activeMarker: marker,
      selectedShop: this.shopCache[props.name],
      showingInfoWindow: true,
    });

  onInfoWindowClose = () =>
    this.setState({
      activeMarker: null,
      selectedShop: emptyShop,
      showingInfoWindow: false,
    });

  onMapClicked = () => {
    if (this.state.showingInfoWindow)
      this.setState({
        activeMarker: null,
        showingInfoWindow: false,
      });
  };

  onSeriesChanged = (event: React.ChangeEvent<HTMLInputElement>) => {
    if (event.target.checked) {
      this.state.series.add(event.target.value);
    } else {
      this.state.series.delete(event.target.value);
    }
    this.state.seriesArray = Array.from(this.state.series);
    this.setState({ series: this.state.series });
  };

  onSearchConditionChanged = (event: React.ChangeEvent<HTMLInputElement>) => {
    if (event.target.value == "and" || event.target.value == "or") {
      this.state.searchCondition = event.target.value;
      this.setState({ searchCondition: this.state.searchCondition });
    }
  };

  render() {
    return (
      <div>
        <div className="flex mt-6">
          <span className={"h-6 font-bold"}>絞り込み</span>
          <SeriesCheckbox
            title="ワッチャプリマジ！"
            value="primagi"
            checked={this.state.series.has("primagi")}
            onChange={this.onSeriesChanged}
          />
          <SeriesCheckbox
            title="キラッとプリ☆チャン"
            value="prichan"
            checked={this.state.series.has("prichan")}
            onChange={this.onSeriesChanged}
          />
          <SeriesCheckbox
            title="プリパラ オールアイドル"
            value="pripara"
            checked={this.state.series.has("pripara")}
            onChange={this.onSeriesChanged}
          />
        </div>
        <div className="flex mt-6">
          <span className={"h-6 font-bold"}>検索条件</span>
          <SearchConditionRadio
            title="OR検索"
            value="or"
            checked={this.state.searchCondition == "or"}
            onChange={this.onSearchConditionChanged}
          />
          <SearchConditionRadio
            title="AND検索"
            value="and"
            checked={this.state.searchCondition == "and"}
            onChange={this.onSearchConditionChanged}
          />
        </div>
        <Map
          // @ts-ignore
          google={this.props.google}
          zoom={this.props.zoom}
          onReady={this.onMapReady}
          onCenterChanged={this.onMapRefresh}
          onBoundsChanged={this.onMapRefresh}
          onZoomChanged={this.onMapRefresh}
          onDragend={this.onMapRefresh}
          onRecenter={this.onMapRefresh}
          onResize={this.onMapRefresh}
          initialCenter={{
            lat: this.props.latitude,
            lng: this.props.longitude,
          }}
          containerStyle={{
            position: "relative",
            width: "100%",
            height: "100%",
          }}
        >
          <Marker
            position={{
              lat: this.state.latitude,
              lng: this.state.longitude,
            }}
            icon={{
              url: "/img/marker_current_position.png",
            }}
            zIndex={10}
          />
          {this.state.shops
            .filter((shop) => {
              switch (this.state.searchCondition) {
                case "or":
                  return this.state.seriesArray.some((series) => {
                    return shop.series.has(series);
                  });
                case "and":
                  return this.state.seriesArray.every((series) => {
                    return shop.series.has(series);
                  });
              }
              return false;
            })
            .map((shop) => {
              const iconUrl = getShopMarkerIconUrl(shop.name);
              return (
                <Marker
                  key={shop.name}
                  onClick={this.onMarkerClick}
                  position={{
                    lat: shop.geography.geopoint.latitude,
                    lng: shop.geography.geopoint.longitude,
                  }}
                  // @ts-ignore
                  name={shop.name}
                  zIndex={1}
                  icon={iconUrl != null ? { url: iconUrl } : null}
                />
              );
            })}

          <InfoWindow
            marker={this.state.activeMarker}
            // @ts-ignore
            onClose={this.onInfoWindowClose}
            visible={this.state.showingInfoWindow}
          >
            <div className={"container mx-auto p-2 md:w-64 lg:w-64 xl:w-64"}>
              <h3 className={"text-sm font-bold"}>
                {this.state.selectedShop.name}
              </h3>
              <dl>
                <dt className={"font-semibold"}>住所</dt>
                <dd>{formatAddress(this.state.selectedShop.address)}</dd>
                <dt className={"font-semibold"}>設置筐体</dt>
                {this.state.selectedShop.series.has("primagi") && (
                  <dd>ワッチャプリマジ！</dd>
                )}
                {this.state.selectedShop.series.has("prichan") && (
                  <dd>キラッとプリ☆チャン</dd>
                )}
                {this.state.selectedShop.series.has("pripara") && (
                  <dd>プリパラ オールアイドル</dd>
                )}
                <dt className={"font-semibold"}>更新日時</dt>
                <dd>{this.state.selectedShop.updated_at?.toLocaleString()}</dd>
              </dl>
            </div>
          </InfoWindow>
        </Map>
      </div>
    );
  }
}

export default GoogleApiWrapper({
  apiKey: process.env.REACT_APP_GOOGLE_BROWSER_API_KEY,
  language: "ja",
  // @ts-ignore
})(MapContainer);
