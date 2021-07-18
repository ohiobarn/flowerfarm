package cmd_test

import (
	"testing"

	"github.com/ohiobarn/flowerfarm/obffctl/cmd"
)

// to run use:  go test -v  ./cmd/...

func TestForecastRun(t *testing.T) {

	forecastFileName := "/Users/tgilkerson/github/flowerfarm/obffctl/test-data/forecast.json"
	productsFileName := "/Users/tgilkerson/github/flowerfarm/obffctl/test-data/products.json"
	productsModifiedFileName := "/Users/tgilkerson/github/flowerfarm/obffctl/test-data/products-modified.json"

	audit := cmd.ForecastRun(forecastFileName, productsFileName, productsModifiedFileName)

	// Check modified count
	want := 2
	got := audit.ModifiedCount
	if want != got {
		t.Errorf("ModifiedCount: want %v, got %v", want, got)
	}

	// Check create count
	want = 1
	got = audit.NewCount
	if want != got {
		t.Errorf("NewCount: want %v, got %v", want, got)
	}

	// Check unchanged count
	want = 1
	got = audit.UnchangedCount
	if want != got {
		t.Errorf("UnchangedCount: want %v, got %v", want, got)
	}
}
