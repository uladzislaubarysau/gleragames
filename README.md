# Gleragames app


To run the application you need to run the command `make run_dc`.<br/>

To run app from source we should run container with databes `make run_db` and change host in `data source name` parameters in file `config/config-sample.json`.<br/>

`database -> 127.0.0.1`<br/>


After this we can use `make run` or just `go run ./cmd/app/main.go`<br/>