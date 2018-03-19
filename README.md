dp-auth-api-stub
================

Remember to update the [README](README.md) and [CHANGELOG](CHANGELOG.md) files.



| Authorization header | Profile | Status | Response Body |
| ---------------------|:--------:| ------:| :-------------
| empty                | No auth token header provided | 401 | `{"message":"not authenticated"}` |
| 939616dc-7599-4ded-9a86-a9c66fbf98e0 | dataset api | 200 | `{"identifier":"dataset-api"}` |
| 61a11670-f7ed-4c20-90e2-a36367ad9119 | valid user | 200 | `{"identifier":"rickSanchez@theCitadelOfRicks.com"}` |


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
