FROM golang:1.18

# copy project files
COPY . /go/src/counter/

# set working directory
WORKDIR /go/src/counter/

# build go app
RUN go build -o ./cmd/main ./cmd/main.go

EXPOSE 8080

# run app
ENTRYPOINT [ "./cmd/main" ]
