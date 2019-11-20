# Glance-API

## Available features:

* full users support, included CRUD, admin's endpoint and some others temporary features
* events features, included CRUD, handling some user events, filtering, etc
* feature based on the maintenance logic
* feature based on the parcels logic
* feature based on the wellbeing requests logic


### Pre-requirements
* Go v1.11.4+
* Docker version 18.09.6+

### How to build?
`$git pull`

`$docker-compose up --build -d`

#### HTTP/HTTPS options

By default, the application is running on the HTTP network scheme, but you can set GLANCE_API_SSL to `true` and then set GLANCE_API_SERVER_CERT_PATH and GLANCE_API_SERVER_CERT_KEY.  
If you don't have SSL certs you can generate self-signed certs for development purpose.

After these actions, your application can be run on HTTPS network scheme without NGINX, Apache, etc. 
#### Environment variables

Authentication configuration
```
GLANCE_AUTHENTICATION_SECRET(required string)
GLANCE_AUTHENTICATION_ALGORITHM(optional, default HS256)
```
DB configuration
```
GLANCE_DATABASE_NAME(required string)
GLANCE_DATABASE_HOST(required string)
GLANCE_DATABASE_PORT(required integer)
GLANCE_DATABASE_USER(required string)
GLANCE_DATABASE_PASSWORD(required string)
GLANCE_DATABASE_SSL(required string)
```
HTTP/HTTP configuration
```
GLANCE_API_HOST(required string)
GLANCE_API_PORT(required integer)
GLANCE_API_SSL(optional boolean)
GLANCE_API_SERVER_CERT_PATH(optional string)
GLANCE_API_SERVER_CERT_KEY(optional string)
```
Logging configuration
```
API_LOG_LEVEL(optional string, default "debug". available options: trace, debug, info, warning, error, fatal, panic)
```

#### Generate self-signed certificates for development purpose
You can generate self-signed certs by using `generate_certs.sh`  

If you have permission problems with this scrip you should run  in the terminal the following command:  
`chmod +x generate_certs`

Pre-requirements:
*  OpenSSL 1.1.1+

### Swagger documentation
Swagger is running via docker-compose up command and deploying on the 9080 port.
### Go code documentation

You can see documentation inside source code option or you can run web version of documentation by using  
``godoc -http=:{any available port on your PC}``