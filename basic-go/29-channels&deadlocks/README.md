# Channels & Deadlocks

[Channels & Deadlocks](https://go.dev/doc/effective_go#channels)

1. run the command to create go.mod file

```
go mod init channels
```

2. run go code

```
go run .
```

### output

```
>> go run main.go
1
true
```

3. run go code with Race Condition

```
go run --race .
```

### output

```
1
true
```