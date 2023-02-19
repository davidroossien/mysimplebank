# Backend Master Class [Golang + Postgres + Kubernetes + gRPC]

* https://okta.udemy.com/course/backend-master-class-golang-postgresql-kubernetes

## Local go development workspace setup
1. I use a mac with brew. 
1. My .bashrc contains:
    1. #for go
    1. export PATH="${PATH}:~/usr/local/opt"
    1. export GO111MODULE=on
    1. export PATH=${PATH}:`go env GOPATH`/bin
1. This project uses both a go workspace and go modules.

## Prerequisites
1. Run "brew install golang-migrate"
1. Run "brew install sqlc"
1. Install docker desktop
1. Install postgres:12-alpine & configure as shown in the videos
1. Lots more, see the videos

## Setup
1. Git clone this to your local machine.
1. Open VSCode and open the root folder.
1. Open a terminal.
1. Go to the project root folder.
1. Run "go work init"
1. Run "go work use ."
1. Run "go mod init"
1. Run "go mod tidy" as needed

## Configure
1. Create an app.env file in the root folder and add the following:
    1.  DB_DRIVER = ""
    1.  DB_SOURCE = ""
