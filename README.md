# HTMX Test - Application Web avec Go

Application web moderne utilisant HTMX pour des interfaces dynamiques, Go (net/http, html/template) pour le backend, et un générateur de messages aléatoires en français.

## Fonctionnalités

- 🚀 **Serveur web** natif Go (net/http)
- 🎨 **Templates** avec le package `html/template`
- ⚡ **HTMX** pour des mises à jour partielles sans JavaScript
- 🎲 **Générateur de messages** français aléatoires (119 messages disponibles)
- 🏗️ **Architecture modulaire** avec séparation claire des composants
- 🎯 **Pages disponibles** :
  - Page d'accueil (`/`) avec générateur de messages
  - Page À propos (`/about`)
  - Page d'administration (`/admin`)
- 📁 **Fichiers statiques** intégrés (CSS Pico, HTMX)

## Prérequis

- Go 1.24 ou supérieur
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

## API Endpoints

- `GET /` - Page d'accueil
- `GET /about` - Page à propos
- `GET /admin` - Page d'administration
- `GET /randommessage` - API pour récupérer un message français aléatoire (format HTMX)

## Structure du projet

```
.
├── go.mod                    # Dépendances Go
├── go.sum                    # Checksums des dépendances
├── main.go                   # Point d'entrée avec serveur HTTP natif
├── check.http                # Tests HTTP
├── static/                   # Fichiers statiques intégrés
│   ├── htmx.min.js          # Bibliothèque HTMX
│   └── pico.zinc.min.css    # Framework CSS Pico
├── internal/
│   └── services/            # Services métier
│       └── messages.go      # Générateur de messages français
├── templates/               # Templates HTML
│   ├── about.go.tmpl        # Page à propos
│   ├── admin.go.tmpl        # Page admin  
│   ├── home.go.tmpl         # Page d'accueil
│   ├── layout.go.tmpl       # Layout principal
│   ├── message.go.tmpl      # Template pour messages
│   ├── title.go.tmpl        # Composant titre
├── bin/                     # Binaires compilés
└── tmp/                     # Fichiers temporaires
```





### Architecture

- **Serveur HTTP natif** : net/http pour la gestion des routes et des requêtes
- **html/template** : Génération de templates HTML sécurisés
- **HTMX** : Interactions dynamiques côté client
- **Embedded Static Files** : Fichiers statiques intégrés au binaire
- **Service Layer** : Séparation de la logique métier (messages)
## Logging

Le serveur journalise chaque requête HTTP au format Apache combined log :

```
IP - - [date] "METHOD PATH PROTO" status bytes "Referer" "User-Agent" duration
```

Exemple :

```
127.0.0.1 - - [22/Jul/2025:14:32:10 +0200] "GET /about HTTP/1.1" 200 512 "http://localhost:8080/" "Mozilla/5.0 ..." 2.3ms
```

## Licence

MIT
