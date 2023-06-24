# Backend Master Class [Golang + Postgres + Kubernetes + gRPC]

* https://www.udemy.com/course/backend-master-class-golang-postgresql-kubernetes/

## Local go development workspace setup
1. I use a mac with brew.
1. My .bashrc contains:
    1. #for go
    1. export PATH="${PATH}:~/usr/local/opt"
    1. export GO111MODULE=on
    1. export PATH=${PATH}:`go env GOPATH`/bin
1. This project uses both a go workspace and go modules.
1. If on an older version that 1.19, then
    1. run "brew upgrade go"
    1. run "brew reinstall go"
    1. restart VSCode
    1. go mod tidy

## Prerequisites
1. Run "brew install golang-migrate"
1. Run "brew install sqlc"
1. Install docker desktop
1. Lots more, see the videos

## Setup
1. Git clone this to your local machine.
1. Open VSCode and open the root folder.
1. Open a terminal.
1. Go to the project root folder.
1. Run "go mod tidy"
1. Run "make postgres"
1. Run "make createdb"
1. Run "make migrateup"
1. Run "make test"

## Configure (for reference)
1. Create an app.env file in the root folder and add the following:
    1.  DB_DRIVER = "postgres"
    1.  DB_SOURCE = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"

## Go commands (for reference)
1. Run "go work init"
1. Run "go work use ."
1. Run "go mod init"
1. Run "go mod tidy" as needed