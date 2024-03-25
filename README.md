# go-blog-aggregator
Service in Go, building a RESTfull Api, PostgreSQL, SQLc, Goose, service worker.

## install dependencies
### goose
```bash
- go install github.com/pressly/goose/v3/cmd/goose@latest
```
### sqlc
```bash
- go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```
### docker-compose
```bash
#create container with postgres and pgadmin
- docker-compose up -d
```

## Commands to exec migrations goose and sqlc
```bash
- goose postgres postgres://postgres:postgres@localhost:5432/blogator up
- sqlc generate

```
## Run the project
```bash
- go build && ./go-blog-aggregator 
```

## Code Snippets
### cURL Users
```bash
curl --location 'http://localhost:8080/v1/users' \
--header 'Content-Type: application/json' \
--data '{
    "name": "carlos"
}'

curl --location 'http://localhost:8080/v1/users' \
--header 'Authorization: ApiKey c3bc53a80b11556312adb6db57dde0b426710ad331cf6d41ead7215884d4beff'

curl --location 'http://localhost:8080/v1/users'
```
### cURL feeds
```bash
curl --location 'http://localhost:8080/v1/feeds' \
--header 'Authorization: ApiKey c3bc53a80b11556312adb6db57dde0b426710ad331cf6d41ead7215884d4beff' \
--header 'Content-Type: application/json' \
--data '{
  "name": "Im with You",
  "url": "https://google.com/index.xml"
}'

curl --location --request GET 'http://localhost:8080/v1/feeds' \
--header 'Authorization: ApiKey c9b2bd699bd2ecee2d12c63303eb5f93127bbeb3684af65ac3402f4877b1e31f' \
--header 'Content-Type: application/json' \
--data '{
  "name": "The Boot.dev Blog",
  "url": "https://blog.boot.dev/index.xml"
}'
```
### cURL feed-follows
```bash
curl --location 'http://localhost:8080/v1/feed-follows' \
--header 'Authorization: ApiKey 5939e02affe1272fe8897f33b795673f829a3c29a58663b15068f45a225a4839' \
--header 'Content-Type: application/json' \
--data '{
  "feed_id": "741bdb1b-ef49-4116-866a-ae4f2df3f723"
}'

curl --location --request DELETE 'http://localhost:8080/v1/feed-follows/d06bc3ab-2b38-4f7a-b8e7-6d1ff5f3bdc0' \
--header 'Authorization: ApiKey 5939e02affe1272fe8897f33b795673f829a3c29a58663b15068f45a225a4839'

curl --location 'http://localhost:8080/v1/feed-follows' \
--header 'Authorization: ApiKey 5939e02affe1272fe8897f33b795673f829a3c29a58663b15068f45a225a4839'
```
### cURL posts
```bash
curl --location 'http://localhost:8080/v1/posts?limit=9' \
--header 'Authorization: ApiKey c9b2bd699bd2ecee2d12c63303eb5f93127bbeb3684af65ac3402f4877b1e31f'

curl --location 'http://localhost:8080/v1/posts' \
--header 'Authorization: ApiKey c9b2bd699bd2ecee2d12c63303eb5f93127bbeb3684af65ac3402f4877b1e31f' \
--header 'Content-Type: application/json' \
--data '{
    "title": "The Rat 0",
    "url": "https://www.boot.dev/assignments2",
    "description": "insert a new post into the database.",
    "feed_id": "2165651a-0113-4a93-9172-e80511f00b3a"
}'

curl --location 'http://localhost:8080/v1/posts/all'
```