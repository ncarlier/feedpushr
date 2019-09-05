# RAKE plugin for Feedpushr

Extract keywords of an article by using Rapid Automatic Keyword Extraction algorithm.

## Configuration

You can provide RAKE configuration in order to tune the RAKE algorithm:

| Property | Default | Description |
|----------|---------|-------------|
| `minCharLength`       | `4` | Each word has at least N characters |
| `maxWordsLength`      | `3` | Each phrase has at most N words |
| `minKeywordFrequency` | `4` | Each keyword appears in the text at least N times |

## Installation

Copy the `feedpushr-rake.so` file into your Feedpushr working directory.

## Usage

```bash
$ feedpushr --log-pretty --plugin ./feedpushr-rake.so
```

---

