package types

type RegisterUserPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	FitrstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}