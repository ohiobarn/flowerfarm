WP_URL := https://ohiobarnflowerfarm.com
YYMMDD := $(shell echo "`date +%Y%m%d`")
<<<<<<< HEAD
VARIETY_LIST=$(shell echo "`yq r tend/csv_export/Crop_plan.json [].Variety`")
=======
>>>>>>> master




####################################################################################
# This makefile is intended to run on the Mile Two network (your PC)
####################################################################################
define help_info
	@echo "\n\n"
	@echo "******************************************************************************"
	@echo "Squarespace Workflow:"
	@echo "******************************************************************************"
	@echo ""
	@echo " todo:"
	@echo "   1 - todo"
	@echo "   2 - todo"
	@echo "   3 - todo ${YYMMDD}"
  @echo ""
	@echo "******************************************************************************"
	@echo "mkdocs Workflow:"
	@echo "******************************************************************************"
	@echo ""
	@echo " 1 - mkdocs serve"
	@echo " 2 - mkdocs build --clean"
	@echo " 3 - mkdocs gh-deploy"
	@echo ""	
	@echo ""
endef

help:
	$(call help_info)

ss-merge:
	# Convert crop plan to json then to yaml
	#rm -rf tend/wrk
	# mkdir -p wrk
	src/merge.sh ${PWD}/exports/airtable/forecast.csv ${PWD}/exports/tend/csv_export/Crop_plan.csv ${PWD}/exports/squarespace/export/products.csv

