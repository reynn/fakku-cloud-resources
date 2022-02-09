#!/usr/bin/env bash

# set -ex

if test ! -x "$(which vault)"; then
  echo "Vault CLI is not available"
  exit 1
fi

echo "Getting $MYSQL_VAULT_SECRET from vault"
MYSQL_DATABASE_INFO=$(vault kv get --format=json "$MYSQL_VAULT_SECRET" | dasel select -r json -c -s '.')

export MYSQL_USERNAME="$(echo $MYSQL_DATABASE_INFO | dasel select -r json --plain -s '.data.data.username')"
export MYSQL_PASSWORD="$(echo $MYSQL_DATABASE_INFO | dasel select -r json --plain -s '.data.data.password')"
MYSQL_HOST="$(echo $MYSQL_DATABASE_INFO | dasel select -r json --plain -s '.data.data.host')"
MYSQL_PORT="$(echo $MYSQL_DATABASE_INFO | dasel select -r json --plain -s '.data.data.port')"
export MYSQL_DATABASE="$(echo $MYSQL_DATABASE_INFO | dasel select -r json --plain -s '.data.data.database')"
export MYSQL_ADDRESS="$MYSQL_HOST:$MYSQL_PORT"

/usr/bin/app $@
