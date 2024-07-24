package token

import (
	config "auth-service/config"
	pb "auth-service/generated/auth_service"

	"github.com/golang-jwt/jwt"
	"log"
	"time"
)

func GenerateJWT(user *pb.LoginResponse) (*pb.LoginResponse, error) {
	accessToken := jwt.New(jwt.SigningMethodHS256)
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	claims := accessToken.Claims.(jwt.MapClaims)
	claims["user_id"] = user.UserId
	claims["Name"] = user.Username
	claims["email"] = user.Email
	claims["iat"] = time.Now().Unix()
	claims["ext"] = time.Now().Add(time.Hour).Unix()

	cfg := config.Load()

	access, err := accessToken.SignedString([]byte(cfg.SIGNING_KEY))
	if err != nil {
		return nil, err
	}

	rftClaims := refreshToken.Claims.(jwt.MapClaims)
	rftClaims["user_id"] = user.UserId
	rftClaims["name"] = user.Username
	rftClaims["email"] = user.Email
	rftClaims["iat"] = time.Now().Unix()
	rftClaims["ext"] = time.Now().Add(time.Hour * 24 * 7).Unix()

	refresh, err := refreshToken.SignedString([]byte(cfg.SIGNING_KEY))
	if err != nil {
		log.Fatalf("Access token is not generated %v", err)
	}

	return nil
}
