package generators

import (
	"github.com/ortizdavid/filemanager/core"
	//"github.com/ortizdavid/AppGenCore/helpers"
	//dbsamples "github.com/ortizdavid/AppGenCore/samples/db"
	phpsamples "github.com/ortizdavid/AppGenCore/samples/php"
	//"github.com/ortizdavid/AppGenCore/samples/scripts"
)

type PhpGenerator struct {}

var phpFileManager *core.FileManager
var phpImport *phpsamples.AppImport


func (php *PhpGenerator) GetProjectTypes() []string {
	types := [] string {"api", "mvc"}
	return types
}