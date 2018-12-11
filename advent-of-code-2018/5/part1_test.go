package main

import "testing"

func Test_isOpposite(t *testing.T) {
	type args struct {
		a byte
		b byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"",
			args{'a', 'A'},
			true,
		},
		{
			"",
			args{'A', 'A'},
			false,
		},
		{
			"",
			args{'b', 'A'},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isOpposite(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("isOpposite() = %v, want %v", got, tt.want)
			}
		})
	}
}
