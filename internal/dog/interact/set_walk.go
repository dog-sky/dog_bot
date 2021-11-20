package interact

import (
	"context"

	desc "github.com/dog-sky/dog_bot/pkg/dog/api"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) SetWalk(ctx context.Context, in *desc.SetWalkRequest) (*desc.SetWalkReply, error) {
	if in.GetAction() == desc.DogAction_UNKNOWN_DOG_ACTION {
		return nil, status.Error(codes.InvalidArgument, "unknown dog action")
	}

	return &desc.SetWalkReply{
		Result: &desc.SetWalkReply_Result{
			Created: true,
		},
	}, nil
}
