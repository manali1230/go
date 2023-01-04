## JSON

[JSON](https://go.dev/blog/json)

This is `json` go code and I am using Visual Studio Code for performing practical. I have installed `go` Extension in visual studio.

1. this command will generate a `go.mod` file.

```
go mod init json
```

while writing code if VS Code is showing install tools related to go then click on install.

2. Run `main.go` file by using.

```
go run main.go
```

## Output

```
Json Data
[{"Name":"Python Course","Price":2000,"CourseName":"Python","Tutor":["Rina","Mina"]},{"Name":"Go Course","Price":3000,"CourseName":"Go","Tutor":["Manoj,Saroj"]}]
```

3. Run `indented-json.go` file by using.

```
go run indented-json.go
```

## Output

```
Json Data
[
        {
                "Name": "Python Course",
                "Price": 2000,
                "CourseName": "Python",
                "Tutor": [
                        "Rina",
                        "Mina"
                ]
        },
        {
                "Name": "Go Course",
                "Price": 3000,
                "CourseName": "Go",
                "Tutor": [
                        "Manoj,Saroj"
                ]
        }
]
```

4. Run `small-key.go` file by using.

```
go run small-key.go
```

## Output

```
Json Data
[
        {
                "name": "Python Course",
                "price": 2000,
                "CourseName": "Python",
                "tutor": [
                        "Rina",
                        "Mina"
                ]
        },
        {
                "name": "Go Course",
                "price": 3000,
                "CourseName": "Go",
                "tutor": [
                        "Manoj,Saroj"
                ]
        }
]
```

5. Run `decode-json.go` file by using.

```
go run decode-json.go
```

## Output

```
Decode Json
Json is Valid

main.course{Name:"Go Course", Price:3000, CourseName:"Go", Tutor:[]string{"Manoj,Saroj"}}

map[string]interface {}{"CourseName":"Go", "name":"Go Course", "price":3000, "tutor":[]interface {}{"Manoj,Saroj"}}
Key : tutor and value : [Manoj,Saroj] => datatype : []interface {}
Key : name and value : Go Course => datatype : string
Key : price and value : 3000 => datatype : float64
Key : CourseName and value : Go => datatype : string
```