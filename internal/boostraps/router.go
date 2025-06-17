package boostraps

import (
	"final_project/internal/application/app/authapp"
	"final_project/internal/application/app/categoryapp"
	"final_project/internal/application/app/commentapp"
	"final_project/internal/application/app/importinvoiceapp"
	"final_project/internal/application/app/interestapp"
	"final_project/internal/application/app/itemapp"
	"final_project/internal/application/app/postapp"
	"final_project/internal/application/app/roleapp"
	"final_project/internal/application/app/transactionapp"
	"final_project/internal/application/app/userapp"
	"final_project/internal/application/app/warehouseapp"
	"final_project/internal/application/worker/chatapp"
	"final_project/internal/domain/auth"
	importinvoice "final_project/internal/domain/import_invoice"
	"final_project/internal/domain/post"
	persistence "final_project/internal/infrastructure/persistence/repo"
	redisapp "final_project/internal/infrastructure/redis"
	"final_project/internal/infrastructure/seeder"
	"final_project/internal/interface/http/handler"
	middlewares "final_project/internal/interface/http/middleware"
	workerHandler "final_project/internal/interface/worker/handler"
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

	stream := "chatstream"
	group := "chatgroup"
	consumer := "worker-chat"

	//redis
	redisRepo := redisapp.NewRedisRepo(redisClient)

	//role permission dependency
	rolePerRepo := persistence.NewRolePerRepoDB(db)
	roleUC := roleapp.NewUseCase(rolePerRepo)
	roleHandler := handler.NewRoleHandler(roleUC)

	//category dependency
	categoryRepo := persistence.NewCategoryRepoDB(db)
	categoryUC := categoryapp.NewUseCase(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryUC)

	//user dependency
	userRepo := persistence.NewUserRepoDB(db)
	userUC := userapp.NewUseCase(userRepo, rolePerRepo, redisRepo)
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

	//interest dependency
	interestRepo := persistence.NewInterestRepoDB(db)
	interestUC := interestapp.NewUseCase(interestRepo, userRepo, postRepo)
	interestHandler := handler.NewInterestHandler(interestUC)

	//transaction dependency
	transactionRepo := persistence.NewTransactionRepoDB(db)
	transactionUC := transactionapp.NewUseCase(transactionRepo, userRepo, interestRepo, itemRepo, postRepo)
	transactionHandler := handler.NewTransactionHandler(transactionUC)

	//import invoice dependency
	importInvoiceService := importinvoice.NewImportInvoiceService()
	importInvoiceRepo := persistence.NewImportInvoiceRepoDB(db)
	importInvoiceUC := importinvoiceapp.NewUseCase(importInvoiceRepo, userRepo, itemRepo, importInvoiceService)
	importInvoiceHandler := handler.NewImportInvoiceHandler(importInvoiceUC)

	//warehouse dependency
	warehouseRepo := persistence.NewWarehouseRepoDB(db)
	warehouseUC := warehouseapp.NewUseCase(warehouseRepo)
	warehouseHandler := handler.NewWarehouseHandler(warehouseUC)

	//message dependency
	commentRepo := persistence.NewCommentRepoDB(db)
	commentUC := commentapp.NewUseCase(commentRepo)
	commentHandler := handler.NewCommentHandler(commentUC)

	//chat dependency
	chatUC := chatapp.NewUseCase(commentRepo)

	//auth dependency
	authService := auth.NewAuthService()
	authRepo := persistence.NewAuthRepoDB(db)
	authUC := authapp.NewUseCase(authRepo, authService, redisRepo, rolePerRepo, userRepo)
	authHandler := handler.NewAuthHandler(authUC)

	redisSeed := redisapp.NewRedisSeeder(
		redisRepo,
		rolePerRepo,
	)

	seed := seeder.NewSeeder(
		rolePerRepo,
		itemRepo,
		userRepo,
		categoryRepo,
		postRepo,
		postService,
		importInvoiceRepo,
		importInvoiceService,
	)

	seed.Seed()
	redisSeed.SeedInitialData()

	//run chat worker
	streamConsumer := redisapp.NewStreamConsumer(redisClient, stream, group, consumer)
	streamConsumer.CreateConsumerGroup()

	chatHandler := workerHandler.NewChatHandler(streamConsumer, chatUC)

	go chatHandler.Run()

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
		//role API
		v1.GET("/roles", roleHandler.GetAll)

		//client API
		v1.GET("/clients", userHandler.GetAllClient)
		v1.GET("/clients/:clientID", userHandler.GetClientByID)
		v1.POST("/clients", userHandler.CreateClient)
		v1.PATCH("/clients/:clientID", userHandler.UpdateClient)
		v1.DELETE("/clients/:clientID", userHandler.DeleteClient)

		//admin API
		v1.GET("/admins", userHandler.GetAllAdmin)
		v1.GET("/admins/:adminID", userHandler.GetAdminByID)
		v1.POST("/admins", userHandler.CreateAdmin)
		v1.PATCH("/admins/:adminID", userHandler.UpdateAdmin)
		v1.DELETE("/admins/:adminID", userHandler.DeleteAdmin)

		//item API
		v1.GET("/items", itemHandler.GetAllItem)
		v1.GET("/items/:itemID", itemHandler.GetItemByID)
		v1.POST("/items", itemHandler.CreateItem)
		v1.PATCH("/items/:itemID", itemHandler.UpdateItem)
		v1.DELETE("/items/:itemID", itemHandler.DeleteItem)

		//client post API
		v1.GET("/client/posts", postHandler.GetAllPost)
		v1.GET("/client/posts/:postID", postHandler.GetPostByID)

		//post API
		v1.GET("/posts", middlewares.AuthGuard, postHandler.GetAllAdminPost)
		v1.GET("/posts/:postID", postHandler.GetPostByID)
		v1.GET("/posts/slug/:postSlug", postHandler.GetPostBySlug)
		v1.POST("/posts", middlewares.AuthGuard, postHandler.CreatePost)
		v1.PATCH("/posts/:postID", postHandler.UpdatePost)

		//category API
		v1.GET("/categories", categoryHandler.GetAll)

		//interest API
		v1.GET("/interests", middlewares.AuthGuard, interestHandler.GetAll)
		v1.GET("/interests/:interestID", middlewares.AuthGuard, interestHandler.GetByID)
		v1.POST("/interests", middlewares.AuthGuard, interestHandler.Create)
		v1.DELETE("/interests/:postID", middlewares.AuthGuard, interestHandler.Delete)

		//transaction API
		v1.GET("/transactions", middlewares.AuthGuard, transactionHandler.GetAll)
		v1.POST("/transactions", middlewares.AuthGuard, transactionHandler.Create)
		v1.PATCH("/transactions/:transactionID", middlewares.AuthGuard, transactionHandler.Update)

		//import invoice API
		v1.POST("/import-invoice", middlewares.AuthGuard, importInvoiceHandler.CreateImportInvoice)
		v1.GET("/import-invoice", middlewares.AuthGuard, importInvoiceHandler.GetAllImportInvoice)

		//warehouse API
		v1.GET("/warehouses", middlewares.AuthGuard, warehouseHandler.GetAll)
		v1.GET("/warehouses/:warehouseID", middlewares.AuthGuard, warehouseHandler.GetByID)
		v1.PATCH("/warehouses/:warehouseID", middlewares.AuthGuard, warehouseHandler.Update)

		//item warehouse API
		v1.GET("/item-warehouses", middlewares.AuthGuard, warehouseHandler.GetAllItem)
		v1.GET("/item-warehouses/:itemCode", middlewares.AuthGuard, warehouseHandler.GetItemByCode)

		//message API
		v1.GET("/messages", middlewares.AuthGuard, commentHandler.GetAll)

		//auth API
		v1.GET("/get-me", middlewares.AuthGuard, authHandler.AdminGetMe)
		v1.GET("/client/get-me", middlewares.AuthGuard, authHandler.ClientGetMe)
		v1.POST("/client/login", authHandler.UserLogin)
		v1.POST("/login", authHandler.AdminLogin)
		v1.POST("/refresh-token", authHandler.GetAccessToken)
		v1.POST("/logout", middlewares.AuthGuard, authHandler.AdminLogout)
		v1.POST("/client/logout", middlewares.AuthGuard, authHandler.ClientLogout)
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
