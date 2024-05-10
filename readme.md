# GO-BNB

A demo application for a system that manages a small Bed & Breakfast.

---

## Running the application

```shell
go run ./cmd/web 
```

## Running tests

```shell
go test -v     
go test -cover
go test -coverprofile=coverage.out && go tool cover -html=coverage.out
```


---

## Dependencies
- Built with the [Go Language](https://go.dev/)
- [Chi Router](https://github.com/go-chi/chi) - Application Router
- [nosurf](https://github.com/justinas/nosurf) - CSRF Tokens
- [SCS: HTTP Session Management for Go](https://github.com/alexedwards/scs) - Session Management
- [`govalidator`](https://github.com/asaskevich/govalidator) - 
- 