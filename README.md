# Private data masking service

Masks private data such as credit card numbers, phone numbers and other identification data.

## Current feature set

- Masks all major credit cards (VISA, Master Card, American Express, Diners Club, Discover, JCB)
- Masks all North-American phone numbers
- Masks Social Security Numbers (SSN)

## Development

The API is written in Golang so you will need the platform installed on your system.
In order to compile a binary, use the build script.
You will need to have Docker installed to make this work for you. The script will kick off
a Docker container that will perform the builds for all operating systems.

## Testing

Currently the API can only be tested through the `api_test.sh` script. The script
makes requests to the API and you will be able to see if the requests go through because you will
see activity in the window where you first ran the API and you will also receive a JSON
response with the private data properly masked.
