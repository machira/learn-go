package maps

import "testing"

func TestSearch(t *testing.T) {
	t.Run("defined word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		_, got := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("undefined word", func(t *testing.T) {
		dictionary := Dictionary{"key": "value"}
		err, _ := dictionary.Search("invalidKey")

		if err != NoSuchWordError {
			t.Errorf("received the wrong error")
		}
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test", "this is just a test")

	want := "this is just a test"
	err, got := dictionary.Search("test")
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test definition"}

	dictionary.Delete(word)

	err, _ := dictionary.Search(word)
	if err != NoSuchWordError {
		t.Errorf("Expected %q to be deleted", word)
	}
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
