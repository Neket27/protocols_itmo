package math

import (
	"errors"
	"testing"
)

// TestDivide — тесты для функции Divide
func TestDivide(t *testing.T) {
	tests := []struct {
		name      string
		a         float64
		b         float64
		want      float64
		wantErr   bool
		errString string
	}{
		{
			name:    "обычное деление",
			a:       10,
			b:       2,
			want:    5,
			wantErr: false,
		},
		{
			name:      "деление на 0",
			a:         10,
			b:         0,
			want:      0,
			wantErr:   true,
			errString: "division by zero",
		},
		{
			name:    "отрицательные числа",
			a:       -10,
			b:       2,
			want:    -5,
			wantErr: false,
		},
		{
			name:    "дробные числа",
			a:       7.5,
			b:       2.5,
			want:    3,
			wantErr: false,
		},
		{
			name:    "деление на отрицательное",
			a:       10,
			b:       -2,
			want:    -5,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Divide(tt.a, tt.b)

			if tt.wantErr {
				if err == nil {
					t.Errorf("Divide() ожидалась ошибка, но не получена")
					return
				}
				if tt.errString != "" && err.Error() != tt.errString {
					t.Errorf("Divide() ошибка = %v, ожидалась %v", err.Error(), tt.errString)
				}
				return
			}

			if err != nil {
				t.Errorf("Divide() неожиданная ошибка = %v", err)
				return
			}

			if got != tt.want {
				t.Errorf("Divide() = %v, ожидалось %v", got, tt.want)
			}
		})
	}
}

// BenchmarkDivide — бенчмарк производительности
func BenchmarkDivide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Divide(100, 5)
	}

}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return float64(a / b), nil
}
