# Starte von einem Golang base image
FROM golang:alpine

# Installiere notwendige Pakete, einschließlich Chromium
RUN apk update && apk add --no-cache \
    chromium \
    chromium-chromedriver \
    nss \
    freetype \
    freetype-dev \
    harfbuzz \
    ca-certificates \
    ttf-freefont \
    fontconfig

# Setze Umgebungsvariablen für Chromium
ENV CHROME_BIN=/usr/bin/chromium-browser \
    CHROME_PATH=/usr/lib/chromium/

# Setze das aktuelle Arbeitsverzeichnis in das Container image
WORKDIR /app

# Kopiere go mod und sum files
COPY go.mod go.sum ./

# Lade die Go dependencies
RUN go mod download

# Kopiere den Quellcode in das Container image
COPY . .

# Baue die Go Anwendung
RUN go build -o main .

# Exponiere Port 8080
EXPOSE 8080

# Führe den kompilierten binären Code aus
CMD ["./main"]
