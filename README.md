# service

### POST https://evening-ocean-12443.herokuapp.com/covidcases


### Request:
```json
  {
    "latitude":19.101053,
    "longitude":75.740677
  }
```
```bash 
curl --location --request POST 'https://pure-fjord-73951.herokuapp.com/covidcases' \
--header 'Content-Type: application/json' \
--data-raw '{
    "latitude":19.101053,
    "longitude":74.740677
}'
```

### Response:
```json
{
    "StateName": "Maharashtra",
    "District": "Beed",
    "ActiveNo": 1678,
    "ConfirmedNo": 99459,
    "DeceasedNo": 2637,
    "RecoveredNo": 95137
}
```
