package main

import (
	"fmt"
	"net/http"
	"time"
	"context"
	"encoding/json"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	log "github.com/Sirupsen/logrus"
)

//type User struct {
//	Id       string `json:"id"`
//	Username string `json:"username"`
//	Password string `json:"password"`
//}
var TimeFunc = time.Now
var jwtSecret []byte = []byte("thepolyglotdeveloper12312")

func GenerateToken(user User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * time.Duration(24)).Unix(), //This token will live for 24 hours
		"iat": time.Now().Unix(),
		"sub": user.ID,
	})
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateJWT(t string) (*jwt.Token, error) {
	if t == "" {
		return nil, errors.New("Authorization token must be present")
	}
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		// dont't forget to validate the alg is want you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})
	if err == nil && token.Valid{
		return token, nil
	}else{
		return nil, errors.New("Invalid authorization token")
	}
}

func authToken(h http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		ctx := r.Context()
		token := r.Header.Get("token")
		jwtToken, err := ValidateJWT(token)
		if err != nil{
			log.Println(err)
			data := make(map[string]interface{})
			data["msg"] = "token error"
			responseJSON, err := json.Marshal(data)
			if err != nil{
				log.Println(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(responseJSON)
			return
		}
		h.ServeHTTP(w, r.WithContext(context.WithValue(ctx, "jwt", jwtToken)))
	})
}


func CreateTokenEndpoint(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST"{
		mail := r.FormValue("mail")
		password := r.FormValue("password")
		for _, user := range users{
			if user.Mail == mail{
				if user.Password == password{
					expireAt := TimeFunc().Unix() + 10
					now := TimeFunc().Unix()
					// validate user from database here and add unique mark for token
					// also set expiration time
					token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
						"exp": expireAt,
						"iat": now,
						"sub": user.ID,
					})
					tokenString, err := token.SignedString(jwtSecret)
					if err != nil {
						fmt.Println(err)
						w.Header().Set("content-type", "application/json")
						w.Write([]byte(`{ "error": "ERROR WHEN CREATE JWT TOKEN" }`))
						return
					}
					w.Header().Set("content-type", "application/json")
					w.Write([]byte(`{ "token": "` + tokenString + `" }`))
					return
				} else{
					w.Header().Set("content-type", "application/json")
					w.Write([]byte(`{ "error": "password is not correct" }`))
					return
				}
			}
		}
		w.Header().Set("content-type", "application/json")
		w.Write([]byte(`{ "error": "user not found" }`))
		return
	} else{
		w.Header().Set("content-type", "application/json")
		w.Write([]byte(`{ "error": "request method no allowed" }`))
		return
	}
}