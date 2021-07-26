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
	"strings"

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

type Audit struct {
	NewCount         int
	ModifiedCount    int
	UnchangedCount   int
	FPDocs           []fpDoc
	ProductsModified []ProductDoc
}
type fpDoc struct {
	compareStatus    string
	isModifed        uint
	forecastDoc      ForecastDoc
	productBeforeDoc ProductDoc
	productAfterDoc  ProductDoc
}

func (audit *Audit) recordNew(fd ForecastDoc, pd ProductDoc) {

	var fpDoc fpDoc
	fpDoc.compareStatus = status.new
	fpDoc.forecastDoc = fd
	fpDoc.productAfterDoc = pd

	audit.FPDocs = append(audit.FPDocs, fpDoc)
	audit.ProductsModified = append(audit.ProductsModified, pd)
	audit.NewCount++

}

func (audit *Audit) recordModified(
	fd ForecastDoc,
	pdOrig ProductDoc,
	pdMod ProductDoc,
	isMod uint) {

	var fpDoc fpDoc
	fpDoc.compareStatus = status.modified
	fpDoc.isModifed = isMod
	fpDoc.forecastDoc = fd
	fpDoc.productAfterDoc = pdMod
	fpDoc.productBeforeDoc = pdOrig

	audit.FPDocs = append(audit.FPDocs, fpDoc)
	audit.ProductsModified = append(audit.ProductsModified, pdMod)
	audit.ModifiedCount++

}

func (audit *Audit) recordUncchanged(fd ForecastDoc, pd ProductDoc) {

	var fpDoc fpDoc
	fpDoc.compareStatus = status.unchanged
	fpDoc.forecastDoc = fd
	fpDoc.productAfterDoc = pd

	audit.FPDocs = append(audit.FPDocs, fpDoc)
	audit.UnchangedCount++

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

const (
	isModifiedTitle = 1 << iota
	isModifiedStock
	isModifiedDescription
	isModifiedPrice
)

const colorReset = string("\033[0m")
const colorRed = string("\033[31m")
const colorGreen = string("\033[32m")
const colorYellow = string("\033[33m")
const colorBlue = string("\033[34m")

// const colorPurple = string("\033[35m")
const colorCyan = string("\033[36m")

//const colorWhite = string("\033[37m")

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
func ForecastRun(forecastFileName string, productsFileName string, productsModifiedFileName string) *Audit {

	forecast := loadForecast(forecastFileName)
	products := loadProducts(productsFileName)

	//forecastProductCollection := make([]ForecastProductDoc, 0)
	var audit Audit
	// set initial size to keep the append from unecessarly resize
	audit.ProductsModified = make([]ProductDoc, 200)

	//
	// updateProductsFromForecast - Inspect forecast documents update and/or add prouct docs as needed
	//
	updateProductsFromForecast(forecast, products, &audit)

	//
	//
	//
	processingReport(&audit)

	//
	// writeProducts - Output the productsModified slice as a json file
	//
	writeProducts(&audit.ProductsModified, productsModifiedFileName)

	fmt.Printf("\n************ DONE ****************\n")
	return &audit
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

func processingReport(audit *Audit) {
	new := 0
	mod := 0
	same := 0
	other := 0
	dmp := diffmatchpatch.New()

	fmt.Printf("\n\nProcessing Report\n\n")

	for _, doc := range audit.FPDocs {
		switch doc.compareStatus {
		case status.new:
			new++
		case status.modified:
			mod++
		case status.unchanged:
			same++
			fmt.Printf("\n%v[Unchanged]%v SKU: %v, Title: %v\n", colorBlue, colorReset, doc.productAfterDoc.SKU, doc.productAfterDoc.Title)
		default:
			other++
		}
	}

	for _, doc := range audit.FPDocs {
		if doc.compareStatus == status.new {
			fmt.Printf("\n%v[New]%v SKU: %v, Title: %v\n", colorYellow, colorReset, doc.productAfterDoc.SKU, doc.productAfterDoc.Title)
		}
	}

	for _, doc := range audit.FPDocs {
		if doc.compareStatus == status.modified {
			fmt.Printf("\n%v[Modified]%v SKU: %v, Title: %v\n", colorCyan, colorReset, doc.productAfterDoc.SKU, doc.productAfterDoc.Title)

			if doc.isModifed&isModifiedDescription == isModifiedDescription {
				diffs := dmp.DiffMain(doc.productBeforeDoc.Description, doc.productAfterDoc.Description, false)
				diffWrapped := wrapWords(dmp.DiffPrettyText(diffs), 15, "\t")
				// beforeWrapped := wrap(doc.productBeforeDoc.Description, 60, "     ")
				// afterWrapped := wrap(doc.productAfterDoc.Description, 60, "     ")
				//diffWrapped := wrap(doc.productBeforeDoc.Description, doc.productAfterDoc.Description, 120, "\t")

				//fmt.Printf("%v  [Diff  ] %v\n%v\n", colorYellow, colorReset, diffWrapped)
				// fmt.Printf("%v  [Before] %v\n%v\n", colorRed, colorReset, beforeWrapped)
				// fmt.Printf("%v  [After ] %v\n%v\n\n", colorGreen, colorReset, afterWrapped)
				fmt.Printf("%v  [Description] %v\n%v\n\n", colorGreen, colorReset, diffWrapped)
			}
		}
	}

	fmt.Printf("\nfpDoc Len: %v\n", len(audit.FPDocs))
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
func updateProductsFromForecast(
	forecast *Forecast,
	products *Products,
	audit *Audit) {

	totalCount := 0

	for _, forecastDoc := range *forecast {

		totalCount++
		fmt.Printf("\n--------------------------------------------\n")
		fmt.Printf("forecast SKU: %s%s%s\n", colorBlue, forecastDoc.SKU, colorReset)

		productDoc := findProductDocBySKU(forecastDoc.SKU, products)
		productBeforeDoc := productDoc

		if productDoc.SKU == "" {

			newProductDoc := createProduct(forecastDoc)
			audit.recordNew(forecastDoc, newProductDoc)

		} else {

			isModified, productNewDoc := doUpdate(forecastDoc, productDoc)
			if isModified > 0 {

				audit.recordModified(forecastDoc, productBeforeDoc, productNewDoc, isModified)

			} else {

				audit.recordUncchanged(forecastDoc, productDoc)

			}
		}

	}

	fmt.Printf("\n--------------------------------------------\n")
	fmt.Printf("%sProduct update count...: %d%s\n", colorYellow, audit.ModifiedCount, colorReset)
	fmt.Printf("%sProduct create count...: %d%s\n", colorRed, audit.NewCount, colorReset)
	fmt.Printf("%sProduct unchanged count: %d%s\n", colorGreen, audit.UnchangedCount, colorReset)
	fmt.Printf("\nVerify:\n")

	// TODO - remove this section once you have more/better unit tests
	var sum int
	var countStatus string
	countStatus = colorRed + "Error" + colorReset
	sum = audit.ModifiedCount + audit.NewCount + audit.UnchangedCount
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

// todo - change pDoc to value and returen a pDocMod
func doUpdate(fDoc ForecastDoc, pDoc ProductDoc) (uint, ProductDoc) {
	var isModified uint = 0
	dmp := diffmatchpatch.New()

	newProductTitle := strings.TrimSpace(fDoc.Crop) + " - " + strings.TrimSpace(fDoc.Variety)
	newProductStock := strings.TrimSpace(fDoc.ThisWeek)
	newProductDescription := buidDescription(fDoc)
	newProductPrice := strings.TrimSpace(fDoc.PricePerBunch)

	if newProductDescription != pDoc.Description {
		isModified = isModified | isModifiedDescription
		diffs := dmp.DiffMain(pDoc.Description, newProductDescription, false)
		fmt.Printf("%sDescription (diff):%s %s\n\n", colorGreen, colorReset, dmp.DiffPrettyText(diffs))
		fmt.Printf("%sDescription (before):%s %s\n", colorGreen, colorReset, pDoc.Description)
		fmt.Printf("%sDescription (after).:%s %s\n\n", colorGreen, colorReset, newProductDescription)

	}

	if newProductTitle != pDoc.Title {
		isModified = isModified | isModifiedTitle
		diffs := dmp.DiffMain(pDoc.Title, newProductTitle, false)
		fmt.Printf("%sTitle:%s %s\n", colorGreen, colorReset, dmp.DiffPrettyText(diffs))
	}

	if newProductStock != pDoc.Stock {
		isModified = isModified | isModifiedStock
		diffs := dmp.DiffMain(pDoc.Stock, newProductStock, false)
		fmt.Printf("%sStock:%s %s\n", colorGreen, colorReset, dmp.DiffPrettyText(diffs))
	}

	if newProductPrice != pDoc.Price {
		isModified = isModified | isModifiedPrice
		diffs := dmp.DiffMain(pDoc.Price, newProductPrice, false)
		fmt.Printf("%sPrice:%s %s\n", colorGreen, colorReset, dmp.DiffPrettyText(diffs))
	}

	// if isModified is greater than zero then something changed
	if isModified > 0 {

		pDoc.Title = newProductTitle
		pDoc.Description = newProductDescription
		pDoc.Stock = newProductStock
		pDoc.Price = newProductPrice
	}

	// Returns true if merge occured
	return isModified, pDoc
}

func buidDescription(fDoc ForecastDoc) string {

	desc := fmt.Sprintf("<p>%s - %s | %s<br>%s<br><hr><b>Forecast:</b> <br>This week: %s <br>Next week: %s <br>Future: %s<br><i>%s stems per bunch</i><hr><br></p>",
		fDoc.Crop,
		fDoc.Variety,
		fDoc.SKU,
		fDoc.ShopDescription,
		fDoc.ThisWeek,
		fDoc.NextWeek,
		fDoc.Future,
		fDoc.StemsPerBunch,
	)

	return strings.TrimSpace(desc)
}

//
// createProduct - Create a default product document from a forecast doc
//                 and add it to the productsModified slice
func createProduct(fDoc ForecastDoc) ProductDoc {

	var pDoc ProductDoc

	pDoc.SKU = fDoc.SKU
	pDoc.Title = fDoc.Crop + " - " + fDoc.Variety
	pDoc.Description = buidDescription(fDoc)
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

// Wraps text at the specified number of columns
func wrap(s1 string, s2 string, limit int, pad string) string {

	if strings.TrimSpace(s1) == "" {
		return s1
	}

	var result string = ""

	for len(s1) >= 1 {
		// but insert \r\n at specified limit
		result = result + pad + s1[:limit] + "\r\n" + pad + s2[:limit] + "\r\n\n"

		// discard the elements that were copied over to result
		s1 = s1[limit:]
		s2 = s2[limit:]

		// change the limit
		// to cater for the last few words in
		//
		if len(s1) < limit {
			limit = len(s1)
		}

	}

	return result

}

func wrapWords(s string, limit int, pad string) string {

	if strings.TrimSpace(s) == "" {
		return pad + s
	}

	var result string = ""

	words := strings.Fields(s)
	if len(words) < limit {
		return pad + s
	}

	for len(words) >= 1 {
		// but insert \r\n at specified limit
		result = result + pad + strings.Join(words[:limit], " ") + "\r\n"

		// discard the elements that were copied over to result
		words = words[limit:]

		// change the limit
		// to cater for the last few words in
		//
		if len(words) < limit {
			limit = len(words)
		}

	}

	return result

}
