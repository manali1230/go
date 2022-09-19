## SLICES

[Slices](https://go.dev/tour/moretypes/6)

This is `Slices` go code and I am using Visual Studio WSL[Ubuntu] for performing practical. I have installed `go` Extension in visual studio.

1. this command will generate a `go.mod` file.

```
go mod tidy slices
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

Slices
The datatype of langList is []string
```

3. Run `append.go` file by using.

```
go run append.go
```

## Output
The output of the above program is - 

```
>> go run append.go 

Append Slices
[Hindi English Marathi Bengali Gujarati]
[English Marathi Bengali Gujarati]
[Marathi Bengali Gujarati]
[Marathi]
```

4. Run `make.go` file by using.

```
go run make.go
```

## Output
The output of the above program is - 

```
>> go run make.go

Slices - make
The datatype of Scores is []int
panic: runtime error: index out of range [2] with length 2

goroutine 1 [running]:
main.main()
        /home/devops/go-lang/go/08-slices/make.go:12 +0xe6
exit status 2
```

In the above output there is error because we have increased the default values number wrt the defined value.

5. Run `make-append.go` file by using.

```
go run make-append.go
```

## Output
The output of the above program is - 

```
>> go run make-append.go

Slices - makeAppend
The datatype of Scores is []int
[56 78 23 45 67]
false
[23 45 56 67 78]
true
```