package main

import (
	"log"
	"net/http"

	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/plugins/admin"
	// "github.com/GoAdminGroup/go-admin/plugins/admin/modules/table" // Ensure this import
	// "github.com/GoAdminGroup/go-admin/template/types"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin
	r := gin.Default()

	// Initialize GoAdmin Engine
	eng := engine.Default()

	// Configure GoAdmin
	cfg := config.Config{
		Databases: config.DatabaseList{
			"default": {
				Driver: "mysql",
				Dsn:    "root:yourpassword@tcp(127.0.0.1:3306)/myapp", // Replace with your MySQL credentials
			},
		},
		UrlPrefix: "admin", // Admin panel URL: http://localhost:8080/admin
		Store: config.Store{
			Path:   "./uploads",
			Prefix: "uploads",
		},
		Theme: "adminlte",
	}

	// Attach GoAdmin to Gin
	err := eng.AddConfig(&cfg).Use(r)
	if err != nil {
		log.Fatal(err)
	}

	// Create Admin Plugin
	adminPlugin := admin.NewAdmin()

	// Add a Custom Page to Admin (Using table.Generator)
	// adminPlugin.AddGenerator("user", func(ctx *table.Context) (types.Panel, error) {
	// 	return types.Panel{
	// 		Title: "User Management",
	// 	}, nil
	// })

	// Register Admin Plugin
	eng.AddPlugins(adminPlugin)

	// Start Server
	r.StaticFS("/uploads", http.Dir("./uploads"))
	r.Run(":8080")
}
