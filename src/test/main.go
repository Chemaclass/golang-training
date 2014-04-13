package test

import (
	"even"
	"testing"
)
func TestEven(t *testing.T){
	if !even.Even(2) {
		t.Log("2 should be even!")
		t.Fail()
	}
}