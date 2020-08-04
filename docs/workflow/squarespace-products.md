# Squarespace Products Sync Workflow


Reference:

* [OBFF website](https://ohiobarnflowerfarm.com/)
* [Airtable](https://airtable.com/)

## Preorder

When customers go to the [OBFF preorder page](https://ohiobarnflowerfarm.com/preorder) and they are presented with an Airtable form. The form will create records in the `Preorder` table.  

> todo - need to implement a notification when preorders are created

The `Preorder` table has a status column. When a preorder is taken care of the status should be moved to `done`.  The record can be deleted at that point.

## Forecast

After a survey of the beds an estimate of the number of stems that can be harvested in a week, two weeks and beyond are created. Estimate week two as if all stems in week one are sold to customers. Do the same for future estimates.  Enter the estimated number of stems by SKU in to the `Forecast` table in Airtable.

An Airtable public view if the information in the `Forecast` table can be viewed by customers on the [OBFF forecast page](https://ohiobarnflowerfarm.com/forecast)

## Square Space Products

The `Square Space Products` table is an import of the products exported from Squarespace.

> todo - I don't think this is needed, confirm and remove if true

## Prerequisites

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
