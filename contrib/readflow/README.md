# Readflow plugin for Feedpushr

Send new articles to [Readflow](https://readflow.app/).

## Configuration

You have to provides Readflow configuration in order to use this plugin:

| Property | Description |
|----------|-------------|
| `url` | Readflow instance URL (by default: https://readflow.app) |
| `apiKey` | API key |

## Installation

Copy the `readflow-reader.so` file into your Feedpushr working directory.

## Usage

```bash
$ feedpushr --log-pretty --plugin ./feedpushr-readflow.so
```

---

