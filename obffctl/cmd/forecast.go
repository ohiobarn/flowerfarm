/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

//
// Forecast
//
type Forecast []ForecastDoc

type ForecastDoc struct {
	SKU                string `json:"SKU #"`
	Image              string `json:"Image"`
	Crop               string `json:"Crop"`
	Variety            string `json:"Variety"`
	Grower             string `json:"Grower"`
	ThisWeek           string `json:"This Week"`
	NextWeek           string `json:"Next Week"`
	Future             string `json:"Future"`
	StemsPerBunch      string `json:"Stems per Bunch"`
	AvailabilityMonths string `json:"Availability Months"`
	Category           string `json:"Category"`
	ThisNextFuture     string `json:"This week | Next week | Future"`
	OpenInShop         string `json:"Open in Shop"`
	Color              string `json:"Color"`
	Tier               string `json:"Tier"`
}

//
// Products
//
type Products []ProductsDoc

type ProductsDoc struct {
	ProductID       string `json:"Product ID [Non Editable]"`
	VariantID       string `json:"Variant ID [Non Editable]""`
	ProductType     string `json:"Product Type [Non Editable]"`
	ProductPage     string `json:"Product Page"`
	ProductURL      string `json:"Product URL"`
	Title           string `json:"Title"`
	Description     string `json:"Description"`
	SKU             string `json:"SKU"`
	OptionName1     string `json:"Option Name 1"`
	OptionValue1    string `json:"Option Value 1"`
	OptionName2     string `json:"Option Name 2"`
	OptionValue2    string `json:"Option Value 2"`
	OptionName3     string `json:"Option Name 3"`
	OptionValue3    string `json:"Option Value 3"`
	OptionName4     string `json:"Option Name 4"`
	OptionValue4    string `json:"Option Value 4"`
	OptionName5     string `json:"Option Name 5"`
	OptionValue5    string `json:"Option Value 5"`
	OptionName6     string `json:"Option Name 6"`
	OptionValue6    string `json:"Option Value 6"`
	Price           string `json:"Price"`
	SalePrice       string `json:"Sale Price"`
	OnSale          string `json:"On Sale"`
	Stock           string `json:"Stock"`
	Categories      string `json:"Categories"`
	Tags            string `json:"Tags"`
	Weight          string `json:"Weight"`
	Length          string `json:"Length"`
	Width           string `json:"Width"`
	Height          string `json:"Height"`
	Visible         string `json:"Visible"`
	HostedImageURLs string `json:"Hosted Image URLs"`
}

// forecastCmd represents the forecast command
var forecastCmd = &cobra.Command{
	Use:   "forecast",
	Short: "Use forecast data to update Square Space inventory records",
	Long: `This command uses the forecast data from AirTable to update the Square Space product records. 
After running this command a new, updatated products file is created and be imported to Square Space.

Example ran from the flowerfarm project root:

 $ obffctl forecast --forecast wrk/forecast.json --products wrk/products.json
`,
	Run: func(cmd *cobra.Command, args []string) {

		forecast := loadForecast(cmd, args)
		products := loadProducts(cmd, args)
		updateProductsWithForecast(forecast, products)
	},
}

func init() {
	rootCmd.AddCommand(forecastCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// forecastCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// forecastCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// forecastCmd.Flags().String("forecast","exports/airtable/forecast.json","Path to the forecast file, must be in json format")
	forecastCmd.Flags().StringP("forecast", "f", "exports/airtable/forecast.json", "Path to the forecast file exported from AirTable, must be in json format")
	forecastCmd.MarkFlagRequired("forecast")

	forecastCmd.Flags().StringP("products", "p", "exports/squarespace/export/products.json", "Path to the products file exported from Square Space, must be in json format")
	forecastCmd.MarkFlagRequired("products")
}

func loadForecast(cmd *cobra.Command, args []string) *Forecast {
	//
	// Open forecastFile
	//
	forecastFileName, _ := cmd.Flags().GetString("forecast")
	fmt.Printf("forcastFilePath: %s\n", forecastFileName)
	forecastFile, err := os.Open(forecastFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Successfully Opened %s", forecastFileName)
	// defer the closing of our forecastFile so that we can parse it later on
	defer forecastFile.Close()

	//
	// Read file into forecast type
	//
	forecastJsonData, _ := ioutil.ReadAll(forecastFile)
	forecast := new(Forecast)
	err = json.Unmarshal(forecastJsonData, &forecast)

	return forecast
}

func loadProducts(cmd *cobra.Command, args []string) *Products {
	//
	// Open productsFile
	//
	productsFileName, _ := cmd.Flags().GetString("products")
	fmt.Printf("productsFilePath: %s\n", productsFileName)
	productsFile, err := os.Open(productsFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Successfully Opened %s", productsFileName)
	// defer the closing of our forecastFile so that we can parse it later on
	defer productsFile.Close()

	//
	// Read file into products type
	//
	productsJsonData, _ := ioutil.ReadAll(productsFile)
	products := new(Products)
	err = json.Unmarshal(productsJsonData, &products)

	return products
}

func updateProductsWithForecast(forecast *Forecast, products *Products) {

	for _, productDoc := range *products {

		fmt.Printf("Processing product sku: %s\n", productDoc.SKU)
		findForecastDocBySKU(productDoc.SKU, forecast)
		fmt.Println("------------------------------------------------------------------------------")
	}

}

func findForecastDocBySKU(sku string, forecast *Forecast) ForecastDoc {

	for _, forecastDoc := range *forecast {
		if forecastDoc.SKU == sku {
			fmt.Printf("\nFound forcast doc sku: %s\n", sku)
			return forecastDoc
		}
	}

	var errorDoc ForecastDoc
	return errorDoc
}
