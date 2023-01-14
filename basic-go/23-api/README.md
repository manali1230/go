# API

[API](https://go.dev/doc/tutorial/web-service-gin)

This is `API` go code and I am using Visual Studio Code for performing practical. I have installed `go` Extension in visual studio.
while writing code if VS Code is showing install tools related to go then click on install.

[SemanticVersioning](https://www.geeksforgeeks.org/introduction-semantic-versioning/)

[GoModulesReference](https://go.dev/ref/mod)

[GorillaMux](https://pkg.go.dev/github.com/gorilla/mux)

1. this command will generate a `go.mod` file.

```
go mod init github.com/manali1230/api
```

2. get gorilla/mux module

```
>> go get -u github.com/gorilla/mux
go: added github.com/gorilla/mux v1.8.0
```

3. build api

```
go build .
```

4. run go

```
>> go run main.go 
CRUD OPERATIONS
```
5. hit the url `http://localhost:1000` and get the result .

<img src="./images/api-web.png" width="400">

api call in thunder client

<img src="./images/api-call-thunder-client.png" width="400">

### Operations output

1. Get list of all courses available.

<img src="./images/getAllCourses.png" width="400">

2. Get only one course

<img src="./images/getOneCourse.png" width="400">

3. Create one course

<img src="./images/createOneCourse.png" width="300">

check course created or not

<img src="./images/getAllCourses-after-createOneCourse.png" width="300">

4. Update one course

<img src="./images/updateOneCourse.png" width="300">

check course updated or not

<img src="./images/getAllCourses-after-updateOneCourse.png" width="300">

5. Delete one course

<img src="./images/deleteOneCourse.png" width="300">

check course deleted or not\

<img src="./images/checkDeleteCourse.png" width="300">