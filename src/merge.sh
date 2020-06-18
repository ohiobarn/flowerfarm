doMerge() {

  echo "Using Forecast: $FORECAST"
  echo "Using Crop plan: $CROP_PLAN"
  echo "Using Products: $PRODUCTS"
  echo "Using Work Dir: $WORK_DIR"

  cd $WORK_DIR
  rm *.json *.txt

  #
  # csv to json
  #
  csvjson $FORECAST | jq > forecast.json
	csvjson $CROP_PLAN | jq > Crop_plan.json
	csvjson $PRODUCTS | jq > products.json

  # 
  # Data Scrub 
  #
  sed -i "" 's/Est. Yield/estYield/g' Crop_plan.json #remove space from json key
  sed -i "" 's/ Stems//g' Crop_plan.json # remove " Stems" from estYield, this will make it a number
  cat products.json  | iconv -f utf-8-mac -t utf-8 | sed 'y/āáǎàēéěèīíǐìōóǒòūúǔùǖǘǚǜĀÁǍÀĒÉĚÈĪÍǏÌŌÓǑÒŪÚǓÙǕǗǙǛ/aaaaeeeeiiiioooouuuuüüüüAAAAEEEEIIIIOOOOUUUUÜÜÜÜ/' > products-scrub.json
  cat Crop_plan.json | iconv -f utf-8-mac -t utf-8 | sed 'y/āáǎàēéěèīíǐìōóǒòūúǔùǖǘǚǜĀÁǍÀĒÉĚÈĪÍǏÌŌÓǑÒŪÚǓÙǕǗǙǛ/aaaaeeeeiiiioooouuuuüüüüAAAAEEEEIIIIOOOOUUUUÜÜÜÜ/' > Crop_plan-scrub.json
  mv products.json products-backup.json
  mv Crop_plan.json Crop_plan-backup.json
  mv products-scrub.json products.json
  mv Crop_plan-scrub.json Crop_plan.json
 
  # init with the start of an array
  echo "[" > products-updated.json
  update_count=0;

  
  # Loop through product records and find the corresponding forcast record
  # update the product record with the info from the forcast record
  # write new product record to products-updated.json
  #
  product_count=$(jq ". | length" products.json)
  for (( p=0; p<$product_count; p++ ))
  do

    # product record
    jq -c .[$p] products.json > product-record.json
    product_sku=$(jq --raw-output .SKU product-record.json)

    #
    # find corresponding forecast record if it exist
    # 
    eval "jq --raw-output '.[] | select (.SKU == \"$product_sku\")' forecast.json > forecast-record.json"

    forecast_sku=$(jq --raw-output  '.SKU' forecast-record.json)
    forecast_variety=$(jq --raw-output  '.Variety' forecast-record.json)
    forecast_plant=$(jq --raw-output  '.Plant' forecast-record.json)
    forecast_week1=$(jq --raw-output  '."This Week"' forecast-record.json)
    forecast_notes=$(jq --raw-output  '.Notes' forecast-record.json)

    #
    # If found update product record with info from forcast record
    #
    if [ "$forecast_sku" == "$product_sku" ]; then
      update_count=$((update_count+1))
      echo " "
      echo "***************************************************"
      echo "Found sku: $product_sku at index $p" update_count $update_count
      echo "***************************************************"
      echo " "
      echo "Using:"
      echo forecast_sku [$forecast_sku]
      echo forecast_variety [$forecast_variety]
      echo forecast_plant [$forecast_plant]
      echo forecast_week1 [$forecast_week1]
      echo forecast_notes [$forecast_notes]
      echo " "

      yq w product-record.json Description "$forecast_plant - $forecast_variety (forcast: $forecast_week1)  Notes: $forecast_notes" --tojson --inplace
    fi

    #
    # write product record out to products-updated.json
    #
    if [ "$p" -gt "0" ]; then
      printf "," >> products-updated.json
    fi
    cat product-record.json >>  products-updated.json
  done

  # close up the array
  echo "]" >>  products-updated.json

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
    ] | @csv' products-updated.json  > products-updated.csv

  #
  # Add header row
  #
  echo '"Product ID [Non Editable]","Variant ID [Non Editable]","Product Type [Non Editable]","Product Page","Product URL","Title","Description","SKU","Option Name 1","Option Value 1","Option Name 2","Option Value 2","Option Name 3","Option Value 3","Price","Sale Price","On Sale","Stock","Categories","Tags","Weight","Length","Width","Height","Visible","Hosted Image URLs"' \
    | cat - products-updated.csv > temp && mv temp products-updated.csv
}

printUsage() {

  echo ""
  echo "This script will merge a tend csv export with a squarespace products export:"
  echo ""
  echo "Usage:"
  echo "  merge.sh [path/to/csv_export/airtable/forecast.csv] [path/to/csv_export/tend/Crop_plan.csv] [path/csv_export/squarespace/products.csv]"
  echo ""
  echo "Example:"
  echo "  $ ./merge.sh tend/cvs_export/Crop_plan.csv squarespace/products.csv"
  echo ""
}   

#  set -x

# Check for input parm
if [[ -z "$1" ]]; then
  echo "Airtable forcast.csv is required"
  printUsage
elif [[ -z "$2" ]]; then
  echo "Tend Crop_plan.csv is required"
  printUsage
elif [[ -z "$3" ]]; then
  echo "Squarespace procucts.csv is required"
  printUsage
else 
  FORECAST=$1
  CROP_PLAN=$2
  PRODUCTS=$3
  WORK_DIR="./wrk"
  mkdir -p $WORK_DIR

  doMerge

fi
