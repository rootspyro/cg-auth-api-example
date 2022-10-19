# AUTH-API-EXAMPLE

## Run Project
```shell 
# validate the swaggger configuration
swagger validate swagger/swagger.yaml

# update the changes
swagger generate server -t gen -f swagger/swagger.yaml --default-scheme http --exclude-main

# If packages are required
go mod tidy

# run the server
go run cmd/main.go --port 3000
```
