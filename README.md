# primap
[![Build Status](https://github.com/sue445/primap/workflows/build/badge.svg?branch=master)](https://github.com/sue445/primap/actions?query=workflow%3Abuild)

## Develop
### Run server
```bash
cp .env.examle .env
vi .env

docker-compose up --build
```

open http://localhost:8000

### Testing
```bash
docker-compose build && docker-compose run --rm app bash -c 'firebase emulators:exec --only firestore "make test"'
```
