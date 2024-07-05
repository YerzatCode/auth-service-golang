package controller

import (
	"net/http"
	"time"

	"github.com/YerzatCode/auth-service/internal/config"
	"github.com/YerzatCode/auth-service/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserReqBody struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRes struct {
	ID        uint `gorm:"primaryKey"`
	Username  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (h handler) Signup(c *gin.Context) {
	body := UserReqBody{}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Incorrect data",
		})
		return
	}
	if body.Username == "" || body.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Incorrect data",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to hash password",
		})

		return
	}

	user := model.UserModel{Email: body.Email, Password: string(hash), Username: body.Username}

	if resault := h.DB.Create(&user); resault.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Failde to create user",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"username": user.Username,
		"email":    user.Email,
		"id":       user.ID,
	})
}

func (h handler) Login(c *gin.Context) {
	body := UserLogin{}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Incorrect data",
		})
		return
	}

	var user model.UserModel

	h.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid email or password",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid email or password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.Cfg.Secret))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create token",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func (h handler) Validate(c *gin.Context) {
	user, _ := c.Get("user")

	users := user.(model.UserModel)

	// var data UserLogin

	data := UserRes{
		ID:        users.ID,
		Email:     users.Email,
		Username:  users.Username,
		CreatedAt: users.CreatedAt,
		UpdatedAt: users.UpdatedAt,
		DeletedAt: users.DeletedAt,
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
