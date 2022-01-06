package interact

import (
	"context"

	"github.com/dog-sky/dog_bot/internal/service/db/models"
	desc "github.com/dog-sky/dog_bot/pkg/dog/api"
)

// StatusList ...
func (i *Implementation) StatusList(ctx context.Context, in *desc.StatusListrequest) (*desc.StatusListReply, error) {
	filter := &models.Filter{}

	if in.Filter.GetActions() != nil && len(in.Filter.Actions) != 0 {
		actions := make([]string, len(in.Filter.Actions))
		for i, a := range in.Filter.Actions {
			actions[i] = a.String()
		}

		filter.Actions = actions
	}

	filter.Date.From = in.Filter.Date.GetFrom()
	filter.Date.To = in.Filter.Date.GetTo()

	res, err := i.db.StatusList(ctx, filter)
	if err != nil {
		return nil, err
	}

	return &desc.StatusListReply{
		Result: res,
	}, nil
}
