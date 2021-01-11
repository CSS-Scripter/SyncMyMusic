#!/bin/bash

# Docker deployment
docker-compose -f staging-docker-compose.yml down
docker-compose -f staging-docker-compose.yml up -d --build
