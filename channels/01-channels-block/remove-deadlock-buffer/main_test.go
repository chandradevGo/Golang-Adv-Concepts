package main

import (
	"reflect"
	"testing"
)

func Test_processChannel(t *testing.T) {
	tests := []struct {
		name string
		want []int
	}{
		{
			name: "Test case 1",
			want: []int{42, 43},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processChannel(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("processChannel() = %v, want %v", got, tt.want)
			}
		})
	}
}
