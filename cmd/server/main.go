package main

import (
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	//Tải các biến môi trường
	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Println("Có lỗi khi tải biến môi trường:", err)
	}

	gin.SetMode(os.Getenv("GIN_MODE"))

	db := GormConnection()

	route := InitRoute(db)

	port := os.Getenv("PORT")

	fmt.Println("Port của bạn: ", port)

	ln, err := net.Listen("tcp", "0.0.0.0:"+port)

	if err != nil {
		panic(err)
	}

	_ = http.Serve(ln, route)
}
