import { Map, InfoWindow, Marker, GoogleApiWrapper } from "google-maps-react";
import React from "react";

type Props = {
  latitude: number;
  longitude: number;
  zoom: number;
};
export class MapContainer extends React.Component<Props, {}> {
  state = {
    activeMarker: {},
    selectedPlace: {},
    showingInfoWindow: false,
  };

  onMarkerClick = (props, marker) =>
    this.setState({
      activeMarker: marker,
      selectedPlace: props,
      showingInfoWindow: true,
    });

  onInfoWindowClose = () =>
    this.setState({
      activeMarker: null,
      showingInfoWindow: false,
    });

  onMapClicked = () => {
    if (this.state.showingInfoWindow)
      this.setState({
        activeMarker: null,
        showingInfoWindow: false,
      });
  };

  render() {
    return (
      <Map
        // @ts-ignore
        google={this.props.google}
        zoom={this.props.zoom}
        initialCenter={{
          lat: this.props.latitude,
          lng: this.props.longitude,
        }}
      >
        <Marker
          onClick={this.onMarkerClick}
          // @ts-ignore
          name={"Current location"}
        />

        <InfoWindow
          // @ts-ignore
          marker={this.state.activeMarker}
          onClose={this.onInfoWindowClose}
          visible={this.state.showingInfoWindow}
        >
          <div>
            <h1>
              {
                // @ts-ignore
                this.state.selectedPlace.name
              }
            </h1>
          </div>
        </InfoWindow>
      </Map>
    );
  }
}

export default GoogleApiWrapper({
  apiKey: process.env.REACT_APP_GOOGLE_BROWSER_API_KEY,
  language: "ja",
  // @ts-ignore
})(MapContainer);
