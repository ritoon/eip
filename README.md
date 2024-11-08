# EIP

## API
CRUD on Users, Games and Addresses

## Front
Login page with credentials and store a jwt token into the **sessionStorage**
Game list page with all pages

## Run locally

```sh
cd front 
make dev
cd ../geocoding
make dev
cd ../api
make dev
docker compose up
```




## Metrics with grafana
![dashboard](doc/grafana-dashboard.jpg)


## Mount data

First register a new user.
```sh
curl --location 'http://localhost:8888/register' \
--header 'Content-Type: application/json' \
--header 'Authorization: ••••••' \
--data '{
    "name":"Jo",
    "email":"toto",
		"address":{
			"city":"Strasbourg",
            "zip":"12345",
            "street":"la tour"
		},
		"games":[
			{"name":"Donjon Dragon"},
			{"name":"Warhammer"}
		],
    "pass":"toto"
}'
```

Then login the user
```sh
curl --location 'http://localhost:8888/login' \
--header 'Content-Type: application/json' \
--data '{
    "email":"toto",
    "pass":"toto"
}'
```
Get the token in the JWT response and copy past to the file `api/scrapper/scrapper.go` line 55.
Finally run the dummy test of this file `api/scrapper/scrapper_test.go`



## RUN

build docker images


## test with vegeta
jq -ncM '{method: "GET", url: "http://localhost:8888/rooms" | @base64, header: {"Authorization": ["Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1dWlkX3VzZXIiOiIwMWI3ZmJhZS01MjdmLTQ4ZTctODgxMi0wNGMxZGJmOGJiZWUiLCJhY2Nlc3NfbGV2ZWwiOiIiLCJlbWFpbCI6ImJpZGVuQHRlc3QuZ28iLCJleHAiOjE3MzA5OTYwMzN9.ndAHWSPiIuniNM-wZQNMqdGVR01rJAlw5AkKVTalfhk"]}}'  | vegeta attack -duration=5s | tee results.bin | vegeta report  vegeta report -type=json results.bin > metrics.json
  cat results.bin | vegeta plot > plot.html
  cat results.bin | vegeta report -type="hist[0,0.04ms,0.05ms,0.07ms,0.08ms,1ms,3ms]"