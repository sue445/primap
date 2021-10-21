package util

import (
	"github.com/stretchr/testify/assert"
	"github.com/sue445/primap/prismdb"
	"testing"
)

func TestAggregateShops(t *testing.T) {
	shops := []*prismdb.Shop{
		{
			Name:       "ふぇすたらんど小野店",
			Prefecture: "兵庫県",
			Address:    "兵庫県小野市王子町８６８－１ イオン小野店２Ｆ",
			Series:     []string{"pripara"},
		},
		{
			Name:       "ふぇすたらんど小野",
			Prefecture: "兵庫県",
			Address:    "兵庫県小野市王子町８６８－１イオン小野店２Ｆ",
			Series:     []string{"prichan"},
		},
		{
			Name:       "モーリーファンタジー唐津",
			Prefecture: "佐賀県",
			Address:    "佐賀県唐津市鏡字立神４６７１ イオン２階",
			Series:     []string{"prichan", "pripara"},
		},
		{
			Name:       "プリズムストーン  新札幌",
			Prefecture: "北海道",
			Address:    "北海道札幌市厚別区厚別中央２条５丁目６－２　ＤＵＯ－１　４階－１イオン小野店２Ｆ",
			Series:     []string{"prichan"},
		},
		{
			Name:       "プリズムストーン札幌",
			Prefecture: "北海道",
			Address:    "北海道札幌市厚別区厚別中央２条５丁目６－２ ＤＵＯ－１ ４階",
			Series:     []string{"pripara"},
		},
		{
			Name:       "ＳＯＹＵＧａｍｅＦｉｅｌｄ湘南",
			Prefecture: "神奈川県",
			Address:    "神奈川県藤沢市辻堂新町四丁目１番１号 湘南モールＦＩＬＬ２Ｆ",
			Series:     []string{"pripara"},
		},
		{
			Name:       "ＳＯＹＵ　Ｇａｍｅ　Ｆｉｅｌｄ湘南店",
			Prefecture: "神奈川県",
			Address:    "神奈川県藤沢市辻堂新町四丁目１番１号　湘南モールＦＩＬＬ２Ｆ",
			Series:     []string{"prichan"},
		},
		{
			Name:       "ビックカメラ新宿西口本店",
			Prefecture: "東京都",
			Address:    "東京都新宿区西新宿１－１１－１",
			Series:     []string{"prichan"},
		},
		{
			Name:       "NICOPAウイングタウン岡崎",
			Prefecture: "愛知県",
			Address:    "愛知県岡崎市羽根町小豆坂3番地",
			Series:     []string{"prichan", "primagi"},
		},
		{
			Name:       "ＮＩＣＯＰＡウイングタウン岡崎",
			Prefecture: "愛知県",
			Address:    "愛知県岡崎市羽根町小豆坂３番地",
			Series:     []string{"pripara"},
		},
		{
			Name:       "モーリーファンタジー・f新潟南",
			Prefecture: "新潟県",
			Address:    "新潟県新潟市江南区下早通柳田１丁目１番１号イオンモール２Ｆ",
			Series:     []string{"prichan", "pripara"},
		},
		{
			Name:       "モーリーファンタジーf新潟南",
			Prefecture: "新潟県",
			Address:    "新潟県新潟市江南区下早通柳田1-1-1",
			Series:     []string{"primagi"},
		},
		{
			Name:       "LABI渋谷",
			Prefecture: "東京都",
			Address:    "東京都渋谷区道玄坂2-29-20",
			Series:     []string{"primagi"},
		},
		{
			Name:       "ヤマダ電機LABI渋谷",
			Prefecture: "東京都",
			Address:    "東京都渋谷区道玄坂２－２９－２０",
			Series:     []string{"prichan"},
		},
		{
			Name:       "プリズムストーン  東京駅",
			Prefecture: "東京都",
			Address:    "東京都千代田区丸の内１－９－１　東京駅一番街　Ｂ１Ｆ　東京キャラクターストリート内",
			Series:     []string{"prichan", "pripara"},
		},
		{
			Name:       "プリズムストーン東京駅",
			Prefecture: "東京都",
			Address:    "東京都千代田区丸の内1-9-1",
			Series:     []string{"primagi"},
		},
		{
			Name:       "SOYUGameField御所野",
			Prefecture: "秋田県",
			Address:    "秋田県秋田市御所野地蔵田一丁目１番地１号イオンモール秋田２Ｆ",
			Series:     []string{"prichan", "pripara"},
		},
		{
			Name:       "ソユーゲームフィールド御所野",
			Prefecture: "秋田県",
			Address:    "秋田県秋田市御所野地蔵田一丁目1番地1号 ｲｵﾝﾓｰﾙ秋田2F",
			Series:     []string{"primagi"},
		},
		{
			Name:       "ＳＯＹＵＧａｍｅＦｉｅｌｄ長野",
			Prefecture: "長野県",
			Address:    "長野県長野市三輪九丁目４３番２４号 イオンタウン長野三輪２Ｆ",
			Series:     []string{"pripara"},
		},
		{
			Name:       "ＳＯＹＵＧａｍｅＦｉｅｌｄ長野三輪",
			Prefecture: "長野県",
			Address:    "長野県長野市三輪九丁目４３番２４号イオンタウン長野三輪２Ｆ",
			Series:     []string{"prichan"},
		},
		{
			Name:       "ＳＯＹＵＺＡＵＲＵＳＷＯＲＬＤ大森",
			Prefecture: "東京都",
			Address:    "東京都大田区大森北二丁目１３－１　イトーヨーカドー大森店３Ｆ",
			Series:     []string{"prichan", "pripara"},
		},
		{
			Name:       "ソユーザウルスワールド 大森",
			Prefecture: "東京都",
			Address:    "東京都大田区大森北二丁目13-1  ｲﾄｰﾖｰｶﾄﾞｰ大森店3F",
			Series:     []string{"primagi"},
		},
		{
			Name:       "タイトーステーション　ＢＩＧＦＵＮ平和島",
			Prefecture: "東京都",
			Address:    "東京都大田区平和島１－１－１ＢＩＧＦＵＮ平和島３Ｆ",
			Series:     []string{"prichan", "primagi"},
		},
		{
			Name:       "タイトーステーションＢＩＧＦＵＮ平和島店",
			Prefecture: "東京都",
			Address:    "東京都大田区平和島１－１－１ＢＩＧＦＵＮ平和島３Ｆ",
			Series:     []string{"pripara"},
		},
	}

	got := AggregateShops(shops)

	want := []*prismdb.Shop{
		{
			Name:       "LABI渋谷",
			Prefecture: "東京都",
			Address:    "東京都渋谷区道玄坂2-29-20",
			Series:     []string{"prichan", "primagi"},
		},
		{
			Name:       "NICOPAウイングタウン岡崎",
			Prefecture: "愛知県",
			Address:    "愛知県岡崎市羽根町小豆坂3番地",
			Series:     []string{"prichan", "primagi", "pripara"},
		},
		{
			Name:       "ふぇすたらんど小野",
			Prefecture: "兵庫県",
			Address:    "兵庫県小野市王子町８６８－１ イオン小野店２Ｆ",
			Series:     []string{"prichan", "pripara"},
		},
		{
			Name:       "ソユーゲームフィールド御所野",
			Prefecture: "秋田県",
			Address:    "秋田県秋田市御所野地蔵田一丁目１番地１号イオンモール秋田２Ｆ",
			Series:     []string{"prichan", "primagi", "pripara"},
		},
		{
			Name:       "ソユーゲームフィールド湘南",
			Prefecture: "神奈川県",
			Address:    "神奈川県藤沢市辻堂新町四丁目１番１号 湘南モールＦＩＬＬ２Ｆ",
			Series:     []string{"prichan", "pripara"},
		},
		{
			Name:       "ソユーゲームフィールド長野三輪",
			Prefecture: "長野県",
			Address:    "長野県長野市三輪九丁目４３番２４号 イオンタウン長野三輪２Ｆ",
			Series:     []string{"prichan", "pripara"},
		},
		{
			Name:       "ソユーザウルスワールド大森",
			Prefecture: "東京都",
			Address:    "東京都大田区大森北二丁目１３－１　イトーヨーカドー大森店３Ｆ",
			Series:     []string{"prichan", "primagi", "pripara"},
		},
		{
			Name:       "タイトーステーションBIGFUN平和島",
			Prefecture: "東京都",
			Address:    "東京都大田区平和島１－１－１ＢＩＧＦＵＮ平和島３Ｆ",
			Series:     []string{"prichan", "primagi", "pripara"},
		},
		{
			Name:       "ビックカメラ新宿西口本店",
			Prefecture: "東京都",
			Address:    "東京都新宿区西新宿１－１１－１",
			Series:     []string{"prichan"},
		},
		{
			Name:       "プリズムストーン札幌",
			Prefecture: "北海道",
			Address:    "北海道札幌市厚別区厚別中央２条５丁目６－２　ＤＵＯ－１　４階－１イオン小野店２Ｆ",
			Series:     []string{"prichan", "pripara"},
		},
		{
			Name:       "プリズムストーン東京駅",
			Prefecture: "東京都",
			Address:    "東京都千代田区丸の内１－９－１　東京駅一番街　Ｂ１Ｆ　東京キャラクターストリート内",
			Series:     []string{"prichan", "primagi", "pripara"},
		},
		{
			Name:       "モーリーファンタジーf新潟南",
			Prefecture: "新潟県",
			Address:    "新潟県新潟市江南区下早通柳田１丁目１番１号イオンモール２Ｆ",
			Series:     []string{"prichan", "primagi", "pripara"},
		},
		{
			Name:       "モーリーファンタジー唐津",
			Prefecture: "佐賀県",
			Address:    "佐賀県唐津市鏡字立神４６７１ イオン２階",
			Series:     []string{"prichan", "pripara"},
		},
	}

	assert.Equal(t, want, got)
}

func Test_normalizeShopName(t *testing.T) {
	type args struct {
		shopName string
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{
				shopName: "モーリーファンタジー唐津",
			},
			want: "モーリーファンタジー唐津",
		},
		{
			args: args{
				shopName: "ＳＯＹＵＧａｍｅＦｉｅｌｄ湘南",
			},
			want: "ソユーゲームフィールド湘南",
		},
		{
			args: args{
				shopName: "ニコパウイングタウン岡崎",
			},
			want: "NICOPAウイングタウン岡崎",
		},
		{
			args: args{
				shopName: "モーリーファンタジー・f新潟南",
			},
			want: "モーリーファンタジーf新潟南",
		},
		{
			args: args{
				shopName: "ヤマダ電機LABI渋谷",
			},
			want: "LABI渋谷",
		},
		{
			args: args{
				shopName: "ＳＯＹＵＺＡＵＲＵＳＷＯＲＬＤ大森",
			},
			want: "ソユーザウルスワールド大森",
		},
		{
			args: args{
				shopName: "SOYUFamilyGameField防府",
			},
			want: "ソユーファミリーゲームフィールド防府",
		},
		{
			args: args{
				shopName: "SOYU Family Game Field花巻",
			},
			want: "ソユーファミリーゲームフィールド花巻",
		},
		{
			args: args{
				shopName: "THE3RDPLANETBiVi京都二条",
			},
			want: "THE 3RD PLANET BiVi京都二条",
		},
		{
			args: args{
				shopName: "THE3RDPLANETフレスポ国分",
			},
			want: "THE 3RD PLANETフレスポ国分",
		},
		{
			args: args{
				shopName: "タイトーステーション　ＢＩＧＦＵＮ平和島",
			},
			want: "タイトーステーションBIGFUN平和島",
		},
	}
	for _, tt := range tests {
		t.Run(tt.args.shopName, func(t *testing.T) {
			got := normalizeShopName(tt.args.shopName)
			assert.Equal(t, tt.want, got)
		})
	}
}
