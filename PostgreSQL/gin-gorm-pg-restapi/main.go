package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string `json:"name" gorm:"name"`
	Email string `json:"email" gorm:"email"`
}

var db *gorm.DB

// The GET /users endpoint retrieves all users from the database
func getUsers(c *gin.Context) {
	var users []User
	result := db.Find(&users)
	fmt.Println("frist time: ", result.RowsAffected) // returns count of records found
	fmt.Println("first time: ", result.Error)        // returns error or nil
	c.JSON(200, users)
}

// The POST /users endpoint creates a new user
func createUser(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatus(400)
		return
	}

	// Get first matched record
	r1 := db.Where("name = ?", user.Name).First(&user)
	// SELECT * FROM users WHERE name = 'jinzhu' ORDER BY id LIMIT 1;

	var users []User
	// Get all matched records
	r2 := db.Where("email = ?", user.Email).Find(&users)
	// SELECT * FROM users WHERE email = 'jinzhu';

	if r1.RowsAffected == 0 && r2.RowsAffected == 0 {
		result := db.Create(&user) // pass pointer of data to Create
		fmt.Println(result.RowsAffected)
		fmt.Println(result.Error)
		c.JSON(201, user)
	} else if r1.RowsAffected != 0 {
		c.JSON(400, "name da ton tai")
	} else {
		c.JSON(400, "email da ton tai")
	}
}

// The GET /users/:id endpoint fetches a user by ID
func getUserByID(c *gin.Context) {
	id := c.Param("id")
	var user User

	result := db.First(&user, id)
	if result.RowsAffected == 1 {
		c.JSON(200, user)
	} else {
		c.JSON(404, "not found user in database")
	}
}

// The UPDATE /users/:id endpoint update a user by ID
func updateUserByID(c *gin.Context) {
	var user_update User

	if err := c.BindJSON(&user_update); err != nil {
		c.AbortWithStatus(400)
		return
	}

	id := c.Param("id")
	var user User
	result := db.First(&user, id)

	// check user by id is already in database
	if result.Error != nil {
		c.AbortWithStatus(500)
		return
	}
	if result.RowsAffected == 0 {
		c.AbortWithStatus(404)
		return
	}

	// check name and email update is in the database???
	var user1 User
	r := db.Where("(name = ? AND id <> ?) OR (email = ? AND id <> ?)", user_update.Name, id, user_update.Email, id).First(&user1)
	
	fmt.Println("error: ", r.Error)
	fmt.Println("count: ", r.RowsAffected)

	if r.RowsAffected != 0 {
		c.JSON(400, "duplicate name or email")
		return
	}

	user.Name = user_update.Name
    user.Email = user_update.Email
    db.Save(&user)

    c.JSON(200, gin.H{
        "success": true,
        "message": "Cập nhật thành công",
    })
}

// The DELETE /users/:id endpoint delete a user by ID
func deleteUserByID(c *gin.Context) {
	id := c.Param("id")
	var user User

	result := db.First(&user, id)

	if result.RowsAffected == 1 {
		db.Delete(&user)
		c.JSON(200, user)
	} else {
		c.JSON(404, "user not found in database")
	}
}

func main() {
	var err error
	dsn := "host=localhost user=postgres password=quan1234 dbname=gin-gorm-pg port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("successfully connected database")
	}

	// Auto migrate the User model
	db.AutoMigrate(&User{})

	r := gin.Default()

	// Routes
	r.GET("/users", getUsers)
	r.GET("/users/:id", getUserByID)
	r.POST("/users", createUser)
	r.PUT("/users/:id", updateUserByID)
	r.DELETE("/users/:id", deleteUserByID)

	// start the server
	r.Run(":8080")
}
