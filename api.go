package users

import (
	"context"
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
		Factory *factory.Factory
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

func (c *Client)AddNode(ctx context.Context, node models.NodeRegReq) error {
	panic("implement me")
}

func (c *Client)GetNode(ctx context.Context, id string) (models.Node, error) {
	panic("implement me")
}

func (c *Client)ListNodes(ctx context.Context) ([]models.Node, error) {
	panic("implement me")
}

func (c *Client)AddRegion(ctx context.Context, req models.RegionRegReq) error {
	panic("implement me")
}

func (c *Client)GetRegion(ctx context.Context, id string) (models.Region, error) {
	panic("implement me")
}

func (c *Client)ListRegions(ctx context.Context) ([]models.Region, error) {
	panic("implement me")
}

func (c *Client)Register(ctx context.Context, req models.AdminRegReq) (admin models.Admin, err error) {
	panic("implement me")
}

func (c *Client)Login(ctx context.Context, id, password string) (token string, err error) {
	panic("implement me")
}

