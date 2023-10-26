package event

import (
	"reflect"
	"testing"
)

func TestParseTemplateArgs(t *testing.T) {
	var StringPtr *string
	var BoolPtr *bool

	type Result struct {
		TemplateArgs map[string]any
		UnparsedArgs []string
	}

	type args struct {
		args []string
	}
	var tests = []struct {
		name    string
		args    args
		want    Result
		wantErr bool
	}{
		{
			name: "no args",
			args: args{args: []string{}},
			want: Result{
				TemplateArgs: map[string]any{},
				UnparsedArgs: []string{},
			},
		},
		{
			name: "no template args",
			args: args{
				args: []string{"/path/to/file"},
			},
			want: Result{
				TemplateArgs: map[string]any{},
				UnparsedArgs: []string{"/path/to/file"},
			},
			wantErr: false,
		},
		{
			name: "test",
			args: args{
				args: []string{
					"--state=@State:string",
					"--ok=@Ok:bool",
				},
			},
			want: Result{
				TemplateArgs: map[string]any{
					"State": StringPtr,
					"Ok":    BoolPtr,
				},
				UnparsedArgs: []string{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTemplate, gotUnparsed, err := ParseTemplateArgs(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ParseTemplateArgs() error = %v", err)
			}

			got := Result{TemplateArgs: gotTemplate, UnparsedArgs: gotUnparsed}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseTemplateArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
