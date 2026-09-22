package main

// Vérifie si le personnage connaît déjà "Boule de Feu".
// Проверяет, знает ли персонаж уже "Boule de Feu".
func spellBook(player *Character) {
	for _, skill := range player.Skill {
		if skill == "Boule de Feu" {
			return
		}
	}

	// Ajoute "Boule de Feu" à la liste des compétences du personnage.
	// Добавляет "Boule de Feu" в список умений персонажа.
	player.Skill = append(player.Skill, "Boule de Feu")
}
