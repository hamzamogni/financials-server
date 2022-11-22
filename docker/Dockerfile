FROM golang:1.19.3 as development

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# reflex package used to run command when files change
RUN go install github.com/cespare/reflex@latest

COPY . .

EXPOSE 8080

CMD reflex -r "\.go$" go run app/main.go --start-service
