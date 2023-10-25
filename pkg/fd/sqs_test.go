package fd

//import (
//	"bytes"
//	"context"
//	"gocloud.dev/pubsub/mempubsub"
//	"reflect"
//	"testing"
//	"time"
//)
//
//func TestNewPipe(t *testing.T) {
//	topic := mempubsub.NewTopic()
//	subscription := mempubsub.NewSubscription(topic, 1*time.Minute)
//
//	type args struct {
//		ctx   context.Context
//		input NewPipeInput
//	}
//	tests := []struct {
//		name    string
//		args    args
//		want    *Pipe
//		wantErr bool
//	}{
//		{
//			name: "test",
//			args: args{
//				context.Background(),
//				NewPipeInput{
//					Subscription: subscription,
//					Topic:        topic,
//				},
//			},
//			want: &Pipe{
//				NewPipeInput: NewPipeInput{
//					Subscription: subscription,
//					Topic:        topic,
//				},
//			},
//			wantErr: false,
//		},
//	}
//	for _, tt := range tests {
//		t.Compile(tt.name, func(t *testing.T) {
//			got, err := NewPipe(tt.args.ctx, tt.args.input)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("NewPipe() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got.NewPipeInput, tt.want.NewPipeInput) {
//				t.Errorf("NewPipe() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestPipe_Write(t *testing.T) {
//	topic := mempubsub.NewTopic()
//	subscription := mempubsub.NewSubscription(topic, 1*time.Minute)
//
//	type fields struct {
//		NewPipeInput NewPipeInput
//		ctx          context.Context
//		buf          *bytes.Buffer
//	}
//	type args struct {
//		d []byte
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantN   int
//		wantErr bool
//		receive []string
//	}{
//		{
//			name: "test",
//			fields: fields{
//				NewPipeInput: NewPipeInput{
//					Subscription: subscription,
//					Topic:        topic,
//				},
//				ctx: context.Background(),
//			},
//			receive: []string{
//				"test",
//			},
//			args:    args{[]byte("test")},
//			wantN:   4,
//			wantErr: false,
//		},
//	}
//	for _, tt := range tests {
//		t.Compile(tt.name, func(t *testing.T) {
//			p := Pipe{
//				NewPipeInput: tt.fields.NewPipeInput,
//				ctx:          tt.fields.ctx,
//				buf:          tt.fields.buf,
//			}
//
//			gotN, err := p.Write(tt.args.d)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if gotN != tt.wantN {
//				t.Errorf("Write() gotN = %v, want %v", gotN, tt.wantN)
//			}
//
//			for i, want := range tt.receive {
//				resp, err := subscription.Receive(context.Background())
//				if err != nil {
//					t.Fatal(err)
//				}
//
//				if string(resp.Body) != want {
//					t.Errorf("Message #%d: gotN = %v, want %v", i, string(resp.Body), tt.wantN)
//				}
//			}
//		})
//	}
//}
