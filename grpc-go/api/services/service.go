package services

import (
	"github.com/4x2Hach1/learning/grpc-go/api/models"
	"github.com/4x2Hach1/learning/grpc-go/api/pb"
	"github.com/jmoiron/sqlx"
)

type serviceInfr struct {
	db *models.Sql
}

type serviceUnimplemented struct {
	pb.UnimplementedServerServiceServer
}

type service struct {
	*serviceUnimplemented
	*serviceHello
	*serviceUser
}

func NewService(db *sqlx.DB) *service {
	sql := models.NewSqlDB(db)
	infra := &serviceInfr{sql}

	return &service{
		serviceHello: &serviceHello{serviceInfr: infra},
		serviceUser:  &serviceUser{serviceInfr: infra},
	}
}
