dp-auth-api-stub
================

Service is to assist with developing digital publishing (dp) services and to aide acceptance tests, by stubbing the zebedee endpoints used for authenticating a user or service.

### Configuration

An overview of the configuration options available, either as a table of
environment variables, or with a link to a configuration guide.

| Environment variable | Default | Description
| -------------------- | ------- | -----------
| BIND_ADDR            | :8082   | The host and port to bind to

### Create new service accounts
To add new services to the list of authenticated ones, follow the steps below:

* Extend config.yml to include new service name by appending the service name to the current list under the key `services`
* Add a scenario for new service, if the response status is 200 it would bewise to use the default `SERVICE_AUTH_TOKEN` in that service as the `authorization-token` here.

### Mocking creation of a new service account
To mimic the creation of a new service account send the following request body with the valid service account name (see below):

```
{
    "name":"VALID_ACCOUNT_NAME"
}
```
Beware that this does not actually change the `config.yml` file and the current scenarios will be unchanged.

### Create service account response examples

***Successfully requests will always return the same token value***

| Request Body | Description | Status | Response Body |
| -------------| ----------- | ------ | --------------| 
| `{"name":"dataset-api"}` | successfully create new service account | 201 | `{"name":"dataset-api","token":"939616dc-7599-4ded-9a86-a9c66fbf98e0"}` |
| `{"name":""}` | no service name in request body | 400 | `{"message":"bad request"}` |
| `{"name":"Morty Smith"}` | service name is not in the list of valid options | 500 | `{"message":"internal server error"}` |

### Get Identity response examples (used for testing)

| Authorization header | X-Florence-Token | Profile | Status | Response Body |
| ---------------------|----------------- | --------| -------| --------------| 
| EMPTY                | EMPTY | unauthorized service and user | 401 | `{"message":"not authenticated"}` |
| 939616dc-7599-4ded-9a86-a9c66fbf98e0 | EMPTY | authorized service: and no user | 200 |  `{"identifier":"dataset-api"}` |
| 939616dc-7599-4ded-9a86-a9c66fbf98e0 | 85c718c3-9ba4-4f31-99bb-3e4eaabb2cc1 | authorized service and authorized user | 200 | `{"identifier":"rickSanchez@theCitadelOfRicks.com"}` |
| 939616dc-7599-4ded-9a86-a9c66fbf98e0 | 043a90de-a5b9-42b1-9af4-ad865dda1405 | authorized service and unauthorized user | 401 | `{"message":"not authenticated"}` |
| 58a37231-d534-4f63-b013-25b2f5d59c1a | EMPTY | service token only, internal server error | 500 | `{"message":"internal server error"}` |
| 58a37231-d534-4f63-b013-25b2f5d59c1a | 2741ac66-2d65-4112-a314-a613123f4677 | service and user token, internal server error | 500 | `{"message":"internal server error"}` |

### Contributing

See [CONTRIBUTING](CONTRIBUTING.md) for details.

### License

Copyright © 2016-2018, Office for National Statistics (https://www.ons.gov.uk)

Released under MIT license, see [LICENSE](LICENSE.md) for details.
