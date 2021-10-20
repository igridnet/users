package api

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/igridnet/users"
	"github.com/igridnet/users/factory"
	"github.com/igridnet/users/models"
)

func NewClient(db *pg.DB, f *factory.Factory, tokenizer users.Tokenizer, hasher users.Hasher) *Client {
	return &Client{
		Tokenizer: tokenizer,
		Hasher:    hasher,
		Factory:   f,
		Db:        db,
	}
}

var (
	_ users.Service       = (*Client)(nil)
	_ users.RegionService = (*Client)(nil)
	_ users.NodeService   = (*Client)(nil)
)

type Client struct {
	Tokenizer users.Tokenizer
	Hasher    users.Hasher
	Factory   *factory.Factory
	Db        *pg.DB
}

func (c *Client) AddNode(ctx context.Context, node models.NodeRegReq) error {
	newNode, err := c.Factory.NewNode(node)
	if err != nil {
		return err
	}
	_, err = c.Db.Model(&newNode).Returning("*").Insert()
	if err != nil {
		return fmt.Errorf("could not insert new node %v due to: %w\n", node, err)
	}

	return nil
}

func (c *Client) GetNode(ctx context.Context, id string) (models.Node, error) {
	node := new(models.Node)
	err := c.Db.Model(node).Where("id = ?", id).Select()
	if err != nil {
		return models.Node{}, fmt.Errorf("could not retrieve node of id %s: %w", id, err)
	}

	tt := *node
	return tt, nil
}

func (c *Client) ListNodes(ctx context.Context) ([]models.Node, error) {
	var nodes []models.Node
	err := c.Db.Model(&nodes).Select()
	if err != nil {
		return nil, err
	}
	return nodes, nil
}

func (c *Client) AddRegion(ctx context.Context, req models.RegionRegReq) error {
	newRegion, err := c.Factory.NewRegion(req)
	if err != nil {
		return err
	}
	_, err = c.Db.Model(&newRegion).Returning("*").Insert()
	if err != nil {
		return fmt.Errorf("could not insert new region %v due to: %w\n", newRegion, err)
	}

	return nil
}

func (c *Client) GetRegion(ctx context.Context, id string) (models.Region, error) {
	region := new(models.Region)
	err := c.Db.Model(region).Where("id = ?", id).Select()
	if err != nil {
		return models.Region{}, fmt.Errorf("could not retrieve node of id %s: %w", id, err)
	}

	tt := *region
	return tt, nil
}

func (c *Client) ListRegions(ctx context.Context) ([]models.Region, error) {
	var regions []models.Region
	err := c.Db.Model(&regions).Select()
	if err != nil {
		return nil, err
	}
	return regions, nil
}

func (c *Client) Register(ctx context.Context, req models.AdminRegReq) (admin models.Admin, err error) {
	newAdmin, err := c.Factory.NewAdmin(req)
	if err != nil {
		return admin, err
	}
	_, err = c.Db.Model(&newAdmin).Returning("*").Insert()
	if err != nil {
		return admin, fmt.Errorf("could not insert new region %v due to: %w\n", newAdmin, err)
	}

	return newAdmin, nil
}

func (c *Client) Login(ctx context.Context, id, password string) (token string, err error) {
	admin := new(models.Admin)
	err = c.Db.Model(admin).Where("id = ?", id).Select()
	if err != nil {
		return "", fmt.Errorf("could not retrieve user of id %s: %w", id, err)
	}

	hashedPassword := admin.Password

	//compare passwords
	err = c.Hasher.Compare(hashedPassword, password)
	if err != nil {
		return "", err
	}

	key := users.NewKey(admin.ID, "admin")

	return c.Tokenizer.Issue(key)

}

