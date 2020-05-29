doMerge() {

  echo "Using Crop plan: $CROP_PLAN"
  echo "Using Products: $PRODUCTS"
  echo "Using Work Dir: $WORK_DIR"

  # Convert crop plan csv to json
	csvjson $CROP_PLAN | jq > $WORK_DIR/Crop_plan.json

  # Convert products csv to json
	csvjson $PRODUCTS | jq > $WORK_DIR/products.json

  # get length of array in each file
  variety_count=$(jq '. | length' $WORK_DIR/Crop_plan.json)
  product_count=$(jq '. | length' $WORK_DIR/products.json)
  echo "Number of varieties: $variety_count"
  echo "Number of products: $product_count"

  # loop through each product and update with infrormation from crop plan
  for (( v=0; v<$variety_count; v++ ))
  do
    # hack to remove space from json key
    sed -i "" 's/Est. Yield/estYield/g' wrk/Crop_plan.json

    crop=$(jq --raw-output .[$v].Crop $WORK_DIR/Crop_plan.json)
    est_yield=$(jq --raw-output .[$v].estYield $WORK_DIR/Crop_plan.json)
    variety_full=$(jq --raw-output .[$v].Variety $WORK_DIR/Crop_plan.json)
    variety_short=$(echo $variety_full | cut -d"|" -f1)
    variety_sku=$(echo $variety_full | cut -d"|" -f2)

    # trim whitespace
    variety_sku=$(echo $variety_sku | tr -d '[:space:]')

    # echo "variety $v"
    # echo variety_sku $variety_sku
    # echo "crop $crop"
    # echo "variety_full $variety_full"
    # echo "variety_short $variety_short"
    # echo "est_yield $est_yield"
    # echo "crop $crop"
    # echo "------------------------------------------"

    # find product by sku
    for (( p=0; p<$product_count; p++ ))
    do
      product_sku=$(jq --raw-output .[$p].SKU $WORK_DIR/products.json)
      # echo compare [$variety_sku] to [$product_sku]
      if [ "$variety_sku" == "$product_sku" ]; then
        echo "Found sku: $product_sku"
      fi
    done

  done

}

printUsage() {

  echo ""
  echo "This script will merge a tend csv export with a squarespace products export:"
  echo ""
  echo "Usage:"
  echo "  merge.sh [path/to/tend/csv_export/Crop_plan.csv] [path/squarespace/export/products.csv]"
  echo ""
  echo "Example:"
  echo "  $ ./merge.sh tend/cvs_export/Crop_plan.csv squarespace/products.csv"
  echo ""
}

# Check for input parm
if [[ -z "$1" ]]; then
  echo "Tend Crop_plan.csv is required"
  printUsage
elif [[ -z "$2" ]]; then
  echo "Squarespace procucts.csv is required"
  printUsage
else 
  CROP_PLAN=$1
  PRODUCTS=$2
  WORK_DIR="./wrk"
  mkdir -p $WORK_DIR

  doMerge

fi
