package test

import (
	"testing"

	"github.com/ortizdavid/AppGenCore/generators"
)

var python *generators.PythonGenerator

func TestCreatePythonMvcPostgres(t *testing.T) {
	python.GenerateApp("PyMvcPostgres", "mvc", "postgres");
}

func TestCreatePythonApiPostgres(t *testing.T) {
	python.GenerateApp("PyApiPostgres", "api", "postgres");
}

func TestCreatePythonMvcMySQL(t *testing.T) {
	python.GenerateApp("PyMvcMySQL", "mvc", "mysql");
}

func TestCreatePythonApiMySQL(t *testing.T) {
	python.GenerateApp("PyApiMySQL", "api", "mysq");
}

