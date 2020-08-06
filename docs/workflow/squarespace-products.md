# Squarespace Products Sync Workflow


Reference:

* [OBFF website](https://ohiobarnflowerfarm.com/)
* [Airtable](https://airtable.com/)

## AirTable Public Views

This section describes how to update the records in public views. These views are hosted in airtable and are shared or embedded on the OBFF website.

| View                                                             | Location                             | Referenced
|------------------------------------------------------------------|--------------------------------------|------------
[Estimated Yield (public)](https://airtable.com/shr68zqWOk0BYXgMj) | `FlowerDB2020:Forecast`              | [OBFF web](https://ohiobarnflowerfarm.com/forecast#estimated-yield)
[Avail Plan (public)](https://airtable.com/shrbTArfM48DQykzV)      | `FlowerDB2020:Tend Crop Plan Ext`    | [OBFF web](https://ohiobarnflowerfarm.com/forecast#availability-at-a-glance)

### Estimated Yield (public)

1. **Survey** - After a survey of the beds an estimate of the number of stems that can be harvested in a week, two weeks and beyond are created. Estimate week two as if all stems in week one are sold to customers. Do the same for future estimates.  
1. **Edit Records**  - Edit or crate the records in the [Forecast(Back Office)](https://airtable.com/tblUwgrBBwsPngBVO/viwC45pNTVuBV6yOk?blocks=hide) table in Airtable. 
    * Select the estimated number of stems by SKU in the `This Week`, `Next Week` and `Future` columns
    * Adjust the `Stems per Bunch` if needed
    * The `Starting` date is for your reference, it is not shown to the customer
    * The `Notes` will show on the product detail page on the OBFF website.  
1. **Done** - The updated records will automatically show in in  the public views

### Avail Plan (public)

**todo** - need to document the work flow for exporting from tend and importing into airtable

---

## Square Space Inventory

Todo - need to describe better the process that is something like:

* export [squarespace products](https://ohiobarnflowerfarm.squarespace.com/config/commerce/inventory) 
* export Tend `Crop_plan`
* export airtable 'Forecast.Back Office`
* run make ss-merge
* import  [squarespace products](https://ohiobarnflowerfarm.squarespace.com/config/commerce/inventory) 


### Prerequisites

This workflow uses the following tools for working with CSV files.

See the [csvkit doc here](https://csvkit.readthedocs.io/en/latest/)

```bash
sudo pip install csvkit
```

Below is example usage:

```bash
csvjson tend/csv_export/Crop_plan.csv | jq > test.json
in2csv test.json 
```
