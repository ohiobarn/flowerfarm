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
	forecastProductCollection := cmd.ForecastRun(forecastFileName, productsFileName, productsModifiedFileName)

	// Todo - add properties to the ForecastProductDoc for counters.
	//      - update the logic to use the new counters
	//      - then update this test to assert the counters
	t.Logf("forecastProductCollection: %v", forecastProductCollection)

	var want int = 1
	var got int = 1
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}
