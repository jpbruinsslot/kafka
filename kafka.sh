#!/bin/bash

execute() {
    local cmd=$1
    docker exec -it kafka bin/$cmd.sh "$@"
}

function usage() {
    echo "Usage: $0 <command> [args]"
}

function main() {
    local cmd=$1

    if [[ -z "$cmd" ]]; then
        usage
        exit 1
    fi

    execute "$@"
}

main "$@"
