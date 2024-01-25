# Run application
* run docker-compose `make build-composer`
* run migration for db `make migration`
* run app `make app-start`
# Test application
* for test app run `make test`


#Example urls
```shell script
curl --location --request POST 'http://127.0.0.1:8090/payment/account' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name":"Alex"
}
'
```

```shell script
curl --location --request POST 'http://127.0.0.1:8090/payment/account/1' \
--header 'Source-Type: game' \
--header 'Content-Type: application/json' \
--data-raw '{
    "state": "lost", 
    "amount": "30.15", 
    "transactionId": "eac43bed-82a3-420d-8b09-c5820cf9dda2"
}
'
```
# Swagger Docs
- **[documentation](http://127.0.0.1:8090/documentation)**

# Prometheus Metrics
- **[metrics](http://127.0.0.1:8090/metrics)**

# Evn files
* .env - for main application

### ENV description
```
# APP ...
APP_PORT=8080 - Port for app
APP_VERSION=1.0
TIME_ROLLBACK_TX=120 # value in seconds for rollbaclk transaction 
DB_POSTGRE_HOST=postgres
DB_POSTGRE_PORT=5432
DB_POSTGRE_USER=admin
DB_POSTGRE_PASSWORD=secret
DB_POSTGRE_DATABASE=postgres
``` 

