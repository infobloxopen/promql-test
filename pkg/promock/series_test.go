package promock

import (
	"testing"
)

func TestSeries_Eval(t *testing.T) {

	type args struct {
		ts int64
	}
	tests := []struct {
		name    string
		sc      SeriesConfig
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "simple",
			sc: SeriesConfig{
				ID:         "simple",
				MetricName: "dummy",
				Expr:       "float(ts) - 10",
				Step:       15,
			},
			args: args{
				ts: 1000,
			},
			want:    990,
			wantErr: false,
		},
		{
			name: "no metric name",
			sc: SeriesConfig{
				ID:   "simple",
				Expr: "float(ts) - 10",
				Step: 15,
			},
			args: args{
				ts: 1000,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := NewSeries(
				tt.sc,
			)
			if err != nil {
				t.Fatalf("NewSeries unexpectedly failed: %s", err)
			}
			got, err := s.Eval(tt.args.ts)
			if (err != nil) != tt.wantErr {
				t.Errorf("Series.Eval() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Series.Eval() = %v, want %v", got, tt.want)
			}
		})
	}
}
