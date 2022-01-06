package interact

import (
	"context"
	"fmt"

	desc "github.com/dog-sky/dog_bot/pkg/dog/api"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// SetStatus ...
func (i *Implementation) SetStatus(ctx context.Context, in *desc.SetStatusRequest) (*desc.SetStatusReply, error) {
	if in.GetAction() == desc.DogAction_UNKNOWN_DOG_ACTION {
		return nil, status.Error(codes.InvalidArgument, "unknown dog action")
	}

	if err := i.db.SetStatus(ctx, in.Action.String()); err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("can't set action: %s", err.Error()))
	}

	return &desc.SetStatusReply{
		Result: &desc.SetStatusReply_Result{
			Created: true,
		},
	}, nil
}
