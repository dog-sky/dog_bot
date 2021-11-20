package interact

import (
	desc "github.com/dog-sky/dog_bot/pkg/dog/api"

	"google.golang.org/grpc"
)

type Implementation struct {
	desc.UnimplementedDogServer
}

func New(s *grpc.Server) {
	desc.RegisterDogServer(s, &Implementation{})
}
