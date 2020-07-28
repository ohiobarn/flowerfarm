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

  
  # Loop through product records and find the corresponding forecast record
  # update the product record with the info from the forecast record
  # write new product record to products-updated.json
  #
  product_count=$(jq ". | length" products.json)
  for (( p=0; p<$product_count; p++ ))
  do
    printf "."

    # product record
    jq -c .[$p] products.json > product-record.json
    product_sku=$(jq --raw-output .SKU product-record.json)
    product_title=$(jq --raw-output .Title product-record.json)
    product_description=$(jq --raw-output .Description product-record.json)
    product_visible=$(jq --raw-output .Visible product-record.json)
    product_page=$(jq --raw-output '."Product Page"' product-record.json)

    #
    # find corresponding forecast record if it exist
    # 
    eval "jq --raw-output '.[] | select (.SKU == \"$product_sku\")' forecast.json > forecast-record.json"

    forecast_sku=$(jq --raw-output  '.SKU' forecast-record.json)
    forecast_variety=$(jq --raw-output  '.Variety' forecast-record.json)
    forecast_plant=$(jq --raw-output  '.Plant' forecast-record.json)
    forecast_week1=$(jq --raw-output  '."This Week"' forecast-record.json)
    forecast_week2=$(jq --raw-output  '."Next Week"' forecast-record.json)
    forecast_week3=$(jq --raw-output  '."Future"' forecast-record.json)
    forecast_notes=$(jq --raw-output  '.Notes' forecast-record.json | sed 's/null//g')
    forecast_comment=$(jq --raw-output  '.Comment' forecast-record.json)
    forecast_show=$(jq --raw-output  '.Show' forecast-record.json) 
    forecast_stems_per_bunch=$(jq --raw-output  '."Stems per Bunch"' forecast-record.json)

    #
    # If found update product record with info from forecast record
    #
    if [ "$forecast_sku" == "$product_sku" ]; then
      update_count=$((update_count+1))
      echo " "
      echo "***************************************************"
      echo "Found sku: $product_sku at index $p" update_count $update_count
      echo "***************************************************"

      echo " "
      #
      # Set new product field values
      #
      # desc_part=$(echo $product_description | cut -d'|' -f 1)
      # forecast_part=$(echo $product_description | cut -d'|' -f 2)
      
      new_product_title=$(echo "$forecast_plant - $forecast_variety")

      # if [[ "$product_description" == *"|"* ]]; then
        new_product_description=$(echo "<p>$new_product_title | $forecast_sku <br><hr><b>Forecast:</b><br>This week: $forecast_week1 / Next week: $forecast_week2 / Future: $forecast_week3<br> $forecast_stems_per_bunch stems per bunch<br>$forecast_notes</p>")
      # else
      #   echo "######################################## WARNING ################################"
      #   echo "NO PIPE found skipping forecast for this product, product description not changed..."
      #   new_product_description=$product_description        
      # fi
      echo "Descriptions - before/after:"
      echo "BEFORE: $product_description"
      echo "AFTER.: $new_product_description"
      echo " "
  
      if [ "$forecast_show" == "true" ]; then
        new_product_visible="true"
      else
        new_product_visible="false"
      fi
      echo "Visible - before/after:"
      echo "BEFORE: $product_visible"
      echo "AFTER.: $new_product_visible"
      echo " "

      echo "Title - before/after:"
      echo "BEFORE: $product_title"
      echo "AFTER.: $new_product_title"
      echo " "

      yq w product-record.json Title "$new_product_title" --tojson --inplace
      yq w product-record.json Description "$new_product_description" --tojson --inplace
      yq w product-record.json Visible "$new_product_visible" --tojson --inplace

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
  echo "Airtable forecast.csv is required"
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
