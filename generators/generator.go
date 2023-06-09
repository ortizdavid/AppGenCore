package generators

import (
	"fmt"
	"github.com/ortizdavid/AppGenCore/helpers"
)

type IGenerator interface {
	GetApplicationTypes() []string
	GenerateConfig(rootDir string, db string)
	GenerateMain(rootDir string, appType string)
	GenerateModels(rootDir string)
	GenerateViews(rootDir string)
	GenerateStaticFiles(rootDir string)
	GenerateApiControllers(rootDir string)
	GenerateMvcControllers(rootDir string)
	GenerateHelpers(rootDir string)
	GenerateMySqlDB(rootDir string)
	GeneratePostgresDB(rootDir string)
	GenerateUploadsDir(rootDir string)
	InstallDeps(rootDir string)
	GenerateReadme(rootDir string, db string, appType string)
	GenerateGitIgnore(rootDir string, appType string)
	GenerateApp(appName string, appType string, db string)
}

type ApplicationGenerator struct {}


func (application *ApplicationGenerator) Generate(appName string, lang string, appType string, db string) {
	fmt.Printf(helpers.CREATING_APPLICATION, appName)

	switch lang {
	case "python":
		var python *PythonGenerator
		python.GenerateApp(appName, appType, db)
	}
}


func (application *ApplicationGenerator) GetCommands() []string {
	commands := [] string {"-name", "-lang", "-type", "-db"}
	return commands
}


func (application *ApplicationGenerator) GetDatabases() []string {
	databases := [] string {"postgres", "mysql"}
	return databases
}


func (application *ApplicationGenerator) GetLanguages() []string {
	languages := [] string {"python", "go"}
	return languages
}
