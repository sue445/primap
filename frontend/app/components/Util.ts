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
