// Package runner contains the Runner model.
package runner

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrRunnerNameCannotBeEmpty = errors.New("name cannot be empty")
)

// Runner Model that represents the Runner
type Runner struct {
	id           uuid.UUID
	name         string
	emailAddress EmailAddress
	createdAt    time.Time
}

// NewRunner Creates a new Runner
func NewRunner(name, emailAddress string) (Runner, error) {

	//validate the name
	if name == "" {
		return Runner{}, ErrRunnerNameCannotBeEmpty
	}

	//validate the email address
	email, err := NewEmailAddress(emailAddress)
	if err != nil {
		return Runner{}, err
	}

	//create a new UUID for the runner
	id := uuid.New()

	return Runner{
		id:           id,
		name:         name,
		emailAddress: email,
		createdAt:    time.Now().UTC(),
	}, nil
}

// ID Returns the ID of the runner
func (r Runner) ID() uuid.UUID {
	return r.id
}

// Name Returns the name of the runner
func (r Runner) Name() string {
	return r.name
}

// EmailAddress Returns the email address of the runner
func (r Runner) EmailAddress() string {
	return r.emailAddress.String()
}
