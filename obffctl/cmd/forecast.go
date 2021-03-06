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
	AvailabilityMonths string `json:"Availability Months"`
	Category           string `json:"Category"`
	ThisNextFuture     string `json:"This week | Next week | Future"`
	ShopDescription    string `json:"Shop Description"`
	OpenInShop         string `json:"Open in Shop"`
	Color              string `json:"Color"`
	Tier               string `json:"Tier"`
	StemsPerBunch      string `json:"Stems per Bunch"`
	PricePerBunch      string `json:"Price per Bunch"`
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

type ForecastProductDoc struct {
	compareStatus string
	forecastDoc   ForecastDoc
	productDoc    ProductDoc
}

type compareStatus struct {
	new       string
	modified  string
	unchanged string
}

var status = compareStatus{
	new:       "New",
	modified:  "Modified",
	unchanged: "Unchanged",
}

const colorReset = string("\033[0m")
const colorRed = string("\033[31m")
const colorGreen = string("\033[32m")
const colorYellow = string("\033[33m")
const colorBlue = string("\033[34m")

// const colorPurple = string("\033[35m")
// const colorCyan = string("\033[36m")
// const colorWhite = string("\033[37m")

// forecastCmd represents the forecast command
var forecastCmd = &cobra.Command{
	Use:   "forecast",
	Short: "Use forecast data to update Square Space inventory records",
	Long: `This command uses the forecast data from AirTable to update the Square Space product records. 
After running this command a new, updatated products file is created and be imported to Square Space.

Example ran from the flowerfarm project root:

 $ obffctl forecast --forecast wrk/forecast.json --products wrk/products.json
`,
	Run: forecastRun,
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

	forecastCmd.Flags().StringP("products-modified", "m", "exports/squarespace/export/products-modified.json", "Output path to the new modified products file")
	forecastCmd.MarkFlagRequired("products-modified")
}

func forecastRun(cmd *cobra.Command, args []string) {
	forecastFileName, _ := cmd.Flags().GetString("forecast")
	productsFileName, _ := cmd.Flags().GetString("products")
	productsModifiedFileName, _ := cmd.Flags().GetString("products-modified")

	ForecastRun(forecastFileName, productsFileName, productsModifiedFileName)
}
func ForecastRun(forecastFileName string, productsFileName string, productsModifiedFileName string) *[]ForecastProductDoc {

	forecast := loadForecast(forecastFileName)
	products := loadProducts(productsFileName)

	productsModified := make([]ProductDoc, 0)
	forecastProductCollection := make([]ForecastProductDoc, 0)

	//
	// updateProductsFromForecast - Inspect forecast documents update and/or add prouct docs as needed
	//
	updateProductsFromForecast(forecast, products, &productsModified, &forecastProductCollection)

	//
	//
	//
	processingReport(&forecastProductCollection)

	//
	// writeProducts - Output the productsModified slice as a json file
	//
	writeProducts(&productsModified, productsModifiedFileName)

	fmt.Printf("\n************ DONE ****************\n")
	return &forecastProductCollection
}

func loadForecast(forecastFileName string) *Forecast {
	//
	// Open forecastFile
	//
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

func loadProducts(productsFileName string) *Products {
	//
	// Open productsFile
	//
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

func writeProducts(productsModified *[]ProductDoc, productsModifiedFileName string) {

	fmt.Printf("\nWrite productsMofified as json file: %s\n", productsModifiedFileName)

	data, _ := json.Marshal(productsModified)
	_ = ioutil.WriteFile(productsModifiedFileName, data, 0644)

}

func processingReport(fpColl *[]ForecastProductDoc) {
	new := 0
	mod := 0
	same := 0
	other := 0

	for _, fpDoc := range *fpColl {
		switch fpDoc.compareStatus {
		case status.new:
			new++
		case status.modified:
			mod++
		case status.unchanged:
			same++
		default:
			other++
		}
	}
	fmt.Printf("New: %v, Mod: %v, Same: %v, Other: %v, Total: %v\n", new, mod, same, other, new+mod+same+other)
}

//
// updateProductsFromForecast
//   For each forecast doc do:
//     - find its cooresponding product doc, by matching on SKU
//     - compare the two and if different use forcase to update product
//     - if no cooresponding product doc is found creat a new product doc
//     - add the new/modified product doc to the `productsModified` slice
//
func updateProductsFromForecast(forecast *Forecast, products *Products, productsModified *[]ProductDoc, forecastProductCollection *[]ForecastProductDoc) {

	productUpdateCount := 0
	productCreateCount := 0
	productUnchangedCount := 0
	totalCount := 0

	for _, forecastDoc := range *forecast {

		totalCount++
		fmt.Printf("\n--------------------------------------------\n")
		fmt.Printf("forecast SKU: %s%s%s\n", colorBlue, forecastDoc.SKU, colorReset)

		productDoc := findProductDocBySKU(forecastDoc.SKU, products)
		if productDoc.SKU == "" {

			newProductDoc := createProduct(forecastDoc)

			*productsModified = append(*productsModified, newProductDoc)

			var fpDoc ForecastProductDoc
			fpDoc.compareStatus = status.new
			fpDoc.forecastDoc = forecastDoc
			fpDoc.productDoc = newProductDoc
			*forecastProductCollection = append(*forecastProductCollection, fpDoc)

			fmt.Printf("%s[ Create ] %s", colorYellow, colorReset)
			fmt.Printf("Forecast has no matching Product, Product will be created from forecast\n")
			productCreateCount++

		} else {

			if doUpdate(forecastDoc, productDoc) {

				*productsModified = append(*productsModified, productDoc)

				var fpDoc ForecastProductDoc
				fpDoc.compareStatus = status.modified
				fpDoc.forecastDoc = forecastDoc
				fpDoc.productDoc = productDoc
				*forecastProductCollection = append(*forecastProductCollection, fpDoc)

				fmt.Printf("%s[ Modified ] %s", colorYellow, colorReset)
				fmt.Printf("Forecast and Product are different, Product will be updated from forecast\n")
				productUpdateCount++

			} else {

				if false {
					fmt.Printf("%s[ No change ] %s", colorYellow, colorReset)
					fmt.Printf("Forecast and Product are the same, do nothing\n")
				}

				var fpDoc ForecastProductDoc
				fpDoc.compareStatus = status.unchanged
				fpDoc.forecastDoc = forecastDoc
				fpDoc.productDoc = productDoc
				*forecastProductCollection = append(*forecastProductCollection, fpDoc)

				productUnchangedCount++
			}
		}
	}

	fmt.Printf("\n--------------------------------------------\n")
	fmt.Printf("%sProduct update count...: %d%s\n", colorYellow, productUpdateCount, colorReset)
	fmt.Printf("%sProduct create count...: %d%s\n", colorRed, productCreateCount, colorReset)
	fmt.Printf("%sProduct unchanged count: %d%s\n", colorGreen, productUnchangedCount, colorReset)
	fmt.Printf("\nVerify:\n")

	var sum int
	var countStatus string
	countStatus = colorRed + "Error" + colorReset
	sum = productUpdateCount + productCreateCount + productUnchangedCount
	if (sum == totalCount) && (sum == len(*forecast)) {
		countStatus = colorGreen + "OK" + colorReset
	}

	fmt.Printf(" [%s] loop total, count sum and forecast slice size must all match: [%d / %d / %d]\n", countStatus, totalCount, sum, len(*forecast))

}

//
// doUpdate
//   Use forecast doc to genererate a new product doc.  Then compare
//   the new product with existing product to see if any change is needed
//
func doUpdate(fDoc ForecastDoc, pDoc ProductDoc) bool {
	doUpdate := false
	dmp := diffmatchpatch.New()

	newProductTitle := fDoc.Crop + " - " + fDoc.Variety
	newProductStock := fDoc.ThisWeek
	newProductDescription := buidTitle(fDoc)
	newProductPrice := fDoc.PricePerBunch

	if newProductDescription != pDoc.Description {
		doUpdate = true
		diffs := dmp.DiffMain(pDoc.Description, newProductDescription, false)
		fmt.Printf("%sDescription:%s %s\n", colorGreen, colorReset, dmp.DiffPrettyText(diffs))
	}

	if newProductTitle != pDoc.Title {
		doUpdate = true
		diffs := dmp.DiffMain(pDoc.Title, newProductTitle, false)
		fmt.Printf("%sTitle:%s %s\n", colorGreen, colorReset, dmp.DiffPrettyText(diffs))
	}

	if newProductStock != pDoc.Stock {
		doUpdate = true
		diffs := dmp.DiffMain(pDoc.Stock, newProductStock, false)
		fmt.Printf("%sStock:%s %s\n", colorGreen, colorReset, dmp.DiffPrettyText(diffs))
	}

	if newProductPrice != pDoc.Price {
		doUpdate = true
		diffs := dmp.DiffMain(pDoc.Price, newProductPrice, false)
		fmt.Printf("%sPrice:%s %s\n", colorGreen, colorReset, dmp.DiffPrettyText(diffs))
	}

	if doUpdate {

		pDoc.Title = newProductTitle
		pDoc.Description = newProductDescription
		pDoc.Stock = newProductStock
		pDoc.Price = newProductPrice
	}

	// Returns true if merge occured
	return doUpdate
}

func buidTitle(fDoc ForecastDoc) string {

	title := fmt.Sprintf("<p>%s - %s | %s<br>%s<br><hr><b>Forecast:</b> <br>This week: %s <br>Next week: %s <br>Future: %s<br><i>%s stems per bunch</i><hr><br></p>",
		fDoc.Crop,
		fDoc.Variety,
		fDoc.SKU,
		fDoc.ShopDescription,
		fDoc.ThisWeek,
		fDoc.NextWeek,
		fDoc.Future,
		fDoc.StemsPerBunch,
	)

	return title
}

//
// createProduct - Create a default product document from a forecast doc
//                 and add it to the productsModified slice
func createProduct(fDoc ForecastDoc) ProductDoc {

	var pDoc ProductDoc

	pDoc.SKU = fDoc.SKU
	pDoc.Title = fDoc.Crop + " - " + fDoc.Variety
	pDoc.Description = buidTitle(fDoc)
	pDoc.Price = fDoc.PricePerBunch
	pDoc.Stock = fDoc.ThisWeek
	pDoc.Tags = "mrfc"
	pDoc.ProductType = "PHYSICAL"
	pDoc.ProductPage = "mrfc"
	pDoc.Visible = "true"

	fmt.Printf("Title: %s\n", pDoc.Description)
	fmt.Printf("Description: %s\n", pDoc.Description)
	fmt.Printf("Stock: %s\n", pDoc.Description)

	return pDoc

}

func findProductDocBySKU(forecastSKU string, products *Products) ProductDoc {

	for _, productsDoc := range *products {
		if productsDoc.SKU == forecastSKU {
			return productsDoc
		}
	}
	var errorDoc ProductDoc
	return errorDoc
}
