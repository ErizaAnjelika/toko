package main

import (
	"toko/configs"
	"toko/handlers"
	"toko/middlewares"
	"toko/migrations"

	"github.com/gin-gonic/gin"

	"net/http"
	_ "net/http/pprof"
)

func main() {
	db, err := configs.InitDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	migrations.Migrate(db)

	router := gin.Default()

	// Tambahkan prefix API
	api := router.Group("/api/v1")

	// Rute CRUD Produk
	api.GET("/products", middlewares.AuthMiddleware("admin"), handlers.ListProducts(db))
	api.GET("/products/:id", middlewares.AuthMiddleware("admin"), handlers.GetProduct(db))
	api.POST("/products", middlewares.AuthMiddleware("admin"), handlers.CreateProduct(db))
	api.PUT("/products/:id", middlewares.AuthMiddleware("admin"), handlers.UpdateProduct(db))
	api.DELETE("/products/:id", middlewares.AuthMiddleware("admin"), handlers.DeleteProduct(db))

	// Rute CRUD Product Category
	api.GET("/product-categories", middlewares.AuthMiddleware("admin"), handlers.ListProductCategories(db))
	api.GET("/product-categories/:id", middlewares.AuthMiddleware("admin"), handlers.GetProductCategory(db))
	api.POST("/product-categories", middlewares.AuthMiddleware("admin"), handlers.CreateProductCategory(db))
	api.PUT("/product-categories/:id", middlewares.AuthMiddleware("admin"), handlers.UpdateProductCategory(db))
	api.DELETE("/product-categories/:id", middlewares.AuthMiddleware("admin"), handlers.DeleteProductCategory(db))

	// rute CRUD StokProduct
	api.GET("/StokProducts", middlewares.AuthMiddleware(), handlers.ListStok(db))
	api.GET("/StokProducts/:id", middlewares.AuthMiddleware(), handlers.GetStokProduct(db))
	api.POST("/StokProducts", middlewares.AuthMiddleware(), handlers.CreateStokProduct(db))
	api.PUT("/StokProducts/:id", middlewares.AuthMiddleware(), handlers.UpdateStokProduct(db))
	api.DELETE("/StokProducts/:id", middlewares.AuthMiddleware(), handlers.DeleteStokProduct(db))

	// Rute Transaksi
	api.POST("/transactions", middlewares.AuthMiddleware(), handlers.CreateTransaction(db))
	api.GET("/transactions/:id", middlewares.AuthMiddleware(), handlers.GetTransactionWithItems(db))
	api.GET("/transactions", middlewares.AuthMiddleware(), handlers.ListTransactions(db))

	// Rute Chart
	api.GET("/chart/sales-by-date", middlewares.AuthMiddleware(), handlers.GetSalesByDate(db))
	api.GET("/chart/top-selling-products", middlewares.AuthMiddleware(), handlers.GetTopSellingProducts(db))

	// rute user
	api.GET("/users", middlewares.AuthMiddleware("admin"), handlers.ListUsers(db))
	api.POST("/register", middlewares.AuthMiddleware("admin"), handlers.Register(db))

	// Rute Auth
	api.POST("/login", handlers.Login(db))

	router.GET("/debug/pprof/*pprof", gin.WrapH(http.DefaultServeMux))

	router.Run(":8080")
}
