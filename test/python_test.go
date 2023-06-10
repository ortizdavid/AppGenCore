package test

import (
	"testing"
	"github.com/ortizdavid/AppGenCore/generators"
)

var python *generators.ApplicationGenerator

func TestCreatePythonMvc(t *testing.T) {
	testCases := []applicationTest {
		{ appName: appTestDir+"PythonMVC1", appLang: "python", appType: "mvc", appDb: "postgres", expected: true, },
		{ appName: appTestDir+"PythonMVC2", appLang: "python", appType: "mvc", appDb: "mysql", expected: true, },
	}
	for _, test := range testCases {
		got := python.Generate(test.appName, test.appLang, test.appType, test.appDb)
		if test.expected != got {
			t.Errorf("Expected %t, Got %t", test.expected, got)
		}
	}
}

func TestCreatePythonApi(t *testing.T) {
	testCases := []applicationTest {
		{ appName: appTestDir+"PythonAPI1", appLang: "python", appType: "api", appDb: "postgres", expected: true, },
		{ appName: appTestDir+"PythonAPI2", appLang: "python", appType: "api", appDb: "mysql", expected: true, },
	}
	for _, test := range testCases {
		got := python.Generate(test.appName, test.appLang, test.appType, test.appDb)
		if test.expected != got {
			t.Errorf("Expected %t, Got %t", test.expected, got)
		}
	}
}

func TestCreatePythonErrorApplication(t *testing.T) {
	testCases := []applicationTest {
		{ appName: appTestDir+"", appLang: "", appType: "", appDb: "", expected: false, },
		{ appName: appTestDir+"PythonP1", appLang: "", appType: "mvc", appDb: "mysql", expected: false, },
		{ appName: appTestDir+"PythonP2", appLang: "python", appType: "cqrs", appDb: "postgres", expected: false, },
		{ appName: appTestDir+"PythonP3", appLang: "python", appType: "api", appDb: "oracle", expected: false, },
	}
	for _, test := range testCases {
		got := python.Generate(test.appName, test.appLang, test.appType, test.appDb)
		if test.expected != got {
			t.Errorf("Expected %t, Got %t", test.expected, got)
		}
	}
}