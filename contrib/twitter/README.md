# Twitter plugin for Feedpushr

Send new articles to a Twitter timeline.

## Configuration

You have to provides Twitter configuration in order to use this plugin:

| Property | Description |
|----------|-------------|
| `consumerKey` | Consumer key |
| `consumerSecret` | Consumer secret |
| `accessToken` | Access token |
| `accessTokenSecret` | Access token secret |
| `format` | Tweet [format](https://github.com/ncarlier/feedpushr#output-format) (by default: `{{.Title}}\n{{.Link}}`) | 

## Installation

Copy the `feedpushr-twitter.so` file into your Feedpushr working directory.

## Usage

```bash
$ feedpushr --log-pretty --plugin ./feedpushr-twitter.so
```

---

