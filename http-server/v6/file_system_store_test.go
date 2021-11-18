package main

import (
	"strings"
	"testing"
)

func TestFileSystemStorage(t *testing.T) {
	t.Run("test GetLeague", func(t *testing.T) {
		database := strings.NewReader(`
			[
				{"Name": "Cleo", "Wins": 10},
				{"Name": "Chris", "Wins": 33}
			]`)

		store := FileSystemPlayerStore{database}
		got := store.GetLeague()

		want := []Player{
			{"Cleo", 10},
			{"Chris", 33},
		}
		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database := strings.NewReader(`[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
		]`)

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
