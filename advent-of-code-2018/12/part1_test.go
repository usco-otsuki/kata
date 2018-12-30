package main

import (
	"reflect"
	"testing"
)

//

func TestAddDots(t *testing.T) {
	type args struct {
		state []byte
	}
	tests := []struct {
		name  string
		args  args
		want  []byte
		want1 int
	}{
		{"test1", args{[]byte("..#.#.")}, []byte("....#.#...."), 2},
		{"test2", args{[]byte("....#.#.")}, []byte("....#.#...."), 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := AddDots(tt.args.state)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddDots() got = %s, want %s", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("AddDots() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
