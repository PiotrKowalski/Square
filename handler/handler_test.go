package handler

import (
	"github.com/PiotrKowalski/square/config"
	pb "github.com/PiotrKowalski/square/proto"
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
