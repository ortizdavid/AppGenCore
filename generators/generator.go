package generators

import (
	"fmt"
	"github.com/ortizdavid/AppGenCore/helpers"
)

type IGenerator interface {
	generateConfig(rootDir string, db string)
	generateMain(rootDir string, appType string)
	generateModels(rootDir string)
	generateViews(rootDir string)
	generateStaticFiles(rootDir string)
	generateApiControllers(rootDir string)
	generateMvcControllers(rootDir string)
	generateHelpers(rootDir string)
	generateMySqlDB(rootDir string)
	generatePostgresDB(rootDir string)
	generateUploadsDir(rootDir string)
	installDeps(rootDir string)
	generateReadme(rootDir string, db string, appType string)
	generateGitIgnore(rootDir string, appType string)
	GenerateApp(appName string, appType string, db string)
}

type ApplicationGenerator struct {}


func (application *ApplicationGenerator) Generate(appName string, appLang string, appType string, appDb string) bool {

	if helpers.Contains(application.getLanguages(), appLang) == false {
		fmt.Printf(helpers.APPLICATION_ERROR)
		fmt.Printf(helpers.UNSUPPORTED_LANGUAGE, appLang)
		helpers.ListLanguages()
		return false
	} else if helpers.Contains(application.getTypes(), appType) == false {
		fmt.Printf(helpers.APPLICATION_ERROR)
		fmt.Printf(helpers.UNSUPPORTED_TYPE, appType)
		helpers.ListTypes()
		return false
	} else if helpers.Contains(application.getDatabases(), appDb) == false {
		fmt.Printf(helpers.APPLICATION_ERROR)
		fmt.Printf(helpers.UNSUPPORTED_DATABASE, appDb)
		helpers.ListDataBases()
		return false
	} else {
		fmt.Printf(helpers.CREATING_APPLICATION, appName)
		switch appLang {
		case "python":
			var python *PythonGenerator
			python.GenerateApp(appName, appType, appDb)
		case "php":
			var python *PhpGenerator
			python.GenerateApp(appName, appType, appDb)
		case "go":
			var python *GolangGenerator
			python.GenerateApp(appName, appType, appDb)
		}
		fmt.Printf(helpers.APPLICATION_CREATED, appName)
		return true
	}
}


func (application *ApplicationGenerator) getTypes() []string {
	types := [] string {"api", "mvc"}
	return types
}


func (application *ApplicationGenerator) getDatabases() []string {
	databases := [] string {"postgres", "mysql"}
	return databases
}


func (application *ApplicationGenerator) getLanguages() []string {
	languages := [] string {"python", "php", "go"}
	return languages
}
