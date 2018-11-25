#!/usr/bin/env bash

openssl req -newkey rsa:2048 -nodes -keyout server.key -x509 -days 365 -out server.crt
head -1 /dev/random -c32 > symmetric.key && chmod 600 symmetric.key
