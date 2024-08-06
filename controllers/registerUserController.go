package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go-mvc/initializers"
	"go-mvc/models"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strconv"
	"time"
)

func Register(c *fiber.Ctx) error {
	// first we need to get the data from the post request
	// variable to get that data

	var postData map[string]string
	// pass the data as a reference // this returns an error for us to handle -- using a shorthand of the if statement err handling
	if err := c.BodyParser(&postData); err != nil {
		return err
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(postData["password"]), 14)
	// getting the data from the postData to insert into the database
	user := models.User{
		Name:     postData["name"],
		Email:    postData["Email"],
		Password: password,
		Role:     postData["Role"],
		Active:   postData["Active"],
	}
	// inserting the user into the database
	initializers.DB.Create(&user)
	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	SecretKey := []byte(os.Getenv("SecretKey"))

	var postData map[string]string
	// pass the data as a reference // this returns an error for us to handle -- using a shorthand of the if statement err handling
	if err := c.BodyParser(&postData); err != nil {
		return err

	} // next get the email and password of the user
	var user models.User
	initializers.DB.Where("email = ?", postData["Email"]).First(&user)

	if user.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"Message": "User Not Found",
		})
	} // not check the password -- case matters with this ... new naming convention just dropped all lowercase
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(postData["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"Message": "Incorrect Email and/or Password",
		})
	}

	//claims := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
	//	Issuer:    strconv.Itoa(int(user.ID)),                         // assign the token to the user
	//	ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // one day token
	//})
	stringUserid := strconv.Itoa(int(user.ID))

	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // one day token
		ID:        stringUserid,
	}

	//token, err := claims.SignedString([]byte(SecretKey))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// ss is the assigned token variable
	ss, err := token.SignedString(SecretKey)

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"Message": "Could not login",
		})
	}

	// create the cookie to store the JWT
	cookie := fiber.Cookie{
		Name:     "JWT",
		Value:    ss,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
	} // cookie is set up to only be readable by the frontend i believe

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"Message": "Success",
	})
}

func User(c *fiber.Ctx) error {
	SecretKey := []byte(os.Getenv("SecretKey"))

	//get the cookie
	cookie := c.Cookies("JWT")

	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"Message": "Unauthenticated",
		})
	}

	// get the claims from the cookie -- converting from type claims to type registered-claims
	claims := token.Claims.(*jwt.RegisteredClaims)

	var user models.User

	initializers.DB.Where("id = ?", claims.ID).First(&user)

	return c.JSON(user)
}
