package boostraps

import (
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

	repo := persistence.NewUserRepoDB(db)
	uc := userapp.NewUseCase(repo)
	h := handler.NewUserHandler(uc)

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

	r.POST("/users", h.GetAllUser)

	v1 := r.Group("/api/v1")
	{
		v1.GET("/users", h.GetAllUser)
		v1.GET("/users/:userID", h.GetUserByID)
		v1.POST("/users", h.CreateUser)
		v1.PUT("/users", h.UpdateUser)
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
