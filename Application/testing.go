//test stubs for users

package poker

import (
	"fmt"
	"io"
	"testing"
	"time"
)

type StubPlayerStore struct {
	Scores   map[string]int
	WinCalls []string
	League   League
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.Scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.WinCalls = append(s.WinCalls, name)
}

func (s *StubPlayerStore) GetLeague() League {
	return s.League
}

func AssertPlayerWin(t testing.TB, store *StubPlayerStore, winner string) {
	t.Helper()
	if len(store.WinCalls) != 1 {
		t.Fatalf("got %d calls to record wins wante %d", len(store.WinCalls), 1)
	}
	if store.WinCalls[0] != winner {
		t.Errorf("did not store correct winner got %q wanted %q", store.WinCalls[0], winner)

	}
}

// spy blinder allows you to spy on scheduled alerts

type SpyBlindAlerter struct {
	Alerts []SchedueledAlert
}

type SchedueledAlert struct {
	At     time.Duration
	Amount int
}

func (s *SpyBlindAlerter) ScheduleAlertAt(at time.Duration, amount int, to io.Writer) {
	s.Alerts = append(s.Alerts, SchedueledAlert{at, amount})
}

func (s SchedueledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.Amount, s.At)
}
