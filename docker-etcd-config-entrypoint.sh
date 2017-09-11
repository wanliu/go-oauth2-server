#!/bin/sh

set -e

# to make sure etcd is ready (election ended and leader elected)
while ! etcdctl endpoint health &>/dev/null; do :; done

exec etcdctl put /config/go_oauth2_server.json '{
  "Database": {
    "Type": "postgres",
    "Host": "127.0.0.1",
    "Port": 5432,
    "User": "postgres",
    "Password": "postgres",
    "DatabaseName": "go_oauth2_server",
    "MaxIdleConns": 5,
    "MaxOpenConns": 5
  },
  "Oauth": {
    "AccessTokenLifetime": 3600,
    "RefreshTokenLifetime": 1209600,
    "AuthCodeLifetime": 3600
  },
  "Session": {
      "Secret": "test_secret",
      "Path": "/",
      "MaxAge": 604800,
      "HTTPOnly": true
  },
  "AssetsMappings": [
    {
      "Dir": "public/uploads",
      "Host": "http://localhost:8080/uploads"
    }
  ],
  "IsDevelopment": true
}'
