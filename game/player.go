package game

type Player struct {
	Name  string
	Score []int
	Bid   int
	Trick int
}

func (p *Player) CurrentScore() int {
	return p.Score[len(p.Score)-1]
}

func (p *Player) calculateScore() {
	currentScore := p.CurrentScore()
	switch {
	case p.Bid == p.Trick:
		if p.Bid == 0 {
			currentScore += 5
		} else {
			currentScore += 10 * p.Trick
		}
	case p.Bid < p.Trick:
		currentScore += p.Trick - p.Bid
	default:
		currentScore -= 10 * (p.Bid - p.Trick)
	}
	p.Score = append(p.Score, currentScore)
	p.Bid = 0
	p.Trick = 0
}
