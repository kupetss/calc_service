package calculator

import "testing"

func TestCalculate(t *testing.T) {
	tests := []struct {
		expression string
		expected   float64
		hasError   bool
	}{
		{"2+2", 4, false},
		{"2+2*2", 6, false},
		{"(2+2)*2", 8, false},
		{"10/2", 5, false},
		{"10/0", 0, true},
		{"2++2", 0, true},
	}

	for _, tt := range tests {
		result, err := Calculate(tt.expression)
		if (err != nil) != tt.hasError {
			t.Errorf("unexpected error state for %s: got %v, expected error %v", tt.expression, err, tt.hasError)
		}
		if !tt.hasError && result != tt.expected {
			t.Errorf("unexpected result for %s: got %f, expected %f", tt.expression, result, tt.expected)
		}
	}
}
