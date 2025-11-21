#!/bin/sh
set -e

TARGET_URL="${TARGET_URL:-https://amaze-api.curiola.com/hello}"

echo "[$(date)] helloworld container starting..."
echo "Calling endpoint: $TARGET_URL"
echo "API_TOKEN: $API_TOKEN"

set +e

if [ -n "$API_TOKEN" ]; then
  RESP=$(curl -sS -X POST \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $API_TOKEN" \
    -d "{\"message\":\"hello from container\",\"hostname\":\"$(hostname)\"}" \
    "$TARGET_URL" 2>&1)
else
  RESP=$(curl -sS -X POST \
    -H "Content-Type: application/json" \
    -d "{\"message\":\"hello from container\",\"hostname\":\"$(hostname)\"}" \
    "$TARGET_URL" 2>&1)
fi

STATUS=$?
set -e

echo "curl exit status: $STATUS"
echo "Response:"
echo "$RESP"

tail -f /dev/null
