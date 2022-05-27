# Orders Service b

Stores orders in its database. Provides an endpoint to receive batches of orders and then persist them on its database.

The service has two other endpoints for

- Getting the list of orders paginated and with optional filters
  - per country, per date, and weight limit
- Get the total number of orders per country
- Get the total weight of orders per country
