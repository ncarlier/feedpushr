module github.com/ncarlier/feedpushr-contrib/rake

require (
	github.com/k3a/html2text v0.0.0-20180923223239-2cdb1fac5429
	github.com/ncarlier/feedpushr v0.0.0-00010101000000-000000000000
)

replace github.com/ncarlier/feedpushr => ../..

go 1.13
