package power

import "testing"

func TestIsPowerOfThree(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "27",
			args: args{n: 27},
			want: true,
		},
		{
			name: "0",
			args: args{n: 0},
			want: false,
		},
		{
			name: "9",
			args: args{n: 9},
			want: true,
		},
		{
			name: "18",
			args: args{n: 18},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPowerOfThree(tt.args.n); got != tt.want {
				t.Errorf("IsPowerOfThree() = %v, want %v", got, tt.want)
			}
		})
	}
}
