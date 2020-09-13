# primap

## Develop
### Testing
```bash
docker-compose build
docker-compose run --rm app bash -c 'firebase emulators:exec --only firestore "make test"'
```
