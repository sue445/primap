package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SubtractSlice(t *testing.T) {
	src := []string{"a", "b", "c"}
	sub := []string{"c", "d", "e"}

	got := SubtractSlice(src, sub)

	assert.Equal(t, []string{"a", "b"}, got)
}

func TestContains(t *testing.T) {
	slice := []string{"a", "b", "c"}

	type args struct {
		item string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Contains",
			args: args{
				item: "a",
			},
			want: true,
		},
		{
			name: "Not Contains",
			args: args{
				item: "d",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Contains(slice, tt.args.item)
			assert.Equal(t, tt.want, got)
		})
	}
}
