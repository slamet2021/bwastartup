package main

import (
	"bwastartup/handler"
	"bwastartup/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	//fmt.Println("Connection to database is success")

	//var users []user.User
	//db.Find(&users)

	//for _, user := range users {
	//	fmt.Println(user.Name)
	//	fmt.Println(user.Email)
	//	fmt.Println(user.Occupation)
	//	fmt.Println("=================")
	//}
	/* Membuat router untuk mengakses json api */
	//	router := gin.Default()
	//	router.GET("/handler", handler)
	//	router.Run() */

	// Repository, save struct User ke db
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	//	input := user.LoginInput{
	//		Email:    "agung@email.com",
	//		Password: "password,
	//	}

	//	user, err := userService.Login(input)
	//	if err != nil {
	//		fmt.Println("Terjadi kesalahan")
	//		fmt.Println(err.Error())
	//	}

	//	fmt.Println(user.Email)
	//	fmt.Println(user.Name)

	// Pengecekan login
	//	userByEmail, err := userRepository.FindByEmail("agung@email.com")
	//	if err != nil {
	//		fmt.Println(err.Error())
	//	}

	//	if userByEmail.ID == 0 {
	//		fmt.Println("User tidak ditemukan")
	//	} else {
	//		fmt.Println(userByEmail.Name)
	//	}

	// Route
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)

	router.Run()

	// test register user
	//userInput := user.RegisterUserInput{}
	//userInput.Name = "Tes simpan dari service"
	//userInput.Email = "contoh@gmail.com"
	//userInput.Occupation = "Petani"
	//userInput.Password = "password"

	//userService.RegisterUser(userInput)
}

// handler gin untuk API u/ test json
//func handler(c *gin.Context) {
//	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

///	if err != nil {
//		log.Fatal(err.Error())
//	}

//	var users []user.User
//	db.Find(&users)

//	c.JSON(http.StatusOK, users)

// input dari user
// handler, mapping input ke struct
// service, mapping ke struct User
// db
//}
