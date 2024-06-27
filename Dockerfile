FROM golang:1.23rc1-alpine3.20

WORKDIR /opt/app

COPY . .

CMD ["go", "run", "main.go"]