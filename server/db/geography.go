package db

import "google.golang.org/genproto/googleapis/type/latlng"

// Geography represents (geohash and geopoint)
type Geography struct {
	GeoHash  string         `firestore:"geohash"  json:"geohash"`
	GeoPoint *latlng.LatLng `firestore:"geopoint" json:"geopoint"`
}
