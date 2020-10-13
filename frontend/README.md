# primap-frontend
Frontend of primap

## Requirement API keys
Register followings from https://console.cloud.google.com/apis/credentials

* `REACT_APP_GOOGLE_BROWSER_API_KEY`
  * Application restrictions: HTTP referrers (web sites)
  * Website restrictions: `localhost:8080/*` (local), `primap.web.app/*` (production)
  * API restrictions: Cloud Firestore API, Maps JavaScript API, Firebase Management API, Firebase Installations API

## Develop
### Setup
At first, install https://github.com/direnv/direnv

```bash
cp .envrc.examle .envrc
vi .envrc
direnv allow

npm install
```

### Run server
```bash
npm start
```

open http://localhost:8080
