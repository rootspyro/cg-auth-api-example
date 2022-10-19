# AUTH-API-EXAMPLE

## Run Project
```shell 
# update the changes
swagger generate -t gen -f swagger/swagger.yaml --default-scheme http --exclude-main

# run the server
go run cmd/main.go --port 3000
```
