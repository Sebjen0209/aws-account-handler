package middleware

//extracting the request headers
//extracting our claims
//valdating everything

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/golang-jwt/jwt/v5"
)

//inspiration from: https://medium.com/@cheickzida/golang-implementing-jwt-token-authentication-bba9bfd84d60

func ValidateJWTMiddleware(next func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)) func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		// extract the headers from our token
		tokenString := extractTokenFromHeaders(request.Headers)
		if tokenString == "" {
			return events.APIGatewayProxyResponse{
				Body:       "Missing Auth token",
				StatusCode: http.StatusUnauthorized,
			}, nil
		}
		// parse the token for our claims

		claims, err := parseToken(tokenString)
		if err != nil {
			return events.APIGatewayProxyResponse{
				Body:       "User Unauthorized",
				StatusCode: http.StatusUnauthorized,
			}, err
		}

		// did this token expire
		expires := int64(claims["expires"].(float64))
		if time.Now().Unix() > expires {
			return events.APIGatewayProxyResponse{
				Body:       "token expired",
				StatusCode: http.StatusUnauthorized,
			}, nil
		}

		return next(request)
	}

}

func extractTokenFromHeaders(headers map[string]string) string {
	authHeader, ok := headers["Authorization"]
	if !ok {
		return ""
	}

	splitToken := strings.Split(authHeader, "Bearer ")
	if len(splitToken) != 2 {
		return ""
	}

	return splitToken[1]
}

func parseToken(tokenString string) (jwt.MapClaims, error) {
	fmt.Printf("Parsing token: %s\n", tokenString) // Log the token

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		secret := "secret"
		return []byte(secret), nil
	})

	if err != nil {
		fmt.Printf("Error parsing token: %v\n", err)
		return nil, fmt.Errorf("unauthorized: %v", err)
	}

	if !token.Valid {
		fmt.Println("Token is not valid")
		return nil, fmt.Errorf("token is not valid - unauthorized")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("Token claims are not of type MapClaims")
		return nil, fmt.Errorf("token of unrecognized type - unauthorized")
	}

	fmt.Printf("Token claims: %v\n", claims) // Log the claims
	return claims, nil
}
