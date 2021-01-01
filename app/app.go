package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/onurkybsi/rester/app/handler"
	"github.com/onurkybsi/rester/app/model"
	"github.com/onurkybsi/rester/config"
)

// App has router and db instances
type App struct {
	Router *mux.Router
}

// Init Sets the necessary configurations for the server
func (a *App) Init(config *config.Config) {
	a.Router = mux.NewRouter()

	a.post("/loadtest/reqSeq", handler.ReqSequential)
	a.post("/loadtest/reqSimultaneously", handler.ReqSequential)
}

// Run Runs the server on the specified port
func (a *App) Run(port string) {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), a.Router))
}

func (a *App) get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods(model.GetMethod)
}

func (a *App) post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods(model.PostMethod)
}

func (a *App) put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods(model.PutMethod)
}

func (a *App) delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods(model.DeleteMethod)
}
