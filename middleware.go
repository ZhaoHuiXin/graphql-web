package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
)

// logging middleware
func logged(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now().UTC()
		next.ServeHTTP(w, r)
		log.WithFields(log.Fields{
			"path":    r.RequestURI,
			"IP":      r.RemoteAddr,
			"elapsed": time.Now().UTC().Sub(start),
		}).Info()
	})
}
// verify authorization middleware
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
