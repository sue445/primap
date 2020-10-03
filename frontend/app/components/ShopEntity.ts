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
export class Geography {
  geohash: string;
  geopoint: LatLng;

  static createFrom(source: any) {
    if ("string" === typeof source) source = JSON.parse(source);
    const result = new Geography();
    result.geohash = source["geohash"];
    result.geopoint = source["geopoint"]
      ? LatLng.createFrom(source["geopoint"])
      : null;
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
  geography: Geography;
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
    result.geography = source["geography"]
      ? Geography.createFrom(source["geography"])
      : null;
    result.deleted = source["deleted"];
    return result;
  }
}
