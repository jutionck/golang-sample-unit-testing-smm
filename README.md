## Sample Golang Unit Test

### Run Test
#### First Step
```
go test -v
```

#### Second Step
```
go test ./... -v
```

#### Coverage Test
```
go test ./... -v -coverprofile=cover.out
go tool cover -html=cover.out
```
