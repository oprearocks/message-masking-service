#!/bin/bash

for ((i = 1; i <= 1000; i++)); do
  curl -XPOST -H "Content-Type: application/json" -H "Accept:application/json" -H "Authorization: Basic YWRtaW46YWRtaW4=" -d '{"Locale": "en_US", "Text": "This is some text that 4023600421224531 to be masked. And more to come de 7123777722227431 ce? This is some text that 4023600421224531 to be masked. And more to come de 7123777722227431 ce?"}' http://localhost:8080/mask
done
