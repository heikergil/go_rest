# Go: SPIKE

## Instructions: 
- clone the repository
- install go following this steps https://go.dev/doc/install
-  run $go run main.go
- interact with the api with these end points on postman:
    GET /users
    POST /users
    send body: {
        "name":"any name"
    }
    GET /users/:id
    PATCH /users/:id
    DELETE /user/:id
