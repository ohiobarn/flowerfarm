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

// SKU #,Image,Crop,Variety,Grower,SKU,This Week,Next Week,Future,Stems per Bunch,Availability Months,Category,This week | Next week | Future,Open in Shop,Shop Description,Color,Tier,Preorder Table,Preorder

type Forecast []ForecastRec

type ForecastRec struct {
	SKU string `json:"SKU #"`
}

// forecastCmd represents the forecast command
var forecastCmd = &cobra.Command{
	Use:   "forecast",
	Short: "Use forecast data to update Square Space inventory records",
	Long: `This command uses the forecast data from AirTable to update the Square Space product records. 
After running this command a new, updatated products file is created and be imported to Square Space.

Example ran from the flowerfarm project root:

 $ obff forecast --forecast wrk/forecast.json --products wrk/products.json
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("forecast called")
		forecastFile, _ := cmd.Flags().GetString("forecast")
		fmt.Printf("forcastFile: %s\n", forecastFile)

		// Open our jsonFile
		jsonFile, err := os.Open(forecastFile)
		// if we os.Open returns an error then handle it
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("Successfully Opened %s", forecastFile)
		// defer the closing of our jsonFile so that we can parse it later on
		defer jsonFile.Close()

		json_data, _ := ioutil.ReadAll(jsonFile)
		forecast := new(Forecast)
    err = json.Unmarshal(json_data, &forecast)

		fmt.Printf("debub: \n%v", forecast)


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

// func updateProductsFromForecast(){

// }
