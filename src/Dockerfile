FROM golang as builder

WORKDIR /app

COPY src .

RUN go mod download

RUN go get .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# deployment image
FROM golang:latest  

LABEL author="Eduardo Akira"

WORKDIR /root/
COPY --from=builder /app .

CMD [ "./app" ]

EXPOSE 8080