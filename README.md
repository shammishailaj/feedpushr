# feedpushr

[![Build Status](https://travis-ci.org/ncarlier/feedpushr.svg?branch=master)](https://travis-ci.org/ncarlier/feedpushr)
[![Image size](https://images.microbadger.com/badges/image/ncarlier/feedpushr.svg)](https://microbadger.com/images/ncarlier/feedpushr)
[![Docker pulls](https://img.shields.io/docker/pulls/ncarlier/feedpushr.svg)](https://hub.docker.com/r/ncarlier/feedpushr/)

A simple feed aggregator daemon with sugar on top.

## Features

- Single executable with an embedded database.
- Manage feed subscriptions.
- Import/Export feed subscriptions with [OPML][opml] files.
- Aggressive and tunable aggregation process.
- Manage feed aggregation individually.
- Apply modifications on articles with a pluggable filter system.
- Push new articles to a pluggable output system (STDOUT, HTTP endpoint, ...).
- Support of [PubSubHubbud][pubsubhubbud] the open, simple, web-scale and
  decentralized pubsub protocol.
- REST API with complete [OpenAPI][openapi] documentation.
- Full feature CLI to interact with the daemon's API.
- Metrics production for monitoring.

## Installation

Run the following command:

```bash
$ go get -v github.com/ncarlier/feedpushr
```

**Or** download the binary regarding your architecture:

```bash
$ sudo curl -s https://raw.githubusercontent.com/ncarlier/feedpushr/master/install.sh | bash
```

**Or** use Docker:

```bash
$ docker run -d --name=feedpushr ncarlier/feedpushr
```

## Configuration

You can configure the daemon by setting environment variables:

| Variable | Default | Description |
|----------|---------|-------------|
| `APP_LISTEN_ADDR` | `:8080` | HTTP server address |
| `APP_PUBLIC_URL` | none | Public URL used by PubSubHubbud Hubs. PSHB is disabled if not set. |
| `APP_STORE` | `boltdb://data.db` | Data store location ([BoltDB][boltdb] file) |
| `APP_FILTERS` | none | Filter chain (`foo` or `fetch`) |
| `APP_OUTPUT` | `stdout` | Output destination (`stdout` or HTTP URL) |
| `APP_DELAY` | `1m` | Delay between aggregations (ex: `30s`, `2m`, `1h`, ...) |
| `APP_TIMEOUT` | `5s` | Aggregation timeout (ex: `2s`, `30s`, ...) |
| `APP_CACHE_RETENTION` | `72h` | Cache retention duration (ex: `24h`, `48h`, ...) |
| `APP_LOG_LEVEL` | `info` | Logging level (`debug`, `info`, `warn` or `error`) |
| `APP_LOG_PRETTY` | `false` | Plain text log output format if true (JSON otherwise) |
| `APP_LOG_OUTPUT` | `stdout` | Log output target (`stdout` or `file://sample.log`) |

You can override this settings by using program parameters.
Type `feedpushr --help` to see those parameters.

## Filters

Before being sent articles can be modified through a filter chain.
Currently, there are two built-in filter:

- `foo`:
  This filter will prefix the title of the article with the name of the
  application.
- `fetch`:
  This filter will attempt to extract the content of the article from the source
  URL.

You can chain all the filters you need.

Filters can be extended using plugins.

## Outputs

New articles are sent to an output.
Currently, there are two built-in output providers:

- `stdout`: New articles are sent as JSON document to the standard output of the
  process.
  This can be useful if you want to pipe the command to another shell command.
  *ex: Store the output into a file. Forward the stream via `Netcat`. Use an ETL
  tool such as [Logstash][logstash], etc.*
- `http`: New articles are sent as JSON document to an HTTP endpoint (POST).

Outputs can be extended using plugins.

## Plugins

You can easily extend the application by adding plugins.

A plugin is a compiled library file that must be loaded when the application
starts.

To load a plugin you have to use the `--plugin` parameter. Example:

```bash
$ feedpushr --plugin ./feedpushr-twitter-linux-amd64.so
```

You can find some external plugins (such as for Twitter) into this
[repository][contrib].

## Use cases

### Start the daemon

```bash
$ # Start the daemon with default configuration:
$ feedpushr
$ # Start the daemon and send new articles to a HTTP endpoint:
$ feedpushr --output https://requestb.in/t4gdzct4
$ # Start the daemon with a database initialized
$ # with subscriptions from an OPML file:
$ feedpusrh --import ./my-subscriptions.xml
$ # Start the daemon with custom configuration:
$ export APP_OUTPUT="https://requestb.in/t4gdzct4"
$ export APP_STORE="boltdb:///var/opt/feedpushr.db"
$ export APP_DELAY=20s
$ export APP_LOG_LEVEL=warn
$ feedpushr
```
### Add feeds

```bash
$ # Add feed with the CLI
$ feedpushr-ctl create feed --url http://www.hashicorp.com/feed.xml
$ # Add feed with cURL
$ curl -XPOST http://localhost:8080/v1/feeds?url=http://www.hashicorp.com/feed.xml
$ # Import feeds from an OPML file
$ curl -XPOST http://localhost:8080/v1/opml -F"file=@subscriptions.opml"
```

### Manage feeds

```bash
$ # List feeds
$ feedpushr-ctl list feed
$ # Get a feed
$ feedpushr-ctl get feed --id=9090dfac0ccede1cfcee186826d0cc0d
$ # Remove a feed
$ feedpushr-ctl delete feed --id=9090dfac0ccede1cfcee186826d0cc0d
$ # Stop aggregation of a feed
$ feedpushr-ctl stop feed --id=9090dfac0ccede1cfcee186826d0cc0d
$ # Start aggregation of a feed
$ feedpushr-ctl start feed --id=9090dfac0ccede1cfcee186826d0cc0d
```

### Misc

```bash
$ # Get OpenAPI JSON
$ curl  http://localhost:8080/swagger.json
$ # Get runtime vars
$ curl  http://localhost:8080/v1/vars
```

## For development

To be able to build the project you will need to:

- Install `makefiles` external helpers:
  ```bash
  $ git submodule init
  $ git submodule update
  ```
- Install `dep` and `goa`:
  ```bash
  $ go get -u github.com/golang/dep/cmd/dep
  $ go get -u github.com/goadesign/goa/...
  ```

Then you can build the project using make:

```bash
$ make
```

Type `make help` to see other possibilities.

---

[opml]: https://en.wikipedia.org/wiki/OPML
[openapi]: https://www.openapis.org/
[pubsubhubbud]: https://github.com/pubsubhubbub/
[boltdb]: https://github.com/coreos/bbolt
[logstash]: https://www.elastic.co/fr/products/logstash
[contrib]: https://github.com/ncarlier/feedpushr-contrib/

