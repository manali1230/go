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
![Input](img/images/api-web.png)

api call in thunder client
![Input](img/images/api-call-thunder-client.png)

### Operations output

1. Get list of all courses available.
![Input](img/images/getAllCourses.png)

2. Get only one course
![Input](img/images/getOneCourse.png)

3. Create one course
![Input](img/images/createOneCourse.png)

check course created or not
![Input](img/images/getAllCourses-after-createOneCourse.png)

4. Update one course
![Input](img/images/updateOneCourse.png)

check course updated or not
![Input](img/images/getAllCourses-after-updateOneCourse.png)

5. Delete one course
![Input](img/images/deleteOneCourse.png)

check course deleted or not
![Input](img/images/checkDeleteCourse.png)

<style type="text/css">
    img {
        width: 250px;
    }
</style>