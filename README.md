# primap
[![Build Status](https://github.com/sue445/primap/workflows/build-server/badge.svg?branch=master)](https://github.com/sue445/primap/actions?query=workflow%3Abuild-server)
[![Build Status](https://github.com/sue445/primap/workflows/build-frontend/badge.svg?branch=master)](https://github.com/sue445/primap/actions?query=workflow%3Abuild-frontend)
[![Build Status](https://github.com/sue445/primap/workflows/deploy-server/badge.svg?branch=master)](https://github.com/sue445/primap/actions?query=workflow%3Adeploy-server)
[![Build Status](https://github.com/sue445/primap/workflows/deploy-frontend/badge.svg?branch=master)](https://github.com/sue445/primap/actions?query=workflow%3Adeploy-frontend)

## Develop
### Requirement API keys
Register followings from https://console.cloud.google.com/apis/credentials

* `REACT_APP_GOOGLE_BROWSER_API_KEY`
  * Application restrictions: HTTP referrers (web sites)
  * Website restrictions: `localhost:5000/*` (local), `primap.web.app/*` (production)
  * API restrictions: Cloud Firestore API, Maps JavaScript API

### Run server
```bash
cp .env.examle .env
vi .env
```

open http://localhost:8000
