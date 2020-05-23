# RDBMS plugin for Feedpushr

Send new articles to a RDBMS (Mysql, PostgreSQL or sqlite3).

## Configuration

You have to provides your database parameters in order to use this plugin:

| Property | Description |
|----------|-------------|
| `driver` | database driver (available: sqlite3, mysql, postgres) |
| `database` | database name |
| `host` | database host |
| `port` | database port |
| `username` | databse username |
| `password` | database password |
| `verbose` | verbose query activity |

## Installation

Copy the `feedpushr-rdbms.so` file into your Feedpushr working directory.

## Usage

```bash
$ feedpushr --log-pretty --plugin ./feedpushr-rdbms.so
```

---


