package payment

import "testing"

func TestRedeStrategy_Process(t *testing.T) {
	type args struct {
		p Payment
		s Source
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"process with success",
			args{
				Payment{},
				Source{},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewRedeStrategy()
			if err := c.Process(tt.args.p, tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("RedeStrategy.Process() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
