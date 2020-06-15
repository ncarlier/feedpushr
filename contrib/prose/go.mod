module github.com/ncarlier/feedpushr/v3/contrib/prose

require (
	github.com/deckarep/golang-set v1.7.1 // indirect
	github.com/iancoleman/strcase v0.0.0-20191112232945-16388991a334
	github.com/k3a/html2text v0.0.0-20191003111652-62431c4a3ba5
	github.com/mingrammer/commonregex v1.0.1 // indirect
	github.com/ncarlier/feedpushr/v3 v3.0.0-00010101000000-000000000000
	gonum.org/v1/gonum v0.7.0 // indirect
	gopkg.in/jdkato/prose.v2 v2.0.0-20190814032740-822d591a158c
	gopkg.in/neurosnap/sentences.v1 v1.0.6 // indirect
)

replace github.com/ncarlier/feedpushr/v3 => ../..

go 1.13
