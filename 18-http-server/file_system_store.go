package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"slices"
	"sort"
)

type FileSystemPlayerStore struct {
	database *os.File
}

func (f FileSystemPlayerStore) GetMany() ([]Player, error) {
	f.database.Seek(0, io.SeekStart)

	var players []Player
	err := json.NewDecoder(f.database).Decode(&players)
	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)
	}

	sort.Slice(players, func(a, b int) bool {
		return players[b].Score < players[a].Score
	})

	// Or
	// ---
	// slices.SortFunc(players, func(a, b Player) int {
	// 	return b.Score - a.Score
	// })

	return players, err
}

func (f FileSystemPlayerStore) GetPlayerScore(name string) (int, bool) {
	f.database.Seek(0, io.SeekStart)

	player, ok := f.GetPlayer(name)

	return player.Score, ok
}

func (f FileSystemPlayerStore) GetPlayer(name string) (Player, bool) {
	f.database.Seek(0, io.SeekStart)

	players, _ := f.GetMany()

	i := slices.IndexFunc(players, func(p Player) bool {
		return p.Name == name
	})

	if i == -1 {
		return Player{}, false
	}

	return players[i], true
}

func (f FileSystemPlayerStore) Add(name string) error {
	players, _ := f.GetMany()

	i := slices.IndexFunc(players, func(p Player) bool {
		return p.Name == name
	})

	if i == -1 {
		player := Player{name, 0}
		players = append(players, player)
		i = len(players) - 1
	}

	mu.Lock()
	players[i].Score++
	mu.Unlock()

	f.database.Truncate(0)
	f.database.Seek(0, io.SeekStart)

	err := json.NewEncoder(f.database).Encode(&players)
	if err != nil {
		return fmt.Errorf("problem encoding, %v", err)
	}

	return nil
}
