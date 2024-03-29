FROM golang as builder

LABEL author="Eduardo Akira"

WORKDIR /app

COPY . .

## RUN go mod download

## RUN go get .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

CMD [ "./app" ]

EXPOSE 8080