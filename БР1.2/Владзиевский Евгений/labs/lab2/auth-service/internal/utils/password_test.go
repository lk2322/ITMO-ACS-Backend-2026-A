package utils

import "testing"

func TestHashAndCheckPassword(t *testing.T) {
	const password = "Passw0rd!"

	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("хеширование не удалось: %v", err)
	}

	if hash == password {
		t.Error("хеш совпадает с исходным паролем — пароль не захеширован")
	}
	if !CheckPasswordHash(password, hash) {
		t.Error("верный пароль не прошёл проверку")
	}
	if CheckPasswordHash("wrong-password", hash) {
		t.Error("неверный пароль прошёл проверку")
	}
}

func TestHashPasswordIsSalted(t *testing.T) {
	const password = "Passw0rd!"

	first, err := HashPassword(password)
	if err != nil {
		t.Fatalf("хеширование не удалось: %v", err)
	}
	second, err := HashPassword(password)
	if err != nil {
		t.Fatalf("хеширование не удалось: %v", err)
	}

	// bcrypt добавляет случайную соль: два хеша одного пароля различаются,
	// но оба проходят проверку
	if first == second {
		t.Error("хеши одного пароля совпали — соль не применяется")
	}
	if !CheckPasswordHash(password, second) {
		t.Error("второй хеш не проходит проверку")
	}
}
