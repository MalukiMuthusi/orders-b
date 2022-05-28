# Orders Service b

Stores orders in its database. Provides an endpoint to receive batches of orders and then persist them on its database.

The service has two other endpoints for

- Getting the list of orders paginated and with optional filters
  - per country, per date, and weight limit
- Get the total number of orders per country
- Get the total weight of orders per country

## Configs

```sh
# database credentials
export ORDERS_ENV="dev"
export ORDERS_DB_USER="postgres"
export ORDERS_DB_PWD="password"
export ORDERS_DB_NAME="orders"
export ORDERS_DB_PORT="5432"
export ORDERS_DB_CLOUD=false
export ORDERS_DB_HOST="127.0.0.1"
export ORDERS_DB_TIMEZONE="Africa/Nairobi"
export ORDERS_DB_INSTANCE_CONNECTION_NAME="theta-outrider-342406:us-central1:wallet"

# optionally overwrite
export PORT=8080
```
