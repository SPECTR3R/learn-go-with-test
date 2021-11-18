package main

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func createTempFile(t testing.TB, initialData string) (io.ReadWriteSeeker, func()) {
	t.Helper()
	tmpfile, err := ioutil.TempFile("", "db")
	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}
	tmpfile.Write([]byte(initialData))
	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}
	return tmpfile, removeFile
}

func TestFileSystemStorage(t *testing.T) {
	t.Run("test GetLeague", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t,
			`[
				{"Name": "Cleo", "Wins": 10},
				{"Name": "Chris", "Wins": 33}
			]`)
		defer cleanDatabase()

		store := FileSystemPlayerStore{database}
		got := store.GetLeague()

		want := []Player{
			{"Cleo", 10},
			{"Chris", 33},
		}
		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t,
			`[
				{"Name": "Cleo", "Wins": 10},
				{"Name": "Chris", "Wins": 33}
			]`)
		defer cleanDatabase()
		store := FileSystemPlayerStore{database}

		got := store.GetPlayerScore("Chris")
		assertScoreEquals(t, got, 33)
		got = store.GetPlayerScore("Cleo")
		assertScoreEquals(t, got, 10)
	})
}

func assertScoreEquals(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
