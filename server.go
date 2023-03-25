package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	db := connect()
	err := db.AutoMigrate(&User{})
	if err != nil {
		fmt.Printf("Failed to migrate: \n%v\n\n", err)
	}

	app := fiber.New()

	app.Get("/", indexHandler)

	app.Post("/", postHandler)

	app.Put("/update", putHandler)

	app.Delete("/delete", deleteHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}

func indexHandler(context *fiber.Ctx) error {
	return context.SendString("Hello")
}
func postHandler(context *fiber.Ctx) error {
	return context.SendString("Hello")
}
func putHandler(context *fiber.Ctx) error {
	return context.SendString("Hello")
}

func deleteHandler(context *fiber.Ctx) error {
	return context.SendString("Hello")
}

func deleteFileHandler(context *fiber.Ctx) {
	connect().Delete(&User{}, "id = ?", context.Params("id"))
}

/*
GORM Methods:
   	- Create(value interface{}) *DB: Insert a new record into the database.
   	- Save(value interface{}) *DB: Insert or update a record in the database.
   	- Updates(values interface{}) *DB: Update the specified columns of a record in the database.
	- Delete(value interface{}, opts ...interface{}) *DB: Delete a record from the database.
   	- Find(dest interface{}, conds ...interface{}) *DB: Query the database for records that match the specified
	conditions and store the results in the dest variable.
   	- First(dest interface{}, conds ...interface{}) *DB: Query the database for the first record that matches the
	specified conditions and store the result in the dest variable.
   	- Raw(sqlxstring, values ...interface{}) *DB: Execute a raw SQL query and store the results in the dest variable.
   	- Where(query interface{}, args ...interface{}) *DB: Add a WHERE clause to a query based on the specified
	conditions.
	- Order(value interface{}) *DB: Specify the order in which the query results should be returned.
   	- Limit(n int) *DB: Limit the number of results returned by a query.
   	- Offset(n int) *DB: Offset the starting position of the results returned by a query.
*/

func connect() *gorm.DB {
	dsn := "postgresql://collin:-m1M5vOeQBlV5Z_88OFBQQ@brisk-shade-2330.g95.cockroachlabs.cloud:26257/defaultdb?sslmode=verify-full"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	return db
}
