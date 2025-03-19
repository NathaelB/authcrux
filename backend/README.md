# AuthCrux Backend

Ce projet est un backend d'authentification structuré selon l'architecture hexagonale (ou architecture ports et adaptateurs) en Go.

## Structure du projet

L'architecture hexagonale divise l'application en trois couches principales :

```
backend/
├── domain/            # Le cœur de l'application (business logic)
│   ├── model/         # Entités métier
│   └── port/          # Interfaces (ports) pour communiquer avec le domaine
│
├── application/       # Implémentation des cas d'utilisation 
│   ├── handler/       # Handlers HTTP (adaptateurs d'entrée)
│   ├── service/       # Services qui implémentent la logique métier
│   └── http.go        # Configuration et gestion du serveur HTTP
│
├── infrastructure/    # Adaptateurs pour les technologies externes
│   ├── database/      # Configuration de la base de données
│   └── repository/    # Implémentations des repositories (adaptateurs de sortie)
│       └── gorm_model/ # Modèles spécifiques à GORM
```

## Technologies utilisées

- **Go** - Langage de programmation
- **Fiber** - Framework HTTP rapide et flexible
- **GORM** - ORM (Object Relational Mapper) pour Go
- **PostgreSQL** - Base de données relationnelle

## Avantages de l'architecture hexagonale

1. **Séparation des préoccupations** - Le code est organisé par responsabilité
2. **Testabilité** - Facile à tester grâce aux interfaces et à l'inversion de dépendance
3. **Flexibilité** - On peut facilement remplacer les adaptateurs (ex: changer de base de données)
4. **Maintenabilité** - Le code métier est protégé des changements d'infrastructure

## Organisation technique

- **Domain** (cœur) : Contient les modèles et interfaces (ports)
- **Application** (cas d'utilisation) : Contient les services métier, la gestion HTTP et les handlers 
- **Infrastructure** (adaptateurs) : Contient les implémentations techniques des repositories

Cette organisation place les handlers HTTP dans la couche application car ils sont considérés comme faisant partie des cas d'utilisation, tout en gardant la gestion de la base de données dans l'infrastructure.

## Démarrage rapide

### Prérequis

- Go 1.18 ou supérieur
- PostgreSQL

### Variables d'environnement

```
# Base de données
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=authcrux
DB_SSLMODE=disable

# Serveur
PORT=3000
ENV=development  # ou production
```

### Installation et exécution

1. Cloner le dépôt
2. Installer les dépendances: `go mod download`
3. Démarrer PostgreSQL
4. Lancer l'application: `go run main.go`

Pour le mode production:
```bash
ENV=production go run main.go
```

## API Endpoints

L'API propose les endpoints suivants:

### Authentification

- `POST /api/register` - Inscription d'un nouvel utilisateur
- `POST /api/login` - Connexion utilisateur

### Gestion des utilisateurs

- `GET /api/users` - Liste des utilisateurs (avec pagination)
- `GET /api/users/:id` - Détails d'un utilisateur
- `PUT /api/users/:id` - Mise à jour d'un utilisateur
- `DELETE /api/users/:id` - Suppression d'un utilisateur

## Évolutions possibles

- Ajout d'authentification JWT
- Validation plus complète des données
- Gestion des rôles et permissions
- Tests unitaires et d'intégration
- Documentation OpenAPI 