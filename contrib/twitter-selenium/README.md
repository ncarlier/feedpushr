# Twitter-Selenium plugin for Feedpushr

Send new articles to a Twitter timeline with Selenium.

## Configuration

You have to provides your Twitter credentials in order to use this plugin:

| Property | Description |
|----------|-------------|
| `username` | username/account |
| `password` | password |
| `format` | Tweet [format](https://github.com/ncarlier/feedpushr#output-format) (by default: `{{.Title}}\n{{.Link}}`) | 

## Installation

Copy the `feedpushr-twitter-selenium.so` file into your Feedpushr working directory.

## Usage

```bash
$ feedpushr --log-pretty --plugin ./feedpushr-twitter-selenium.so
```

---

