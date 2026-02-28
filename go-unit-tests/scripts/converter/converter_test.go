package converter

import (
	"reflect"
	"testing"
)

func Test_StructToMap(t *testing.T) {
	type args struct {
		data any
	}
	tests := []struct {
		name string
		args args
		want map[string]any
	}{
		{
			name: "ok",
			args: args{
				data: struct {
					A int
					B string
				}{
					A: 1,
					B: "bloop",
				},
			},
			want: map[string]any{
				"A": float64(1), // float64, т.к. преобразуем в json.
				"B": "bloop",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StructToMap(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("structToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}