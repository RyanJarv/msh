package local

import (
	"os"
	"testing"
)

func TestRun(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				args: []string{"cat", "../../../test/data/test"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		// Close stdin so cat doesn't get stuck waiting for input.
		os.Stdin.Close()

		t.Run(tt.name, func(t *testing.T) {
			if err := Run(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
