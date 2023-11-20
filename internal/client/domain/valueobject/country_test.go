package valueobject

import "testing"

func TestNewCountry(t *testing.T) {
	type args struct {
		c string
	}
	tests := []struct {
		name    string
		args    args
		want    Country
		wantErr bool
	}{
		{"normal", args{"japan"}, Country("japan"), false},
		{"normal", args{"china"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCountry(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCountry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewCountry() = %v, want %v", got, tt.want)
			}
		})
	}
}
