package main

import "fmt"

// Équipe un objet sur le personnage.
// Надевает предмет на персонажа.
func equipItem(c *Character, item string) {

	// Choisit l'action selon le nom de l'objet.
	// Выбирает действие в зависимости от названия предмета.
	switch item {

	case "Chapeau de l'aventurier":
		// 1. Si un chapeau est déjà équipé, on le remet dans l'inventaire
		// 1. Если шляпа уже надета, возвращаем её обратно в инвентарь.
		if c.Equip.Head != "" {

			// Remet l'ancien chapeau dans l'inventaire.
			// Возвращает старую шляпу в инвентарь.
			c.Inventory = append(c.Inventory, c.Equip.Head)

			c.MaxHP -= 10 // On retire le bonus du chapeau précédent
			// Убираем бонус старой шляпы к максимальному здоровью.
		}

		// 2. On équipe le nouveau chapeau
		// 2. Надеваем новую шляпу.
		c.Equip.Head = item

		// Ajoute 10 PV maximum.
		// Добавляет 10 к максимальному здоровью.
		c.MaxHP += 10

		fmt.Println("\nVous équipez : Chapeau de l'aventurier (+10 PV Max) !")

	case "Tunique de l'aventurier":

		// Vérifie si une tunique est déjà équipée.
		// Проверяет, надета ли уже туника.
		if c.Equip.Body != "" {

			// Remet l'ancienne tunique dans l'inventaire.
			// Возвращает старую тунику в инвентарь.
			c.Inventory = append(c.Inventory, c.Equip.Body)

			c.MaxHP -= 25 // On retire le bonus de la tunique précédente
			// Убираем бонус старой туники к максимальному здоровью.
		}

		// Équipe la nouvelle tunique.
		// Надевает новую тунику.
		c.Equip.Body = item

		// Ajoute 25 PV maximum.
		// Добавляет 25 к максимальному здоровью.
		c.MaxHP += 25

		fmt.Println("\nVous équipez : Tunique de l'aventurier (+25 PV Max) !")

	case "Bottes de l'aventurier":

		// Vérifie si des bottes sont déjà équipées.
		// Проверяет, надеты ли уже ботинки.
		if c.Equip.Feet != "" {

			// Remet les anciennes bottes dans l'inventaire.
			// Возвращает старые ботинки в инвентарь.
			c.Inventory = append(c.Inventory, c.Equip.Feet)

			c.MaxHP -= 15 // On retire le bonus des bottes précédentes
			// Убираем бонус старых ботинок к максимальному здоровью.
		}

		// Équipe les nouvelles bottes.
		// Надевает новые ботинки.
		c.Equip.Feet = item

		// Ajoute 15 PV maximum.
		// Добавляет 15 к максимальному здоровью.
		c.MaxHP += 15

		fmt.Println("\nVous équipez : Bottes de l'aventurier (+15 PV Max) !")

	default:

		// Affiche un message si l'objet ne peut pas être équipé.
		// Показывает сообщение, если предмет нельзя надеть.
		fmt.Println("\nCet objet ne peut pas être équipé.")

		// Arrête la fonction.
		// Останавливает функцию.
		return
	}

	// On retire l'équipement de l'inventaire une fois équipé
	// После того как предмет надет, удаляем его из инвентаря.
	removeInventory(c, item)
}
