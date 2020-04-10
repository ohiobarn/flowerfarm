WP_URL := https://ohiobarnflowerfarm.com
YYMMDD := $(shell echo "`date +%Y%m%d`")




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

squarespace-todo:
	echo todo

