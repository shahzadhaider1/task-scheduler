package runtime

import (
	"reflect"
	"testing"
)

func TestNewRuntime(t *testing.T) {
	tests := []struct {
		name    string
		want    *Runtime
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRuntime()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRuntime() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRuntime() got = %v, want %v", got, tt.want)
			}
		})
	}
}
