# Kafka plugin for Feedpushr

Send new articles to Kafka.

## Configuration

You have to provides Kafka configuration in order to use this plugin:

| Property | Description |
|----------|-------------|
| `brokers` | Comma separated list of broker host |
| `topic` | Target topic |
| `format` | Payload [format](https://github.com/ncarlier/feedpushr#output-format) (internal JSON format if not provided) | 

## Installation

Copy the `feedpushr-kafka.so` file into your Feedpushr working directory.

## Usage

```bash
$ feedpushr --log-pretty --plugin ./feedpushr-kafka.so
```
## Testing

Start Kafka Docker stack:

```bash
$ docker-compose up
```

Consume `test` topic with [kaf](https://github.com/birdayz/kaf):

```bash
$ go get github.com/birdayz/kaf/cmd/kaf
$ kaf consume test
```

Start Feedpushr and configure Kafka output:

- Brokers: `localhost:9092`
- Topic: `test`

Kaf should consume new records.

---

