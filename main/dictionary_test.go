package main

import "testing"

func TestSearch(t *testing.T) {
	const k = "test"
	const v = "this is just a test"
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

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
