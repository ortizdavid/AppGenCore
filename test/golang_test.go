package test

import (
	"testing"
	"github.com/ortizdavid/AppGenCore/generators"
)

var golang *generators.ApplicationGenerator


func TestCreateGolangMvc(t *testing.T) {
	testCases := []applicationTest {
		{ appName: appTestDir+"GolangMVC1", appLang: "go", appType: "mvc", appDb: "postgres", expected: true, },
		{ appName: appTestDir+"GolangMVC2", appLang: "go", appType: "mvc", appDb: "mysql", expected: true, },
	}
	for _, test := range testCases {
		got := golang.Generate(test.appName, test.appLang, test.appType, test.appDb)
		if test.expected != got {
			t.Errorf("Expected %t, Got %t", test.expected, got)
		}
	}
}

func TestCreateGolangApi(t *testing.T) {
	testCases := []applicationTest {
		{ appName: appTestDir+"GolangAPI1", appLang: "go", appType: "api", appDb: "postgres", expected: true, },
		{ appName: appTestDir+"GolangAPI2", appLang: "go", appType: "api", appDb: "mysql", expected: true, },
	}
	for _, test := range testCases {
		got := golang.Generate(test.appName, test.appLang, test.appType, test.appDb)
		if test.expected != got {
			t.Errorf("Expected %t, Got %t", test.expected, got)
		}
	}
}

func TestCreateGolangErrorApplication(t *testing.T) {
	testCases := []applicationTest {
		{ appName: appTestDir+"", appLang: "", appType: "", appDb: "", expected: false, },
		{ appName: appTestDir+"GolangP1", appLang: "", appType: "mvc", appDb: "mysql", expected: false, },
		{ appName: appTestDir+"GolangP2", appLang: "go", appType: "cqrs", appDb: "postgres", expected: false, },
		{ appName: appTestDir+"GolangP3", appLang: "go", appType: "api", appDb: "oracle", expected: false, },
	}
	for _, test := range testCases {
		got := golang.Generate(test.appName, test.appLang, test.appType, test.appDb)
		if test.expected != got {
			t.Errorf("Expected %t, Got %t", test.expected, got)
		}
	}
}