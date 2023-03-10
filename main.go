package main

import (
	"fmt"
	"housy/database"
	"housy/pkg/mysql"
	// "housy/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// initial DB
	mysql.DatabaseInit()

	// migration
	database.RunMigration()

	r := mux.NewRouter()

	// routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	fmt.Println("server running localhost: 5000")
	http.ListenAndServe("localhost:5000", r)
}
