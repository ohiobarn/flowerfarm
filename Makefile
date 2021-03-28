WP_URL := https://ohiobarnflowerfarm.com
NOW := $(shell echo "`date +%Y-%m-%d`")


####################################################################################
# This makefile is intended to run on the Mile Two network (your PC)
####################################################################################
define help_info
	@echo "\n\n"
	@echo "******************************************************************************"
	@echo "Squarespace Workflow:"
	@echo "******************************************************************************"
	@echo ""
	@echo " see:"
	@echo "   1 - https://ohiobarn.github.io/flowerfarm/workflow/squarespace-products/"
	@echo "   2 - make ss-merge"
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
	echo "Backing up products..."
	cp "${PWD}/exports/squarespace/export/products.csv" "${PWD}/exports/squarespace/export/products-${NOW}.csv"
	src/merge.sh ${PWD}/exports/airtable/forecast.csv ${PWD}/exports/squarespace/export/products.csv
