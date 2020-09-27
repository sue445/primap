# primap
[![Build Status](https://github.com/sue445/primap/workflows/build-server/badge.svg?branch=master)](https://github.com/sue445/primap/actions?query=workflow%3Abuild-server)
[![Build Status](https://github.com/sue445/primap/workflows/build-frontend/badge.svg?branch=master)](https://github.com/sue445/primap/actions?query=workflow%3Abuild-frontend)
[![Build Status](https://github.com/sue445/primap/workflows/deploy/badge.svg?branch=master)](https://github.com/sue445/primap/actions?query=workflow%3Adeploy)

## Develop
### Requirement API keys
Register followings from https://console.cloud.google.com/apis/credentials

* `GOOGLE_MAPS_API_KEY`
  * Application restrictions: None
  * API restrictions: Geocoding API
* `REACT_APP_GOOGLE_BROWSER_API_KEY`
  * Application restrictions: HTTP referrers (web sites)
  * Website restrictions: `localhost:5000/*` (local), `primap.web.app/*` (production)
  * API restrictions: Cloud Firestore API, Maps JavaScript API

### Run server
```bash
cp .env.examle .env
vi .env

docker-compose up --build
```

open http://localhost:8000

### Run firebase hosting
```bash
firebase serve --only hosting
```

### Testing
```bash
docker-compose build server && docker-compose run --rm server bash -c 'firebase --project test emulators:exec --only firestore,pubsub "make test"'
```
