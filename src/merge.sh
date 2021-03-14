doMerge() {

  echo "Using Forecast: $FORECAST"
  echo "Using Products: $PRODUCTS"
  echo "Using Work Dir: $WORK_DIR"

  cd $WORK_DIR
  rm *.json *.txt

  #
  # csv to json
  #
  csvjson $FORECAST | jq > forecast.json
	csvjson $PRODUCTS | jq > products.json

  # 
  # Data Scrub 
  #
  cat products.json  | iconv -f utf-8-mac -t utf-8 | sed 'y/āáǎàēéěèīíǐìōóǒòūúǔùǖǘǚǜĀÁǍÀĒÉĚÈĪÍǏÌŌÓǑÒŪÚǓÙǕǗǙǛ/aaaaeeeeiiiioooouuuuüüüüAAAAEEEEIIIIOOOOUUUUÜÜÜÜ/' > products-scrub.json
  mv products.json products-backup.json
  mv products-scrub.json products.json
 
  # init with the start of an array
  echo "[" > products-updated.json
  update_count=0;

  #####################################################################################
  #                   ###  Forecast UPDATES   ###
  # Loop through product records and find the corresponding forecast record
  # update the product record with the info from the forecast record
  # write new product record to products-updated.json
  #####################################################################################
  
  product_count=$(jq ". | length" products.json)
  printf "\nLook at $product_count product records, check for updates...\n"
  for (( p=0; p<$product_count; p++ ))
  do
    # product record
    jq -c .[$p] products.json > product-record.json
    product_sku=$(jq --raw-output .SKU product-record.json)
    product_title=$(jq --raw-output .Title product-record.json)
    product_description=$(jq --raw-output .Description product-record.json)
    product_visible=$(jq --raw-output .Visible product-record.json)

    #
    # find corresponding forecast record if it exist
    # 
    eval "jq --raw-output '.[] | select (.SKU == \"$product_sku\")' forecast.json > forecast-record.json"

    forecast_sku=$(jq --raw-output  '.SKU' forecast-record.json)
    forecast_variety=$(jq --raw-output  '.Variety' forecast-record.json)
    forecast_crop=$(jq --raw-output  '.Crop' forecast-record.json)
    forecast_week1=$(jq --raw-output  '."This Week"' forecast-record.json)
    forecast_week2=$(jq --raw-output  '."Next Week"' forecast-record.json)
    forecast_week3=$(jq --raw-output  '."Future"' forecast-record.json)
    forecast_notes=$(jq --raw-output  '.Notes' forecast-record.json | sed 's/null//g')
    forecast_show=$(jq --raw-output  '.Show' forecast-record.json) 
    forecast_stems_per_bunch=$(jq --raw-output  '."Stems per Bunch"' forecast-record.json)
    forecast_grower=$(jq --raw-output  '.Grower' forecast-record.json) 

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
      new_product_title=$(echo "$forecast_crop - $forecast_variety")
      
      dtitle="$new_product_title | $forecast_sku"
      spb="<i>$forecast_stems_per_bunch stems per bunch</i>"
      dforecast="<hr><b>Forecast:</b> <br>This week: $forecast_week1 <br>Next week: $forecast_week2 <br>Future: $forecast_week3<br>$spb<hr>"
      dgrower="Grower: $forecast_grower"
      dnotes="$forecast_notes"
      new_product_description=$(printf "<p>%s<br>%s<br>%s<br>%s</p>" "$dtitle" "$dforecast" "$dnotes" "$dgrower")


      if [ "$forecast_show" == "checked" ]; then
        new_product_visible="true"
      else
        new_product_visible="false"
      fi

      if [ "$product_description" != "$new_product_description" ]; then
        printf "CANGE - Description:\n"
        printf "   BEFORE: $product_description\n"
        printf "   AFTER.: $new_product_description \n"
      fi 

      if [ "$product_visible" != "$new_product_visible" ]; then
        printf "CANGE - Visible (before|after): "
        printf "$product_visible | $new_product_visible \n"
      fi

      if [ "$product_title" != "$new_product_title" ]; then
        printf "CANGE - Title (before|after): "
        printf "$product_title | $new_product_title \n"
      fi

      yq eval ".Title = \"${new_product_title}\"" product-record.json --tojson --inplace
      yq eval ".Description = \"$new_product_description\"" product-record.json --tojson --inplace
      yq eval ".Visible = \"$new_product_visible\"" product-record.json --tojson --inplace

    fi

    #
    # write product record out to products-updated.json
    #
    if [ "$p" -gt "0" ]; then
      printf "," >> products-updated.json
    fi
    cat product-record.json >>  products-updated.json

    # Hack, remove " [Non Editable]" from key names, I cant deal with it
    sed -i '' "s/ \[Non Editable\]//g" products-updated.json 

  done


  #####################################################################################
  #                   ###   NEW Products   ###
  # Loop through each forecast record and if not in prooducts then it needs added
  ######################################################################################
  
  forecast_count=$(jq ". | length" forecast.json)
  printf "\nLook at $forecast_count forecast records, check if any are new...\n"
  for (( fc=0; fc<$forecast_count; fc++ ))
  do
    jq -c .[$fc] forecast.json > forecast-record.json
    forecast_sku=$(jq --raw-output  '.SKU' forecast-record.json)
    forecast_variety=$(jq --raw-output  '.Variety' forecast-record.json)
    forecast_crop=$(jq --raw-output  '.Crop' forecast-record.json)
    forecast_week1=$(jq --raw-output  '."This Week"' forecast-record.json)
    forecast_week2=$(jq --raw-output  '."Next Week"' forecast-record.json)
    forecast_week3=$(jq --raw-output  '."Future"' forecast-record.json)
    forecast_notes=$(jq --raw-output  '.Notes' forecast-record.json | sed 's/null//g')
    forecast_show=$(jq --raw-output  '.Show' forecast-record.json) 
    forecast_stems_per_bunch=$(jq --raw-output  '."Stems per Bunch"' forecast-record.json)

    #
    # find corresponding product record if it exist
    # 
    eval "jq --raw-output '.[] | select (.SKU == \"$forecast_sku\")' products.json > product-record.json"
    product_sku=$(jq --raw-output .SKU product-record.json)

    if [ "$forecast_sku" != "$product_sku" ]; then
      echo New product found and will be added sku: $forecast_sku
      update_count=$((update_count+1))
      #
      # Set new product field values
      #
      new_product_sku=$(echo "$forecast_sku")
      new_product_title=$(echo "$forecast_crop - $forecast_variety")
      new_product_description=$(echo "This is a new product. The stock and price need set. This description will be replaced with the actual description the next update run.")
      
      # hardcode initial values
      new_product_price=999
      new_product_stock=999
      new_product_tags=mrfc


      if [ "$forecast_show" == "checked" ]; then
        new_product_visible="true"
      else
        new_product_visible="false"
      fi

      #
      # Add to product record
      #
      yq eval -n ".SKU = \"${new_product_sku}\"" --tojson > product-record.json
      yq eval ".Title = \"${new_product_title}\"" product-record.json --tojson --inplace
      yq eval ".Description = \"${new_product_description}\"" product-record.json --tojson --inplace
      yq eval ".Visible = \"${new_product_visible}\"" product-record.json --tojson --inplace
    
      yq eval '.Product~Type = "PHYSICAL"' product-record.json --tojson --inplace
      yq eval '.Product~Page = "mrfc"' product-record.json --tojson --inplace
      yq eval ".Product~URL = \"${new_product_sku}\"" product-record.json --tojson --inplace
      
      yq eval ".Price = \"${new_product_price}\"" product-record.json --tojson --inplace
      yq eval '.Sale~Price = "0"' product-record.json --tojson --inplace
      yq eval '.On~Sale = "FALSE"' product-record.json --tojson --inplace

      yq eval ".Stock = \"${new_product_stock}\"" product-record.json --tojson --inplace
      yq eval ".Tags = \"${new_product_tags}\"" product-record.json --tojson --inplace

      # hack, I don't know how to handle spaces in keys
      sed -i '' 's/~/ /g' product-record.json

      #
      # write product record out to products-updated.json
      # note, always prepend , assume records exist from first loop
      #
      printf "," >> products-updated.json
      cat product-record.json >>  products-updated.json
    fi


  done

  # close up the array
  echo "]" >>  products-updated.json

  #
  # json to csv
  #
  jq -r '.[] | [
      ."Product ID", 
      ."Variant ID",
      ."Product Type",
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
  echo "This script will merge an airtable forecast export with a squarespace products export:"
  echo ""
  echo "Usage:"
  echo "  merge.sh [path/to/csv_export/airtable/forecast.csv] [path/csv_export/squarespace/products.csv]"
  echo ""
  echo "Example:"
  echo "  $ ./merge.sh airtable/forecast.csv squarespace/products.csv"
  echo ""
}   

#  set -x

# Check for input parm
if [[ -z "$1" ]]; then
  echo "Airtable forecast.csv is required"
  printUsage
elif [[ -z "$2" ]]; then
  echo "Squarespace procucts.csv is required"
  printUsage
else 
  FORECAST=$1
  PRODUCTS=$2
  WORK_DIR="./wrk"
  mkdir -p $WORK_DIR

  doMerge

fi
