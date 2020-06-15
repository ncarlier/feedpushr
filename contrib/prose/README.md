# Prose plugin for Feedpushr

Extract Named Entity from text and convert them to hashtags..

## Configuration

You can provide Prose configuration in order to tune the recognition algorithm:

| Property | Default | Description |
|----------|---------|-------------|
| `filter` | `all` | Filter entity by label (available: all,person,gpe) |
| `format`      | `hashtag`  | Format each named entity in a specific type (available: hashtag,keyword,none |
| `separator`   | `space`    | Separate each named entity with a defined chararcter (available: space,tab,comma,semi-colon,pipe) |
| `minCharLength`      | `1`  | Each entity has at least N characters |
| `maxCharLength`      | `15` | Each entity has a maxium of N characters |

## Installation

Copy the `feedpushr-prose.so` file into your Feedpushr working directory.

## Usage

```bash
$ feedpushr --log-pretty --plugin ./feedpushr-prose.so
```

---

