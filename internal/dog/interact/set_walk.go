package interact

import (
	"context"
	"fmt"

	desc "github.com/dog-sky/dog_bot/pkg/dog/api"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) SetWalk(ctx context.Context, in *desc.SetWalkRequest) (*desc.SetWalkReply, error) {
	if in.GetAction() == desc.DogAction_UNKNOWN_DOG_ACTION {
		return nil, status.Error(codes.InvalidArgument, "unknown dog action")
	}

	err := i.db.SetStatus(ctx, in.Action.String())
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("can't set action: %s", err.Error()))
	}

	return &desc.SetWalkReply{
		Result: &desc.SetWalkReply_Result{
			Created: true,
		},
	}, nil
}
