# WHOAMi

[![Build Status](https://travis-ci.com/nhoclove/whoami.svg?branch=master)](https://travis-ci.com/nhoclove/whoami)
[![codecov](https://codecov.io/gh/nhoclove/whoami/branch/master/graph/badge.svg)](https://codecov.io/gh/nhoclove/whoami)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/nhoclove/whoami/blob/master/LICENSE.md)


# Description
A simple HTTP service

# Features
- Exposes 2 REST APIs
  - POST: /whoami
  - GET:/ whoami
- Responses in JSON format

# How to
- Run local:
  
    `go run main.go --addr=localhost:3000`
- Run with Docker:
    ```
    $ docker build -t whoami:latest .
    $ docker run --rm -it --name=whoami -p 3000:3000 whoami:latest
    ```
  
