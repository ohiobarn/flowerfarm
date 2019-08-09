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