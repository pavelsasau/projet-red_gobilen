# projet-red_gobilen

Mini-RPG développé en Go dans le cadre de notre formation à Paris Ynov Campus.

Le jeu fonctionne entièrement dans le terminal.

## Présentation

Projet RED est un jeu de rôle en console dans lequel le joueur crée un personnage, gère son inventaire, achète et fabrique des objets, apprend des sorts et combat des monstres.

Le joueur peut choisir entre plusieurs classes et améliorer son personnage grâce aux équipements et aux objets disponibles dans le jeu.

## Fonctionnalités

### Personnage

- Création du personnage
- Choix du nom
- Choix de la classe :
  - Humain : 100 PV
  - Elfe : 80 PV
  - Nain : 120 PV
- Niveau du personnage
- Gestion des PV
- Gestion de l'or
- Affichage des équipements

### Inventaire

- Inventaire de base de 10 emplacements
- Ajout et suppression d'objets
- Utilisation des objets
- Potions de vie
- Potions de poison
- Livres de sorts
- Équipements
- Possibilité d'agrandir l'inventaire
- Maximum de 3 améliorations d'inventaire

### Marchand

Le joueur peut acheter différents objets avec son or :

- Potion de vie
- Potion de poison
- Livre de Sort : Boule de Feu
- Fourrure de Loup
- Peau de Troll
- Cuir de Sanglier
- Plume de Corbeau
- Amélioration de l'inventaire

### Forgeron

Le joueur peut utiliser les matériaux achetés pour fabriquer des équipements :

- Chapeau de l'aventurier
- Tunique de l'aventurier
- Bottes de l'aventurier

Chaque fabrication demande des matériaux ainsi que de l'or.

### Équipements

Les équipements permettent d'augmenter les PV maximum du personnage :

- Chapeau de l'aventurier : +10 PV
- Tunique de l'aventurier : +25 PV
- Bottes de l'aventurier : +15 PV

### Sorts

Le personnage commence avec :

- Coup de poing

Il peut ensuite apprendre :

- Boule de Feu

Pendant un combat, le joueur peut choisir et utiliser ses compétences.

### Combat

Le jeu possède un système de combat au tour par tour.

Pendant son tour, le joueur peut :

- effectuer une attaque basique
- utiliser un sort
- ouvrir son inventaire
- utiliser une potion

Le Gobelin d'entraînement possède :

- 40 PV
- 5 points d'attaque

Le gobelin effectue une attaque normale à chaque tour et une attaque plus puissante tous les 3 tours.

### Potion de poison

La potion de poison peut être utilisée pendant un combat.

Elle inflige :

- 10 points de dégâts
- pendant 3 tours

## Menu principal

Le menu principal permet d'accéder à :

1. Informations du personnage
2. Inventaire
3. Marchand
4. Forgeron
5. Entraînement
0. Quitter

## Technologies utilisées

- Go
- Visual Studio Code
- Git
- GitHub

## Intelligence artificielle

Nous avons utilisé ChatGPT et Gemini comme outils d'aide pendant le développement.

Ils nous ont aidés à :

- mieux comprendre certaines parties du code
- expliquer certaines fonctions
- rechercher des erreurs
- corriger certains problèmes
- réfléchir à l'organisation du projet

Le code a ensuite été intégré et testé dans notre projet.

## Installation

Cloner le dépôt :

```bash
git clone https://github.com/pavelsasau/projet-red_gobilen.git