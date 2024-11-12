package route

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("pescms-rent")

func CreateJwt(username string, userinfo string) (string, error) {
	auth := jwt.MapClaims{}
	auth["exp"] = time.Now().Add(time.Hour * 30).Unix() // JWT的过期时间
	auth["iat"] = time.Now().Unix()                     // 表示JWT的签发时间
	auth["nbf"] = time.Now().Unix()                     // 表示JWT在何时之前不可被接受处理的时间。这可以用于防止JWT在预定的生效时间之前被使用。
	auth["iss"] = "pescms-rent"                         // JWT的实体或服务的标识
	auth["sub"] = "pescms-rent"                         // JWT所面向的用户或实体的标识
	auth["username"] = username
	auth["userinfo"] = userinfo
	// 创建 token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, auth)

	// 签名 token
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// fmt.Println("Signed Token:", signedToken)

	return signedToken, nil
}

func ValidateJwt(signedToken string) (jwt.MapClaims, error) {
	parsedToken, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		// 确认签名算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// 返回签名密钥
		return secretKey, nil
	})
	if err != nil {
		return jwt.MapClaims{}, err
	}

	// 确认 token 有效
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		// fmt.Println(claims["exp"])
		// fmt.Printf("%T %v", claims["iat"])

		return claims, err
	} else {
		return jwt.MapClaims{}, fmt.Errorf("invalid token")
	}
}
