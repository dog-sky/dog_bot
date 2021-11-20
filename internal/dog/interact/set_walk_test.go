package interact

import (
	"context"
	"reflect"
	"testing"

	desc "github.com/dog-sky/dog_bot/pkg/dog/api"
)

func TestImplementation_SetWalk(t *testing.T) {
	t.Parallel()

	mockImp := new(Implementation)

	type args struct {
		ctx context.Context
		in  *desc.SetWalkRequest
	}
	tests := []struct {
		name    string
		i       *Implementation
		args    args
		want    *desc.SetWalkReply
		wantErr bool
	}{
		{
			name: "popis",
			i:    mockImp,
			args: args{
				context.TODO(),
				&desc.SetWalkRequest{
					Action: desc.DogAction_POPIS,
				},
			},
			want: &desc.SetWalkReply{
				Result: &desc.SetWalkReply_Result{
					Created: true,
				},
			},
			wantErr: false,
		},
		{
			name: "error",
			i:    mockImp,
			args: args{
				context.TODO(),
				&desc.SetWalkRequest{
					Action: desc.DogAction_UNKNOWN_DOG_ACTION,
				},
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.i.SetWalk(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Implementation.SetWalk() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Implementation.SetWalk() = %v, want %v", got, tt.want)
			}
		})
	}
}
