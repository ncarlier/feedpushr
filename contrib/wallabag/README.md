# Wallabag plugin for Feedpushr

Send new articles to a Wallabag instance.

## Configuration

You have to provide credentials for your wallabag instance, see the [API docs](https://doc.wallabag.org/en/developer/api/readme.html).

| Property | Description |
|----------|-------------|
| `url` | Base URL of the Wallabag instance (https://app.wallabag.it unless self-hosting)
| `username` | Username
| `password` | Password
| `clientId` | Client ID for the Wallabag API client
| `clientSecret` | Client secret
| `includeContent` | Whether to send the include the article content in the request to the API. If not, wallabag will use its own fetcher to get the content from the article URL.

## Installation

Copy the `feedpushr-wallabag.so` file into your Feedpushr working directory.

## Usage

```bash
$ feedpushr --log-pretty --plugin ./feedpushr-wallabag.so
```

---

