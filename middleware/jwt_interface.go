package middleware

type JwtInterface interface {
	GenerateJWT(userID int, name string) (string, error)
}
