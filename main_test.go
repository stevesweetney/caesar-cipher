package main

import "testing"

var path = "./benchFiles/words.txt"
var a Lines = []string{}

func init() {
	for i := 0; i < 1000000; i++ {
		a = append(a, "The quick brown fox jumps over the lazy dog")
	}
}

func TestReadFile(t *testing.T) {
	_, err := ReadFile(path)

	if err != nil {
		t.Error("Error Reading file")
	}
}

func BenchmarkCipherAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a.CipherAll(13)
	}
}
