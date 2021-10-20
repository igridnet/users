package users

import (
	"context"
	"github.com/igridnet/users/models"
)

type (
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

