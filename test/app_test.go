package test

type applicationTest struct {
	appName string
	appLang string
	appType string
	appDb string
	expected bool
}

var appTestDir string = "../applications/"