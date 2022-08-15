package services

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

// Struct responsável por guardar as informações confidencias
// para a criação de um token JWT
type jwtService struct {
	secretKey string
	issure    string
}

// TODO: usar o .env para informar o secretKey e o issure
func NewJWTService() *jwtService {

	err := godotenv.Load()
	if err != nil {
		return nil
	}

	return &jwtService{
		secretKey: os.Getenv("JWT_SECRET_KEY"),
		issure:    os.Getenv("JWT_ISSURE"),
	}
}

// Struct responsável por guardar as informações em si de quem
// criou o token JWT
type Claim struct {
	Sub string `json:"sub"`
	jwt.StandardClaims
}

// Método que gera um token JWT
func (s *jwtService) GenerateToken(id string) (string, error) {
	claim := &Claim{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    s.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	t, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}

// Função que valida um token JWT passado como parâmetro
func (s *jwtService) ValidateToken(token string) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token: %v", token)
		}

		return []byte(s.secretKey), nil
	})

	return err == nil
}
