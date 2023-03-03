package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`

	Notes []Note `json:"notes" gorm:"foreignKey:UID"`
}

type Note struct {
	gorm.Model
	Title string `json:"title"`
	Body  string `json:"body"`

	UID uint `json:"uid"`
}

type Claims struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`

	jwt.StandardClaims
}

func HashStr(str string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(str), 14)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func ValidateHash(clear string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(clear))
	return err == nil
}

func CreateUser(db *gorm.DB, user User) error {
	res := db.Create(&user)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func AuthUser(db *gorm.DB, c *gin.Context) (User, Claims, error) {
	cookie, err := c.Cookie("jwt")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no cookie",
		})
		return User{}, Claims{}, err
	}

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(cookie, claims,
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
	if err != nil || !token.Valid {
		log.Println("Invalid token.")
		c.AbortWithStatus(http.StatusUnauthorized)
		return User{}, Claims{}, err
	}

	var user User
	res := db.First(&user, claims)

	if res.Error != nil {
		log.Println(res.Error)
		c.AbortWithStatus(http.StatusInternalServerError)
		return User{}, Claims{}, res.Error
	}

	return user, *claims, nil
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	db, err := gorm.Open(sqlite.Open("./test.sqlite"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	err = db.AutoMigrate(&User{}, &Note{})
	if err != nil {
		log.Fatalln(err)
	}

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		AllowWildcard:    true,
	}))

	api := r.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		api.POST("/note/create", func(c *gin.Context) {
			user, _, err := AuthUser(db, c)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": err.Error(),
				})
				return
			}

			var note Note
			err = c.BindJSON(&note)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}

			err = db.Model(&user).Association("Notes").Append(&note)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "note created!",
			})
		})
		api.POST("/user/create", func(c *gin.Context) {
			var body struct {
				Name     string `json:"name"`
				Email    string `json:"email"`
				Password string `json:"password"`
			}

			err := c.ShouldBindJSON(&body)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}

			user := User{
				Name:     body.Name,
				Email:    body.Email,
				Password: HashStr(body.Password),
			}

			err = CreateUser(db, user)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "user created!",
			})
		})
		api.POST("/user/login", func(c *gin.Context) {
			var body struct {
				Email    string `json:"email"`
				Password string `json:"password"`
			}

			err := c.ShouldBindJSON(&body)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}

			var user User
			db.Where("email = ?", body.Email).First(&user)

			if !ValidateHash(body.Password, user.Password) {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": "invalid credentials",
				})
				return
			}

			claims := &Claims{
				Name:  user.Name,
				Email: user.Email,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
				},
			}

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenStr, err := token.SignedString(jwtKey)

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			}

			c.Header("Content-Type", "application/json")
			c.SetCookie(
				"jwt",
				tokenStr,
				3600,
				"/",
				"localhost",
				false,
				true)
			c.JSON(http.StatusOK, gin.H{
				"message": "success",
			})
		})
		api.GET("/user/notes", func(c *gin.Context) {
			user, _, err := AuthUser(db, c)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": err.Error(),
				})
				return
			}

			var notes []Note
			err = db.Model(&user).Association("Notes").Find(&notes)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"notes": notes,
			})
		})
	}

	err = r.Run(":7100")
	if err != nil {
		log.Fatalln(err)
	}
}
