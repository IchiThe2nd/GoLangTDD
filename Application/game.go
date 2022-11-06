package poker

//game maanages state of a game
type Game interface {
	Start(numberOfPlayers int)
	Finish(winner string)
}
