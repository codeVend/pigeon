package auth

type User struct {
	Photo string `json:"photo"`
	About string `json:"about"`
	Username string `json:"username"`
	Password string `json:"password"` //Нужно захэшить
	Number string `json:"number"`
	ID int `json:"id"`
	Email string `json:"email"`
}
func NewUser(Username string, Number string, Password string, ID int, Email string) {
}