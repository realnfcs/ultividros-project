package middlewares

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/realnfcs/ultividros-project/api/infra/services"
)

// Middleware do fiber para validação do token JWT (por meio de fluxo de cookies)
func JWTAuthCookie() func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {

		token := ctx.Cookies("jwt_token")
		if token == "" {
			return ctx.Status(fiber.StatusNetworkAuthenticationRequired).JSON(fiber.Map{
				"error": "JWT token don't exist",
			})
		}

		if !services.NewJWTService().ValidateToken(token) {
			return ctx.Status(fiber.StatusNetworkAuthenticationRequired).JSON(fiber.Map{
				"error": "invalid token",
			})
		}

		godotenv.Load()

		claims := jwt.MapClaims{}
		t, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})

		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err,
			})
		}

		claim := t.Claims.(jwt.MapClaims)

		ctx.Locals("id", claim["sub"])

		return ctx.Next()
	}
}
