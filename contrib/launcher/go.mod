module github.com/ncarlier/feedpushr-contrib/launcher

go 1.12

require (
	github.com/getlantern/golog v0.0.0-20190830074920-4ef2e798c2d7 // indirect
	github.com/getlantern/systray v0.0.0-20190727060347-6f0e5a3c556c
	github.com/ncarlier/feedpushr/v2 v2.0.0-00010101000000-000000000000
)

replace github.com/ncarlier/feedpushr/v2 => ../..
