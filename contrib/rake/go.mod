module github.com/ncarlier/feedpushr/v3/contrib/rake

require (
	github.com/k3a/html2text v0.0.0-20191003111652-62431c4a3ba5
	github.com/ncarlier/feedpushr/v3 v3.0.0-00010101000000-000000000000
)

require (
	github.com/antonmedv/expr v1.8.9 // indirect
	github.com/gopherjs/gopherjs v0.0.0-20190910122728-9d188e94fb99 // indirect
	github.com/rs/zerolog v1.23.0 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
)

replace github.com/ncarlier/feedpushr/v3 => ../..

go 1.17
