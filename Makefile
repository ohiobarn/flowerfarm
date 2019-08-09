REMOTE_HOST := ftp.ohiobarnflowerfarm.com
REMOTE_WWW_ROOT := ./www
REMOTE_USER := ohiobar1
WP_URL := https://ohiobarnflowerfarm.com
SNAPSHOT_FOLDER := $(shell echo "duplicator-work/flower-snapshot-`date +%Y%m%d`")



####################################################################################
# This makefile is intended to run on the Mile Two network (your PC)
####################################################################################
define install_instructions
	@echo "\n\n"
	@echo "******************************************************************************"
	@echo "Workflow:"
	@echo "******************************************************************************"
	@echo ""
	@echo " 1 - Run 'make dev' and do web development"
	@echo " 2 - Run 'make publish' to publish your change"
	@echo ""
	@echo "\n"
endef

help:
	$(call install_instructions)

dev:
	cp -R data data-backup
	docker-compose up &
	open http://localhost:8080/

publish:
	@mkdir -p ${SNAPSHOT_FOLDER}
	@echo "Select 'duplicator' in the left menu and click the 'Create New' package button"
	@echo "When the package is built select the 'one-click download'"
	@echo "Save in this folder: ${SNAPSHOT_FOLDER}"
	@echo "do 'make install-package' next"

test:
	ssh ${REMOTE_USER}@${REMOTE_HOST} "mv ~/www/wp-config.php ~/www/wp-config.php.back"

install-package: 
	scp -r ${SNAPSHOT_FOLDER} ${REMOTE_USER}@${REMOTE_HOST}:~/${SNAPSHOT_FOLDER}
	ssh ${REMOTE_USER}@${REMOTE_HOST} "mv ~/www/wp-config.php ~/www/wp-config.php.back"
	ssh ${REMOTE_USER}@${REMOTE_HOST} "cp ~/${SNAPSHOT_FOLDER}/* ./www/"
	open ${WP_URL}/installer.php

ssh-remote:
	ssh ${REMOTE_USER}@${REMOTE_HOST}

dsp-db-info:
	ssh ${REMOTE_USER}@${REMOTE_HOST} "cat ./www/wp-config.php | grep DB_"








