# primap-function
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
```bash
firebase --project test emulators:exec --only firestore,pubsub "make test"
```

or

```bash
docker-compose up --build
```
