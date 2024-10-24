package game

import "math/rand"

type Game struct {
	Players []Player
	Round   int
	Rounds  []int
	Dealer  *Player
}

func (g *Game) InitGame(players []string) {
	g.Players = []Player{}
	g.Round = 0
	g.Rounds = []int{}
	g.Dealer = nil
	for i := range players {
		g.addPlayer(players[i])
	}
	g.Dealer = &g.Players[rand.Intn(len(g.Players))]
	g.generateRounds()
}

func (g *Game) GetNoCards() int {
	return g.Rounds[g.Round]
}

func (g *Game) NextRound() bool {
	if g.Round == len(g.Rounds)-1 {
		return false
	}
	for i := range g.Players {
		g.Players[i].calculateScore()
	}
	g.newDealer()
	g.Round++

	return true
}

func (g *Game) GetWinner() *Player {
	var winner *Player
	maxScore := -1
	for i := range g.Players {
		currentScore := g.Players[i].CurrentScore()
		if currentScore > maxScore {
			maxScore = currentScore
			winner = &g.Players[i]
		}
	}

	return winner
}

func (g *Game) GetSumOfTriks() int {
	sum := 0
	for _, p := range g.Players {
		sum += p.Trick
	}

	return sum
}

func (g *Game) GetSumOfBids() int {
	sum := 0
	for _, p := range g.Players {
		sum += p.Bid
	}

	return sum
}

func (g *Game) addPlayer(n string) {
	p := Player{
		Name:  n,
		Score: []int{0},
	}
	g.Players = append(g.Players, p)
}

func (g *Game) generateRounds() {
	n := len(g.Players)
	max := 36 / n
	for i := 1; i <= max; i++ {
		g.Rounds = append(g.Rounds, i)
	}
	for i := 0; i < n; i++ {
		g.Rounds = append(g.Rounds, max)
	}
	for i := max - 1; i >= 1; i-- {
		g.Rounds = append(g.Rounds, i)
	}
	for i := 0; i < n; i++ {
		g.Rounds = append(g.Rounds, max)
	}
}

func (g *Game) newDealer() {
	for i, player := range g.Players {
		if player.Name == g.Dealer.Name {
			g.Dealer = &g.Players[(i+1)%len(g.Players)]
			break
		}
	}
}
