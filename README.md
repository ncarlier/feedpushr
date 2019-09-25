# feedpushr

[![Build Status](https://travis-ci.org/ncarlier/feedpushr.svg?branch=master)](https://travis-ci.org/ncarlier/feedpushr)
[![Go Report Card](https://goreportcard.com/badge/github.com/ncarlier/feedpushr)](https://goreportcard.com/report/github.com/ncarlier/feedpushr)
[![Docker pulls](https://img.shields.io/docker/pulls/ncarlier/feedpushr.svg)](https://hub.docker.com/r/ncarlier/feedpushr/)
[![Donate](https://img.shields.io/badge/Donate-PayPal-green.svg)](https://www.paypal.me/nunux)

A simple feed aggregator service with sugar on top.

![Logo](feedpushr.svg)

## Features

- Single executable with an embedded database.
- Manage feed subscriptions.
- Import/Export feed subscriptions with [OPML][opml] files.
- Aggressive and tunable aggregation process.
- Manage feed aggregation individually.
- Apply modifications on articles with a pluggable filter system.
- Push new articles to a pluggable output system (STDOUT, HTTP endpoint, ...).
- Use tags to customize the pipeline.
- Support of [PubSubHubbud][pubsubhubbud] the open, simple, web-scale and
  decentralized pubsub protocol.
- REST API with complete [OpenAPI][openapi] documentation.
- Full feature Web UI and CLI to interact with the API.
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

You can configure the service by setting environment variables:

| Variable | Default | Description |
|----------|---------|-------------|
| `APP_ADDR` | `:8080` | HTTP server address |
| `APP_PUBLIC_URL` | none | Public URL used by PubSubHubbud Hubs. PSHB is disabled if not set. |
| `APP_STORE` | `boltdb://data.db` | Data store location ([BoltDB][boltdb] file) |
| `APP_DELAY` | `1m` | Delay between aggregations (ex: `30s`, `2m`, `1h`, ...) |
| `APP_TIMEOUT` | `5s` | Aggregation timeout (ex: `2s`, `30s`, ...) |
| `APP_CACHE_RETENTION` | `72h` | Cache retention duration (ex: `24h`, `48h`, ...) |
| `APP_LOG_LEVEL` | `info` | Logging level (`debug`, `info`, `warn` or `error`) |
| `APP_LOG_PRETTY` | `false` | Plain text log output format if true (JSON otherwise) |
| `APP_LOG_OUTPUT` | `stdout` | Log output target (`stdout` or `file://sample.log`) |

You can override this settings by using program parameters.
Type `feedpushr --help` to see those parameters.

## Filters

Before being sent, articles can be modified through a filter chain.

Currently, there are some built-in filter:

| Filter | Properties | Description |
|----------|---------|-------------|
| `title`  | `prefix` (default: `foo:`)| This filter will prefix the title of the article with a given value. |
| `fetch`  | None       | This filter will attempt to extract the content of the article from the source URL. |
| `minify` | None       | This filter will minify the HTML content of the article. |

Filters can be extended using [plugins](#plugins).

## Tags

Tags are used to customize the pipeline.

You can define tags on feeds using the Web UI or the API:

```bash
$ curl -XPOST http://localhost:8080/v1/feeds?url=http://www.hashicorp.com/feed.xml&tags=foo,bar
```

Tags can also be imported/exported in OPML format. When using OMPL, tags are stored into the [category attribute][opml-category]. OPML category is a string of comma-separated slash-delimited category strings.
For example, this OMPL attribute `<category>/test,foo,/bar/bar</category>` will be converted to the following tag list: `test, foo, bar_bar`.

Once feeds are configured with tags, each new article will inherit these tags and be pushed out with them.

Tags are also used by filters and outputs to manage their activation.
If you have a filter or an output using tags, only articles corresponding to these tags will be processed by this filter or output.

Example: If you add a `title` filter with `foo,bar` as tags, only new articles with tags `foo` and `bar` will have their title modified with a prefix.

## Outputs

New articles are sent to outputs.

Currently, there are two built-in output providers:

| Output | Properties | Description |
|----------|---------|-------------|
| `stdout` | None    | New articles are sent as JSON documents to the standard output of the process. This can be useful if you want to pipe the command to another shell command. *ex: Store the output into a file. Forward the stream via `Netcat`. Use an ETL tool such as [Logstash][logstash], etc.* |
| `http` | `url` | New articles are sent as JSON documents to an HTTP endpoint (POST). |
| `readflow` | - `url` (default: [official API][readflow-api] <br>- `apiKey` | New articles are sent to [readflow][readflow] instance. |

JSON document format:

```json
{
	"title": "Article title",
	"description": "Article description",
	"content": "Article HTML content",
	"link": "Article URL",
	"updated": "Article update date (String format)",
	"updatedParsed": "Article update date (Date format)",
	"published": "Article publication date (String format)",
	"publishedParsed": "Article publication date (Date format)",
	"guid": "Article feed GUID",
	"meta": {
		"key": "Metadata keys and values added by filters"
	},
	"tags": ["list", "of", "tags"]
}
```

Outputs can be extended using [plugins](#plugins).

## Plugins

You can easily extend the application by adding plugins.

A plugin is a compiled library file that must be loaded when the application starts.

Plugins inside `$PWD` are automaticaly loaded.
You can also load a plugin using the `--plugin` parameter.

Example:

```bash
$ feedpushr --plugin ./feedpushr-twitter.so
```

You can find some external plugins (such as for Twitter) into this
[directory](./contrib).

## Agent

Feedpushr can be started in "desktop mode" thanks to an agent. The purpose of `feedpushr-agent` is to start the daemon and add an icon to your taskbar. This icon allows you to control the daemon and quickly access the user interface.

## User Interface

You can access Web UI on http://localhost:8080/ui

![Screenshot](screenshot.png)

## Use cases

### Start the service

```bash
$ # Start service with default configuration:
$ feedpushr
$ # Start service with custom configuration:
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
$ # Here a quick ETL shell pipeline:
$ # Send transformed articles to HTTP endpoint using shell tools (jq and httpie)
$ feedpushr \
  | jq -c "select(.title) | {title:.title, content:.description, origin: .link}" \
  | while read next; do echo "$next" | http http://postb.in/b/i1J32KdO; done
```

## For development

To be able to build the project you will need to:

- Install `makefiles` external helpers:
  ```bash
  $ git submodule init
  $ git submodule update
  ```
- Install `goa`:
  ```bash
  $ go get -u github.com/goadesign/goa/...
  ```

Then you can build the project using make:

```bash
$ make
```

Type `make help` to see other possibilities.

## License

GNU General Public License v3.0

See [LICENSE](./LICENSE) to see the full text.

---

[opml]: https://en.wikipedia.org/wiki/OPML
[openapi]: https://www.openapis.org/
[pubsubhubbud]: https://github.com/pubsubhubbub/
[boltdb]: https://github.com/coreos/bbolt
[logstash]: https://www.elastic.co/fr/products/logstash
[opml-category]: http://dev.opml.org/spec2.html#otherSpecialAttributes
[readflow]: https://about.readflow.app
[readflow-api]: https://api.readflow.app
