package io_models

import "github.com/golang-jwt/jwt"

type Claims struct {
	Username string `json:"username"`
	XXLG string `json:"xxlg"` // device uuid
	jwt.StandardClaims
}
