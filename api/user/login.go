package user

import (
	"encoding/json"
	"errors"
	"time"

	"go-simple-web/container"
	"go-simple-web/util"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// LoginForm
type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login
//
//	@param ctx
func Login(ctx *gin.Context) {
	loginForm := &LoginForm{}
	if err := ctx.ShouldBindJSON(loginForm); err != nil {
		util.Failure(ctx, -1, err.Error())
		return
	}
	users := []User{}

	cfg := container.GetAppConfig()
	if err := json.Unmarshal([]byte(cfg.Users), &users); err != nil {
		util.Failure(ctx, -2, "user list error")
		return
	}
	var (
		pass bool
	)
	for _, item := range users {
		if item.Username == loginForm.Username && item.Password == loginForm.Password {
			pass = true
			break
		}
	}
	if !pass {
		util.Failure(ctx, -3, "user password not right")
		return
	}
	domain := util.GetCookieDomain(ctx, cfg.Domain)
	expireTs := time.Now().Unix() + 3600*24*30
	token, err := generateJwt(loginForm.Username, expireTs)
	if err != nil {
		util.Failure(ctx, -2, err.Error())
		return
	}
	ctx.SetCookie(cfg.LoginKey, token, 3600*24*30, "/", domain, false, false)
	util.Success(ctx, map[string]interface{}{
		"username": loginForm.Username,
		"token":    token,
	})
}

func generateJwt(username string, expireTs int64) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":  username,
		"expire_at": expireTs,
	})
	cfg := container.GetAppConfig()
	return jwtToken.SignedString([]byte(cfg.Salt))
}

type JwtData struct {
	Username string `json:"username"`
	ExpireAt int64  `json:"expire_at"`
}

func parseJwt(jwtToken string) (*JwtData, error) {
	var data *JwtData = &JwtData{}
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		cfg := container.GetAppConfig()
		return []byte(cfg.Salt), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		data.Username = claims["username"].(string)
		data.ExpireAt = int64(claims["expire_at"].(float64))
	} else {
		return nil, errors.New("not valid")
	}
	return data, nil
}
