REMOTE_HOST := ftp.ohiobarnflowerfarm.com
REMOTE_WWW_ROOT := ./www
REMOTE_USER := ohiobar1
WP_URL := https://ohiobarnflowerfarm.com
SNAPSHOT_FOLDER := $(shell echo "duplicator-work/flower-snapshot-`date +%Y%m%d`")
DUMP_FOLDER := $(shell echo "~/db-dump")
DUMP_NAME := $(shell echo "dump-`date +%Y%m%d`.sql")
PROD_DB_NAME := ohiobar1_prod
STAGING_DB_NAME := ohiobar1_staging



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

mk-snapshot-folder:
	mkdir -p ${SNAPSHOT_FOLDER}
	

publish:
	@mkdir -p ${SNAPSHOT_FOLDER}
	@echo "Select 'duplicator' in the left menu and click the 'Create New' package button"
	@echo "When the package is built select the 'one-click download'"
	@echo "Save in this folder: ${SNAPSHOT_FOLDER}"
	@echo "do 'make install-package' next"

install-package: 
	scp -r ${SNAPSHOT_FOLDER} ${REMOTE_USER}@${REMOTE_HOST}:~/${SNAPSHOT_FOLDER}
	ssh ${REMOTE_USER}@${REMOTE_HOST} "mv ~/www/wp-config.php ~/www/wp-config.php.back"
	ssh ${REMOTE_USER}@${REMOTE_HOST} "cp ~/${SNAPSHOT_FOLDER}/* ./www/"
	open ${WP_URL}/installer.php

cp-staging-prod:
	ssh ${REMOTE_USER}@${REMOTE_HOST} "mv ~/www/wp-config.php ~/www/wp-config.php.back"
	ssh ${REMOTE_USER}@${REMOTE_HOST} "cp ~/www/staging/wp-snapshots/*.zip ./www/"
	ssh ${REMOTE_USER}@${REMOTE_HOST} "cp ~/www/staging/wp-snapshots/*installer.php ./www/installer.php"

ssh-remote:
	ssh ${REMOTE_USER}@${REMOTE_HOST}

dsp-db-info:
	@echo "Prod:"
	ssh ${REMOTE_USER}@${REMOTE_HOST} "cat ./www/wp-config.php | grep DB_"
	@echo "Staging:"
	ssh ${REMOTE_USER}@${REMOTE_HOST} "cat ./www/staging/wp-config.php | grep DB_"

local-db-backup:
	rm -rf data-backup/data
	cp -R data data-backup

dump-db:
	ssh ${REMOTE_USER}@${REMOTE_HOST} "mysqldump ${PROD_DB_NAME} > ${DUMP_FOLDER}/${PROD_DB_NAME}-${DUMP_NAME}"
	ssh ${REMOTE_USER}@${REMOTE_HOST} "mysqldump ${STAGING_DB_NAME} > ${DUMP_FOLDER}/${STAGING_DB_NAME}-${DUMP_NAME}"
	ssh ${REMOTE_USER}@${REMOTE_HOST} "sed 's/ohiobarnflowerfarm.com\/staging/ohiobarnflowerfarm.com\/prod/g' ${DUMP_FOLDER}/${STAGING_DB_NAME}-${DUMP_NAME} > ${DUMP_FOLDER}/migrage-staging-to-prod-${DUMP_NAME}"
	ssh ${REMOTE_USER}@${REMOTE_HOST} "sed 's/ohiobarnflowerfarm.com\/prod/ohiobarnflowerfarm.com\/staging/g' ${DUMP_FOLDER}/${PROD_DB_NAME}-${DUMP_NAME}    > ${DUMP_FOLDER}/migrage-prod-to-staging-${DUMP_NAME}"


staging-to-prod-db:
	ssh ${REMOTE_USER}@${REMOTE_HOST} "mysql -h localhost -D ${PROD_DB_NAME} < ${DUMP_FOLDER}/migrage-staging-to-prod-${DUMP_NAME}"
	ssh ${REMOTE_USER}@${REMOTE_HOST} "cp -R ~/www/staging/* ~/www/prod/"
	ssh ${REMOTE_USER}@${REMOTE_HOST} "sed -i 's/ohiobar1_staging/ohiobar1_prod/g' ~/www/staging/wp-config.php"

prod-to-staging:
	ssh ${REMOTE_USER}@${REMOTE_HOST} "mysql -h localhost -D ${STAGING_DB_NAME} < ${DUMP_FOLDER}/migrage-prod-to-staging-${DUMP_NAME}"
	ssh ${REMOTE_USER}@${REMOTE_HOST} "cp -R ~/www/prod/* ~/www/staging/"
	ssh ${REMOTE_USER}@${REMOTE_HOST} "sed -i 's/ohiobar1_prod/ohiobar1_staging/g' ~/www/staging/wp-config.php"







