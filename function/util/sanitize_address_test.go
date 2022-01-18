package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SanitizeAddress(t *testing.T) {
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
		{
			args: args{
				address: "宮城県仙台市太白区７丁目２０−３",
			},
			want: "宮城県仙台市太白区7-20-3",
		},
		{
			args: args{
				address: "千葉県流山市おおたかの森南１丁目5ｰ1 流山おおたかの森SC ３Ｆ",
			},
			want: "千葉県流山市おおたかの森南1-5-1",
		},
		{
			args: args{
				address: "栃木県宇都宮市インターパーク６－１－１ＦＫＤショピングモールインターパーク店２階",
			},
			want: "栃木県宇都宮市インターパーク6-1-1",
		},
		{
			args: args{
				address: "千葉県成田市赤坂２ー１ー１０ボンベルタ成田３ＦＦｏｒｋｉｄｓ’ｂｙこぐま成田店",
			},
			want: "千葉県成田市赤坂2-1-10",
		},
		{
			args: args{
				address: "京都府京都市右京区西院追分町２５－１イオンモール京都五条３Ｆ",
			},
			want: "京都府京都市右京区西院追分町25-1",
		},
		{
			args: args{
				address: "福島県いわき市小名浜港背後地土地区画整理事業地内 ｲｵﾝ4F",
			},
			want: "福島県いわき市小名浜港背後地土地区画整理事業地内",
		},
		{
			args: args{
				address: "福島県いわき市小名浜港背後地震災復興土地区画整理事業地内イオンスタイルいわき小名浜４Ｆ",
			},
			want: "福島県いわき市小名浜港背後地震災復興土地区画整理事業地内",
		},
		{
			args: args{
				address: "福島県 いわき市小名浜港背後地震災復興土地区画整理事業地内イオンモールいわき小名浜イオンスタイルいわき小名浜4階",
			},
			want: "福島県 いわき市小名浜港背後地震災復興土地区画整理事業地内",
		},
		{
			args: args{
				address: "福島県福島市南矢野目字西荒田50-の17 ｲｵﾝ福島店3階",
			},
			want: "福島県福島市南矢野目字西荒田50-17",
		},
	}
	for _, tt := range tests {
		t.Run(tt.args.address, func(t *testing.T) {
			got := SanitizeAddress(tt.args.address)
			assert.Equal(t, tt.want, got)
		})
	}
}
