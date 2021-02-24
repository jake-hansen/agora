#!/bin/sh

migrate -path=/migrations -database ${DATABASE_TYPE}://${USER}:${PASSWORD}@${PROTOCOL}\(${HOST}:${PORT}\)/agora up
