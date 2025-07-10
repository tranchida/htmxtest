# Application Web HTMX avec Go

Une application web moderne utilisant HTMX pour des interfaces dynamiques et Go (avec le framework Echo) pour le backend.

## Fonctionnalités

- 🚀 **Routing** avec Echo
- 🎨 **Templating** avec le package `templ`
- ⚡ **HTMX** pour des mises à jour partielles sans JavaScript
- 🏗️ **Architecture modulaire** avec séparation claire des templates
- 🎯 **Pages disponibles** :
  - Page d'accueil (`/`)
  - Page À propos (`/about`)
  - Page d'administration (`/admin`)

## Prérequis

- Go 1.20 ou supérieur
- Un navigateur web moderne

## Installation

1. Cloner le dépôt :
   ```bash
   git clone [URL_DU_REPO]
   cd htmxtest
   ```

2. Installer les dépendances :
   ```bash
   go mod download
   ```

## Lancement

```bash
go run main.go
```

L'application sera disponible à l'adresse : [http://localhost:8080](http://localhost:8080)

## Structure du projet

```
.
├── go.mod           # Dépendances Go
├── main.go          # Point d'entrée de l'application
├── static/          # Fichiers statiques (CSS, JS, images)
└── internal/
    └── templates/   # Templates HTML avec templ
        ├── home.templ
        ├── about.templ
        ├── admin.templ
        ├── layout.templ
        └── title.templ
```

## Développement

### Rechargement automatique

Pour un développement plus fluide, vous pouvez utiliser [air](https://github.com/cosmtrek/air) pour le rechargement automatique :

```bash
go install github.com/cosmtrek/air@latest
air
```

### Génération des templates

Après avoir modifié un fichier `.templ`, exécutez :

```bash
templ generate
```

## Licence

MIT
