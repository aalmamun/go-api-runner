#!/bin/bash

NAME=$1

echo "Calling API with name: $NAME"

# Public API: agify.io (predicts age from name)
RESPONSE=$(curl -s "https://api.agify.io?name=${NAME}")

echo "API Response:"
echo "$RESPONSE"
