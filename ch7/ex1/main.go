package main

type Team struct {
	Name    string
	Players []string
}

type League struct {
	Teams []Team
	Wins  map[string]int
}

func main() {
	teams := []Team{
		{"Team1", []string{"Player1", "Player2"}},
		{"Team2", []string{"Player3", "Player4"}},
	}
	wins := map[string]int{
		"Team1": 3,
		"Team2": 4,
	}
	league := League{teams, wins}
	println(league.Teams[0].Name)
	println(league.Wins["Team1"])
}
