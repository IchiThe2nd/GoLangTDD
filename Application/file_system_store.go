package poker

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type FileSystemPlayerStore struct {
	database *json.Encoder
	league   League
}

func FileSystemPlayerStoreFromFilePath(path string) (*FileSystemPlayerStore, func(), error) {
	db, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		return nil, nil, fmt.Errorf("Problem oipening %s,%v", path, err)
	}
	closeFunc := func() {
		db.Close()
	}

	store, err := NewFileSystemPlayerStore(db)

	if err != nil {
		return nil, nil, fmt.Errorf("problem crrwatuing file system player store , %v ", err)

	}
	return store, closeFunc, nil
}

func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {
	//needed in case database was previuously used and readwriuteseeker was not at origin of file (this is probably kinda wrong)
	err := initialisePlayerDBFile(file)
	if err != nil {
		return nil, fmt.Errorf("problem initialising player db file, %v", err)
	}

	league, err := NewLeague(file)

	if err != nil {
		return nil, fmt.Errorf("problem loading the player store from file %s , %v", file.Name(), err)
	}
	return &FileSystemPlayerStore{
		database: json.NewEncoder(&tape{file}),
		league:   league,
	}, nil
}

// returns a league which is type  []Player
func (f *FileSystemPlayerStore) GetLeague() League {
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})
	return f.league
}
func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.league.Find(name)

	if player != nil {
		return player.Wins
	}
	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	player := f.league.Find(name)
	// player is a copy...but we are writing the copy back so doesnt matter
	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}
	//f.database.Seek(0, 0) nop longer needed as part of tape.write
	f.database.Encode(f.league)

}

func initialisePlayerDBFile(file *os.File) error {
	file.Seek(0, 0)
	info, err := file.Stat()

	if err != nil {
		return fmt.Errorf("Problem getting file info from %s,%v", file.Name(), err)
	}

	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, 0)
	}
	return nil
}

/* REVIEW BELOW when You get a chance
//range through  the League( returned from league) and increment wins
	league := f.GetLeague()
	for i, player := range league {
		if player.Name == name {
			//player.Name is a copy from league..(why cant I us e &player.Wins++)
			league[i].Wins++
			//&player.Wins = &player.Wins + 1
		}
	}
	//reset db seeker back
	f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(league)

*/
