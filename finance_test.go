package humanize

import (
	"math"
	"testing"
)

func TestFinance(t *testing.T) {
	tests := []struct {
		name string
		arg  float64
		want string
	}{
		{"Quadrillions", 2_475_260_494_216_000, "$2.47Q"},
		{"Trillions", 2_475_260_494_216, "$2.47T"},
		{"Billions", 2_475_260_494, "$2.47B"},
		{"Millions", 2_475_260, "$2.47M"},
		{"Thousands", 2_475, "$2.47K"},
		{"Hundreds", 247, "$247"},
		{"Tens", 24, "$24"},
		{"NaN", math.NaN(), "NaN"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Finance(tt.arg); got != tt.want {
				t.Errorf("Finance() = %v, want %v", got, tt.want)
			}
		})
	}
}
