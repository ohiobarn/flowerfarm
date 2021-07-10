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

	"github.com/sergi/go-diff/diffmatchpatch"
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
	ShopDescription    string `json:"Shop Description"`
	OpenInShop         string `json:"Open in Shop"`
	Color              string `json:"Color"`
	Tier               string `json:"Tier"`
}

//
// Products
//
type Products []ProductDoc

type ProductDoc struct {
	ProductID       string `json:"Product ID [Non Editable]"`
	VariantID       string `json:"Variant ID [Non Editable]"`
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

var productsUpdated []ProductDoc

const colorReset = string("\033[0m")
const colorRed = string("\033[31m")
const colorGreen = string("\033[32m")
const colorYellow = string("\033[33m")
const colorBlue = string("\033[34m")
const colorPurple = string("\033[35m")
const colorCyan = string("\033[36m")
const colorWhite = string("\033[37m")

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

		//debug
		// fmt.Println("Check for updates 1")
		// l := len(productsUpdated)
		// fmt.Println("Check for updates 2")
		// if l == 0 {
		// 	fmt.Println("No products need updated")
		// } else {
		// 	printSlice(productsUpdated)
		// }

		fmt.Println("******** DONE ************")
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

	//
	// Other init stuff
	//
	productsUpdated = make([]ProductDoc, 0)
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
	json.Unmarshal(forecastJsonData, &forecast)

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
	json.Unmarshal(productsJsonData, &products)

	return products
}

func updateProductsWithForecast(forecast *Forecast, products *Products) {
	fmt.Printf("-----------------------------------------------------------------\n")
	fmt.Printf("*** Update products from forecast\n")
	fmt.Printf("-----------------------------------------------------------------\n")

	productUpdateCount := 0

	for _, productDoc := range *products {

		fmt.Printf("\n<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<\n")
		fmt.Printf("Processing product sku: %s%s%s\n", colorBlue, productDoc.SKU, colorReset)

		forecastDoc := findForecastDocBySKU(productDoc.SKU, forecast)
		if forecastDoc.SKU == "" {

			fmt.Printf("%sNo matching forcast found%s\n", colorYellow, colorReset)

		} else {

			fmt.Printf("Found matching forcast doc...\n")
			if doMerge(&forecastDoc, &productDoc) {
				productUpdateCount++
			}

		}
		fmt.Printf(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>\n")
	}
	fmt.Printf("%sTotal products updated [%d/%d]%s\n", colorRed, productUpdateCount, len(*products), colorReset)
}

func findForecastDocBySKU(sku string, forecast *Forecast) ForecastDoc {

	for _, forecastDoc := range *forecast {
		if forecastDoc.SKU == sku {
			return forecastDoc
		}
	}
	var errorDoc ForecastDoc
	return errorDoc
}

//
// doMerge
//   Use forecast doc to genererate a new product doc.  Then compare
//   the new product with existing product to see if any change is needed
//
func doMerge(f *ForecastDoc, p *ProductDoc) bool {
	needToMerge := false
	dmp := diffmatchpatch.New()

	newProductTitle := f.Crop + " - " + f.Variety
	newProductStock := f.ThisWeek
	newProductDescription := fmt.Sprintf("<p>%s | %s<br>%s<br><hr><b>Forecast:</b> <br>This week: %s <br>Next week: %s <br>Future: %s<br><i>%s stems per bunch</i><hr><br></p>",
		newProductTitle, f.SKU,
		f.ShopDescription,
		f.ThisWeek,
		f.NextWeek,
		f.Future,
		f.StemsPerBunch,
	)

	if newProductDescription != p.Description {
		needToMerge = true
		diffs := dmp.DiffMain(p.Description, newProductDescription, false)
		fmt.Printf("%sDescription:%s %s\n", colorGreen, colorReset, dmp.DiffPrettyText(diffs))
	}

	if newProductTitle != p.Title {
		needToMerge = true
		diffs := dmp.DiffMain(p.Title, newProductTitle, false)
		fmt.Printf("Title: %s\n", dmp.DiffPrettyText(diffs))
	}

	if newProductStock != p.Stock {
		needToMerge = true
		diffs := dmp.DiffMain(p.Stock, newProductStock, false)
		fmt.Printf("Stock: %s\n", dmp.DiffPrettyText(diffs))
	}

	if needToMerge {
		productDocUpdated := p
		productDocUpdated.Title = newProductTitle
		productDocUpdated.Description = newProductDescription
		productDocUpdated.Stock = newProductStock

		productsUpdated = append(productsUpdated, *productDocUpdated)
	}

	// Returns true if merge occured
	return needToMerge

}

func printSlice(p Products) {
	s := p
	fmt.Println("PRINT SLICE:")
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
