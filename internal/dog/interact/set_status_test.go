package interact

import (
	"context"
	"testing"

	"github.com/dog-sky/dog_bot/internal/service/db/mocks"

	desc "github.com/dog-sky/dog_bot/pkg/dog/api"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestImplementation_SetStatus(t *testing.T) {
	mc := minimock.NewController(t)

	type args struct {
		ctx context.Context
		in  *desc.SetStatusRequest
	}
	tests := []struct {
		name string
		init func() *Implementation

		args func(t *testing.T) args

		want       *desc.SetStatusReply
		wantErr    bool
		inspectErr func(err error, t *testing.T)
	}{
		{
			name: "unknown action",
			init: func() *Implementation {
				db := mocks.NewDBMock(mc)

				return New(db)
			},
			args: func(t *testing.T) args {
				return args{
					context.Background(),
					&desc.SetStatusRequest{
						Action: desc.DogAction_UNKNOWN_DOG_ACTION,
					},
				}
			},
			want:    nil,
			wantErr: true,
			inspectErr: func(err error, t *testing.T) {
				st, _ := status.FromError(err)

				assert.Equal(t, codes.InvalidArgument, st.Code())
			},
		},
		{
			name: "popis action",
			init: func() *Implementation {
				db := mocks.NewDBMock(mc)
				db.SetStatusMock.Expect(context.Background(), desc.DogAction_POPIS.String()).Return(nil)

				return New(db)
			},
			args: func(t *testing.T) args {
				return args{
					context.Background(),
					&desc.SetStatusRequest{
						Action: desc.DogAction_POPIS,
					},
				}
			},
			want: &desc.SetStatusReply{
				Result: &desc.SetStatusReply_Result{
					Created: true,
				},
			},
			wantErr: false,
			inspectErr: func(err error, t *testing.T) {
				assert.NoError(t, err)
			},
		},
		{
			name: "walk action",
			init: func() *Implementation {
				db := mocks.NewDBMock(mc)
				db.SetStatusMock.Expect(context.Background(), desc.DogAction_WALK.String()).Return(nil)

				return New(db)
			},
			args: func(t *testing.T) args {
				return args{
					context.Background(),
					&desc.SetStatusRequest{
						Action: desc.DogAction_WALK,
					},
				}
			},
			want: &desc.SetStatusReply{
				Result: &desc.SetStatusReply_Result{
					Created: true,
				},
			},
			wantErr: false,
			inspectErr: func(err error, t *testing.T) {
				assert.NoError(t, err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			receiver := tt.init()
			got, err := receiver.SetStatus(tArgs.ctx, tArgs.in)

			assert.Equal(t, got, tt.want)

			if tt.wantErr {
				assert.Error(t, err)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}
