package hash

import "golang.org/x/crypto/bcrypt"

// HashPassword is for hasing password into bcrypt hashed string
func HashPassword(password string) (hashed string) {
	crypt, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		hashed = ""
	} else {
		hashed = string(crypt)
	}

	return
}

// ComparePassword is for compare hashed password input and with the password stored in database
func ComparePassword(passwordInput, passwordDB string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(passwordDB), []byte(passwordInput))
	if err != nil {
		return
	}

	return
}
