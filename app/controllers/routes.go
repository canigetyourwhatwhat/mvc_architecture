package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (server *Server) initializeRoutes() {
	server.Router = echo.New()
	server.Router.Router().Add(http.MethodGet, "/", server.Home)

	//server.Router.HandleFunc("/", server.Home).Methods("GET")
	//server.Router.HandleFunc("/products", server.Products).Methods("GET")
	//server.Router.HandleFunc("/products/{slug}", server.GetProductBySlug).Methods("GET")
	//
	//server.Router.HandleFunc("/carts", server.GetCart).Methods("GET")
	//server.Router.HandleFunc("/carts", server.AddItemToCart).Methods("POST")
	//server.Router.HandleFunc("/carts/update", server.UpdateCart).Methods("POST")
	//server.Router.HandleFunc("/carts/remove/{id}", server.RemoveItemByID).Methods("GET")
	//
	//staticFileDirectory := http.Dir("./assets/")
	//staticFileHandler := http.StripPrefix("/public/", http.FileServer(staticFileDirectory))
	//server.Router.PathPrefix("/public/").Handler(staticFileHandler).Methods("GET")
}
