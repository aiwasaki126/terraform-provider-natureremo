package valueobject

import (
	"testing"
)

func Test_hasValidIncrementedValue(t *testing.T) {
	type args struct {
		v          float64
		increment  float64
		valueRange *valueRange
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"normal", args{-2, 1, &valueRange{-5, 10}}, false},
		{"left edge", args{-5, 1, &valueRange{-5, 10}}, false},
		{"right edge", args{10, 1, &valueRange{-5, 10}}, false},
		{"invalid increment", args{-1.5, 1, &valueRange{-5, 10}}, true},
		{"decimal increment", args{-2.5, 0.5, &valueRange{-5, 10}}, false},
		{"invalid increment", args{-1.5, 1, &valueRange{-5, 10}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := hasValidIncrementedValue(tt.args.v, tt.args.increment, tt.args.valueRange); (err != nil) != tt.wantErr {
				t.Errorf("hasValidIncrementedValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
