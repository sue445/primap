# primap

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
