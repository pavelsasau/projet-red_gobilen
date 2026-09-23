package main

type Monster struct {
	Name      string
	MaxHP     int
	CurrentHP int
	Attack    int
}

func initGoblin() Monster {
	return Monster{
		Name:      "Gobelin d'entrainement",
		MaxHP:     40,
		CurrentHP: 40,
		Attack:    5,
	}
}
