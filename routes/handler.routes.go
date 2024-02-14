package routes

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/Thomazoide/go_postsql_crud/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type SessionCheck struct {
	Token string
}

type SessionPayload struct {
	Email    string
	Password string
}

type SvResponse struct {
	Mensaje string
	Cuerpo  any
}

type handler struct {
	DB *gorm.DB
}

type Claims struct {
	User models.Usuario
	jwt.StandardClaims
}

func NewHandler(db *gorm.DB) handler {
	return handler{db}
}

func HashPassword(password string) (string, error) {
	hPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hPass), err
}

func CheckPasswordHash(password string, hsh string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hsh), []byte(password))
	return err == nil
}

func GenerateToken(User models.Usuario) (string, error) {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error al cargar variables de entorno...")
		return "", errEnv
	}
	var jwtKey = []byte(os.Getenv("TKN_SECRET"))
	expirationTime := time.Now().Add(time.Hour * 24)
	claim := &Claims{
		User: User,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, errTkn := token.SignedString(jwtKey)
	return tokenString, errTkn
}

func ValidateToken(tokenString string) (*Claims, error) {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error al cargar las variables de entorno...")
		return nil, errEnv
	}
	var jwtKey = []byte(os.Getenv("TKN_SECRET"))
	token, errTkn := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if errTkn != nil {
		log.Fatal("Error al validar token...")
		return nil, errTkn
	}
	claim, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.New("token invalido")
	}
	if claim.ExpiresAt < time.Now().Unix() {
		return nil, errors.New("token expirado")
	}
	return claim, nil
}
