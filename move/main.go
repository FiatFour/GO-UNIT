package main

import (
	"errors"
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Factorial(n int) (result int) {
	if n == 0 {
		return 1
	}

	if n < 0 {
		return 0
	}

	return n * Factorial(n-1)
}

func Add(a int, b int) int {
	return a + b
}

// ! FIBER
var validate = validator.New()

// User struct with validation tags.
type UserFiber struct {
	Email    string `json:"email" validate:"required,email"`
	Fullname string `json:"fullname" validate:"required,fullname"`
	Age      int    `json:"age" validate:"required,numeric,min=1"`
}

// setup function initializes the Fiber app.
func setupFiber() *fiber.App {
	app := fiber.New()

	// Register the custom validation function for 'fullname'
	validate.RegisterValidation("fullname", validateFullname)

	app.Post("/users", func(c *fiber.Ctx) error {
		user := new(UserFiber)

		if err := c.BodyParser(user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
		}

		if err := validate.Struct(user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		return c.Status(fiber.StatusOK).JSON(user)
	})

	return app
}

// validateFullname checks if the value contains only alphabets and spaces.
func validateFullname(fl validator.FieldLevel) bool {
	return regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString(fl.Field().String())
}

// ! GORM
type UserGorm struct {
	gorm.Model
	Fullname string
	Email    string `gorm:"unique"`
	Age      int
}

// InitializeDB initializes the database and automigrates the User model.
func InitializeDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&UserGorm{})
	return db
}

// AddUser adds a new user to the database.
func AddUser(db *gorm.DB, fullname, email string, age int) error {
	user := UserGorm{Fullname: fullname, Email: email, Age: age}

	// Check if email already exists
	var count int64
	db.Model(&UserGorm{}).Where("email = ?", email).Count(&count)
	if count > 0 {
		return errors.New("email already exists")
	}

	// Save the new user
	result := db.Create(&user)
	return result.Error
}

func main() {
	//* Fiber
	// app := setupFiber()
	// app.Listen(":8000")

	//! GORM
	// db := InitializeDB()
	// Your application code
	// err := AddUser(db, "John Doe", "jane.doe@example.com", 30)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
}
