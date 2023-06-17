package helpers

import (
	"context"
	"errors"
	"fmt"
	"github.com/bostigger/restaurant-management-api/database"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	options2 "go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

type signedDetails struct {
	Email          string
	Username       string
	UserId         string
	UserType       string
	StandardClaims jwt.StandardClaims
}

func (s *signedDetails) Valid() error {
	if time.Now().Unix() > s.StandardClaims.ExpiresAt {
		return errors.New("token has expired")
	}
	return nil
}

var userCollection *mongo.Collection = database.GetCollection(database.Client, "users")
var SecretKey string = os.Getenv("SECRET_KEY")

func GenerateAllToken(email string, username string, userId string, userType string) (signedToken string, signedRefreshToken string, err error) {
	claims := &signedDetails{
		Email:    email,
		Username: username,
		UserId:   userId,
		UserType: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}
	refreshClaims := &signedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration((168))).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SecretKey))
	if err != nil {
		return "", "", err
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SecretKey))

	if err != nil {
		return "", "", err
	}
	return token, refreshToken, nil

}

func UpdateAllTokens(signedToken string, signedRefreshToken string, userId string) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var updateObj primitive.D
	updateObj = append(updateObj, bson.E{"token", signedToken})
	updateObj = append(updateObj, bson.E{"refreshToken", signedRefreshToken})
	updatedAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateObj = append(updateObj, bson.E{"updatedAt", updatedAt})

	upsert := true
	options := options2.UpdateOptions{
		Upsert: &upsert,
	}
	filter := bson.M{"userId": userId}
	_, err := userCollection.UpdateOne(ctx, filter, bson.D{{"$set", updateObj}}, &options)
	if err != nil {
		log.Fatal(err)
		return
	}
	return

}

func ValidateToken(clientToken string) (claims *signedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		clientToken,
		&signedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})
	if err != nil {
		msg = err.Error()
		return
	}
	claims, ok := token.Claims.(*signedDetails)
	if !ok {
		msg = fmt.Sprintf("token is invalid")
		return
	}
	return claims, msg
}
