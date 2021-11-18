package main

import (
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadSeeker
}

type ReadSeeker interface {
	Seeker
}

type Seeker interface {
	Seek(offset int64, whence int) (int64, error)
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	f.database.Seek(0, 0)
	league, _ := NewLeague(f.database)
	return league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	var wins int
	f.database.Seek(0, 0)
	for _, player := range f.GetLeague() {
		if player.Name == name {
			wins = player.Wins
			break
		}
	}
	return wins
}
