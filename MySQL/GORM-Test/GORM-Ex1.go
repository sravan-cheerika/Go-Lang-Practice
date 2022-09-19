package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product3 struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:root@tcp(127.0.0.1:3306)/MyDB?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error ", err)
	} else {
		fmt.Println("GORM Connected ", db)
	}
	//////////////////////
	// Migrate the schema
	db.AutoMigrate(&Product3{})
	// Create
	db.Create(&Product3{Code: "D42", Price: 100})
	// Read
	var product Product3
	db.First(&product, 1)                 // find product with integer primary key
	db.First(&product, "code = ?", "F42") // find product with code D42
	// Update - updat'e product's price to 200
	db.Model(&product).Update("Price", 300)
	// Update - update multiple fields
	db.Model(&product).Updates(Product3{Price: 50, Code: "sravan"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 150, "Code": "reddy"})
	// Delete - delete product
	//db.Delete(&product, 1)
}
