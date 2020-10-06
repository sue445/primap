# primap
[![Build Status](https://github.com/sue445/primap/workflows/function-build/badge.svg?branch=master)](https://github.com/sue445/primap/actions?query=workflow%3Afunction-build)
[![Build Status](https://github.com/sue445/primap/workflows/frontend-deploy/badge.svg?branch=master)](https://github.com/sue445/primap/actions?query=workflow%3Afrontend-deploy)
[![Build Status](https://github.com/sue445/primap/workflows/frontend-build/badge.svg?branch=master)](https://github.com/sue445/primap/actions?query=workflow%3Afrontend-build)
[![Build Status](https://github.com/sue445/primap/workflows/function-depoy/badge.svg?branch=master)](https://github.com/sue445/primap/actions?query=workflow%3Afunction-depoy)

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
