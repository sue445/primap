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
		name string
		args args
		want string
	}{
		{
			name: "東京都新宿区新宿３－２６－７ 玩具売場",
			args: args{
				address: "東京都新宿区新宿３－２６－７ 玩具売場",
			},
			want: "東京都新宿区新宿3-26-7",
		},
		{
			name: "東京都新宿区新宿３－２６－７",
			args: args{
				address: "東京都新宿区新宿３－２６－７",
			},
			want: "東京都新宿区新宿3-26-7",
		},
		{
			name: "東京都新宿区新宿３－２６",
			args: args{
				address: "東京都新宿区新宿３－２６",
			},
			want: "東京都新宿区新宿3-26",
		},
		{
			name: "東京都新宿区新宿３",
			args: args{
				address: "東京都新宿区新宿３",
			},
			want: "東京都新宿区新宿3",
		},
		{
			name: "東京都新宿区新宿",
			args: args{
				address: "東京都新宿区新宿",
			},
			want: "東京都新宿区新宿",
		},
		{
			name: "東京都新宿区西新宿１－５－１ ハルク５Ｆ トイズコーナー",
			args: args{
				address: "東京都新宿区西新宿１－５－１ ハルク５Ｆ トイズコーナー",
			},
			want: "東京都新宿区西新宿1-5-1",
		},
		{
			name: "福岡県福岡市西区徳永１１３－１　玩具売場",
			args: args{
				address: "福岡県福岡市西区徳永１１３－１　玩具売場",
			},
			want: "福岡県福岡市西区徳永113-1",
		},
		{
			name: "福岡県福岡市西区徳永１１３　玩具売場",
			args: args{
				address: "福岡県福岡市西区徳永１１３　玩具売場",
			},
			want: "福岡県福岡市西区徳永113",
		},
		{
			name: "福島県いわき市平６丁目６番地２　　イトーヨーカドー平店内　プレビプレイランドコーナー　こころっこ",
			args: args{
				address: "福島県いわき市平６丁目６番地２　　イトーヨーカドー平店内　プレビプレイランドコーナー　こころっこ",
			},
			want: "福島県いわき市平6-6-2",
		},
		{
			name: "福島県いわき市平６丁目６番地　　イトーヨーカドー平店内　プレビプレイランドコーナー　こころっこ",
			args: args{
				address: "福島県いわき市平６丁目６番地　　イトーヨーカドー平店内　プレビプレイランドコーナー　こころっこ",
			},
			want: "福島県いわき市平6-6",
		},
		{
			name: "岡山県高梁市中原町１０８４番地の１ポルカ天満屋ハッピータウン内２階",
			args: args{
				address: "岡山県高梁市中原町１０８４番地の１ポルカ天満屋ハッピータウン内２階",
			},
			want: "岡山県高梁市中原町1084-1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sanitizeAddress(tt.args.address)
			assert.Equal(t, tt.want, got)
		})
	}
}
