package Tutorials

import "testing"

func TestPrime(t *testing.T) {
	i := 0
	if isPrime(i) {
		t.Errorf("%d is a Not Prime", i)
	}
	i = 1
	if isPrime(i) {
		t.Errorf("%d is a Not Prime", i)
	}
	i = 2
	if !isPrime(i) {
		t.Errorf("%d is a Prime", i)
	}
	i = 3
	if !isPrime(i) {
		t.Errorf("%d is a Prime", i)
	}
	i = 5
	if !isPrime(i) {
		t.Errorf("%d is a Prime", i)
	}
	i = 7
	if !isPrime(i) {
		t.Errorf("%d is a Prime", i)
	}
	i = 11
	if !isPrime(i) {
		t.Errorf("%d is a Prime", i)
	}
	i = 13
	if !isPrime(i) {
		t.Errorf("%d is a Prime", i)
	}
	i = 17
	if !isPrime(i) {
		t.Errorf("%d is a Prime", i)
	}
	i = 19
	if !isPrime(i) {
		t.Errorf("%d is a Prime", i)
	}
	i = 23
	if !isPrime(i) {
		t.Errorf("%d is a Prime", i)
	}
	i = 29
	if !isPrime(i) {
		t.Errorf("%d is a Prime", i)
	}
	i = 31
	if !isPrime(i) {
		t.Errorf("%d is a Prime", i)
	}

}
