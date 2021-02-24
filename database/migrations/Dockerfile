FROM migrate/migrate

COPY .  /migrations

ENTRYPOINT [ "" ]
CMD ["sh", "-c", "migrate -path=/migrations -database ${DATABASE_CONNECTION_URL} up"]
