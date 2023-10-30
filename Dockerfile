# Utilisez l'image de base Golang version 1.19
FROM golang:1.19

# Crée un répertoire appelé /app dans le conteneur
RUN mkdir /app

# Copie le contenu du répertoire actuel dans /app du conteneur
ADD . /app

# Définit /app comme répertoire de travail, où les commandes ultérieures seront exécutées
WORKDIR /app

# Indique que le conteneur écoutera sur le port 8080
EXPOSE 8080

#La commande correcte pour installer Bash est "apt-get install -y bash"
RUN apt-get update && apt-get install -y bash

# Compile l'application Go dans le répertoire actuel et la nomme "main"
RUN go build -o main .

# Définit la commande qui sera exécutée lorsque le conteneur sera lancé
CMD  ["/app/main"]
