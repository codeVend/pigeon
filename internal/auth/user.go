package auth

type User struct {
	Photo string
	About string
	Username string
	Password string //Нужно захэшить
	Number string
	ID int
	Email string
}