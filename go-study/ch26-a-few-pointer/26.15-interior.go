package main

import "fmt"

type stats struct {
	level             int
	endurance, health int
}

type character struct {
	name  string
	stats stats
}

func levelUp(s *stats) {
	s.level++
	s.endurance = 42 + (14 * s.level)
	s.health = 5 * s.endurance
}

func main() {
	player := character{name: "Mattias"}
	levelUp(&player.stats)
	fmt.Printf("%+v\n", player.stats) // {level:1 endurance:56 health:280}
}
