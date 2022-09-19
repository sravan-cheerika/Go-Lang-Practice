package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
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
	db.AutoMigrate(&Product2{})
	// Create
	db.Create(&Product2{Code: "D42", Price: 100})
	// Read
	var product Product2
	/*db.First(&product2, 1)                 // find product with integer primary key
	db.First(&product2, "code = ?", "D42") // find product with code D42
	// Update - update product's price to 200
	db.Model(&Product2).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	// Delete - delete product
	db.Delete(&product, 1) */
}
