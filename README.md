# An example project to reproduce a problem with nesting http serve mux

## Description

Suppose you have restful api like  POST for  `/api/users` and POST for `/api/posts`

And suppose we  would like to implement these restful api's by nesting the httpServeMux like in the snipper below:

```go
userHandler := userHandler{}
postHandler := postHandler{}

httpServeMux := http.NewServeMux()
httpServeMux.Handle("/api/users/", http.StripPrefix("/api/users", userHandler.Handler())
httpServeMux.Handle("/api/posts/", http.StripPrefix("/api/posts", postHandler.Handler()
```

where, for example, `userHandler.Handler()`  implementation looks like:

```go
func (u *userHandler) Handler() http.Handler {
    httpServeMux := http.NewServeMux()
    httpServeMux.HandleFunc("POST /", func(w http.ResponseWriter, req *http.Request) { /* should handle POST /api/users */ }
    return httpServeMux
}
```

The implementation of `postHandler.Handler()` is similar.

The problem with this implementation is that calling `POST /api/users`(NO trailing slash) result in `405 Method not Allowed`.

Please note that it WORKS FINE when calling  `POST /api/users/`(note the trailing slash)

Additionally, if we have an additional GET request handler, the `userHandler.Handler()` would then look like:

```go
func (u *userHandler) Handler() http.Handler {
    httpServeMux := http.NewServeMux()
    httpServeMux.HandleFunc("GET /", func(w http.ResponseWriter, req *http.Request) { /* should handle GET /api/users */ }
    httpServeMux.HandleFunc("POST /", func(w http.ResponseWriter, req *http.Request) { /* should handle POST /api/users */ }
    return httpServeMux
}
```

then calling `GET /api/users`(NO trailing slash) and `GET /api/users/`(note the trailing slash) are working.

Calls to `POST /api/users/`(note the trailing slash) are working fine. 

But calls  `POST /api/users`(NO trailing slash)
are actually handled by the GET handler, what's going on?

Is this a bug or an intended behavior?

## Instructions to reproduce

1. Clone repo and navigate to the root of this repo.

2. Run the program by

```bash
go run .
```

3. Test with Postman(a collection with v2.1 format is included in the repo). 
For some reason, I am unable to test any of the endpoints with cURL because it always results with `301 Moved Permanently`
when testing without trailing slash. Still, here are the cURL requests:

```bash
curl http://localhost:8080/api/users
```

```bash
curl -X POST http://localhost:8080/api/users
```

4. These cURL command work same both with cURL and Postman.

```bash
curl http://localhost:8080/api/users/
```

```bash
curl -X POST http://localhost:8080/api/users/
```
