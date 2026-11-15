package utils

import (
	"testing"
	"time"
)

const testSecret = "test-secret-key-for-unit-tests"

func TestGenerateAndValidateToken(t *testing.T) {
	token, err := GenerateToken(42, "user@example.com", "owner", testSecret, time.Hour)
	if err != nil {
		t.Fatalf("не удалось выпустить токен: %v", err)
	}

	claims, err := ValidateToken(token, testSecret)
	if err != nil {
		t.Fatalf("валидный токен не прошёл проверку: %v", err)
	}

	if claims.UserID != 42 {
		t.Errorf("UserID: получили %d, ожидали 42", claims.UserID)
	}
	if claims.Email != "user@example.com" {
		t.Errorf("Email: получили %q, ожидали user@example.com", claims.Email)
	}
	if claims.Role != "owner" {
		t.Errorf("Role: получили %q, ожидали owner", claims.Role)
	}
}

func TestValidateTokenWrongSecret(t *testing.T) {
	token, err := GenerateToken(1, "user@example.com", "tenant", testSecret, time.Hour)
	if err != nil {
		t.Fatalf("не удалось выпустить токен: %v", err)
	}

	if _, err := ValidateToken(token, "another-secret"); err == nil {
		t.Error("токен, подписанный другим ключом, не должен проходить проверку")
	}
}

func TestValidateTokenExpired(t *testing.T) {
	// отрицательный срок жизни: токен истёк в момент выпуска
	token, err := GenerateToken(1, "user@example.com", "tenant", testSecret, -time.Minute)
	if err != nil {
		t.Fatalf("не удалось выпустить токен: %v", err)
	}

	if _, err := ValidateToken(token, testSecret); err == nil {
		t.Error("истёкший токен не должен проходить проверку")
	}
}

func TestValidateTokenGarbage(t *testing.T) {
	if _, err := ValidateToken("не-токен-вовсе", testSecret); err == nil {
		t.Error("произвольная строка не должна проходить проверку")
	}
}
