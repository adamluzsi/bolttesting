package bolttesting_test

import (
	"testing"

	test "github.com/adamluzsi/testing"
)

func TestTestEqIntsNotTheSame(t *testing.T) {

	a := []int{1, 2, 3}
	b := []int{1, 2, 3, 4}

	if test.TestEqInts(a, b) {
		t.Log("expectation failed")
		t.Fail()
	}

}

func TestTestEqIntsEverythingIsEq(t *testing.T) {

	a := []int{1, 2, 3}
	b := []int{1, 2, 3}

	if !test.TestEqInts(a, b) {
		t.Log("expectation failed")
		t.Fail()
	}

}

func TestTestEqIntsSliceEqButNotInTheSameOrder(t *testing.T) {

	a := []int{1, 2, 3}
	b := []int{1, 3, 2}

	if !test.TestEqInts(a, b) {
		t.Log("expectation failed")
		t.Fail()
	}

}

func TestTestEqBytesByteSliceTheSame(t *testing.T) {

	a := []byte{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100, 33}
	b := []byte{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100, 33}

	if !test.TestEqBytes(a, b) {
		t.Log("expectation failed")
		t.Fail()
	}

}

func TestTestEqByteSliceIsDifferent(t *testing.T) {

	a := []byte{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100, 33}
	b := []byte{72, 101, 108, 108, 32, 111, 87, 111, 114, 108, 100, 33}

	if test.TestEqBytes(a, b) {
		t.Log("expectation failed")
		t.Fail()
	}

}
