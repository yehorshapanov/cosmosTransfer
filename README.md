### How to run this:

First run docker image:<br />
```azure
docker run --name some-postgres -p 5432:5432 -e \
POSTGRES_PASSWORD=mysecretpassword -d postgres
```

Then run scripts to populate the Database:<br />
```azure
go build cosmosTransfer/cmd/setup_database && go run
```

Then run parser:
```azure
go build main.go
```