#!/bin/bash

if [ "$(docker ps -a -q -f name=task-scheduler-mongo-db)" ]; then
  sudo docker stop task-scheduler-mongo-db
  sudo docker rm -f task-scheduler-mongo-db
fi

if [ "$(docker ps -a -q -f name=task-scheduler-mysql-db)" ]; then
  sudo docker stop task-scheduler-mysql-db
  sudo docker rm -f task-scheduler-mysql-db
fi
