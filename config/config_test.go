package config

import (
	"reflect"
	"testing"
)

func TestGetConfig(t *testing.T) {
	tests := []struct {
		name    string
		want    Config
		envs    map[string]string
		wantErr bool
	}{
		{
			name: "TC_01_Correct_env_loaded",
			want: Config{
				GrpcPort:  "30000",
				RestPort:  "40000",
				EchoUrl:   "https://postman-echo.com/get",
				ServeHTTP: true,
			},
			envs: map[string]string{
				"GRPC_PORT":  "30000",
				"REST_PORT":  "40000",
				"ECHO_URL":   "https://postman-echo.com/get",
				"SERVE_HTTP": "true",
			},
		},
		{
			name: "TC_02_GRPC_PORT_missing_default_30000",
			want: Config{
				GrpcPort:  "30000",
				RestPort:  "40000",
				EchoUrl:   "https://postman-echo.com/get",
				ServeHTTP: true,
			},
			envs: map[string]string{
				"REST_PORT":  "40000",
				"ECHO_URL":   "https://postman-echo.com/get",
				"SERVE_HTTP": "true",
			},
			wantErr: false,
		},
		{
			name: "TC_03_REST_PORT_missing_default_40000",
			want: Config{
				GrpcPort:  "30000",
				RestPort:  "40000",
				EchoUrl:   "https://postman-echo.com/get",
				ServeHTTP: true,
			},
			envs: map[string]string{
				"GRPC_PORT":  "30000",
				"ECHO_URL":   "https://postman-echo.com/get",
				"SERVE_HTTP": "true",
			},
		},
		{
			name: "TC_04_ECHO_URL_missing",
			want: Config{
				GrpcPort:  "30000",
				RestPort:  "40000",
				ServeHTTP: true,
			},
			envs: map[string]string{
				"GRPC_PORT":  "30000",
				"REST_PORT":  "40000",
				"SERVE_HTTP": "true",
			},
		},
		{
			name: "TC_05_SERVE_HTTP_missing",
			want: Config{
				GrpcPort:  "30000",
				RestPort:  "40000",
				EchoUrl:   "https://postman-echo.com/get",
				ServeHTTP: false,
			},
			envs: map[string]string{
				"GRPC_PORT": "30000",
				"REST_PORT": "40000",
				"ECHO_URL":  "https://postman-echo.com/get",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.envs {
				t.Setenv(k, v)
			}

			got, err := GetConfig()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}
