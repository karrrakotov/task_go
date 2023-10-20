FROM golang:latest

# init workdir
WORKDIR /backend_task

# copy go app
COPY . .

# build go app
RUN go mod download
RUN go build cmd/main/main.go

# run app
CMD ["./main"]