package uuid

import (
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/igridnet/igridnet"
)

// ErrGeneratingID indicates error in generating UUID
var ErrGeneratingID = errors.New("generating id failed")

var _ igridnet.IDProvider = (*uuidProvider)(nil)

type uuidProvider struct{}

// New instantiates a UUID provider.
func New() igridnet.IDProvider {
	return &uuidProvider{}
}

func (up *uuidProvider) ID() (string, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return "", fmt.Errorf("%s: %w", ErrGeneratingID.Error(), err)
	}

	return id.String(), nil
}
