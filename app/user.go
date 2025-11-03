package app

import (
	"net/http"
	"strconv"
	"time"

	"base1-blog/infrastructure/web"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const key string = "metanode"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Email    string
}

type UserInfo struct {
	ID        uint
	Username  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRegister struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email"`
}

type UserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Register 用户注册，密码做bcrypt后保存
func Register(v UserRegister) {
	if Exists(v.Username) {
		panic("username exists")
	}
	user := User{}
	user.Username = v.Username
	hashedID, err := bcrypt.GenerateFromPassword([]byte(v.Password), bcrypt.DefaultCost)
	if err != nil {
		panic("注册失败")
	}
	user.Password = string(hashedID)
	user.Email = v.Email
	db.Create(&user)
}

// Exists 根据用户名查询用户是否存在
func Exists(Username string) bool {
	var count int64
	db.Model(&User{}).Where("username = ?", Username).Count(&count)
	return count > 0
}

// getByUsername 根据用户名查询用户，仅在登录时使用
func getByUsername(Username string) *User {
	user := &User{}
	result := db.Model(&User{}).Where("username = ?", Username).First(user)
	if result.Error != nil {
		panic("未找到用户")
	}
	return user
}

func Get(id uint) *UserInfo {
	user := &UserInfo{}
	result := db.Debug().Model(&User{}).Where("id = ?", id).First(user)
	if result.Error != nil {
		panic("未找到用户")
	}

	return user
}

func Login(v UserLogin) string {
	user := getByUsername(v.Username)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(v.Password))
	if err != nil {
		panic("密码错误")
	}
	claims := jwt.RegisteredClaims{Subject: strconv.Itoa(int(user.ID)), ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour))}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte(key))
	if err != nil {
		panic("认证失败")
	}
	return token
}

// Logined 检查后续接口是否已登录
func Logined() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := web.GetToken(c)
		claims := &jwt.RegisteredClaims{}
		_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(key), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "message": "Unauthorized"})
			return
		}
		id, err := strconv.ParseUint(claims.Subject, 0, 10)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "message": "Unauthorized"})
			return
		}
		c.Set("OperatorId", uint(id))
		c.Next()
	}
}
