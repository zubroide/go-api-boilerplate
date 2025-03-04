FROM golang:1.23

WORKDIR /go/src/go-api-boilerplate

RUN go mod init  \
    && go get github.com/steinbacher/goose/cmd/goose \
    && go install -mod=mod github.com/steinbacher/goose/cmd/goose

RUN apt update \
    && apt install -y apt-transport-https gnupg curl debian-keyring debian-archive-keyring \
    && curl -1sLf 'https://dl.cloudsmith.io/public/go-swagger/go-swagger/gpg.2F8CB673971B5C9E.key' | gpg --dearmor -o /usr/share/keyrings/go-swagger-go-swagger-archive-keyring.gpg \
    && curl -1sLf 'https://dl.cloudsmith.io/public/go-swagger/go-swagger/config.deb.txt?distro=debian&codename=any-version' | tee /etc/apt/sources.list.d/go-swagger-go-swagger.list \
    && apt update \
    && apt install -y swagger

COPY ./docker/go/entrypoint.sh ./docker/wait-for-it.sh /root/
RUN chmod 755 /root/entrypoint.sh /root/wait-for-it.sh

COPY go.mod go.sum ./
RUN go mod vendor \
    && go mod download

# Project files
COPY . .
COPY ./docker/go/dbconf.yml ./db/dbconf.yml
RUN touch .env

RUN swagger generate spec -o doc/swagger.json
RUN go mod vendor \
    && go install

ENTRYPOINT ["/root/entrypoint.sh"]
