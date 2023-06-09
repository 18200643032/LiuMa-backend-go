package handlers

import (
	"LiuMa-backend-go/internal/api/e"
	"LiuMa-backend-go/internal/database"
	"LiuMa-backend-go/internal/models"
	"LiuMa-backend-go/internal/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// 用户注册
func Register(c *gin.Context) {
	var user models.User
	appG := utils.Gin{C: c}
	err := c.BindJSON(&user)
	if err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	var count int64
	database.DB.Model(&models.User{}).Where("account = ?", user.Account).Count(&count)
	if count > 0 {
		appG.Response(http.StatusBadRequest, e.ERROR_EXIST_TAG, nil)
		return
	}
	//密码加密
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "密码服务器错误"})
		return
	}
	// 存储用户信息
	user.Password = string(passwordHash)
	if err = database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"mgs": "注册成功"})

}

// 用户登录
func Login(c *gin.Context) {
	var loginInfo struct {
		Account  string `json:"account" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	err := c.BindJSON(&loginInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误"})
		return
	}

	// 查询用户信息
	var user models.User
	if err = database.DB.Where("account = ?", loginInfo.Account).First(&user).Error; err != nil {
		fmt.Println(err, "----------1")
		c.JSON(http.StatusBadRequest, gin.H{"error": "手机号或密码错误"})
		return
	}

	// 验证密码
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInfo.Password)); err != nil {
		fmt.Println(err, "----------2")
		p1, _ := bcrypt.GenerateFromPassword([]byte("111111"), bcrypt.DefaultCost)
		//e1 := bcrypt.CompareHashAndPassword([]byte("111111"), []byte(p1))
		fmt.Println(p1)

		c.JSON(http.StatusBadRequest, gin.H{"error": "手机号或密码错误"})
		return
	}

	// 生成JWT Token
	token, err := utils.GenerateToken(user.Id, user.Username, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "服务器错误"})
		return
	}

	// 更新Token字段并返回成功提示信息
	if err = database.DB.Model(&user).Update("token", token).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "服务器错误"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "登录成功", "token": token, "data": user})
}

// 用户详情
//func GetUserDetailHandler(c *gin.Context) {
//	idStr := c.Param("userid")
//	userId, _ := strconv.ParseInt(idStr, 10, 64)
//	user := models.User{Id: userId}
//	database.DB.Find(&user)
//	c.JSON(200, gin.H{"users": user})
//}
