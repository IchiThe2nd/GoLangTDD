package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	t.Run("sleep before each print", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}
		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted %v got %v", want, spySleepPrinter.Calls)
		}
	})

	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &SpyCountdownOperations{})
		got := buffer.String()
		want := `3
2
1
GO!`

		if got != want {
			t.Errorf("got %q wanted %q", got, want)

		}
	})

}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v nut slept %v", sleepTime, spyTime.durationSlept)
	}
}
