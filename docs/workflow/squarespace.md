# Squarespace

This article is mostly about synchronizing data between these two data sources.

* Squarespacce:  [Admin](https://www.squarespace.com/) | [OBFF website](https://ohiobarnflowerfarm.com/)
* [Airtable](https://airtable.com/)

## Inventory

How to update the inventory in squarespace with the data from airtable.

1 Export products from squarespace

* Open the [squarespace products](https://ohiobarnflowerfarm.squarespace.com/config/commerce/inventory) page
* Click *EXPORT ALL* and save to: `flowerfarm/exports/squarespace/products.csv`

2 Export Forecast table from airtable

* Open [airtable](https://airtable.com/) and navigate to the ***MRFC*** base
* For each Synced table tab, select ***Sync now*** [***Forecast (OBFF)***, ***Forecast (GRF)***, ***Forecast (TMG)***]
* Select the ***Forecast (MRFC)*** table --> ***ss-inventory*** view
* Open ***APPS*** and run the ***Populate Forecast (MRFC)*** app. This will combine the Synced tables into one.
* From the `...`  menu select ***Download CSV*** and save to: `flowerfarm/exports/airtable/forecast.csv`

3 Run the "squarespace merge" make target: `make ss-merge`

4 Import updated products into squarespace

* Open the [squarespace products](https://ohiobarnflowerfarm.squarespace.com/config/commerce/inventory) page
* Click *IMPORT* and and select the following: `flowerfarm/wrk/products-updated.csv`
* Disable/uncheck the *Update Product Quantities* checkbox and click ***IMPORT***

## Prerequisites

This workflow uses the following tools for working with CSV files.

See the [csvkit doc here](https://csvkit.readthedocs.io/en/latest/)

```bash
sudo pip install csvkit
```

Below is example usage:

```bash
in2csv test.json 
```
