# primap-function
[Cloud Functions](https://cloud.google.com/functions) for primap

## Overview
```mermaid
flowchart TD
  scheduler["Cloud Scheduler"] -- Runs once a day --> cron["CronUpdateShops"]

  cron --> queue1["QueueSaveShop"]
  cron --> queue2["QueueSaveShop"]
  cron --> queue3["QueueSaveShop"]

  queue1 --> firestore[(Cloud Firestore)]
  queue2 --> firestore
  queue3 --> firestore
```

There are the following functions.

* [CronUpdateShops](handler_cron_update_shops.go): Runs once a day with [Cloud Scheduler](https://cloud.google.com/scheduler), and get shops from [PrismDB](https://prismdb.takanakahiko.me/)
* [QueueSaveShop](handler_queue_save_shop.go): Get geography from shop address and save to [Cloud Firestore](https://firebase.google.com/docs/firestore)

## Requirement API keys
Register followings from https://console.cloud.google.com/apis/credentials

* `GOOGLE_MAPS_API_KEY`
  * Application restrictions: None
  * API restrictions: Geocoding API
  
## Variables
Register following keys to [Secret Manager](https://cloud.google.com/secret-manager)

* `GOOGLE_MAPS_API_KEY` **(required)**
* `SENTRY_DSN` (optional)

## Development
### Setup
```bash
cp .env.examle .env
vi .env
```

### Testing
Run one of the following

1. `firebase --project test emulators:exec --only firestore,pubsub "make test"`
    * Requires [Firebase CLI](https://firebase.google.com/docs/cli)
2. `docker-compose up --build`
