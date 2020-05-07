package app

import (
	"fmt"
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/dylan-clarry/blogapi/config"
	_ "github.com/lib/pq"
)

// App struct
type App struct {
	Router *mux.Router
}

// Initialize func
func (app *App) Initialize(config *config.Config) {
	fmt.Println("hello from app")
	fmt.Println(config.DB.User)

	// create db connection
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.DB.Host, config.DB.Port, config.DB.User, config.DB.Pass, config.DB.DBName)
	db, err := sql.Open(config.DB.Dialect, psqlInfo)

	if err != nil {
		println("connection error:")
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("connected to the database")
}