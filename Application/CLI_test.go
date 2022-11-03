//test CLI allows input ans stores as PlayerStore

package poker

import "testing"

func TestCLI(t *testing.T) {
	playerStore := &StubPlayerStore{}
	cli := &CLI{playerStore} //CLI is going to be struct with fdield of playerStore?
	cli.PlayPoker()          //call a method that awaits input and sends to updates playerstore field of cli
	if len(playerStore.winCalls) != 1 {
		t.Fatal("expected a win call but did not get any")
	}
}
