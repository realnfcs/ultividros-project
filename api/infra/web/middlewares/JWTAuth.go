package middlewares

import (
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

// Middleware do fiber responsável pela validação do token JWT
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

// Middleware do fiber responsável por pegar os dados do token JWT
// e passar para o próximo handler
// Para pegar os dados, usa-se ctx.Locals("<campo do token>")
func JWTdata() func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		user := ctx.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		ctx.Locals("id", claims["sub"])
		return ctx.Next()
	}
}
