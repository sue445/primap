package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_subtractSliceWithContains(t *testing.T) {
	src := []string{"a", "b", "c"}
	sub := []string{"c", "d", "e"}

	got := subtractSliceWithContains(src, sub)

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

func Test_subtractSliceWithSet(t *testing.T) {
	src := []string{"a", "b", "c"}
	sub := []string{"c", "d", "e"}

	got := subtractSliceWithSet(src, sub)

	assert.Equal(t, []string{"a", "b"}, got)
}

func TestSortedSlice(t *testing.T) {
	type args struct {
		src []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "arg is sorted",
			args: args{
				src: []string{"1", "2"},
			},
			want: []string{"1", "2"},
		},
		{
			name: "arg isn't sorted",
			args: args{
				src: []string{"2", "1"},
			},
			want: []string{"1", "2"},
		},
		{
			name: "arg is empty",
			args: args{
				src: []string{},
			},
			want: []string{},
		},
		{
			name: "arg contains 1 element",
			args: args{
				src: []string{"1"},
			},
			want: []string{"1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SortedSlice(tt.args.src)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_UniqueSlice(t *testing.T) {
	src := []string{"a", "b", "c", "b"}

	got := UniqueSlice(src)

	assert.Equal(t, []string{"a", "b", "c"}, got)
}
