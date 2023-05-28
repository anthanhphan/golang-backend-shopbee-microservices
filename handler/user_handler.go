package handler

import (
	"errors"
	"net/http"
	"shopbee/model"
	"shopbee/model/req"
	"shopbee/security"

	"context"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	uuid "github.com/google/uuid"
)

func Signup(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.UserCreation

		if err := c.ShouldBind(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"[ERROR]": err.Error(),
			})
			return
		}

		userId, err := uuid.NewUUID()

		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"[ERROR]": err.Error(),
			})
		}

		user.Id = userId.String()
		user.Password = security.HashAndSalt([]byte(user.Password))

		if err := db.Table("USER").Create(&user).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"[ERROR]": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": user})
	}
}

func CheckLogin(db *gorm.DB, context context.Context, loginReq req.ReqSignIn) (model.User, error) {
	var user = model.User{}

	if err := db.Table("USER").Where("EMAIL = ?", loginReq.Email).First(&user).Error; err != nil {
		return user, errors.New("user not found")
	}

	return user, nil
}

func Signin(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := req.ReqSignIn{}

		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"[ERROR]": err.Error(),
			})
			return
		}

		user, err := CheckLogin(db, c.Request.Context(), req)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"[ERROR]": err.Error(),
			})
			return
		}

		isTheSame := security.ComparePasswords(user.Password, []byte(req.Password))

		if !isTheSame {
			c.JSON(http.StatusBadRequest, gin.H{
				"[ERROR]": "Password is not correct!",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": user})
	}
}
