package tenniskata

const (
	loveScore = "Love"
	fifteenScore = "Fifteen"
	thirtyScore = "Thirty"
	fortyScore = "Forty"
)

type score struct {
	value int
}

func (s* score) increase() {
	s.value += 1
}

func (s score) print() string {
	switch s.value {
	case 0:
		return loveScore
	case 1:
		return fifteenScore
	case 2:
		return thirtyScore
	case 3:
		return fortyScore
	default:
		return ""
	}
}

type player struct {
	score score
	result string
	name string
}

type tennisGame2 struct {
	player1 player
	player2 player
}

func TennisGame2(player1Name string, player2Name string) TennisGame {
	return &tennisGame2{
		player1: player{name: player1Name},
		player2: player{name: player2Name},
	}
}

func (game *tennisGame2) GetScore() string {
	if game.player1Score() == game.player2Score() {
		if game.player1Score() >= 3 {
			return "Deuce"
		} else {
			return game.printPlayer1Score() + "-All"
		}
	}

	if game.shouldPrintOverallScore() {
		return game.printOverallScore()
	}


	if score := compareScores("player1", game.player1Score(), game.player2Score()); score == "" {
		return compareScores("player2", game.player2Score(), game.player1Score())
	} else {
		return score
	}
}

func compareScores(playerName string, score1, score2 int ) string {
	scoreDifference := score1 - score2
	if score1 >= 4  {
		if scoreDifference >= 2 {
			return  "Win for " + playerName
		}
		if scoreDifference >= 1 {
			return "Advantage " + playerName
		}
	}
	return ""
}

func (game *tennisGame2) shouldPrintOverallScore() bool {
	return game.player1Score() < 4 && game.player2Score() < 4
}

func (game *tennisGame2) printOverallScore() string {
	return game.getPlayer1Result() + "-" + game.getPlayer2Result()
}

func (game *tennisGame2) player2Score() int {
	return game.player2.score.value
}

func (game *tennisGame2) player1Score() int {
	return game.player1.score.value
}

func (game *tennisGame2) WonPoint(player string) {
	if player == "player1" {
		game.player1.score.increase()
	} else {
		game.player2.score.increase()
	}
}

func (game *tennisGame2) getPlayer1Result() string {
	return game.player1.score.print()
}

func (game *tennisGame2) getPlayer2Result() string {
	return game.player2.score.print()
}

func (game tennisGame2) printPlayer1Score() string {
	return game.player1.score.print()
}

func (game tennisGame2) printPlayer2Score() string {
	return game.player2.score.print()
}
