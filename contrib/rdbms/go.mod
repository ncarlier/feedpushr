module github.com/ncarlier/feedpushr/v3/contrib/rdbms

replace github.com/ncarlier/feedpushr/v3 => ../..

go 1.17

require (
	github.com/abadojack/whatlanggo v1.0.1
	github.com/go-sql-driver/mysql v1.6.0
	github.com/jinzhu/gorm v1.9.16
	github.com/ncarlier/feedpushr/v3 v3.0.0-00010101000000-000000000000
)

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/mattn/go-sqlite3 v2.0.1+incompatible // indirect
)
