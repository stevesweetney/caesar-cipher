package main

import "testing"

var path = "./benchFiles/words.txt"

func TestReadFile(t *testing.T) {
	_, err := ReadFile(path)

	if err != nil {
		t.Error("Error Reading file")
	}
}

func BenchmarkCipherAll(b *testing.B) {
	linesBench, _ := ReadFile(path)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		linesBench.CipherAll(13)
	}
}
