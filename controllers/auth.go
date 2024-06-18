package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/vikram761/backend/models"
	"golang.org/x/crypto/bcrypt"
)

type authController struct {
	Db *sql.DB
}

type AuthController interface {
	Register(*gin.Context)
	Login(*gin.Context)
	CheckAuth(*gin.Context)
	ValidateUser(*gin.Context)
    Validate(*gin.Context)
}

func NewAuthController(db *sql.DB) AuthController {
	return &authController{Db: db}
}

func (a *authController) Register(ctx *gin.Context) {
	var user models.User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{
			"status":   "failed",
			"response": err.Error(),
		})
		return
	}

	if err := user.ValidateUser(); err != nil {
		ctx.JSON(400, gin.H{
			"status":   "failed",
			"response": err.Error(),
		})
		return
	}

	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(400, gin.H{
			"status":   "failed",
			"response": err.Error(),
		})
		return
	}
	hashedPassword := string(hashedPasswordBytes)
	_, err = a.Db.Exec("INSERT INTO USERS(NAME, EMAIL, PASSWORD, ROLE) VALUES($1, $2, $3, $4)", user.Name, user.Email, hashedPassword, strings.ToUpper(user.Role))

	if err != nil {
		ctx.JSON(400, gin.H{
			"status":   "failed",
			"response": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":   "success",
		"response": "User added successfully",
	})
}

func (a *authController) Login(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{
			"status":   "failed",
			"response": err.Error(),
		})
	}

	query := a.Db.QueryRow("SELECT * FROM USERS WHERE EMAIL = $1", user.Email)
	var actualUser models.User
	if err := query.Scan(&actualUser.Id, &actualUser.Name, &actualUser.Email, &actualUser.Password, &actualUser.Role); err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(400, gin.H{
				"status":   "failed",
				"response": fmt.Sprintf("The user with email %s not found.", user.Email),
			})
			return
		}
		ctx.JSON(400, gin.H{
			"status":   "failed",
			"response": err.Error(),
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(actualUser.Password), []byte(user.Password)); err != nil {
		ctx.JSON(400, gin.H{
			"status":   "failed",
			"response": "Password didn't match",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   actualUser.Id,
		"exp":  time.Now().Add(time.Hour * 24 * 30).Unix(),
		"role": actualUser.Role,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		ctx.JSON(400, gin.H{
			"status":   "failed",
			"response": "Error during token generation.",
		})
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", tokenString, 3600*24*30, "/", "", true, true)
	ctx.JSON(200, gin.H{
		"status":   "success",
		"response": "User logged in successfully",
	})
}

func (a *authController) CheckAuth(ctx *gin.Context) {
	tokenString, err := ctx.Cookie("Authorization")

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":   "failed",
			"response": "Authorization cookie not found.",
		})
		return
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			ctx.AbortWithStatusJSON(400, gin.H{
				"status":   "failed",
				"response": "Jwt token is expired",
			})
			return
		}
		query := a.Db.QueryRow("SELECT * FROM USERS WHERE ID = $1", claims["id"])

		var user models.User
		err := query.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Role)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status":   "failed",
				"response": err.Error(),
			})
			return
		}

		ctx.Set("user", user)
		ctx.Next()

	} else {
		ctx.AbortWithStatusJSON(400, gin.H{
			"status":   "failed",
			"response": "Error with jwt claims",
		})
		return
	}

}

func (a *authController) Validate(ctx *gin.Context) {
	tokenString, err := ctx.Cookie("Authorization")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":   "failed",
			"response": "Authorization cookie not found.",
		})
		return
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			ctx.AbortWithStatusJSON(400, gin.H{
				"status":   "failed",
				"response": "Jwt token is expired",
			})
			return
		}
        ctx.JSON(200 , gin.H {
          "status"  : "success",
          "response" : "Validated successfully",
        })	

	} else {
		ctx.AbortWithStatusJSON(400, gin.H{
			"status":   "failed",
			"response": "Error with jwt claims",
		})
		return
	}
}

func (a *authController) ValidateUser(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(200, gin.H{
		"status":   "success",
		"response": user,
	})
}
