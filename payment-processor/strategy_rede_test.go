package payment

import (
	"context"
	"testing"
)

func TestRedeStrategy_Process(t *testing.T) {
	type args struct {
		ctx context.Context
		p   Payment
		s   Source
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"process with success",
			args{
				context.Background(),
				Payment{},
				Source{},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewRedeStrategy()
			if err := c.Process(tt.args.ctx, tt.args.p, tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("RedeStrategy.Process() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
