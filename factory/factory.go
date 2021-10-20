package factory

import (
	"fmt"
	"github.com/igridnet/users"
	"github.com/igridnet/users/hasher"
	"github.com/igridnet/users/models"
	"github.com/igridnet/users/random"
	"github.com/igridnet/users/uuid"
	"time"
)

type (
	Factory struct {
		Hasher     users.Hasher
		IDS        users.IDProvider
		Randomizer users.Randomizer
	}
)

func NewFactory() *Factory {
	return &Factory{
		Hasher:     hasher.New(),
		IDS:        uuid.New(),
		Randomizer: random.New([]rune("123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyx")),
	}
}

func (f *Factory) NewRegion(req models.RegionRegReq) (models.Region, error) {
	id, err := f.IDS.ID()
	if err != nil {
		return models.Region{}, fmt.Errorf("could not create region: %w\n", err)
	}

	return models.Region{
		ID:      id,
		Name:    req.Name,
		Desc:    req.Desc,
		Created: time.Now().UnixNano(),
	}, nil
}

func (f *Factory) NewNode(addr, name string, typ int, region, latd, longd string, master string) (models.Node, error) {
	id, err := f.IDS.ID()
	if err != nil {
		return models.Node{}, fmt.Errorf("could not create node: %w\n", err)
	}

	key := f.Randomizer.Get(36)

	node := &models.Node{
		UUID:    id,
		Addr:    addr,
		Key:     key,
		Name:    name,
		Type:    typ,
		Region:  region,
		Latd:    latd,
		Long:    longd,
		Created: time.Now().UnixNano(),
		Master:  master,
	}

	valid, err := node.Valid()

	if err != nil {
		return models.Node{}, fmt.Errorf("could not create node: %w\n", err)
	}

	if !valid {
		return models.Node{}, fmt.Errorf("could not create node")
	}
	retValue := *node
	return retValue, nil

}

func (f *Factory) NewAdmin(req models.AdminRegReq) (models.Admin, error) {
	id, err := f.IDS.ID()
	if err != nil {
		return models.Admin{}, fmt.Errorf("could not create user: %w\n", err)
	}

	hashd, err := f.Hasher.Hash(req.Password)

	if err != nil {
		return models.Admin{}, fmt.Errorf("could not create user: %w\n", err)
	}
	user := models.Admin{
		ID:       id,
		Name:     req.Name,
		Email:    req.Email,
		Password: hashd,
		Created:  time.Now().UnixNano(),
	}
	return user, nil
}
