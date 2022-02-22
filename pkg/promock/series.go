package promock

import (
	"fmt"
	"math"
	"regexp"

	"github.com/Knetic/govaluate"
	//"github.com/prometheus/prometheus/model/labels"
)

var (
	nan = math.NaN()
)

type SeriesConfig struct {
	ID         string            `yaml:"id"`
	Labels     map[string]string `yaml:"labels"`
	MetricName string            `yaml:"metric_name"`
	Expr       string            `yaml:"expr"`
	Step       uint              `yaml:"step"`
}

func ValidateSeries(s *SeriesConfig) error {
	if s.MetricName == "" {
		return fmt.Errorf("metric name is empty")
	}
	if !metricRegex.MatchString(s.MetricName) {
		return fmt.Errorf("metric name is invalid")
	}
	return nil
}

func NewSeries(sc SeriesConfig) (*Series, error) {
	s := &Series{
		sc:     &sc,
		params: make(map[string]interface{}),
	}
	// copy labels
	newLabels := make(map[string]string)
	for k, v := range sc.Labels {
		newLabels[k] = v
	}
	s.sc.Labels = newLabels
	if err := s.Validate(); err != nil {
		return nil, err
	}
	e, err := govaluate.NewEvaluableExpressionWithFunctions(s.sc.Expr, commonFuncs)
	if err != nil {
		return nil, err
	}
	s.compiledExpr = e
	return s, s.Validate()
}

// Series is a configuration for generating a time serices. The name is
// a user configurable value for tracking the series is generated output
// and log lines. Labels are the key value pairs attached to the generated
// metric. Expr is a math expression to calculate the value of the metric at
// time t. Step is the interval (in seconds) the metric should be evaluated at.
type Series struct {
	sc           *SeriesConfig
	compiledExpr *govaluate.EvaluableExpression
	params       map[string]interface{}
}

var (
	metricRegex = regexp.MustCompile("^[a-zA-Z_:][a-zA-Z0-9_:]*$")
)

func (s *Series) Validate() error {
	return ValidateSeries(s.sc)
}

// Eval will calculate the datapoint at the given ts (in milliseconds) using
// series spec.
func (s *Series) Eval(ts int64) (float64, error) {
	s.params["ts"] = ts
	result, err := s.compiledExpr.Evaluate(s.params)
	if err != nil {
		return nan, err
	}
	return convertToFloat64(result)
}
