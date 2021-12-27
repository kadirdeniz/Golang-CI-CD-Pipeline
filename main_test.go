package main

import "testing"

func anasayfaTest(t *testing.T) {
	expect := "Anasayfa"
	got := anasayfa()

	if expect != got {
		t.Errorf("Error")
	}
}
