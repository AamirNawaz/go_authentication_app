package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"go_authentication_app/db"
	"go_authentication_app/models"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strconv"
	"time"
)

var JwtSecreteKey = []byte("myjwtsecret")

func Users(c *fiber.Ctx) error {
	return c.SendString("Get all users")
}

func RegisterUser(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		log.Fatalf("registeration error in post request %v", err)
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
		Status:   data["status"],
	}

	db.DB.Create(&user)

	return (c.JSON(user))
}
func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User
	db.DB.Where("email=?", data["email"]).First(&user)

	//check if email exist
	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"status":  "error",
			"Message": "Email not found",
		})
	}

	//if email exist then match password
	err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"]))
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "Incorrect password",
		})
	}

	//generate jwt token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
	})

	tokenString, err := claims.SignedString([]byte(JwtSecreteKey))

	//fmt.Println("TOKEN:", tokenString)

	if err != nil {
		c.Status(fiber.StatusForbidden)
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "Token Expired or invalid",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})

}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "successfully logout",
	})
}

func GetUser(c *fiber.Ctx) error {
	jwtFromCookie := c.Cookies("jwt")
	//return c.JSON(jwtFromCookie)

	token, err := jwt.ParseWithClaims(jwtFromCookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtSecreteKey), nil
	})

	//fmt.Println("parse token:::", token)
	if err != nil {
		c.Status(fiber.StatusForbidden)
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "UnAuthenticated user",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	//fmt.Println("claims::::", claims)
	var user models.User

	db.DB.Where("id = ?", claims.Issuer).First(&user)
	return c.JSON(user)
}
