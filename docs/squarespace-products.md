# Squarespace Products Workflow

The workflow for updating the products in Squarespace from Tend and Airtable.

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
