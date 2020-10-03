/* Do not change, this code is generated from Golang structs */

export class LatLng {
  latitude?: number;
  longitude?: number;

  static createFrom(source: any) {
    if ("string" === typeof source) source = JSON.parse(source);
    const result = new LatLng();
    result.latitude = source["latitude"];
    result.longitude = source["longitude"];
    return result;
  }
}
export class Time {
  static createFrom(source: any) {
    if ("string" === typeof source) source = JSON.parse(source);
    const result = new Time();
    return result;
  }
}
export class ShopEntity {
  name: string;
  prefecture: string;
  address: string;
  series: string[];
  created_at: Time;
  updated_at: Time;
  location: LatLng;
  latitude: number;
  longitude: number;
  deleted: boolean;

  static createFrom(source: any) {
    if ("string" === typeof source) source = JSON.parse(source);
    const result = new ShopEntity();
    result.name = source["name"];
    result.prefecture = source["prefecture"];
    result.address = source["address"];
    result.series = source["series"];
    result.created_at = source["created_at"]
      ? Time.createFrom(source["created_at"])
      : null;
    result.updated_at = source["updated_at"]
      ? Time.createFrom(source["updated_at"])
      : null;
    result.location = source["location"]
      ? LatLng.createFrom(source["location"])
      : null;
    result.latitude = source["latitude"];
    result.longitude = source["longitude"];
    result.deleted = source["deleted"];
    return result;
  }
}
