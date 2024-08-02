package main

import (
	"fmt"
	"slices"
)

type Team struct {
	Name    string
	Players []string
}

type League struct {
	Teams []Team
	Wins  map[string]int
}

func (l *League) MatchResult(team1 string, score1 int, team2 string, score2 int) {
	if score1 == score2 {
		panic("ties are not allowed")
	}
	if score1 > score2 {
		l.Wins[team1]++
	} else {
		l.Wins[team2]++
	}
}

func (l League) Ranking() []string {
	teamNames := make([]string, 0, len(l.Teams))
	for _, team := range l.Teams {
		teamNames = append(teamNames, team.Name)
	}
	slices.SortFunc(teamNames, func(a, b string) int { return l.Wins[b] - l.Wins[a] })
	return teamNames
}

func main() {
	teams := []Team{
		{"Team3", []string{"Player5", "Player6"}},
		{"Team2", []string{"Player3", "Player4"}},
		{"Team1", []string{"Player1", "Player2"}},
	}
	league := League{
		Teams: teams,
		Wins:  map[string]int{},
	}
	league.MatchResult("Team1", 3, "Team2", 2)
	league.MatchResult("Team2", 2, "Team3", 1)
	league.MatchResult("Team3", 2, "Team1", 4)
	ranking := league.Ranking()
	fmt.Println(ranking)
}
