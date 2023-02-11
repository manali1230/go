# Race Condition

[waitGroups & Mutex](https://pkg.go.dev/sync)

1. run the command to create go.mod file

```
go mod init racecondition
```

2. run go code

```
go run .
```

### output

```
Third GoRoutine
Second GoRoutine
First GoRoutine
[0 3 2 1]
```

3. run go code with Race Condition

```
go run --race .
```

### output

```
First GoRoutine
Second GoRoutine
Third GoRoutine
[0 1 2 3]
```