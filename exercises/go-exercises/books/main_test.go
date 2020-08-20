package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCompare(t *testing.T) {

	tests := []struct {
		MapA     map[string]int
		MapB     map[string]int
		Expected map[string]int
	}{
		{
			MapA:     map[string]int{"cat": 1},
			MapB:     map[string]int{"cat": 1},
			Expected: map[string]int{"cat": 2},
		},
		{
			MapA:     map[string]int{"cat": 1},
			MapB:     map[string]int{"cat": 1, "dog": 2},
			Expected: map[string]int{"cat": 2},
		},
		{
			MapA:     map[string]int{"cat": 1},
			MapB:     map[string]int{"dog": 1},
			Expected: map[string]int{},
		},
	}

	for indx, tt := range tests {
		t.Run(fmt.Sprintf("%d", indx), func(t *testing.T) {
			found := findCommon(tt.MapA, tt.MapB)
			if !reflect.DeepEqual(found, tt.Expected) {
				fmt.Println(found)
				fmt.Println(tt.Expected)
				t.Fail()
			}
		})
	}

	// t.Fail()
}
