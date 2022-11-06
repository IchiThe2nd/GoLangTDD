package poker

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type CLI struct {
	in   *bufio.Scanner
	out  io.Writer
	game Game
}

func NewCLI(in io.Reader, out io.Writer, game Game) *CLI {
	return &CLI{
		in:   bufio.NewScanner(in),
		out:  out,
		game: game,
	}
}

// PlayerPrompt is the text asking the user for the number of players.
const PlayerPrompt = "Please enter the number of players: "

// BadPlayerInputErrMsg is the text telling the user they did bad things.
const BadPlayerInputErrMsg = "Bad value received for number of players, please try again with a number"

// BadWinnerInputMsg is the text telling the user they declared the winner wrong.
const BadWinnerInputMsg = "invalid winner input, expect format of 'PlayerName wins'"

func (cli *CLI) PlayPoker() {
	fmt.Fprint(cli.out, PlayerPrompt)
	numberOfPlayersInput := cli.readLine()
	numberOfPlayers, err := strconv.Atoi(strings.Trim(numberOfPlayersInput, "\n")) //we dont need errors!

	if err != nil {
		fmt.Fprint(cli.out, BadPlayerInputErrMsg)
		return
	}

	cli.game.Start(numberOfPlayers)

	winnerInput := cli.readLine()
	cli.game.Finish(extractWinner(winnerInput))
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}
