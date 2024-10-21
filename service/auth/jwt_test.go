package auth

import "testing"

func TestCreateJwt(t *testing.T) {
	secret := []byte("secret")

	token, err := CreateJwt(secret, 1)
	if err != nil {
		t.Errorf("Error creating JWT: %v", err)
	}

	if token == "" {
		t.Errorf("Expected token to not be empty")
	}
}
