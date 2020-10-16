export function correctLongitude(longitude: number): number {
  longitude += 180;
  while (longitude < 0) {
    longitude += 360;
  }
  return (longitude % 360) - 180;
}
