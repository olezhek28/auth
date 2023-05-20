package auth

import "context"

func (s *service) GetRefreshToken(ctx context.Context, username string, password string) (string, error) {
	// Лезем в кеш

	//Лезем в базу за данными пользователя
	// Проверяем пароль
	// Генерируем токен

	return "", nil
}
