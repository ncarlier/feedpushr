module github.com/ncarlier/feedpushr-contrib/rake

require (
	github.com/k3a/html2text v0.0.0-20191003111652-62431c4a3ba5
	github.com/ncarlier/feedpushr/v3 v3.0.0-00010101000000-000000000000
)

replace github.com/ncarlier/feedpushr/v3 => ../..

go 1.13
