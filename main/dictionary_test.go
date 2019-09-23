package main

import "testing"

const k = "test"
const v = "this is just a test"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{k: v}

	t.Run("known word", func(t *testing.T) {
		value, _ := dictionary.Search(k)
		want := v

		assertStrings(t, value, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add(k, v)
	want := v
	got, err := dictionary.Search(k)

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if want != got {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
