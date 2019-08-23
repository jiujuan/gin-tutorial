package main

import (
    "errors"
    "fmt"
    jwt "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
    "net/http"
    "time"
)

const (
    ErrorServerBusy = "server is busy"
    ErrorReLogin = "relogin"
)

type JWTClaims struct {
    jwt.StandardClaims
    UserID int `json:"user_id"`
    Password string `json:"password"`
    Username string `json:"username"`
}

var (
    Secret = "123#111"  //salt
    ExpireTime = 36000  //token expire time
)

func main() {
    r := gin.Default()
    r.GET("/login/:username/:password", login)
    r.GET("/verify/:token", verify)
    r.GET("/refresh/:token", refresh)
    r.GET("/sayHello/:token", sayHello)
    _ = r.Run(":8000")
}

//generate jwt token
func genToken(claims *JWTClaims) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    signedToken, err := token.SignedString([]byte(Secret))
    if err != nil {
        return "", errors.New(ErrorServerBusy)
    }
    return signedToken, nil
}

//登录，获取jwt token
func login(c *gin.Context) {
    username := c.Param("username")
    password := c.Param("password")
    claims := &JWTClaims{
        UserID: 1,
        Username: username,
        Password: password,
    }
    claims.IssuedAt = time.Now().Unix()
    claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()
    singedToken, err := genToken(claims)
    if err != nil {
        c.String(http.StatusNotFound, err.Error())
        return
    }
    c.String(http.StatusOK, singedToken)
}

//验证jwt token
func verifyAction(strToken string) (*JWTClaims, error) {
    token, err := jwt.ParseWithClaims(strToken, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(Secret), nil
    })
    if err != nil {
        return nil, errors.New(ErrorServerBusy)
    }

    claims, ok := token.Claims.(*JWTClaims)
    if !ok {
        return nil, errors.New(ErrorReLogin)
    }
    if err := token.Claims.Valid(); err != nil {
        return nil, errors.New(ErrorReLogin)
    }

    fmt.Println("verify")
    return claims, nil
}

func sayHello(c *gin.Context) {
    strToken := c.Param("token")
    claim, err := verifyAction(strToken)
    if err != nil {
        c.String(http.StatusNotFound, err.Error())
    }
    c.String(http.StatusOK, "hello, ", claim.Username)
}

func verify(c *gin.Context) {
    strToken := c.Param("token")
    claim, err := verifyAction(strToken)
    if err != nil {
        c.String(http.StatusNotFound, err.Error())
        return
    }
    c.String(http.StatusOK, "verify: ", claim.Username)
}

func refresh(c *gin.Context) {
    strToken := c.Param("token")
    claims, err := verifyAction(strToken)
    if err != nil {
        c.String(http.StatusNotFound, err.Error())
        return
    }
    claims.ExpiresAt = time.Now().Unix() + (claims.ExpiresAt - claims.IssuedAt)
    signedToken, err := genToken(claims)
    if err != nil {
        c.String(http.StatusNotFound, err.Error())
        return
    }
    c.String(http.StatusOK, signedToken, ", ", claims.ExpiresAt)
}
//进入到 demo1 目录，运行 go run main.go
