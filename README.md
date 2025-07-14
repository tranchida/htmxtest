# HTMX Test - Application Web avec Go

Application web moderne utilisant HTMX pour des interfaces dynamiques, Go avec le framework Echo pour le backend, et un générateur de messages aléatoires en français.

## Fonctionnalités

- 🚀 **Serveur web** avec Echo Framework
- 🎨 **Templates** avec le package `templ` 
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
├── main.go                   # Point d'entrée avec serveur Echo
├── rundev.sh                 # Script de développement
├── check.http                # Tests HTTP
├── static/                   # Fichiers statiques intégrés
│   ├── htmx.min.js          # Bibliothèque HTMX
│   └── pico.zinc.min.css    # Framework CSS Pico
├── internal/
│   ├── services/            # Services métier
│   │   └── messages.go      # Générateur de messages français
│   └── templates/           # Templates templ
│       ├── about.templ      # Page à propos
│       ├── admin.templ      # Page admin  
│       ├── home.templ       # Page d'accueil
│       ├── layout.templ     # Layout principal
│       ├── message.templ    # Template pour messages
│       ├── title.templ      # Composant titre
│       └── *_templ.go       # Fichiers générés par templ
├── bin/                     # Binaires compilés
└── tmp/                     # Fichiers temporaires
```

## Développement

### Script de développement

Un script `rundev.sh` est fourni pour faciliter le développement. Il s'occupe de :
- Générer automatiquement les templates
- Démarrer le serveur avec rechargement automatique
- Fournir un environnement de développement optimisé

Pour l'utiliser :

1. Rendez le script exécutable (si ce n'est pas déjà fait) :
   ```bash
   chmod +x rundev.sh
   ```

2. Lancez le script :
   ```bash
   ./rundev.sh
   ```

Le serveur de développement sera accessible à l'adresse : [http://localhost:8080](http://localhost:8080)

### Génération manuelle des templates

Si nécessaire, vous pouvez générer manuellement les templates avec :

```bash
go tool templ generate
```

### Architecture

- **Echo Framework** : Serveur HTTP rapide et léger
- **Templ** : Génération de templates HTML type-safe
- **HTMX** : Interactions dynamiques côté client
- **Embedded Static Files** : Fichiers statiques intégrés au binaire
- **Service Layer** : Séparation de la logique métier (messages)

## Licence

MIT
