# benthos-cached-restful
HTTP POST --> Redis Cache -NO-> Sentiment Process -->Hit PostgreSql
                    |YES                                   |
                    |----------------------------->HTTP Response
## steps1:
```
docker-compose up
```

## step2
```
go run . -c ./config.yaml
```


## test
```
curl "http://localhost:4195/post" -d "any words" -v
```
