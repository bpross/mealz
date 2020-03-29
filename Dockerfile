FROM golang

ENV GO111MODULE=on

WORKDIR /mealz

COPY go.mod .
COPY go.sum .
COPY . .

RUN go get -u golang.org/x/lint/golint
RUN go get github.com/onsi/ginkgo/ginkgo
RUN go get github.com/golang/mock/mockgen
RUN go get -u github.com/pressly/goose/cmd/goose
