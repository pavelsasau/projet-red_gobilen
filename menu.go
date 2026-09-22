package main

import "fmt"

// Affiche les options du menu principal.
// Показывает пункты главного меню.
func menu(player *Character) {
	for {
		fmt.Printf("\nBienvenue %s le %s !\n", player.Name, player.Class)
		fmt.Println("=== MENU PRINCIPAL ===")
		fmt.Println("1. Informations du personnage")
		fmt.Println("2. Inventaire")
		fmt.Println("3. Marchand")
		fmt.Println("4. Forgeron")
		fmt.Println("0. Quitter")
		fmt.Print("Votre choix : ")

		// Lit le choix de l'utilisateur et l'enregistre dans choice.
		// Считывает выбор пользователя и сохраняет его в choice.
		var choice int
		fmt.Scan(&choice)

		// Exécute une action selon le choix de l'utilisateur.
		// Выполняет действие в зависимости от выбора пользователя.
		switch choice {
		case 1:
			fmt.Println()
			displayInfo(*player)

		case 2:
			fmt.Println()
			accessInventory(*player)

		case 3:
			fmt.Println()
			merchantMenu(player)

		case 4:
			fmt.Println()
			blacksmithMenu(player)
		case 0:
			return

		default:
			fmt.Println("Choix invalide")
		}
	}
}
