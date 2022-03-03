package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := "could not find the word you were looking for"

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertStrings(t, err.Error(), want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Add("test", "this is just a test")
		want := "this is just a test"
		assertError(t, err, nil)
		assertDefinition(t, dictionary, "test", want)
	})

	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		err := dictionary.Add("test", "this is just a test")
		want := "word already exists"
		if err == nil {
			t.Fatal("expected to get an error.")
		}
		assertError(t, err, ErrWordExists)
		assertStrings(t, err.Error(), want)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		err := dictionary.Update("test", "new definition")
		want := "new definition"
		assertError(t, err, nil)
		assertDefinition(t, dictionary, "test", want)
	})

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Update("test", "new definition")
		// want := "could not find the word you were looking for"
		assertError(t, err, ErrWordDoesNotExist)
		// assertStrings(t, err.Error(), want)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test definition"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)

	assertError(t, err, ErrNotFound)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dict Dictionary, key, value string) {
	t.Helper()
	got, err := dict.Search(key)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	if got != value {
		t.Errorf("got %q want %q", got, value)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
