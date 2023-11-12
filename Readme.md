## Examples

Некоторые примеры запросов
- [Добавление пользователя](#create-user)


### Добавление пользователя <a name="create-user"></a>

```curl
curl -X 'POST' \
  'http://localhost:9000/signup' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "login" : "Ilon",
  "wealth": 100000
}'
```

```curl
curl --header "Token: ba098c25f5a8321e9940447f551b5126f244e6ed" -X 'GET' \
  'http://localhost:9000/api/info'
}'
```