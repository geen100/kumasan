FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ /app/

RUN go build -o main .
RUN chmod +x main

EXPOSE 8080

CMD [ "/app/main" ]
