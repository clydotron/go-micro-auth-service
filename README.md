# Authentication Microservice 
Basic authentication microservice.
Uses PostgreSQL DB to store credentials.
Two ways to interact with:
via http (soon to be s)
gRPC

Currently does not provide interface to add register new users.


### External dependencies
// go get github.com/go-chi/chi/v5
// go get github.com/go-chi/cors
// go get github.com/jackc/pgconn
// go get github.com/jackc/pgx/v4