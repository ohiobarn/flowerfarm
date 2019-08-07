REMOTE_HOST := ftp.ohiobarnflowerfarm.com
REMOTE_WWW_ROOT := ./www
REMOTE_USER := ohiobar1
Wp_URL := https://ohiobarnflowerfarm.com



####################################################################################
# This makefile is intended to run on the Mile Two network (your PC)
####################################################################################
define install_instructions
	@echo "\n\n"
	@echo "******************************************************************************"
	@echo "todo:"
	@echo "******************************************************************************"
	@echo ""
	@echo "Step 1. todo
	@echo "\n"
endef

help:
	$(call install_instructions)

ssh-remote:
	ssh ${REMOTE_USER}@${REMOTE_HOST}

dsp-db-info:
	ssh ${REMOTE_USER}@${REMOTE_HOST} "cat ./www/wp-config.php | grep DB_"

put-duplicator-package: 
	#scp -r duplicator/* ${REMOTE_USER}@${REMOTE_HOST}:${REMOTE_WWW_ROOT}
	ssh ${REMOTE_USER}@${REMOTE_HOST} "mv ./www/wp-config.php ./www/wp-config.php.back"
	@echo "now open ${WP_URL}/installer.php"







