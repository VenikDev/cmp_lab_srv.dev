### redis
```
docker pull redis
```
run
```
docker run --name redis -p 6379:6379 redis
```
### deploy
```
docker pull venikshow/cmp-lab-srv
```

run
```
docker run --name cmp-lab-srv -p 80:8080 --link redis:redis venikshow/cmp-lab-srv
```