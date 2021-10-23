import { LatLng } from "./ShopEntity";

export function correctLongitude(longitude: number): number {
  if (-180 < longitude && longitude < 180) {
    return longitude;
  }

  longitude += 180;
  while (longitude < 0) {
    longitude += 360;
  }
  return (longitude % 360) - 180;
}

export function formatAddress(address: string): string {
  if (!address) {
    return "";
  }
  return address.replace(/　+/g, "　");
}

export function getShopMarkerIconUrl(shopName: string): string {
  if (shopName.startsWith("プリズムストーン")) {
    return "/img/marker_prismstone.png";
  }

  // Use default icon
  return null;
}

export function getGoogleMapUrl(geopoint?: LatLng): string {
  if (geopoint) {
    return `https://www.google.com/maps/search/?api=1&query=${geopoint.latitude},${geopoint.longitude}`;
  }
  return "";
}
