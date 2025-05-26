package boostraps

import (
	"final_project/internal/application/adminapp"
	"final_project/internal/application/itemapp"
	"final_project/internal/application/userapp"
	"final_project/internal/infrastructure/persistence"
	"final_project/internal/interface/http/handler"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	_ "final_project/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoute(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	//user dependency
	userRepo := persistence.NewUserRepoDB(db)
	userUC := userapp.NewUseCase(userRepo)
	userHandler := handler.NewUserHandler(userUC)

	//item dependency
	itemRepo := persistence.NewItemRepoDB(db)
	itemUC := itemapp.NewUseCase(itemRepo)
	itemHandler := handler.NewItemHandler(itemUC)

	//admin dependency
	adminRepo := persistence.NewAdminRepoDB(db)
	adminUC := adminapp.NewUseCase(adminRepo)
	adminHandler := handler.NewAdminHandler(adminUC)

	r.Use(func(c *gin.Context) {
		// Thêm header CORS cho mỗi request
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")                                       // Cho phép tất cả các origin
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS") // Các phương thức HTTP cho phép
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")                                      // Các header cho phép

		// Xử lý preflight OPTIONS
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	})

	url := ginSwagger.URL("/swagger/doc.json")

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	v1 := r.Group("/api/v1")
	{
		//user CRUD API
		v1.GET("/users", userHandler.GetAllUser)
		v1.GET("/users/:userID", userHandler.GetUserByID)
		v1.POST("/users", userHandler.CreateUser)
		v1.PUT("/users", userHandler.UpdateUser)
		v1.DELETE("/users/:userID", userHandler.DeleteUser)

		//item CRUD API
		v1.GET("/items", itemHandler.GetAllItem)
		v1.GET("/items/:itemID", itemHandler.GetItemByID)
		v1.POST("/items", itemHandler.CreateItem)
		v1.PUT("/items", itemHandler.UpdateItem)
		v1.DELETE("/items/:itemID", itemHandler.DeleteItem)

		//admin CRUD API
		v1.GET("/admins", adminHandler.GetAllAdmins)
		v1.GET("/admins/:adminID", adminHandler.GetAdminByID)
		v1.POST("/admins", adminHandler.CreateAdmin)
		v1.PUT("/admins", adminHandler.UpdateAdmin)
		v1.DELETE("/admins/:adminID", adminHandler.DeleteAdmin)
	}

	// r.Static("/public/images", "./public/images")

	//test
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return r
}
