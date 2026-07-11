package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Team struct {
	Name  string
	Power int
}

func main() {
	rand.Seed(time.Now().UnixNano())

	teams := []Team{
		{"T1", 98},
		{"Gen.G", 97},
		{"Bilibili Gaming", 95},
		{"Anyone's Legend", 93},
		{"G2 Esports", 91},
		{"Top Esports", 90},
		{"Hanwha Life Esports", 94},
		{"CTBC Flying Oyster", 88},
	}

	var champion Team
	best := -1

	for _, team := range teams {
		score := team.Power + rand.Intn(10)
		fmt.Printf("%-20s -> %d điểm\n", team.Name, score)

		if score > best {
			best = score
			champion = team
		}
	}

	fmt.Println("\n============================")
	fmt.Printf("🏆 Dự đoán vô địch CKTG 2026: %s\n", champion.Name)
	fmt.Println("============================")
}