package app

import "github.com/gorilla/mux"

type App struct {
	Router *mux.Router
}

func New() *App {
	a := &App{
		Router: mux.NewRouter(),
	}

	a.initRoutes()

	return a
}

func (a *App) initRoutes() {
	a.Router.HandleFunc("/", a.IndexHandler()).Methods("GET")
	a.Router.HandleFunc("/blockSwaps", a.swapsHandler()).Methods("GET")
	a.Router.HandleFunc("/assetPools", a.assetPoolsHandler()).Methods("GET")
	a.Router.HandleFunc("/assetVolume", a.assetVolumeHandler()).Methods("GET")
	a.Router.HandleFunc("/allAssetsSwapped", a.allAssetsSwapped()).Methods("GET")
}
