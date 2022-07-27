package middlewares

import (
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/joho/godotenv"
)

func JWTAuth() func(ctx *fiber.Ctx) error {
	err := godotenv.Load()
	if err != nil {
		return func(ctx *fiber.Ctx) error {
			return ctx.Status(fiber.StatusNetworkAuthenticationRequired).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}

	jwtSecret := os.Getenv("JWT_SECRET_KEY")

	return jwtware.New(jwtware.Config{
		SigningKey: []byte(jwtSecret),
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(fiber.StatusNetworkAuthenticationRequired).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
	})
}
