package randomstring

import (
	"reflect"
	"testing"
)

func TestRandomString(t *testing.T) {
	tests := []struct {
		name string
		want []byte
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandomString(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RandomString() = %v, want %v", got, tt.want)
			}
		})
	}
}
