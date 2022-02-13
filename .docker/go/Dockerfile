FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o go-language-course ./main.go

CMD ["./go-language-course"]