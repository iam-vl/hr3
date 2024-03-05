package types

import (
	"fmt"
	"regexp"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const (
	BCRYPT_COST   = 12
	MIN_FNAME_LEN = 2
	MIN_LNAME_LEN = 2
	MIN_PWD_LEN   = 7
)

type UserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func (params UserParams) Validate() map[string]string {
	errors := map[string]string{}
	if len(params.FirstName) < MIN_FNAME_LEN {
		errors["firstName"] = fmt.Sprintf("First name must be at least %d characters", MIN_FNAME_LEN)
	}
	if len(params.LastName) < MIN_FNAME_LEN {
		errors["lastName"] = fmt.Sprintf("Last name must be at least %d characters", MIN_LNAME_LEN)
	}
	if len(params.Password) < MIN_FNAME_LEN {
		errors["password"] = fmt.Sprintf("Password must be at least %d characters", MIN_PWD_LEN)
	}
	if !isEmailValid(params.Email) {
		errors["email"] = fmt.Sprintf("Please enter a correct email address")
	}
	return errors
}

type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName         string             `bson:"firstName" json:"firstName"`
	LastName          string             `bson:"lastName" json:"lastName"`
	Email             string             `bson:"email" json:"email"`
	EncryptedPassword string             `bson:"encryptedPassword" json:"-"`
}

func NewUserFromParams(params UserParams) (*User, error) {
	encPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), BCRYPT_COST)
	if err != nil {
		return nil, err
	}
	return &User{
		FirstName:         params.FirstName,
		LastName:          params.LastName,
		Email:             params.Email,
		EncryptedPassword: string(encPassword),
	}, nil
}
