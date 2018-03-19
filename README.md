dp-auth-api-stub
================

Remember to update the [README](README.md) and [CHANGELOG](CHANGELOG.md) files.

### Mocking creation of a new service account
To create a new service account send the following request body with the valid service account name (see below):

```
{
    "name":"VALID_ACCOUNT_NAME"
}
```
#### Valid service account names
- dataset-api

| Request Body | Description | Status | Response Body |
| -------------| ----------- | ------ | --------------| 
| `{"name":"dataset-api"}` | successfully create new service account | 201 | `{"name":"dataset-api","token":"939616dc-7599-4ded-9a86-a9c66fbf98e0"}` |
| `{"name":""}` | no service name in request body | 400 | `{"message":"bad request"}` |
| `{"name":"Morty Smith"}` | service name is not in the list of valid options | 500 | `{"message":"internal server error"}` |

### Get Identity responses

| Authorization header | X-Florence-Token | Profile | Status | Response Body |
| ---------------------|----------------- | --------| -------| --------------| 
| EMPTY                | EMPTY | unauthorized service and user | 401 | `{"message":"not authenticated"}` |
| 939616dc-7599-4ded-9a86-a9c66fbf98e0 | EMPTY | authorized service: and no user | 200 |  `{"identifier":"dataset-api"}` |
| 939616dc-7599-4ded-9a86-a9c66fbf98e0 | 85c718c3-9ba4-4f31-99bb-3e4eaabb2cc1 | authorized service and authorized user | 200 | `{"identifier":"rickSanchez@theCitadelOfRicks.com"}` |
| 939616dc-7599-4ded-9a86-a9c66fbf98e0 | 043a90de-a5b9-42b1-9af4-ad865dda1405 | authorized service and unauthorized user | 401 | `{"message":"not authenticated"}` |
| 58a37231-d534-4f63-b013-25b2f5d59c1a | EMPTY | service token only, internal server error | 500 | `{"message":"internal server error"}` |
| 58a37231-d534-4f63-b013-25b2f5d59c1a | 2741ac66-2d65-4112-a314-a613123f4677 | service and user token, internal server error | 500 | `{"message":"internal server error"}` |


### Configuration

An overview of the configuration options available, either as a table of
environment variables, or with a link to a configuration guide.

| Environment variable | Default | Description
| -------------------- | ------- | -----------
| BIND_ADDR            | :8082   | The host and port to bind to

### Contributing

See [CONTRIBUTING](CONTRIBUTING.md) for details.

### License

Copyright © 2016-2018, Office for National Statistics (https://www.ons.gov.uk)

Released under MIT license, see [LICENSE](LICENSE.md) for details.
