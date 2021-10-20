package uuid

import (
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/igridnet/users"
)

// ErrGeneratingID indicates error in generating UUID
var ErrGeneratingID = errors.New("generating id failed")

var _ users.IDProvider = (*uuidProvider)(nil)

type uuidProvider struct{}

// New instantiates a UUID provider.
func New() users.IDProvider {
	return &uuidProvider{}
}

func (up *uuidProvider) ID() (string, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return "", fmt.Errorf("%s: %w", ErrGeneratingID.Error(), err)
	}

	return id.String(), nil
}
