# Mastodon plugin for Feedpushr

Send new articles as a "Toot" to a [Mastodon](https://joinmastodon.org/) instance.

## Configuration

You have to provides Mastodon configuration in order to use this plugin:

| Property | Description |
|----------|-------------|
| `url` | Mastodon instance URL (by default: https://mastodon.social) |
| `token` | Access token |
| `visibility` | Toot visibility ("direct", "private", "unlisted" or by default "public") |

You can create your access token by using [this type of tool](https://takahashim.github.io/mastodon-access-token/).

Note that an access token of mastodon never expires.
However, you can manage your access tokens using the settings page.

## Installation

Copy the `feedpushr-mastodon.so` file into your Feedpushr working directory.

## Usage

```bash
$ feedpushr --log-pretty --plugin ./feedpushr-mastodon.so
```

---

