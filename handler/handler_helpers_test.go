package handler

import (
	pb "github.com/PiotrKowalski/square/proto"
	"reflect"
	"testing"
)

func Test_getVersion(t *testing.T) {
	tests := []struct {
		name    string
		want    *pb.PingResponse
		wantErr bool
	}{
		{
			name:    "",
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getVersion()
			if (err != nil) != tt.wantErr {
				t.Errorf("getVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getVersion() got = %v, want %v", got, tt.want)
			}
		})
	}
}
