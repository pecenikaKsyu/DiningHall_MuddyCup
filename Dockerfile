FROM golang:alpine

WORKDIR /app

#COPY go.mod ./
COPY . .
RUN go mod download

COPY *.go ./

RUN go build -o dininghall .

CMD "/app/dininghall"

EXPOSE 8001