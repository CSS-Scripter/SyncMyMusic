#!/bin/bash

# Docker deployment
docker-compose -f production-docker-compose.yml down
docker-compose -f production-docker-compose.yml up -d --build
