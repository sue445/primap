# primap
[![Build Status](https://github.com/sue445/primap/workflows/build/badge.svg?branch=master)](https://github.com/sue445/primap/actions?query=workflow%3Abuild)
[![Build Status](https://github.com/sue445/primap/workflows/deploy/badge.svg?branch=master)](https://github.com/sue445/primap/actions?query=workflow%3Adeploy)

## Develop
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
docker-compose build && docker-compose run --rm app bash -c 'firebase --project test emulators:exec --only firestore,pubsub "make test"'
```
