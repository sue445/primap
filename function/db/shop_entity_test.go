package db

import (
	"github.com/stretchr/testify/assert"
	_ "github.com/tkrajina/typescriptify-golang-structs/typescriptify"
	"testing"
)

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
	}
	for _, tt := range tests {
		t.Run(tt.args.address, func(t *testing.T) {
			got := sanitizeAddress(tt.args.address)
			assert.Equal(t, tt.want, got)
		})
	}
}
