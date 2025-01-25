module github.com/ncarlier/feedpushr/v3/contrib/prose

require (
	github.com/iancoleman/strcase v0.2.0
	github.com/jdkato/prose/v2 v2.0.0
	github.com/k3a/html2text v1.0.8
	github.com/ncarlier/feedpushr/v3 v3.0.0-00010101000000-000000000000
)

require (
	github.com/antonmedv/expr v1.15.3 // indirect
	github.com/deckarep/golang-set v1.7.1 // indirect
	github.com/gopherjs/gopherjs v0.0.0-20190910122728-9d188e94fb99 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mingrammer/commonregex v1.0.1 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/neurosnap/sentences v1.0.7 // indirect
	github.com/rs/zerolog v1.33.0 // indirect
	github.com/smarty/assertions v1.16.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	gonum.org/v1/gonum v0.7.0 // indirect
	gopkg.in/neurosnap/sentences.v1 v1.0.6 // indirect
)

replace github.com/ncarlier/feedpushr/v3 => ../..

go 1.23

toolchain go1.23.4
