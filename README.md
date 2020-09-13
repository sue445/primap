# primap

## Develop
### Testing
```bash
docker-compose build
docker-compose run app bash -c 'firebase emulators:exec --only firestore "make test"'
```
