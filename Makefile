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
	@echo "Wordpress Workflow:"
	@echo "******************************************************************************"
	@echo ""
	@echo " Normally:"
	@echo "   1 - Make changes in staging ${WP_URL}/stag/wp-login.php"
	@echo "   2 - When ready run 'make staging-to-prod' "
	@echo "   3 - Test in prod ${WP_URL}/prod/wp-login.php"
  @echo ""
	@echo " Local Development:"
	@echo "   Overview: "
	@echo "     - cp staging to local then do development"
	@echo "     - cp local to staging then test"
	@echo "     - cp staging to prod "
	@echo " " 
	@echo "   1 - Run 'make make-stag-duplicator-package' and follow instructions to install it locally"
	@ech0 "   2 - Make changes locally"
	@echo "   3 - Run 'make make-local-duplicator-package' and follow instructions to install it in staging"
	@echo "   4 - Test in staging"
	@echo     5 - Run 'staging-to-prod' to publish to prod"
	@echo ""
	@echo "******************************************************************************"
	@echo "mkdocs Workflow:"
	@echo "******************************************************************************"
	@echo ""
	@echo " 1 - mkdocs serve"
	@echo " 2 - mkdocs build --clean"
	@echo " 3 - mkdocs gh-deploy"
	@echo ""	
	@echo "\n"
endef

help:
	$(call install_instructions)

start-dev:
	cp -R data data-backup
	docker-compose up
	open http://localhost:8080/

make-stag-duplicator-package:
	mkdir -p ${SNAPSHOT_FOLDER}
	@echo ""
	@echo " - Open prod ${WP_URL}/stag/wp-login.php"
	@echo " - Select 'duplicator' in the left menu and click the 'Create New' package button"
	@echo " - When the package is built select the 'one-click download'"
	@echo " - Save in this folder: ${SNAPSHOT_FOLDER}"
	@echo " - Run 'make install-local-duplicator-package' and follow the instructions"

make-local-duplicator-package:
	mkdir -p ${SNAPSHOT_FOLDER}
	@echo ""
	@echo " - Select 'duplicator' in the left menu and click the 'Create New' package button"
	@echo " - When the package is built select the 'one-click download'"
	@echo " - Save in this folder: ${SNAPSHOT_FOLDER}"
	@echo " - Run 'make install-stag-duplicator-package' and follow the instructions"

install-local-duplicator-package:
	cp ${SNAPSHOT_FOLDER}/* wordpress/
	@echo ""
	@echo " - Run 'make start-dev'"
	@echo " - Open http://localhost:8080/installer.php (uses db values from compose file)"

install-stag-duplicator-package: 
	scp -r ${SNAPSHOT_FOLDER} ${REMOTE_USER}@${REMOTE_HOST}:~/${SNAPSHOT_FOLDER}
	ssh ${REMOTE_USER}@${REMOTE_HOST} "mv ~/www/stag/wp-config.php ~/www/stag/wp-config.php.back"
	ssh ${REMOTE_USER}@${REMOTE_HOST} "cp ~/${SNAPSHOT_FOLDER}/* ./www/stag"
  @echo " - Open ${WP_URL}/stag/installer.php (uses db values from compose file)"

ssh-remote:
	ssh ${REMOTE_USER}@${REMOTE_HOST}

dsp-db-info:
	@echo "Prod:"
	ssh ${REMOTE_USER}@${REMOTE_HOST} "cat ./www/prod/wp-config.php | grep DB_"
	@echo "Staging:"
	ssh ${REMOTE_USER}@${REMOTE_HOST} "cat ./www/stag/wp-config.php | grep DB_"

local-db-backup:
	rm -rf data-backup/data
	cp -R data data-backup

dump-sites:
	ssh ${REMOTE_USER}@${REMOTE_HOST} "mkdir -p ${MIGRATE_FOLDER}/${YYMMDD}/stag"
	ssh ${REMOTE_USER}@${REMOTE_HOST} "mkdir -p ${MIGRATE_FOLDER}/${YYMMDD}/prod"
	ssh ${REMOTE_USER}@${REMOTE_HOST} "mysqldump ${PROD_DB_NAME} > ${MIGRATE_FOLDER}/${YYMMDD}/${PROD_DB_NAME}.sql"
	ssh ${REMOTE_USER}@${REMOTE_HOST} "mysqldump ${STAGING_DB_NAME} > ${MIGRATE_FOLDER}/${YYMMDD}/${STAGING_DB_NAME}.sql"
	ssh ${REMOTE_USER}@${REMOTE_HOST} "cp -R ~/www/stag/* ${MIGRATE_FOLDER}/${YYMMDD}/stag/"
	ssh ${REMOTE_USER}@${REMOTE_HOST} "cp -R ~/www/prod/* ${MIGRATE_FOLDER}/${YYMMDD}/prod/"
	ssh ${REMOTE_USER}@${REMOTE_HOST} "sed 's/\/stag/\/prod/g' ${MIGRATE_FOLDER}/${YYMMDD}/${STAGING_DB_NAME}.sql > ${MIGRATE_FOLDER}/${YYMMDD}/migrage-staging-to-prod.sql"
	ssh ${REMOTE_USER}@${REMOTE_HOST} "sed 's/\/prod/\/stag/g' ${MIGRATE_FOLDER}/${YYMMDD}/${PROD_DB_NAME}.sql    > ${MIGRATE_FOLDER}/${YYMMDD}/migrage-prod-to-staging.sql"


staging-to-prod: dump-sites
	ssh ${REMOTE_USER}@${REMOTE_HOST} "mysql -h localhost -D ${PROD_DB_NAME} < ${MIGRATE_FOLDER}/${YYMMDD}/migrage-staging-to-prod.sql"
	ssh ${REMOTE_USER}@${REMOTE_HOST} "cp -R ~/www/stag/* ~/www/prod/"
	ssh ${REMOTE_USER}@${REMOTE_HOST} "sed -i 's/ohiobar1_staging/ohiobar1_prod/g' ~/www/prod/wp-config.php"

prod-to-staging: dump-sites
	ssh ${REMOTE_USER}@${REMOTE_HOST} "mysql -h localhost -D ${STAGING_DB_NAME} < ${MIGRATE_FOLDER}/${YYMMDD}/migrage-prod-to-staging.sql"
	ssh ${REMOTE_USER}@${REMOTE_HOST} "cp -R ~/www/prod/* ~/www/stag/"
	ssh ${REMOTE_USER}@${REMOTE_HOST} "sed -i 's/ohiobar1_prod/ohiobar1_staging/g' ~/www/stag/wp-config.php"







