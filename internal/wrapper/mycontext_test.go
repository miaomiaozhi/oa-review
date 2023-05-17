package wrapper

import (
	"github.com/kataras/iris/v12"
	"reflect"
	"testing"
)

func TestAcquire(t *testing.T) {
	type args struct {
		original iris.Context
		login    bool
	}
	tests := []struct {
		name string
		args args
		want *Context
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Acquire(tt.args.original, tt.args.login); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Acquire() = %v, want %v", got, tt.want)
			}
		})
	}
}
