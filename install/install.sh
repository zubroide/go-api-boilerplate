#!/bin/bash

sudo apt install -y apt-transport-https gnupg curl debian-keyring debian-archive-keyring
sudo curl -1sLf 'https://dl.cloudsmith.io/public/go-swagger/go-swagger/gpg.2F8CB673971B5C9E.key' | sudo gpg --dearmor -o /usr/share/keyrings/go-swagger-go-swagger-archive-keyring.gpg
sudo curl -1sLf 'https://dl.cloudsmith.io/public/go-swagger/go-swagger/config.deb.txt?distro=debian&codename=any-version' | sudo tee /etc/apt/sources.list.d/go-swagger-go-swagger.list

go get github.com/steinbacher/goose/cmd/goose
go install -mod=mod github.com/steinbacher/goose/cmd/goose

curl -sSfL https://raw.githubusercontent.com/air-verse/air/master/install.sh | sh -s
