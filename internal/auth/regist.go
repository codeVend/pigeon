package auth

import (
	"fmt"
)

func CheckCookies() {
	
}
type Service struct{}

func (s *Service) Register(user *User) error {
	if user.Username == "" {
		return fmt.Errorf("Имя пустое")
	} else {
		fmt.Printf("Пользователь %s успешно зарегистрирован", user.Username)
		return nil
	}
}
func NewUser(Username string, Number string, Password string, ID int, Email string) {

	

}