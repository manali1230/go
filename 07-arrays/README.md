## ARRAYS

[Array](https://go.dev/tour/moretypes/6)

This is `Array` go code and I am using Visual Studio WSL[Ubuntu] for performing practical. I have installed `go` Extension in visual studio.

1. this command will generate a `go.mod` file.

```
go mod tidy arrays
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

Arrays
Value of language list is [Hindi  English]
Length of language list is 3
Value of Computer language list is [C C++]
Length of Computer language list is 2
```

In the above output, the values passed in arrays are two but langList variable length is 3 and compLangList is 2. Length defined by you is considered as length of array and so length is not calculated on the basis of input/values assigned to an array.