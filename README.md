วิธีใช้งาน

1. build docker image
    docker build -t api --no-cache .

2. start docker container
    docker run --rm -p 80:8080 -d api

3. API pay with scb
    curl -XPOST http://localhost/payment/internet-banking-scb -d '{"amount":20.12}'

    Response:
    {"id":7,"status":"pending"}
    
4. API get payment transaction detail
    curl -XGET http://localhost/payment/7

    Response:                                   
    {"amount":20.12,"id":7,"status":"pending"}