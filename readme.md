# golang console application sample using erlang calculations, b and c.

# highway tollbooth traffic sim.

> Just some example code.

## Build Setup

```bash
# list the dependencies to see what will get installed

go list -f '{{ join .Imports "\n" }}'
go list -f '{{ join .Deps "\n" }}'

# get all the dependencies
go mod init example.com
go mod tidy

# run the sim
go run example.xyz

# run all tests
go test ./... -v

# run tests with a check for race conditions
go test -race

# run benchmark examples, if any.
go test -bench=.

# build etc
go build
go install
```
