package app

import (
	"fmt"
	"log"
	"net/http"
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"github.com/dylan-clarry/blogapi/config"
	"github.com/dylan-clarry/blogapi/app/controllers"
)

// App struct
type App struct {
	Router 	*mux.Router
	DB 		*sql.DB
}

// Initialize func
func(app *App) Initialize(config *config.Config) {

	// create db connection
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.DB.Host, config.DB.Port, config.DB.User, config.DB.Pass, config.DB.DBName)
	db, err := sql.Open(config.DB.Dialect, psqlInfo)

	if err != nil {
		println("connection error:")
		panic(err)
	}
	defer db.Close()
	fmt.Println("connected to database")

	app.DB = db
	app.Router = mux.NewRouter()
	app.SetUserRoutes()
}

// Run starts the app on a given port
func(app *App) Run(port string) {
	log.Println("Server listening on port", port)
	log.Fatal(http.ListenAndServe(port, app.Router))

}


// ====================
// Routes
// ====================

// user endpoints
func(app *App) SetUserRoutes() {
	fmt.Println("setuserroute")

	app.Router.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
}

// blog post endpoints