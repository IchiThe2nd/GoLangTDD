package poker

import (
	"os"
	"testing"

func createTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()
	tmpfile, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatalf("could not creat temp file %v", err)
	}

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}

func TestFileSystemStore(t *testing.T) {
	/*
		t.Run("League from a reader", func(t *testing.T) {
			database, cleanDatabase := createTempFile(t, `[
				{"Name":"Cleo","Wins":10},
				{"Name":"Chris","Wins":33}]`)
			defer cleanDatabase()

			store, err := NewFileSystemPlayerStore(database)
			if err != nil {
				t.Errorf("error creating filesystemPlayer store %v ", err)
			}

			got := store.GetLeague()

			want := []Player{
				{"Cleo", 10},
				{"Chris", 33},
			}
			assertLeague(t, got, want)

			//read twice
			got = store.GetLeague()
			assertLeague(t, got, want)
		})
	*/
	t.Run("Get Player Score forom file", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name":"Cleo","Wins":10},
			{"Name":"Chris","Wins":33},
			{"Name":"Shitty","Wins":16}
		]`)

		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		if err != nil {
			t.Errorf("error creating filesystemPlayer store %v ", err)
		}

		got := store.GetPlayerScore("Chris")
		want := 33

		assertScoreEquals(t, got, want)

	})

	t.Run("store win for existing player", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name":"Cleo","Wins": 10},
			{"Name":"Chris","Wins": 33}
		]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)

		store.RecordWin("Chris")

		got := store.GetPlayerScore("Chris")
		want := 34
		assertScoreEquals(t, got, want)

	})

	t.Run("store win forNEW player", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name":"Cleo","Wins": 10},
			{"Name":"Chris","Wins": 33}
		]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)
		store.RecordWin("Pepper")

		got := store.GetPlayerScore("Pepper")
		want := 1
		assertScoreEquals(t, got, want)

	})

	t.Run("works with an empty file", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, "")
		defer cleanDatabase()

		_, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)
	})

	t.Run("league sorted", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name":"Cleo","Wins":10},
			{"Name":"Chris","Wins":33}
		]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)

		assertNoError(t, err)

		got := store.GetLeague()
		want := []Player{
			{"Chris", 33},
			{"Cleo", 10},
		}

		assertLeague(t, got, want)

		got = store.GetLeague()
		assertLeague(t, got, want)
	})
}

func assertScoreEquals(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d wanted %d", got, want)
	}

}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("Did not expect error buyt got this fucker %v ", err)
	}

}
