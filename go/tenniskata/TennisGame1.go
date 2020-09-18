package tenniskata

type playerGame1 struct {
	name  string
	score int
}

type tennisGame1 struct {
	player1 playerGame1
	player2 playerGame1
}

func TennisGame1(player1Name string, player2Name string) TennisGame {
	return &tennisGame1{
		player1: playerGame1{name: player1Name},
		player2: playerGame1{name: player2Name},
	}
}

func (game *tennisGame1) WonPoint(playerName string) {
	if playerName == "player1" {
		game.player1.score += 1
	} else {
		game.player2.score += 1
	}
}

func (game *tennisGame1) getPlayer1Score() int {
	return game.player1.score
}

func (game *tennisGame1) getPlayer2Score() int {
	return game.player2.score
}

func (game *tennisGame1) GetScore() string {
	if game.isTie() {
		if game.getPlayer1Score() > 2 {
			return "Deuce"
		} else {
			return game.getScoreName(game.getPlayer1Score()) + "-All"
		}
	} else if game.canWin() {
		minusResult := game.getPlayer1Score() - game.getPlayer2Score()
		return game.getDifferenceScore(minusResult)
	} else {
		return game.getScoreName(game.getPlayer1Score()) + "-" + game.getScoreName(game.getPlayer2Score())
	}
}

func (game *tennisGame1) getDifferenceScore(minusResult int) string {
	if minusResult == 1 {
		return "Advantage player1"
	} else if minusResult == -1 {
		return "Advantage player2"
	} else if minusResult >= 2 {
		return "Win for player1"
	} else {
		return "Win for player2"
	}
}

func (game *tennisGame1) canWin() bool {
	return game.getPlayer1Score() >= 4 || game.getPlayer2Score() >= 4
}

func (game *tennisGame1) isTie() bool {
	return game.getPlayer1Score() == game.getPlayer2Score()
}

func (game *tennisGame1) getScoreName(score int) string {
	switch score {
	case 0:
		return "Love"
	case 1:
		return "Fifteen"
	case 2:
		return "Thirty"
	case 3:
		return "Forty"
	default:
		return ""
	}
}
