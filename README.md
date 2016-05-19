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

The API functionality can be tested through the `api_test.sh` script. The script
makes requests to the API and you will be able to see if the requests go through because you will
see activity in the window where you first ran the API and you will also receive a JSON
response with the private data properly masked.  
There is also a Postman collection available at `test/data_masking_service.postman_collection.json`.  
Import it and select the "Local" environment. The collection also provides some tests that work with the request body that is provided in the JSON config.

## Usage
Make a request to `https://[SERVICE URL]/mask` with the following body:

```json
{
    "Locale": "en_US",
    "Text": "This should be masked as it is a credit card 4111111111111111(VISA). This too should be masked as it is a North American phone number 1-(555)-555-5555? The service can also mask Social Security Numbers like this one: 555-55-5555",
    "MaskSymbol": "(hidden)"
}
```

The response for the above request would look like the one below:

```json
{
    "Locale": "en_US",
    "Text": "This should be masked as it is a credit card (hidden)(VISA). This too should be masked as it is a North American phone number (hidden)? The service can also mask Social Security Numbers like this one: (hidden)",
    "MaskSymbol": "(hidden)"
}
```

### Request body definition

- `Locale` &mdash; This is the locale being used to display the message - usually reflects the type of data that is required to be masked.
> This could change in the near future to express the country/area where certain patterns are available

- `Text` &mdash; The actual message body that contains the data supposed to be masked

- `MaskSymbol`(default: `X`) &mdash; This is a custom symbol(character) or a word/series of words that is meant to replace the data that is supposed to be masked.
