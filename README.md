# AUTH-API-EXAMPLE

## Env Variables

```shell
# Example .env file

# server conf 
PORT=3000

# postgres conf
PG_HOST=host
PG_PORT=5432
PG_USER=postgres
PG_PASSW=secure_password_here
PG_DB=dbname
```

## Run Project
```shell 
# validate the swaggger configuration
swagger validate swagger/swagger.yaml

# update the changes
swagger generate server -t gen -f swagger/swagger.yaml --default-scheme http --exclude-main

# If packages are required
go mod tidy

# run the server
go run cmd/main.go
```
