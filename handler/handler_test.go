package handler

import (
	"context"
	"github.com/PiotrKowalski/square/config"
	pb "github.com/PiotrKowalski/square/proto"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestNewService(t *testing.T) {
	type args struct {
		config config.Config
	}
	tests := []struct {
		name string
		args args
		want *Service
	}{
		{
			name: "TC_01_Correct",
			args: args{config: config.Config{
				GrpcPort: "30000",
				RestPort: "40000",
				EchoUrl:  "https://postman-echo.com/get",
			}},

			want: &Service{
				UnimplementedServiceServer: &pb.UnimplementedServiceServer{},
				conf: config.Config{
					GrpcPort: "30000",
					RestPort: "40000",
					EchoUrl:  "https://postman-echo.com/get",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_Ping(t *testing.T) {
	conf := config.Config{
		GrpcPort:  "30000",
		RestPort:  "40000",
		EchoUrl:   "https://postman-echo.com/get",
		ServeHTTP: true,
	}

	type fields struct {
		UnimplementedServiceServer *pb.UnimplementedServiceServer
		conf                       config.Config
	}
	type args struct {
		ctx context.Context
		req *pb.PingRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		//want    *pb.PingResponse
		wantErr bool
	}{
		// Should be mocked
		{
			name: "TC_01_Correct_Timestamp",
			fields: fields{
				UnimplementedServiceServer: &pb.UnimplementedServiceServer{},
				conf:                       conf,
			},
			args: args{
				ctx: context.Background(),
				req: &pb.PingRequest{Message: "timestamp"},
			},
			wantErr: false,
		},
		{
			name: "TC_02_Correct_Env",
			fields: fields{
				UnimplementedServiceServer: &pb.UnimplementedServiceServer{},
				conf:                       conf,
			},
			args: args{
				ctx: context.Background(),
				req: &pb.PingRequest{Message: "env"},
			},
			wantErr: false,
		},
		// Needs mock
		//{
		//	name: "TC_03_Correct_Version",
		//	fields: fields{
		//		UnimplementedServiceServer: &pb.UnimplementedServiceServer{},
		//		conf:                       conf,
		//	},
		//	args: args{
		//		ctx: context.Background(),
		//		req: &pb.PingRequest{Message: "version"},
		//	},
		//	wantErr: false,
		//},
		// Would be better if echo response were mocked
		{
			name: "TC_04_Correct_Echo",
			fields: fields{
				UnimplementedServiceServer: &pb.UnimplementedServiceServer{},
				conf:                       conf,
			},
			args: args{
				ctx: context.Background(),
				req: &pb.PingRequest{Message: "echo"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				UnimplementedServiceServer: tt.fields.UnimplementedServiceServer,
				conf:                       tt.fields.conf,
			}
			got, err := s.Ping(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Ping() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.NotNil(t, got.ReturnMessage)

			//if tt.args.req.GetMessage() == "timestamp" || tt.args.req.GetMessage() == "echo" {
			//
			//} else if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("Ping() got = %v, want %v", got, tt.want)
			//}
		})
	}
}
