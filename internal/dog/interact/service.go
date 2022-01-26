package interact

import (
	"github.com/dog-sky/dog_bot/internal/service/db"
	desc "github.com/dog-sky/dog_bot/pkg/dog/api"
)

type Implementation struct {
	db db.DB

	desc.UnimplementedDogServer
}

func New(db db.DB) *Implementation {
	return &Implementation{
		db,
		desc.UnimplementedDogServer{},
	}
}
