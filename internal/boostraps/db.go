package boostraps

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func GormConnection() *gorm.DB {
	var db *gorm.DB

	var (
		devHostName = os.Getenv("MYSQL_HOST")
		devDbName   = os.Getenv("MYSQL_NAME")
		devUser     = os.Getenv("MYSQL_USER")
		devPassword = os.Getenv("MYSQL_PASSWORD")
		devPort     = os.Getenv("MYSQL_PORT")
	)

	// In ra để kiểm tra nếu cần debug
	fmt.Println("MYSQL_HOST:", devHostName)
	fmt.Println("MYSQL_PORT:", devPort)
	fmt.Println("MYSQL_USER:", devUser)
	fmt.Println("MYSQL_PASSWORD:", devPassword)
	fmt.Println("MYSQL_NAME:", devDbName)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		devUser,
		devPassword,
		devHostName,
		devPort,
		devDbName,
	)

	var err error

	// Retry tối đa 10 lần, mỗi lần cách nhau 5s
	for i := 0; i < 10; i++ {
		db, err = gorm.Open(
			mysql.Open(dsn),
			&gorm.Config{
				NamingStrategy: schema.NamingStrategy{
					SingularTable: true,
				},
			},
		)
		if err == nil {
			fmt.Println("✅ Kết nối MySQL thành công!")
			return db
		}
		fmt.Printf("❌ Không thể kết nối MySQL, thử lại sau 5s (%d/10)...\n", i+1)
		fmt.Println("DSN:", dsn)
		time.Sleep(5 * time.Second)
	}

	// Sau 10 lần vẫn lỗi thì panic
	log.Fatalf("Lỗi khi kết nối MySQL sau nhiều lần thử: %v", err)
	return nil
}
