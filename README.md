# flowerfarm

The Ohio Barn Flower Farm

## Sandbox Setup

* Download from https://wordpress.org/download/

## Reference

[wordpress development workflow](https://www.technouz.com/4613/ultimate-wordpress-website-development-workflow/)


## Publish to Bluehost

```bash
ssh ohiobar1@ftp.ohiobarnflowerfarm.com 

cd wordpress
#scp -r wp-content ohiobar1@ftp.ohiobarnflowerfarm.com:./www

scp -r duplicator/* ohiobar1@ftp.ohiobarnflowerfarm.com:./www
```

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