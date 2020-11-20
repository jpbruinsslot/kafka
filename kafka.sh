#!/bin/bash

main() {
    local cmd=$1
    docker-compose exec kafka bin/$cmd.sh "$@"
}

main "$@"
