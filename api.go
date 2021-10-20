package users

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/igridnet/users/factory"
	"github.com/igridnet/users/models"
)

var (
	_ Service = (*Client)(nil)
	_ RegionService = (*Client)(nil)
	_ NodeService = (*Client)(nil)
)



type (

	Client struct {
		Tokenizer Tokenizer
		hasher    Hasher
		factory *factory.Factory
		db      *pg.DB
	}
	Service interface {
		Register(ctx context.Context, req models.AdminRegReq) (admin models.Admin, err error)
		Login(ctx context.Context, id, password string) (token string, err error)
	}

	RegionService interface {
		AddRegion(ctx context.Context, req models.RegionRegReq) error
		GetRegion(ctx context.Context, id string) (models.Region, error)
		ListRegions(ctx context.Context) ([]models.Region, error)
	}

	NodeService interface {
		AddNode(ctx context.Context, node models.NodeRegReq) error
		GetNode(ctx context.Context, id string) (models.Node, error)
		ListNodes(ctx context.Context) ([]models.Node, error)
	}
)

func NewClient(db *pg.DB, f *factory.Factory, tokenizer Tokenizer, hasher Hasher) *Client {
	return &Client{
		Tokenizer: tokenizer,
		hasher:    hasher,
		factory:   f,
		db:        db,
	}
}

func (c *Client)AddNode(ctx context.Context, node models.NodeRegReq) error {
	newNode, err := c.factory.NewNode(node)
	if err != nil {
		return err
	}
	_, err = c.db.Model(&newNode).Returning("*").Insert()
	if err != nil {
		return fmt.Errorf("could not insert new node %v due to: %w\n", node, err)
	}

	return nil
}

func (c *Client)GetNode(ctx context.Context, id string) (models.Node, error) {
	node := new(models.Node)
	err := c.db.Model(node).Where("id = ?", id).Select()
	if err != nil {
		return models.Node{}, fmt.Errorf("could not retrieve node of id %s: %w", id, err)
	}

	tt := *node
	return tt, nil
}

func (c *Client)ListNodes(ctx context.Context) ([]models.Node, error) {
	var nodes []models.Node
	err := c.db.Model(&nodes).Select()
	if err != nil {
		return nil, err
	}
	return nodes, nil
}

func (c *Client)AddRegion(ctx context.Context, req models.RegionRegReq) error {
	newRegion, err := c.factory.NewRegion(req)
	if err != nil {
		return err
	}
	_, err = c.db.Model(&newRegion).Returning("*").Insert()
	if err != nil {
		return fmt.Errorf("could not insert new region %v due to: %w\n", newRegion, err)
	}

	return nil
}

func (c *Client)GetRegion(ctx context.Context, id string) (models.Region, error) {
	region := new(models.Region)
	err := c.db.Model(region).Where("id = ?", id).Select()
	if err != nil {
		return models.Region{}, fmt.Errorf("could not retrieve node of id %s: %w", id, err)
	}

	tt := *region
	return tt, nil
}

func (c *Client)ListRegions(ctx context.Context) ([]models.Region, error) {
	var regions []models.Region
	err := c.db.Model(&regions).Select()
	if err != nil {
		return nil, err
	}
	return regions, nil
}

func (c *Client)Register(ctx context.Context, req models.AdminRegReq) (admin models.Admin, err error) {
	newAdmin, err := c.factory.NewAdmin(req)
	if err != nil {
		return admin,err
	}
	_, err = c.db.Model(&newAdmin).Returning("*").Insert()
	if err != nil {
		return admin,fmt.Errorf("could not insert new region %v due to: %w\n", newAdmin, err)
	}

	return newAdmin,nil
}

func (c *Client)Login(ctx context.Context, id, password string) (token string, err error) {
	admin := new(models.Admin)
	err = c.db.Model(admin).Where("id = ?", id).Select()
	if err != nil {
		return "", fmt.Errorf("could not retrieve user of id %s: %w", id, err)
	}

	hashedPassword := admin.Password

	//compare passwords
	err = c.hasher.Compare(hashedPassword, password)
	if err != nil {
		return "", err
	}

	key := NewKey(admin.ID,"admin")

	return c.Tokenizer.Issue(key)

}

