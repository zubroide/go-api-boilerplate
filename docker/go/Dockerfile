FROM golang:1.12.5

WORKDIR /go/src/go-api-boilerplate

RUN go get github.com/go-swagger/go-swagger/cmd/swagger \
    github.com/Altoros/gorm-goose/cmd/gorm-goose

COPY ./docker/go/entrypoint.sh ./docker/wait-for-it.sh /root/
RUN chmod 755 /root/entrypoint.sh /root/wait-for-it.sh

COPY go.mod .
COPY go.sum .
RUN go mod download

# Project files
COPY . .
COPY ./docker/go/dbconf.yml ./db/dbconf.yml
RUN swagger generate spec -o doc/swagger.json \
    && go install

ENTRYPOINT ["/root/entrypoint.sh"]
