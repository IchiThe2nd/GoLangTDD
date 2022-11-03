//test CLI allows input ans stores as PlayerStore

package poker_test

import (
	"poker"
	"strings"
	"testing"
)

// CLI_test.go
func TestCLI(t *testing.T) {

	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Chris")
	})
	/*
		t.Run("record cleo win from user input", func(t *testing.T) {
			in := strings.NewReader("Cleo wins\n")
			playerStore := &poker.StubPlayerStore{}

			cli := &poker.CLI{playerStore, in}
			cli.PlayPoker()

			poker.AssertPlayerWin(t, playerStore, "Cleo")
		})
	*/
}
