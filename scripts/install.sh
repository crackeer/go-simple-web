#!/bin/bash

function build() {
    go version
    if [ $? -eq 0 ]; then
        echo "Go is installed"
    else
        echo "Go is not installed"
        exit 1
    fi

    go build -o go-simple-web .
    if [ $? -eq 0 ]; then
        echo "Build success"
    else
        echo "Build failed"
        exit 1
    fi
}

function install() {
    if [ ! -d "/usr/local/go-simple-web" ]; then
        sudo mkdir /usr/local/go-simple-web
    fi

    sudo cp go-simple-web /usr/local/go-simple-web/
    if [ $? -eq 0 ]; then
        echo "Copy success"
    else
        echo "Copy failed"
        exit 1
    fi

    if [ ! -d "/usr/local/go-simple-web/frontend" ]; then
        cp -r ./frontend /usr/local/go-simple-web/
    fi

    if [ ! -f "/usr/local/go-simple-web/config" ]; then
        cp  -r ./config/linux/* /usr/local/go-simple-web/
    fi
    if [ ! -f "/usr/lib/systemd/system/go-simple-web.service" ]; then
        cp go-simple-web.service /usr/lib/systemd/system/
    fi
    systemctl enable realsee_server.service
    systemctl restart realsee_server.service
}

build
install
