# WHOAMi

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
  