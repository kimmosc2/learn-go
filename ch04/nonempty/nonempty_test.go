package nonempty

import (
	"reflect"
	"testing"
)

func Test_nonempty(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{[]string{}}, []string{}},
		{"2", args{[]string{"1", "", "", "2", "", "3", "4", "5"}}, []string{"1", "2", "3", "4", "5"}},
		{"3", args{[]string{"1", "", "2", "", "3", "4", ""}}, []string{"1", "2", "3", "4"}},
		{"4", args{[]string{""}}, []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nonempty(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nonempty() = %v, want %v", got, tt.want)
			}
		})
	}
}
