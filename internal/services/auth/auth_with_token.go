package auth

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"text-to-api/internal/domain"
)

// AuthWithToken authenticates a user by parsing and validating a JWT token.
// It returns an AuthResult containing the client ID and a nil environment if authentication is successful.
//
// Parameters:
//   - ctx: The context for managing request deadlines and cancellation signals.
//   - token: The JWT token used for authentication.
//
// Returns:
//   - A pointer to a domain.AuthResult containing the client ID and a nil environment.
//   - An error if the token is invalid, parsing fails, or any other issue occurs during authentication.
//
// Note: When authenticating with a token, the environment can't be determined from the token itself,
// so it is set to nil in the AuthResult. The environment can be determined by other means, such as
// a middleware reading a specific request header.
func (s *service) AuthWithToken(ctx context.Context, token string) (*domain.AuthResult, error) {

	// Parse and validate the token
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			s.logger.Debug(ctx, "unexpected signing method", "alg", token.Header["alg"])
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.jwtSecret, nil
	})

	// If token is invalid or parsing fails
	if err != nil || !jwtToken.Valid {
		s.logger.Debug(ctx, "invalid token", "error", err)
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		s.logger.Debug(ctx, "invalid token claims")
		return nil, fmt.Errorf("invalid token claims")
	}

	sub, ok := claims["sub"].(string)
	if !ok {
		s.logger.Debug(ctx, "invalid token: subject claim not found")
		return nil, fmt.Errorf("invalid token: subject claim not found")
	}

	authResult := &domain.AuthResult{
		ClientID: sub,
	}

	// When authenticating with token, the environment can't be determined so is set
	// to empty string in the AuthResult. The environment can be determined by other means
	// like a middleware reading a specific request header.
	authResult.Environment = ""

	return authResult, nil
}
