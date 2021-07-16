package jwt

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

// 用户名密码
var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

// 接收http请求的结构体
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// 创建将被编码为JWT的结构。
// 我们将jwt.StandardClaims作为嵌入式类型，以提供到期时间等字段。
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

//密钥
var jwtKey = []byte("my_secret_key")

func Demo1() {
	http.HandleFunc("/signin", Signin)
	http.HandleFunc("/welcome", Welcome)
	http.HandleFunc("/refresh", Refresh)
	// 在8000端口启动服务
	http.ListenAndServe(":8000", nil)

}

//签名
func Signin(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		// 如果主体结构错误，则返回HTTP错误
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// 如果设置的用户密码与我们收到的密码相同，那么我们可以继续。
	// 如果不是，则返回“未经授权”状态。
	expectedPassword, ok := users[creds.Username]
	if !ok || expectedPassword != creds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	//接下去进行签名
	//设置令牌的过期时间
	expirationTime := time.Now().Add(5 * time.Minute)
	// 创建JWT声明，其中包括用户名和有效时间
	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// 使用用于签名的算法和令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 最终生成的签名
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// 如果创建JWT时出错，则返回内部服务器错误
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// 最后，我们将客户端cookie token设置为刚刚生成的JWT
	// 我们还设置了与令牌本身相同的cookie到期时间
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			// 如果未设置cookie，则返回未授权状态
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// 对于其他类型的错误，返回错误的请求状态。
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// 获取jwt令牌
	jwtStr := c.Value
	// 初始化`Claims`实例
	claims := &Claims{}
	// 解析JWT字符串并将结果存储在`claims`中。
	// 请注意，我们也在此方法中传递了密钥。
	// 如果令牌无效（如果令牌已根据我们设置的登录到期时间过期）或者签名不匹配,此方法会返回错误.
	tkn, err := jwt.ParseWithClaims(jwtStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// 最后，将欢迎消息以及令牌中的用户名返回给用户
	w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Username)))
}


//续签令牌
func Refresh(w http.ResponseWriter, r *http.Request) {
	//==========================================================
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tknStr := c.Value
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	//==================================这段与Welcome一样，都是用来校验原来的令牌的有效性


	//判断令牌的过期时间还有多久，只有在令牌只剩30s时才允许续签
	if time.Unix(claims.ExpiresAt,0).Sub(time.Now()) > 30 * time.Second {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// 现在，为当前用户创建一个新令牌，并延长其到期时间
	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 查看用户新的`token` cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}