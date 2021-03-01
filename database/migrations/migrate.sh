#!/bin/sh

urlencode () {
  string=$1
  while [ -n "$string" ]; do
    tail=${string#?}
    head=${string%$tail}
    case $head in
      [-._~0-9A-Za-z]) printf %c "$head";;
      *) printf %%%02x "'$head"
    esac
    string=$tail
  done
  echo
}

ENCODED_PASSWORD=$(urlencode "${PASSWORD}")

migrate -path=/migrations -database ${DATABASE_TYPE}://${USER}:${ENCODED_PASSWORD}@${PROTOCOL}\(${HOST}:${PORT}\)/agora up
