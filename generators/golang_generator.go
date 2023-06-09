package generators

import (
	"github.com/ortizdavid/filemanager/core"
	"github.com/ortizdavid/AppGenCore/helpers"
	//dbsamples "github.com/ortizdavid/AppGenCore/samples/db"
	golangsamples "github.com/ortizdavid/AppGenCore/samples/golang"
	//"github.com/ortizdavid/AppGenCore/samples/scripts"
)

type GolangGenerator struct {}

var goFileManager *core.FileManager
var goImport *golangsamples.AppImport
var readmeGo *helpers.ReadMeGo


func (golang *PhpGenerator) GetApplicationTypes() []string {
	types := [] string {"api", "mvc"}
	return types
}