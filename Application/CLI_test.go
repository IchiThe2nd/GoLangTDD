//test CLI allows input ans stores as PlayerStore

package poker

import (
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {

	t.Run("record chris wins from input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n") //an *io.Reader
		playerStore := &StubPlayerStore{}

		cli := &CLI{playerStore, in} //CLI is going to be struct with fdield of playerStore?
		cli.PlayPoker()              //call a method that awaits input and sends to updates playerstore field of cli

		assertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record Cleo wins from input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n") //an *io.Reader
		playerStore := &StubPlayerStore{}

		cli := &CLI{playerStore, in} //CLI is going to be struct with fdield of playerStore?
		cli.PlayPoker()              //call a method that awaits input and sends to updates playerstore field of cli

		assertPlayerWin(t, playerStore, "Cleo")
	})

}
