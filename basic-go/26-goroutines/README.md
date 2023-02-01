# GOROUTINES

1. Initialize  go - it will create a file names as `go.mod`

```
go mod init goroutines
```

2. Run `main.go`

```
>> go run main.go
Hello
Hello
Hello
Hello
Hello
Hello
World
World
World
World
World
World
```

3. Run `parallel.go`

```
>> go run parallel.go 
World
Hello
Hello
World
World
Hello
Hello
World
Hello
World
World
Hello
```