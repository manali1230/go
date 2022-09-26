## LOOP

[Loop](https://go.dev/tour/flowcontrol/1)

[Continue & Break](https://www.digitalocean.com/community/tutorials/using-break-and-continue-statements-when-working-with-loops-in-go)

[Statements](https://go.dev/ref/spec#Statements)

This is `Loop` go code and I am using Visual Studio WSL[Ubuntu] for performing practical. I have installed `go` Extension in visual studio.

1. this command will generate a `go.mod` file.

```
go mod init loop
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

Loops in GO
[Sunday Monday Tuesday Wednesday Thursday Friday Saturday]
Way 1
Sunday
Monday
Tuesday
Wednesday
Thursday
Friday
Saturday
Way 2
Sunday
Monday
Tuesday
Wednesday
Thursday
Friday
Saturday
Way 3
Index 0 : value Sunday
Index 1 : value Monday
Index 2 : value Tuesday
Index 3 : value Wednesday
Index 4 : value Thursday
Index 5 : value Friday
Index 6 : value Saturday
```

3. Run `while.go` file by using.

```
go run while.go 
```

## Output
The output of the above program is - 

```
>> go run while.go  

type of while loop in go
1
2
3
4
5
6
7
8
9
```

4. Run `continue.go` file by using.

```
go run continue.go
```

## Output
The output of the above program is - 

```
>> go run continue.go  

Continue
1
2
3
4
6
7
8
9
```


5. Run `break.go` file by using.

```
go run break.go
```

## Output
The output of the above program is - 

```
>> go run break.go  

Break
1
2
3
4
```

6. Run `goto.go` file by using.

```
go run goto.go
```

## Output
The output of the above program is - 

```
>> go run goto.go  

GOTO
1
2
3
4
This is GOTO!
```

