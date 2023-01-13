## Pointers

This is `pointers` go code and I am using Visual Studio WSL[Ubuntu] for performing practical. I have installed `go` Extension in visual studio.

1. this command will generate a `go.mod` file.

```
go mod init pointers
```

while writing code if VS Code is showing install tools related to go then click on install.

2. Run `main.go` file by using.

```
go run main.go
```

## Output
The output of the above program is - 

```
>> go run main.go 

Pointers
value of the pointer is <nil>
```

`NOTE - The default value of a pointer is nil.`

3. Run `pointer.go` file by using.

```
go run pointer.go
```

## Output
The output of the above program is - 

```
>> go run pointer.go 

Pointers Usage
pointer value is 0xc000018068
value of pointer is 3
```

4. Run `multiplication-with-pointer.go ` file by using.

```
go run multiplication-with-pointer.go 
```

## Output
The output of the above program is - 

```
>> go run multiplication-with-pointer.go  

Pointers Usage
pointer value is 0xc0000b6010
value of pointer is 3
value of num is 30
```