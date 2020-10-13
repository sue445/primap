# primap-function
[Cloud Functions](https://cloud.google.com/functions) for primap

There are the following functions.

* Get shops from [PrismDB](https://prismdb.takanakahiko.me/)
* Get geography from shop address and save to [Cloud Firestore](https://firebase.google.com/docs/firestore)

## Requirement API keys
Register followings from https://console.cloud.google.com/apis/credentials

* `GOOGLE_MAPS_API_KEY`
  * Application restrictions: None
  * API restrictions: Geocoding API

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
