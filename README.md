# flowerfarm

The Ohio Barn Flower Farm

## Sandbox Setup

* Download from https://wordpress.org/download/

## Development

```bash
docker-compose up
open http://localhost:8080/
```

## Publish to Bluehost

```bash
ssh ohiobar1@ftp.ohiobarnflowerfarm.com 

cd wordpress
scp -r duplicator/* ohiobar1@ftp.ohiobarnflowerfarm.com:./www
```

## Reference

[wordpress development workflow](https://www.technouz.com/4613/ultimate-wordpress-website-development-workflow/)

## Bluehost db settings

To discover db info `cat ~/www/wp-config.php`

```php
// ** MySQL settings ** //
/** The name of the database for WordPress */
define( 'DB_NAME', 'ohiobar1_afc' );

/** MySQL database username */
define( 'DB_USER', 'ohiobar1_afc' );

/** MySQL database password */
define( 'DB_PASSWORD', '<not shown>' );

/** MySQL hostname */
define( 'DB_HOST', 'localhost' );
```


* Please create a new database and user for your staging website and update the details in wp-config file in your staging folder. 
* Then update the home and site URL of the staging website as http://ohiobarnflowerfarm.com/staging in wp-options option in your new database. 
* Once you finish the website you can move the staging files to your main website files and then update the site URL as  http://ohiobarnflowerfarm.com/ in your database 
* You can create the new database and user in My SQL database option
* Please create user and database and give all priviliiages.

Since you are using Ubuntu, all you need to do is just to add a file in your home directory and it will disable the mysqldump password prompting. This is done by creating the file ~/.my.cnf (permissions need to be 600)

```bash
[mysqldump]
user=ohiobar1_dbuser
password=<pwd>

[mysql]
user=ohiobar1_dbuser
password=<pwd>
```
