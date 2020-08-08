package algorithm

import "testing"

// TestFib1 no cache
func TestFib1(t *testing.T) {
	type args struct {
		num int
	}
	var tests = []struct {
		name string
		args
		want int
	}{
		{"1", args{0}, 0},
		{"2", args{1}, 1},
		{"3", args{2}, 1},
		{"4", args{3}, 2},
		{"5", args{10}, 55},
		{"6", args{20}, 6765},
		{"7", args{40}, 102334155},
	}

	for _, test := range tests {
		if fib1(test.num) != test.want {
			t.Errorf("name:%v,fib(%d)=%v,but want:%v", test.name, test.num, fib1(test.num), test.want)
		}
	}
}

// TestFib2 has cache from memory
func TestFib2(t *testing.T) {
	type args struct {
		num int
	}
	var tests = []struct {
		name string
		args
		want int
	}{
		{"1", args{0}, 0},
		{"2", args{1}, 1},
		{"3", args{2}, 1},
		{"4", args{3}, 2},
		{"5", args{10}, 55},
		{"6", args{20}, 6765},
		{"7", args{40}, 102334155},
	}

	for _, test := range tests {
		fibCache = make([]int, test.num+1)
		if fib2(test.num) != test.want {
			t.Errorf("name:%v,fib(%d)=%v,but want:%v", test.name, test.num, fib2(test.num), test.want)
		}
	}
}
