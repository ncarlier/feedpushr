module github.com/ncarlier/feedpushr/v3/contrib/rake

require (
	github.com/k3a/html2text v1.0.8
	github.com/ncarlier/feedpushr/v3 v3.0.0-00010101000000-000000000000
)

require (
	github.com/antonmedv/expr v1.15.3 // indirect
	github.com/gopherjs/gopherjs v0.0.0-20190910122728-9d188e94fb99 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/rs/zerolog v1.31.0 // indirect
	golang.org/x/sys v0.12.0 // indirect
)

replace github.com/ncarlier/feedpushr/v3 => ../..

go 1.19
