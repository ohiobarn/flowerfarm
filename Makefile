REMOTE_HOST := ftp.ohiobarnflowerfarm.com
REMOTE_WWW_ROOT := ./www
REMOTE_USER := ohiobar1
WP_URL := https://ohiobarnflowerfarm.com
SNAPSHOT_FOLDER := $(shell echo "duplicator-work/flower-snapshot-`date +%Y%m%d`")
MIGRATE_FOLDER := $(shell echo "~/migrate-work")
YYMMDD := $(shell echo "`date +%Y%m%d`")
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
	ssh ${REMOTE_USER}@${REMOTE_HOST} "cat ./www/prod/wp-config.php | grep DB_"
	@echo "Staging:"
	ssh ${REMOTE_USER}@${REMOTE_HOST} "cat ./www/staging/wp-config.php | grep DB_"

local-db-backup:
	rm -rf data-backup/data
	cp -R data data-backup

dump-sites:
	ssh ${REMOTE_USER}@${REMOTE_HOST} "mkdir -p ${MIGRATE_FOLDER}/${YYMMDD}/staging"
	ssh ${REMOTE_USER}@${REMOTE_HOST} "mkdir -p ${MIGRATE_FOLDER}/${YYMMDD}/prod"
	ssh ${REMOTE_USER}@${REMOTE_HOST} "mysqldump ${PROD_DB_NAME} > ${MIGRATE_FOLDER}/${YYMMDD}/${PROD_DB_NAME}.sql"
	ssh ${REMOTE_USER}@${REMOTE_HOST} "mysqldump ${STAGING_DB_NAME} > ${MIGRATE_FOLDER}/${YYMMDD}/${STAGING_DB_NAME}.sql"
	ssh ${REMOTE_USER}@${REMOTE_HOST} "cp -R ~/www/staging/* ${MIGRATE_FOLDER}/${YYMMDD}/staging/"
	ssh ${REMOTE_USER}@${REMOTE_HOST} "cp -R ~/www/prod/* ${MIGRATE_FOLDER}/${YYMMDD}/prod/"
	ssh ${REMOTE_USER}@${REMOTE_HOST} "sed 's/ohiobarnflowerfarm.com\/staging/ohiobarnflowerfarm.com\/prod/g' ${MIGRATE_FOLDER}/${YYMMDD}/${STAGING_DB_NAME}.sql > ${MIGRATE_FOLDER}/${YYMMDD}/migrage-staging-to-prod.sql"
	ssh ${REMOTE_USER}@${REMOTE_HOST} "sed 's/ohiobarnflowerfarm.com\/prod/ohiobarnflowerfarm.com\/staging/g' ${MIGRATE_FOLDER}/${YYMMDD}/${PROD_DB_NAME}.sql    > ${MIGRATE_FOLDER}/${YYMMDD}/migrage-prod-to-staging.sql"


staging-to-prod:
	ssh ${REMOTE_USER}@${REMOTE_HOST} "mysql -h localhost -D ${PROD_DB_NAME} < ${MIGRATE_FOLDER}/${YYMMDD}/migrage-staging-to-prod.sql"
	ssh ${REMOTE_USER}@${REMOTE_HOST} "cp -R ~/www/staging/* ~/www/prod/"
	ssh ${REMOTE_USER}@${REMOTE_HOST} "sed -i 's/ohiobar1_staging/ohiobar1_prod/g' ~/www/prod/wp-config.php"

prod-to-staging:
	ssh ${REMOTE_USER}@${REMOTE_HOST} "mysql -h localhost -D ${STAGING_DB_NAME} < ${MIGRATE_FOLDER}/${YYMMDD}/migrage-prod-to-staging.sql"
	ssh ${REMOTE_USER}@${REMOTE_HOST} "cp -R ~/www/prod/* ~/www/staging/"
	ssh ${REMOTE_USER}@${REMOTE_HOST} "sed -i 's/ohiobar1_prod/ohiobar1_staging/g' ~/www/staging/wp-config.php"







