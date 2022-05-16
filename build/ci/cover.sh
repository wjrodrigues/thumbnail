#!/bin/sh

COVER=$(go tool cover -func cover.out | grep total | egrep -o [0-9.]+)
TARGET_COVER="100.0"
if [ $COVER = $TARGET_COVER ]; then
  echo "Success coverage $COVER"
  exit 0
else
  echo "Expected coverage $TARGET_COVER current $COVER"
  exit 1
fi
