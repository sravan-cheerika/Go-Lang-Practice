package main

import "testing"

//func TestSum(t *testing.T){
func TestSum(t *testing.T){

	numbers := []int{1,2,3,4,5}

	got := Sum(numbers)
	want := 19

	if want != got {
		t.Errorf("got %d want %d ", got, want)
	}


}