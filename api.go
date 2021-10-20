package users

import (
	"context"
	"github.com/igridnet/users/models"
)

type (
	Service interface {
		Register(ctx context.Context, email, password string)(admin models.Admin,err error)
		Login(ctx context.Context, id, password string) (token string, err error)
	}

	RegionService interface {
		Add(ctx context.Context)error
		Get(ctx context.Context, id string)(models.Region,error)
		List(ctx context.Context)([]models.Region,error)
	}

	NodeService interface {
		Add(ctx context.Context, node models.Node)error
		Get(ctx context.Context, id string)(models.Node,error)
		List(ctx context.Context)([]models.Node,error)
	}
)

