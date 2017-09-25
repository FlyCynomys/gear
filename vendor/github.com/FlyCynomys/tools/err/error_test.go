package err

import (
	"reflect"
	"testing"
)

func TestErrorCode_Error(t *testing.T) {
	type fields struct {
		Code  int
		error error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ErrorCode{
				Code:  tt.fields.Code,
				error: tt.fields.error,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("ErrorCode.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrorCode_Json(t *testing.T) {
	type fields struct {
		Code  int
		error error
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ErrorCode{
				Code:  tt.fields.Code,
				error: tt.fields.error,
			}
			if got := e.Json(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ErrorCode.Json() = %v, want %v", got, tt.want)
			}
		})
	}
}
