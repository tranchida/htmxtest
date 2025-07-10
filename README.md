# Application Go avec HTMX

Ce projet est une application web simple utilisant Go comme backend et HTMX pour les interactions côté client.

## Prérequis

- Go 1.16 ou supérieur

## Installation

1. Cloner le dépôt :
   ```bash
   git clone [URL_DU_REPO]
   cd htmxtest
   ```

2. Lancer l'application :
   ```bash
   go run main.go
   ```

3. Ouvrir un navigateur à l'adresse : http://localhost:8080

## Structure du projet

- `main.go` : Point d'entrée de l'application
- `templates/` : Contient les templates HTML
  - `layout.html` : Template de base
  - `home.html` : Page d'accueil
  - `about.html` : Page À propos
  - `admin.html` : Page d'administration
- `static/` : Fichiers statiques (CSS, JS, images)

## Fonctionnalités

- Navigation fluide avec HTMX
- Chargement dynamique du contenu
- Interface utilisateur moderne avec Pico CSS

## Dépendances

- HTMX : Pour les interactions côté client
- Pico CSS : Pour le style minimaliste
