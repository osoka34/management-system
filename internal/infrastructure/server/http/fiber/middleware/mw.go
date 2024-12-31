package middleware

import (
	"errors"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"

	"service/pkg/utils"
)

const (
	AccessTokenName  = "Authorization"
	RefreshTokenName = "X-Refresh-Token"
)

func Authorization(c fiber.Ctx) error {
	accessToken := c.Get(AccessTokenName)
	refreshToken := c.Get(RefreshTokenName)

	if accessToken == "" || refreshToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "Unauthorized",
			"details": "Access token or refresh token is missing",
		})
	}

	claimsA, err := utils.ValidateAccessToken(accessToken)
	if err != nil && !errors.Is(err, jwt.ErrTokenExpired) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	claimsR, err := utils.ValidateRefreshToken(refreshToken)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{})
	}

	if claimsA.UID != claimsR.UID || claimsA.Username != claimsR.Username {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "Unauthorized",
			"details": "Access token and refresh token do not match",
		})
	}

	c.Locals(utils.KeyAccessClaims{}, claimsA)
	c.Locals(utils.KeyRefreshClaims{}, claimsR)

	return c.Next()
}
