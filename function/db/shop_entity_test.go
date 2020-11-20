package db

import (
	"context"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/sue445/primap/testutil"
	_ "github.com/tkrajina/typescriptify-golang-structs/typescriptify"
	"google.golang.org/genproto/googleapis/type/latlng"
	"os"
	"testing"
	"time"
)

func TestShopEntity_UpdateAddressWithGeography(t *testing.T) {
	os.Setenv("GOOGLE_MAPS_API_KEY", "DUMMY_API_KEY")
	defer func() {
		os.Unsetenv("GOOGLE_MAPS_API_KEY")
	}()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://maps.googleapis.com/maps/api/geocode/json?address=%E6%9D%B1%E4%BA%AC%E9%83%BD%E6%96%B0%E5%AE%BF%E5%8C%BA%E6%96%B0%E5%AE%BF3-26-7&key=DUMMY_API_KEY",
		httpmock.NewStringResponder(200, testutil.ReadTestData("testdata/geocoding_contains_results.json")))
	httpmock.RegisterResponder("GET", "https://maps.googleapis.com/maps/api/geocode/json?address=DUMMY_ADDRESS&key=DUMMY_API_KEY",
		httpmock.NewStringResponder(200, testutil.ReadTestData("testdata/geocoding_zero_results.json")))

	type fields struct {
		Name             string
		Prefecture       string
		Address          string
		SanitizedAddress string
		Series           []string
		CreatedAt        time.Time
		UpdatedAt        time.Time
		Geography        *Geography
		Deleted          bool
	}
	type args struct {
		ctx     context.Context
		address string
	}
	type wants struct {
		Address          string
		SanitizedAddress string
		Geography        *Geography
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wants  wants
	}{
		{
			name:   "Update SanitizedAddress and Geography",
			fields: fields{},
			args: args{
				ctx:     context.Background(),
				address: "東京都新宿区新宿３－２６－７ 玩具売場",
			},
			wants: wants{
				Address:          "東京都新宿区新宿３－２６－７ 玩具売場",
				SanitizedAddress: "東京都新宿区新宿3-26-7",
				Geography: &Geography{
					GeoHash: "xn774crz0",
					GeoPoint: &latlng.LatLng{
						Latitude:  35.6916892,
						Longitude: 139.7018233,
					},
				},
			},
		},
		{
			name: "ZERO_RESULTS",
			fields: fields{
				Geography: &Geography{
					GeoHash: "xxxxxxxx",
					GeoPoint: &latlng.LatLng{
						Latitude:  0,
						Longitude: 0,
					},
				},
			},
			args: args{
				ctx:     context.Background(),
				address: "DUMMY_ADDRESS",
			},
			wants: wants{
				Address:          "DUMMY_ADDRESS",
				SanitizedAddress: "DUMMY_ADDRESS",
				Geography:        nil,
			},
		},
		{
			name: "Not updated",
			fields: fields{
				Address:          "Address",
				SanitizedAddress: "Address",
				Geography: &Geography{
					GeoHash: "xxxxxxxx",
					GeoPoint: &latlng.LatLng{
						Latitude:  0,
						Longitude: 0,
					},
				},
			},
			args: args{
				ctx:     context.Background(),
				address: "Address",
			},
			wants: wants{
				Address:          "Address",
				SanitizedAddress: "Address",
				Geography: &Geography{
					GeoHash: "xxxxxxxx",
					GeoPoint: &latlng.LatLng{
						Latitude:  0,
						Longitude: 0,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ShopEntity{
				Name:             tt.fields.Name,
				Prefecture:       tt.fields.Prefecture,
				Address:          tt.fields.Address,
				SanitizedAddress: tt.fields.SanitizedAddress,
				Series:           tt.fields.Series,
				CreatedAt:        tt.fields.CreatedAt,
				UpdatedAt:        tt.fields.UpdatedAt,
				Geography:        tt.fields.Geography,
				Deleted:          tt.fields.Deleted,
			}

			err := e.UpdateAddressWithGeography(tt.args.ctx, tt.args.address)
			if assert.NoError(t, err) {
				assert.Equal(t, tt.wants.Address, e.Address)
				assert.Equal(t, tt.wants.SanitizedAddress, e.SanitizedAddress)
				assert.Equal(t, tt.wants.Geography, e.Geography)
			}
		})
	}
}

func Test_sanitizeAddress(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{
				address: "東京都新宿区新宿３－２６－７ 玩具売場",
			},
			want: "東京都新宿区新宿3-26-7",
		},
		{
			args: args{
				address: "東京都新宿区新宿３－２６－７",
			},
			want: "東京都新宿区新宿3-26-7",
		},
		{
			args: args{
				address: "東京都新宿区新宿３－２６",
			},
			want: "東京都新宿区新宿3-26",
		},
		{
			args: args{
				address: "東京都新宿区新宿３",
			},
			want: "東京都新宿区新宿3",
		},
		{
			args: args{
				address: "東京都新宿区新宿",
			},
			want: "東京都新宿区新宿",
		},
		{
			args: args{
				address: "東京都新宿区西新宿１－５－１ ハルク５Ｆ トイズコーナー",
			},
			want: "東京都新宿区西新宿1-5-1",
		},
		{
			args: args{
				address: "福岡県福岡市西区徳永１１３－１　玩具売場",
			},
			want: "福岡県福岡市西区徳永113-1",
		},
		{
			args: args{
				address: "福岡県福岡市西区徳永１１３　玩具売場",
			},
			want: "福岡県福岡市西区徳永113",
		},
		{
			args: args{
				address: "福島県いわき市平６丁目６番地２　　イトーヨーカドー平店内　プレビプレイランドコーナー　こころっこ",
			},
			want: "福島県いわき市平6-6-2",
		},
		{
			args: args{
				address: "福島県いわき市平６丁目６番地　　イトーヨーカドー平店内　プレビプレイランドコーナー　こころっこ",
			},
			want: "福島県いわき市平6-6",
		},
		{
			args: args{
				address: "岡山県高梁市中原町１０８４番地の１ポルカ天満屋ハッピータウン内２階",
			},
			want: "岡山県高梁市中原町1084-1",
		},
		{
			args: args{
				address: "北海道岩見沢市大和４条８丁目１　玩具売場",
			},
			want: "北海道岩見沢市大和4条8-1",
		},
		{
			args: args{
				address: "北海道帯広市西４条南２０丁目１　玩具売場",
			},
			want: "北海道帯広市西4条南20-1",
		},
		{
			args: args{
				address: "京都府長岡京市開田４丁目７番１号",
			},
			want: "京都府長岡京市開田4-7-1",
		},
		{
			args: args{
				address: "北海道帯広市稲田町南８線西１０－１玩具売場",
			},
			want: "北海道帯広市稲田町南8線西10-1",
		},
		{
			args: args{
				address: "滋賀県守山市播磨田町１８５の１",
			},
			want: "滋賀県守山市播磨田町185-1",
		},
		{
			args: args{
				address: "滋賀県守山市播磨田町１８５の１玩具売場",
			},
			want: "滋賀県守山市播磨田町185-1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.args.address, func(t *testing.T) {
			got := sanitizeAddress(tt.args.address)
			assert.Equal(t, tt.want, got)
		})
	}
}
