package controllers

import (
	"github.com/labstack/echo/v4"
)

func (server *Server) initializeRoutes() {
	server.Router = echo.New()

	server.Router.GET("/", server.Home)

	server.Router.POST("/register", server.RegisterUser)
	server.Router.POST("/login", server.Login)

	server.Router.GET("/products/list", server.ListProducts)
	server.Router.GET("/products/:code", server.GetProductByCode)

	server.Router.POST("/cart", server.AddItemToCart)

	//server.Router.HandleFunc("/carts", server.GetCart).Methods("GET")
	//server.Router.HandleFunc("/carts/update", server.UpdateCart).Methods("POST")
	//server.Router.HandleFunc("/carts/remove/{id}", server.RemoveItemByID).Methods("GET")
	//
	//staticFileDirectory := http.Dir("./assets/")
	//staticFileHandler := http.StripPrefix("/public/", http.FileServer(staticFileDirectory))
	//server.Router.PathPrefix("/public/").Handler(staticFileHandler).Methods("GET")
}
