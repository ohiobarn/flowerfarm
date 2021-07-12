#!/bin/bash

goMerge() {
    
  echo "Using Forecast: $FORECAST"
  echo "Using Products: $PRODUCTS"

  #
  # csv to json
  #
  csvjson --no-inference "$FORECAST" | jq > ./wrk/forecast.json
  csvjson --no-inference "$PRODUCTS" | jq > ./wrk/products.json

  #
  # go Merge
  #
  obffctl/obffctl forecast \
    --forecast ./wrk/forecast.json \
    --products ./wrk/products.json \
    --products-modified ./wrk/products-modified.json

  #
  # json to csv
  #
  jq -r '.[] | [
      ."Product ID [Non Editable]", 
      ."Variant ID [Non Editable]",
      ."Product Type [Non Editable]",
      ."Product Page",
      ."Product URL",
      ."Title",
      ."Description",
      ."SKU",
      ."Option Name 1",
      ."Option Value 1",
      ."Option Name 2",
      ."Option Value 2",
      ."Option Name 3",
      ."Option Value 3",     
      ."Price",
      ."Sale Price",
      ."On Sale",
      ."Stock",
      ."Categories",
      ."Tags",
      ."Weight",
      ."Length",
      ."Width",
      ."Height",
      ."Visible",
      ."Hosted Image URLs"    
    ] | @csv' ./wrk/products-modified.json  > ./wrk/products-modified.csv

  #
  # Add header row
  #
  header='"Product ID [Non Editable]","Variant ID [Non Editable]","Product Type [Non Editable]","Product Page","Product URL","Title","Description","SKU","Option Name 1","Option Value 1","Option Name 2","Option Value 2","Option Name 3","Option Value 3","Price","Sale Price","On Sale","Stock","Categories","Tags","Weight","Length","Width","Height","Visible","Hosted Image URLs"'
  echo "$(echo "$header" | cat - ./wrk/products-modified.csv)" > ./wrk/products-modified.csv
  
  echo "**************************"
  echo DONE
  echo "**************************"
  echo "Results written to: $PWD/wrk/products-modified.csv"


}

# Check for input parm
if [[ -z "$1" ]]; then
  echo "Airtable forecast.csv is required"
  printUsage
elif [[ -z "$2" ]]; then
  echo "Squarespace procucts.csv is required"
  printUsage
else
  FORECAST="$1"
  PRODUCTS="$2"

  goMerge

fi