package routes

import (
	_authHandler "group-project-2/delivery/handler/auth"
	_cartHandler "group-project-2/delivery/handler/cart"
	_orderHandler "group-project-2/delivery/handler/order"
	_productHandler "group-project-2/delivery/handler/product"
	_userHandler "group-project-2/delivery/handler/user"
	_middlewares "group-project-2/delivery/middlewares"

	"github.com/labstack/echo/v4"
)

func RegisterAuthPath(e *echo.Echo, ah *_authHandler.AuthHandler) {
	e.POST("/login", ah.LoginHandler())
}
func RegisterPathUser(e *echo.Echo, uh *_userHandler.UserHandler) {
	e.POST("/users", uh.PostUserHandler())
	e.GET("/users/:id", uh.GetUserHandler(), _middlewares.JWTMiddleware())
	e.PUT("/users/:id", uh.PutUserHandler(), _middlewares.JWTMiddleware())
	e.DELETE("/users/:id", uh.DeleteUserHandler(), _middlewares.JWTMiddleware())
}

func RegisterPathProduct(e *echo.Echo, ph *_productHandler.ProductHandler) {
	e.POST("/products", ph.AddProductHandler(), _middlewares.JWTMiddleware())
	e.PUT("/products/:id", ph.UpdateProductByIdHandler(), _middlewares.JWTMiddleware())
	e.DELETE("/products/:id", ph.DeleteProductByIdHandler(), _middlewares.JWTMiddleware())
	e.GET("/products", ph.GetAllProductHandler())
	e.GET("/products/:id", ph.GetProductByIdHandler())
}
func RegisterPathCart(e *echo.Echo, ch *_cartHandler.CartHandler) {
	e.POST("/carts", ch.PostCartHandler(), _middlewares.JWTMiddleware())
	e.GET("/carts", ch.GetCartHandler(), _middlewares.JWTMiddleware())
	e.PUT("/carts/:id", ch.PutCartHandler(), _middlewares.JWTMiddleware())
	e.DELETE("/carts/:id", ch.DeleteCartHandler(), _middlewares.JWTMiddleware())

}
func RegisterPathOrder(e *echo.Echo, oh *_orderHandler.OrderHandler) {
	e.POST("/order", oh.PostOrderHandler(), _middlewares.JWTMiddleware())
	e.GET("/order", oh.GetOrderHandler(), _middlewares.JWTMiddleware())
}
