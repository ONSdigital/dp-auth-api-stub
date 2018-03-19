dp-auth-api-stub
================

Remember to update the [README](README.md) and [CHANGELOG](CHANGELOG.md) files.

| Authorization header | X-Florence-Token | Profile | Status | Response Body |
| ---------------------|:---------------: |:--------:| ------:| :-------------
| EMPTY                | EMPTY | unauthorized service and user | 401 | `{"message":"not authenticated"}` |
| 939616dc-7599-4ded-9a86-a9c66fbf98e0 | EMPTY | authorized service: dataset-api, no user | 200 |  `{"identifier":"dataset-api"}` |
| 939616dc-7599-4ded-9a86-a9c66fbf98e0 | 85c718c3-9ba4-4f31-99bb-3e4eaabb2cc1 | authorized service: dataset-api, authorized user | 200 | `{"identifier":"rickSanchez@theCitadelOfRicks.com"}` |
| 939616dc-7599-4ded-9a86-a9c66fbf98e0 | 043a90de-a5b9-42b1-9af4-ad865dda1405 | authorized service: dataset-api, unauthorized user | 401 | `{"message":"not authenticated"}` |


### Configuration

An overview of the configuration options available, either as a table of
environment variables, or with a link to a configuration guide.

| Environment variable | Default | Description
| -------------------- | ------- | -----------
| BIND_ADDR            | :8080   | The host and port to bind to

### Contributing

See [CONTRIBUTING](CONTRIBUTING.md) for details.

### License

Copyright © 2016-2018, Office for National Statistics (https://www.ons.gov.uk)

Released under MIT license, see [LICENSE](LICENSE.md) for details.
