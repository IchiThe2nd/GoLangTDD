//test CLI allows input ans stores as PlayerStore

package poker_test

import (
	"bytes"
	"io"
	"poker"
	"strings"
	"testing"
)

// CLI_test.go

// dummys for dependencies
var dummyBlindAlerter = &poker.SpyBlindAlerter{}
var dummyPlayerStore = &poker.StubPlayerStore{}
var dummyStdIn = &bytes.Buffer{}
var dummyStdOut = &bytes.Buffer{}

type GameSpy struct {
	StartCalled     bool
	StartCalledWith int

	FinishedCalled   bool
	FinishCalledWith string
}

func (g *GameSpy) Start(numberOfPlayers int, alertsDestination io.Writer) {
	g.StartCalledWith = numberOfPlayers
	g.StartCalled = true
}

func (g *GameSpy) Finish(winner string) {
	g.FinishedCalled = true
	g.FinishCalledWith = winner
}

func userSends(messages ...string) io.Reader {
	return strings.NewReader(strings.Join(messages, "\n"))
}

func TestCLI(t *testing.T) {

	t.Run("it prompts for number of players and starts the game", func(t *testing.T) {
		game := &GameSpy{}
		stdout := &bytes.Buffer{}

		in := userSends("5", "Chris wins")

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt)
		assertGameStartedWith(t, game, 5)
		assertFinishCalledWith(t, game, "Chris")

	})

	t.Run("Start game with 8 players ands Cleo as wins", func(t *testing.T) {
		game := &GameSpy{}
		out := &bytes.Buffer{}

		in := userSends("8", "Cleo wins")

		poker.NewCLI(in, out, game).PlayPoker()

		assertGameStartedWith(t, game, 8)
		assertFinishCalledWith(t, game, "Cleo")

		assertMessagesSentToUser(t, out, poker.PlayerPrompt)
	})

	t.Run("prints an error if non-numeric value entered", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := userSends("Pies")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertGameNotStarted(t, game)
		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)

	})
}

func assertGameStartedWith(t testing.TB, game *GameSpy, numberOfPlayersWanted int) {
	t.Helper()
	if game.StartCalledWith != numberOfPlayersWanted {
		t.Errorf("wanted start called with %d but got %d", numberOfPlayersWanted, game.StartCalledWith)
	}
}

func assertGameNotStarted(t testing.TB, game *GameSpy) {
	t.Helper()
	if game.StartCalled {
		t.Errorf("game should not have started")
	}
}

func assertFinishCalledWith(t testing.TB, game *GameSpy, winner string) {
	t.Helper()
	if game.FinishCalledWith != winner {
		t.Errorf("expected finish called with %q but got %q", winner, game.FinishCalledWith)
	}

}

func assertMessagesSentToUser(t testing.TB, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := stdout.String()
	if got != want {
		t.Errorf("got %q sent to stdout but expected %+v", got, messages)
	}
}

func assertScheduledAlert(t testing.TB, got, want poker.SchedueledAlert) {
	t.Helper()
	if got != want {
		t.Errorf("got alert %v that does not match schedduled %v", got, want)
	}
}
