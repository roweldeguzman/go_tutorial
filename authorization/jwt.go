package authorization

import (
	"api/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var secretKey = "ThisIsRealSecretB*tch"

func GenerateJWT(data map[string]interface{}) (string, error) {
	var mySigningKey = []byte(secretKey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["sub"] = data
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		// fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func IsAuthorized(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		w.WriteHeader(401)
		utils.ErrorChecker(0, json.NewEncoder(w).Encode("No Authorization Found"))
		return
	}

	var mySigningKey = []byte(secretKey)

	authHeaderParts := strings.Fields(authHeader)
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		w.WriteHeader(401)
		utils.ErrorChecker(0, json.NewEncoder(w).Encode("Authorization header format must be Bearer {token}"))
		return
	}
	token, err := jwt.Parse(authHeaderParts[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error in parsing")
		}
		return mySigningKey, nil
	})

	if err != nil {
		w.WriteHeader(401)
		utils.ErrorChecker(0, json.NewEncoder(w).Encode("Your Token has been expired"))
		return
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		next(w, r)
		return
	}

	w.WriteHeader(401)
	utils.ErrorChecker(0, json.NewEncoder(w).Encode("Not Authorized"))
}

func GetClaim(r *http.Request) map[string]interface{} {
	var mySigningKey = []byte(secretKey)
	authHeader := r.Header.Get("Authorization")
	authHeaderParts := strings.Fields(authHeader)
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return nil
	}
	token, err := jwt.Parse(authHeaderParts[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error in parsing")
		}
		return mySigningKey, nil
	})

	if err != nil {
		return nil
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims
	}
	return nil
}
