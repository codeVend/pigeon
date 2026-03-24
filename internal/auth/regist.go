package auth

import (
	"fmt"
)

func CheckCookies() {
	
}

func NewUser(Username string, Number string, Password string) {

	type Service struct{}

	func (s *Service) Register(user *User); error {
		if user.Username == "" {
			return fmt.Errorf("Имя пустое")
		} else {
			fmt.Printf("Пользователь %s успешно зарегистрирован", user.Username)
			return nil
		}
	}

}