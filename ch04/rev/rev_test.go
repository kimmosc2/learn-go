package main

import (
	"reflect"
	"testing"
)

func Test_reverse(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{1, 2, 3, 4, 5}}, []int{5, 4, 3, 2, 1}},
		{"2", args{[]int{1, 2, 3, 4}}, []int{4, 3, 2, 1}},
		{"3", args{[]int{}}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
