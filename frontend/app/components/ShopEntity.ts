/* Do not change, this code is generated from Golang structs */

export class LatLng {
  latitude?: number;
  longitude?: number;

  static createFrom(source: any = {}) {
    return new LatLng(source);
  }

  constructor(source: any = {}) {
    if ("string" === typeof source) source = JSON.parse(source);
    this.latitude = source["latitude"];
    this.longitude = source["longitude"];
  }
}
export class Geography {
  geohash: string;
  geopoint?: LatLng;

  static createFrom(source: any = {}) {
    return new Geography(source);
  }

  constructor(source: any = {}) {
    if ("string" === typeof source) source = JSON.parse(source);
    this.geohash = source["geohash"];
    this.geopoint = this.convertValues(source["geopoint"], LatLng);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ("object" === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ShopEntity {
  name: string;
  prefecture: string;
  address: string;
  sanitized_address: string;
  series: string[];
  created_at: Date;
  updated_at: Date;
  geography?: Geography;
  deleted: boolean;

  static createFrom(source: any = {}) {
    return new ShopEntity(source);
  }

  constructor(source: any = {}) {
    if ("string" === typeof source) source = JSON.parse(source);
    this.name = source["name"];
    this.prefecture = source["prefecture"];
    this.address = source["address"];
    this.sanitized_address = source["sanitized_address"];
    this.series = source["series"];
    this.created_at = source["created_at"].toDate();
    this.updated_at = source["updated_at"].toDate();
    this.geography = this.convertValues(source["geography"], Geography);
    this.deleted = source["deleted"];
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ("object" === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
