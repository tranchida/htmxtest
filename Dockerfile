# Utilise une image officielle de Go pour la compilation
FROM golang:1.24.5 AS builder

WORKDIR /app

# Copie les fichiers go mod et sum, puis télécharge les dépendances
COPY go.mod go.sum ./
RUN go mod download

# Copie le reste du code source
COPY . .

# Compile l'application (remplacez 'htmxtest' par le nom de votre binaire si besoin)
RUN CGO_ENABLED=0 GOOS=linux go build -o /htmxtest main.go

# Utilise une image minimale pour l'exécution
FROM alpine:latest
WORKDIR /root/

# Copie le binaire compilé et les fichiers nécessaires
COPY --from=builder /htmxtest ./htmxtest
COPY static ./static
COPY internal/templates ./internal/templates

# Expose le port utilisé par l'application (à adapter si besoin)
EXPOSE 8080

# Commande de lancement
CMD ["./htmxtest"]
