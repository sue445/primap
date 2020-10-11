import { GoogleApiWrapper, InfoWindow, Map, Marker } from "google-maps-react";
import React from "react";
import { GeoFireClient } from "geofirex";
import { ShopEntity, Time } from "./ShopEntity";

type Props = {
  latitude: number;
  longitude: number;
  zoom: number;
  geo: GeoFireClient;
};

const emptyShop = { series: [], updated_at: new Time() } as ShopEntity;

const shopLimit = 2000;

export class MapContainer extends React.Component<Props, {}> {
  state = {
    activeMarker: {} as google.maps.Marker,
    selectedShop: emptyShop,
    showingInfoWindow: false,
    shops: [] as Array<ShopEntity>,
    latitude: this.props.latitude,
    longitude: this.props.longitude,
    series: new Set(["prichan", "pripara"]),
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
      this.loadShops(map);
    });
  };

  loadShops = (map: google.maps.Map) => {
    const bounds = map.getBounds();
    if (!bounds) {
      return;
    }

    const geo = this.props.geo;
    const center = geo.point(map.getCenter().lat(), map.getCenter().lng());
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
        const shop = ShopEntity.createFrom(data);
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
    this.setState({ series: this.state.series });
  };

  render() {
    return (
      <div>
        <div className="flex mt-6">
          <label className="flex items-center">
            <input
              type="checkbox"
              className="form-checkbox h-6 w-6"
              onChange={this.onSeriesChanged}
              value="prichan"
              checked={this.state.series.has("prichan")}
            />
            <span className="ml-1">プリ☆チャン</span>
          </label>
          <label className="flex items-center">
            <input
              type="checkbox"
              className="form-checkbox h-6 w-6"
              onChange={this.onSeriesChanged}
              value="pripara"
              checked={this.state.series.has("pripara")}
            />
            <span className="ml-1">プリパラ</span>
          </label>
        </div>
        <Map
          // @ts-ignore
          google={this.props.google}
          zoom={this.props.zoom}
          onReady={this.onMapReady}
          onCenterChanged={this.onMapRefresh}
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
          {this.state.shops
            .filter((shop) => {
              return shop.series.some((series) => {
                return this.state.series.has(series);
              });
            })
            .map((shop) => (
              <Marker
                key={shop.name}
                onClick={this.onMarkerClick}
                position={{
                  lat: shop.geography.geopoint.latitude,
                  lng: shop.geography.geopoint.longitude,
                }}
                // @ts-ignore
                name={shop.name}
              />
            ))}

          <InfoWindow
            marker={this.state.activeMarker}
            // @ts-ignore
            onClose={this.onInfoWindowClose}
            visible={this.state.showingInfoWindow}
          >
            <div>
              <h3 className={"text-sm font-bold"}>
                {this.state.selectedShop.name}
              </h3>
              <dl>
                <dt className={"font-semibold"}>住所</dt>
                <dd>{this.state.selectedShop.address}</dd>
                <dt className={"font-semibold"}>設置筐体</dt>
                {this.state.selectedShop.series.includes("prichan") && (
                  <dd>キラッとプリ☆チャン</dd>
                )}
                {this.state.selectedShop.series.includes("pripara") && (
                  <dd>プリパラ オールアイドル</dd>
                )}
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
