#!/bin/bash

if [ ! -d "gen" ]; then
  mkdir gen gen/models gen/restapi
fi

swagger generate server -t gen --exclude-main -A TaskScheduler
swagger generate client -t gen -A TaskScheduler
