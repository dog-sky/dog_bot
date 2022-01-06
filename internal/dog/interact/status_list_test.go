package interact

import (
	"context"
	"testing"

	"github.com/dog-sky/dog_bot/internal/service/db/mocks"

	desc "github.com/dog-sky/dog_bot/pkg/dog/api"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestImplementation_StatusList(t *testing.T) {
	mc := minimock.NewController(t)

	type args struct {
		ctx context.Context
		in  *desc.StatusListrequest
	}
	tests := []struct {
		name string

		init func(inDate *timestamppb.Timestamp) *Implementation

		args func(inDate *timestamppb.Timestamp) args

		want       func(inDate *timestamppb.Timestamp) *desc.StatusListReply
		wantErr    bool
		inspectErr func(err error, t *testing.T)
	}{
		{
			name: "unknown action",
			init: func(inDate *timestamppb.Timestamp) *Implementation {
				db := mocks.NewDBMock(mc)
				db.StatusListMock.Return(
					[]*desc.StatusListReply_Action{
						{
							Date:   inDate,
							Action: desc.DogAction_POPIS,
						},
					}, nil,
				)

				return New(db)
			},
			args: func(inDate *timestamppb.Timestamp) args {
				return args{
					context.Background(),
					&desc.StatusListrequest{
						Filter: &desc.StatusListrequest_Filter{
							Date: &desc.StatusListrequest_Filter_Date{
								To: inDate,
							},
							Actions: []desc.DogAction{desc.DogAction_POPIS},
						},
					},
				}
			},
			want: func(inDate *timestamppb.Timestamp) *desc.StatusListReply {
				return &desc.StatusListReply{
					Result: []*desc.StatusListReply_Action{
						{
							Date:   inDate,
							Action: desc.DogAction_POPIS,
						},
					},
				}
			},
			wantErr:    false,
			inspectErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := timestamppb.Now()
			tArgs := tt.args(d)

			receiver := tt.init(d)
			got, err := receiver.StatusList(tArgs.ctx, tArgs.in)

			assert.Equal(t, got, tt.want(d))

			if tt.wantErr {
				assert.Error(t, err)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}
