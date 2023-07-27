package controllers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"learn-middleware-example/app/models"
	"learn-middleware-example/app/services"
	"time"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(c *gin.Context) {

	var userInput User
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.AbortWithStatusJSON(500, gin.H{"Message": "Wrong JSON Format >:("})
		return
	}

	//checking user from database
	var dbUser User
	if err := models.DB.Where("username = ?", userInput.Username).First(&dbUser).Error; err != nil {
		c.AbortWithStatusJSON(401, gin.H{"Message": "Username Not Found"})
		return
	}

	//if found,do compare password
	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(userInput.Password)); err != nil {
		c.AbortWithStatusJSON(401, gin.H{"Message": "Wrong Password"})
		return
	}

	//if password correct,do implement jwt and session

	//sessions
	session, _ := services.GetSession(c.Request, userInput.Username)
	session.Values["username"] = dbUser.Username
	session.Values["exp"] = time.Now().Add(2 * time.Minute).Unix()

	//save session
	session.Save(c.Request, c.Writer)

	//jwt
	token, _ := services.GenerateJWT(userInput.Username)

	c.JSON(200, gin.H{"Successfully Login": token})

}

func RegisterHandler(c *gin.Context) {
	var userInput User

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.AbortWithStatusJSON(500, err.Error())
		return
	}

	//generate password
	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), 10)

	//convert to stringggg
	userInput.Password = string(hashedPass)

	//save to database
	if err := models.DB.Create(&userInput).Error; err != nil {
		c.AbortWithStatusJSON(500, err.Error())
		return
	}
}

type ChangePasswordRequest struct {
	Username    string `json:"username"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

func ChangePassword(c *gin.Context) {
	var changePass ChangePasswordRequest

	//baca req body json
	if err := c.ShouldBindJSON(&changePass); err != nil {
		c.AbortWithStatusJSON(500, err.Error())
		return
	}

	//checking username
	var user User
	if err := models.DB.Where("username = ?", changePass.Username).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(500, gin.H{"Message": "Username Not Found"})
		return
	}

	//compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(changePass.OldPassword)); err != nil {
		c.AbortWithStatusJSON(500, gin.H{"Message": "Wrong Password"})
		return
	}

	//generate password baru untuk new password
	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(changePass.NewPassword), 10)

	//update password
	if err := models.DB.Model(&user).Where("username = ?", changePass.Username).Update("password", string(hashedPass)).Error; err != nil {
		c.AbortWithStatusJSON(500, err.Error())
		return
	}

	c.JSON(200, gin.H{"Message": "Successfully change password"})
	return
}

func WelcomeHandler(c *gin.Context) {
	//show username when trying to log in
	username := c.GetString("username")
	c.JSON(200, gin.H{"Message": "Welcome " + username})
	return
}
