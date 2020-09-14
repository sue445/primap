# primap
[![Build Status](https://github.com/sue445/primap/workflows/test/badge.svg?branch=master)](https://github.com/sue445/primap/actions?query=workflow%3Atest)

## Develop
### Run server
```bash
docker-compose up --build
```

open http://localhost:8000

### Testing
```bash
docker-compose build && docker-compose run --rm app bash -c 'firebase emulators:exec --only firestore "make test"'
```
