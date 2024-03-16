FROM golang:latest

WORKDIR /app

COPY . .

RUN go get .

EXPOSE 8080
EXPOSE 5432

CMD ["go","run","main.go"]