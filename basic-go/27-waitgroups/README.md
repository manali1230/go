# WaitGroups

[waitGroups & Mutex](https://pkg.go.dev/sync)

1. run the command to create go.mod file

```
go mod init waitgroups
```

2. run go code

```
go run main.go
```

### output

```
https://google.com Endpoint !
https://www.facebook.com Endpoint !
https://www.youtube.com Endpoint !
https://go.dev Endpoint !
200 is the status code for https://www.youtube.com
200 is the status code for https://www.facebook.com
200 is the status code for https://go.dev
200 is the status code for https://google.com
[test signals https://www.youtube.com https://www.facebook.com https://go.dev https://google.com]
```