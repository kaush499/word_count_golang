package main

import (
	"os"
	"testing"
)

func TestGetNumberOfCharacters(t *testing.T) {
	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	got := getNumberOfCharacters(file)
	want := 339290

	if want != got {
		t.Errorf("Wrong number of characters counted. got: %d, want: %d", got, want)
	}
}

func TestGetNumberOfWords(t *testing.T) {
	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	got := getNumberOfWords(file)
	want := 58168

	if want != got {
		t.Errorf("Wrong number of words counted. got: %d, want: %d", got, want)
	}
}


