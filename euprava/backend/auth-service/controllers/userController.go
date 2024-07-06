package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"backend/database"

	helper "backend/helpers"
	"backend/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()
var SECRET_KEY string = os.Getenv("SECRET_KEY")

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}

	return string(bytes)
}

func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = fmt.Sprintf("login or passowrd is incorrect")
		check = false
	}

	return check, msg
}

func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		l := log.New(gin.DefaultWriter, "User controller: ", log.LstdFlags)
		l.Println(c.GetString("Authorization"))

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while checking for the email"})
			return
		}

		password := HashPassword(*user.Password)
		user.Password = &password

		count, err = userCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while checking for the phone number"})
			return
		}

		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "this email or phone number already exists"})
			return
		}

		user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_id = user.ID.Hex()
		token, refreshToken, _ := helper.GenerateAllTokens(*user.Email, *user.First_name, *user.Last_name, *user.User_type, *&user.User_id)
		user.Token = &token
		user.Refresh_token = &refreshToken

		resultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)
		if insertErr != nil {
			l.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": insertErr.Error()})
			return
		}
		c.JSON(http.StatusOK, resultInsertionNumber)
		defer cancel()
	}
}

// func Register() gin.HandlerFunc {
// 	return func(c *gin.Context) {

// 		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 		var user models.User
// 		l := log.New(gin.DefaultWriter, "User controller: ", log.LstdFlags)
// 		l.Println(c.GetString("Authorization"))

// 		if err := c.BindJSON(&user); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

// 			return
// 		}

// 		validationErr := validate.Struct(user)
// 		if validationErr != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
// 			return
// 		}

// 		count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
// 		defer cancel()
// 		if err != nil {
// 			log.Panic(err)
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for the email"})
// 			return
// 		}

// 		password := HashPassword(*user.Password)
// 		user.Password = &password

// 		count, err = userCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
// 		defer cancel()
// 		if err != nil {
// 			log.Panic(err)
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for the phone number"})
// 			return
// 		}

// 		if count > 0 {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "this email or phone number already exists"})
// 			return
// 		}

// 		user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
// 		user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
// 		user.ID = primitive.NewObjectID()
// 		user.User_id = user.ID.Hex()
// 		token, refreshToken, _ := helper.GenerateAllTokens(*user.Email, *user.First_name, *user.Last_name, *user.User_type, *&user.User_id)
// 		user.Token = &token
// 		user.Refresh_token = &refreshToken

// 		jsonData, err := json.Marshal(user)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal JSON data"})
// 			return
// 		}

// 		resp, err := http.Post("http://university-service:8088/students/create", "application/json", bytes.NewBuffer(jsonData))
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to communicate with the university service"})
// 			l.Println(c.GetString(err.Error()))
// 			return
// 		}
// 		defer resp.Body.Close()

// 		l = log.New(gin.DefaultWriter, "MY STATUS CODE IS: "+strconv.Itoa(resp.StatusCode), log.LstdFlags)
// 		if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save student data in the university service"})
// 			return
// 		}
// 		resultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)
// 		if insertErr != nil {
// 			l.Println(err.Error())
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": insertErr.Error()})
// 			return
// 		}
// 		c.JSON(http.StatusOK, resultInsertionNumber)
// 		defer cancel()
// 	}
// }

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Request headers:", c.Errors)
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		var foundUser models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": ""})
			return
		}

		err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		passwordIsValid, _ := VerifyPassword(*user.Password, *foundUser.Password)
		defer cancel()
		if passwordIsValid != true {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect password"})

			return
		}

		if foundUser.Email == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
			return
		}
		token, refreshToken, _ := helper.GenerateAllTokens(*foundUser.Email, *foundUser.First_name, *foundUser.Last_name, *foundUser.User_type, foundUser.User_id)

		helper.UpdateAllTokens(token, refreshToken, foundUser.User_id)
		err = userCollection.FindOne(ctx, bson.M{"user_id": foundUser.User_id}).Decode(&foundUser)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}

		c.JSON(http.StatusOK, foundUser)

	}
}

func Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		l := log.New(gin.DefaultWriter, "User controller: ", log.LstdFlags)
		refreshToken := helper.ExtractRefreshToken(c)
		l.Println("primljen je: ", c.GetString("Authorization"))
		if refreshToken == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "no refersh token found"})
		} else {

			c.JSON(http.StatusOK, gin.H{"message": "User logged out successfully"})
		}
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("user_id")

		if err := helper.MatchUserTypeToUid(c, userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var user models.User

		err := userCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, user)

	}

}

func GetUsers(c *gin.Context) {
	l := log.New(gin.DefaultWriter, "User Controller ", log.LstdFlags)
	authHeader := c.Request.Header["Authorization"]

	if len(authHeader) == 0 {
		c.JSON(http.StatusUnauthorized, "No header")
		return
	}
	authString := strings.Join(authHeader, "")
	tokenString := strings.Split(authString, "Bearer ")[1]

	if len(tokenString) == 0 {
		c.JSON(http.StatusUnauthorized, "Token empty")
		return
	}

	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		l.Println("Parsing token..")
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		return
	}

	l.Println("Extract the claims from the parsed token")
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		l.Println("Token invalid")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token invalid"})
		return
	}

	parsedToken, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		l.Println("Error decoding token without verification:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "error decoding token"})
		return
	}

	l.Println("Token claims:", parsedToken.Claims)

	l.Println("Retrieving user id..")
	userID, ok := claims["Uid"].(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user ID not found in token"})
		return
	}

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var user models.User
	userCollection.FindOne(ctx, bson.M{"user_id": userID}).Decode(&user)
	defer cancel()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user_id": userID, "user": user})
}
