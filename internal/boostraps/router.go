package boostraps

import (
	"final_project/internal/application/authapp"
	"final_project/internal/application/categoryapp"
	"final_project/internal/application/itemapp"
	"final_project/internal/application/postapp"
	"final_project/internal/application/userapp"
	"final_project/internal/domain/auth"
	"final_project/internal/domain/post"
	persistence "final_project/internal/infrastructure/persistence/repo"
	"final_project/internal/infrastructure/redisrepo"
	"final_project/internal/infrastructure/seeder"
	"final_project/internal/interface/http/handler"
	middlewares "final_project/internal/interface/http/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	_ "final_project/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoute(db *gorm.DB, redisClient *redis.Client) *gin.Engine {
	r := gin.Default()

	//redis
	redisRepo := redisrepo.NewRedisRepo(redisClient)

	//role permission dependency
	rolePerRepo := persistence.NewRolePerRepoDB(db)

	//category dependency
	categoryRepo := persistence.NewCategoryRepoDB(db)
	categoryUC := categoryapp.NewUseCase(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryUC)

	//user dependency
	userRepo := persistence.NewUserRepoDB(db)
	userUC := userapp.NewUseCase(userRepo, rolePerRepo)
	userHandler := handler.NewUserHandler(userUC)

	//item dependency
	itemRepo := persistence.NewItemRepoDB(db)
	itemUC := itemapp.NewUseCase(itemRepo)
	itemHandler := handler.NewItemHandler(itemUC)

	//post dependency
	postService := post.NewPostService()
	postRepo := persistence.NewPostRepoDB(db)
	postUC := postapp.NewUseCase(postRepo, userRepo, rolePerRepo, postService, itemRepo, categoryRepo)
	postHandler := handler.NewPostHandler(postUC)

	//auth dependency
	authService := auth.NewAuthService()
	authRepo := persistence.NewAuthRepoDB(db)
	authUC := authapp.NewUseCase(authRepo, authService, redisRepo)
	authHandler := handler.NewAuthHandler(authUC)

	seed := seeder.NewSeeder(
		rolePerRepo,
		itemRepo,
		userRepo,
		categoryRepo,
	)

	seed.Seed()

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

	// url := ginSwagger.URL("/swagger/doc.json")

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	{
		//user CRUD API
		v1.GET("/users", userHandler.GetAllUser)
		v1.GET("/users/:userID", userHandler.GetUserByID)
		v1.POST("/users", userHandler.CreateUser)
		v1.PATCH("/users/:userID", userHandler.UpdateUser)
		v1.DELETE("/users/:userID", userHandler.DeleteUser)

		//item CRUD API
		v1.GET("/items", itemHandler.GetAllItem)
		v1.GET("/items/:itemID", itemHandler.GetItemByID)
		v1.POST("/items", itemHandler.CreateItem)
		v1.PATCH("/items/:itemID", itemHandler.UpdateItem)
		v1.DELETE("/items/:itemID", itemHandler.DeleteItem)

		//post API
		v1.GET("/posts", postHandler.GetAllAdminPost)
		v1.GET("/posts/:postID", postHandler.GetPostByID)
		v1.POST("/posts", middlewares.AuthGuard, postHandler.CreatePost)
		v1.PATCH("/posts/:postID", postHandler.UpdatePost)

		//category API
		v1.GET("/categories", middlewares.AuthGuard, categoryHandler.GetAll)

		//auth API
		v1.POST("/login", authHandler.Login)
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
