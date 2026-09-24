package main

// Fonction principale : point de départ du programme.
// Главная функция: точка запуска программы.
func main() {

	// Crée le personnage et stocke le résultat dans la variable player.
	// Создаёт персонажа и сохраняет результат в переменную player.
	player := characterCreation()

	// Lance le menu principal en passant l'adresse du personnage.
	// Запускает главное меню, передавая адрес персонажа.
	menu(&player)
}
