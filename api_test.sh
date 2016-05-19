#!/bin/bash
REQUEST_BODY='{"Locale": "en_US","Text": "This should be masked as it is a credit card 4111111111111111(VISA). This too should be masked as it is a North American phone number 1-(555)-555-5555? The service can also mask Social Security Numbers like this one: 555-55-5555"}'
for ((i = 1; i <= 1000; i++)); do
  curl -XPOST -H "Content-Type: application/json" -H "Accept:application/json" -H "Authorization: Basic YWRtaW46YWRtaW4=" -d "${REQUEST_BODY}" http://localhost:8080/mask
done
