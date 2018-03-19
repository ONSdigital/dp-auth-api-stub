dp-auth-api-stub
================

Remember to update the [README](README.md) and [CHANGELOG](CHANGELOG.md) files.



| Authorization header | Scenario | Response Status | Response Body |
| ---------------------|:--------:| ---------------:| :-------------
| empty                | No auth token header provided, returning unauthenticated response | 401 | ```{"message":"not authenticated"}``` |


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
