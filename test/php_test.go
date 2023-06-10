package test

import (
	"testing"
	"github.com/ortizdavid/AppGenCore/generators"
)

var php *generators.ApplicationGenerator

func TestCreatePhpMvc(t *testing.T) {
	testCases := []applicationTest {
		{ appName: appTestDir+"PhpMVC1", appLang: "php", appType: "mvc", appDb: "postgres", expected: true, },
		{ appName: appTestDir+"PhpMVC2", appLang: "php", appType: "mvc", appDb: "mysql", expected: true, },
	}
	for _, test := range testCases {
		got := php.Generate(test.appName, test.appLang, test.appType, test.appDb)
		if test.expected != got {
			t.Errorf("Expected %t, Got %t", test.expected, got)
		}
	}
}

func TestCreatePhpApi(t *testing.T) {
	testCases := []applicationTest {
		{ appName: appTestDir+"PhpAPI1", appLang: "php", appType: "api", appDb: "postgres", expected: true, },
		{ appName: appTestDir+"PhpAPI2", appLang: "php", appType: "api", appDb: "mysql", expected: true, },
	}
	for _, test := range testCases {
		got := php.Generate(test.appName, test.appLang, test.appType, test.appDb)
		if test.expected != got {
			t.Errorf("Expected %t, Got %t", test.expected, got)
		}
	}
}

func TestCreatePhpErrorApplication(t *testing.T) {
	testCases := []applicationTest {
		{ appName: appTestDir+"", appLang: "", appType: "", appDb: "", expected: false, },
		{ appName: appTestDir+"PhpP1", appLang: "", appType: "mvc", appDb: "mysql", expected: false, },
		{ appName: appTestDir+"PhpP2", appLang: "php", appType: "cqrs", appDb: "postgres", expected: false, },
		{ appName: appTestDir+"PhpP3", appLang: "php", appType: "api", appDb: "oracle", expected: false, },
	}
	for _, test := range testCases {
		got := php.Generate(test.appName, test.appLang, test.appType, test.appDb)
		if test.expected != got {
			t.Errorf("Expected %t, Got %t", test.expected, got)
		}
	}
}