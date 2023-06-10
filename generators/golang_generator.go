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


func (golang *GolangGenerator) GenerateApp(appName string, appType string, db string) {
	/*
	switch appType {
	case "mvc":
		golang.generateViews(appName)
		golang.generateStaticFiles(appName)
		golang.generateMvcControllers(appName)
	case "api":
		golang.generateApiControllers(appName)
		golang.generatePostmanCollections(appName)
	}
	switch db {
	case "mysql":
		golang.generateMySqlDB(appName)
	case "postgres":
		golang.generatePostgresDB(appName)
	}
	golang.generateModels(appName)
	golang.generateHelpers(appName)
	golang.generateConfig(appName, db)
	golang.generateRequirements(appName, db)
	golang.generateMain(appName, appType)
	golang.generateReadme(appName, db, appType)
	golang.generateGitIgnore(appName, appType)
	golang.generateUploadsDir(appName)
	golang.printInstructions(appName)
	*/
}

/*

func (golang *GolangGenerator) generateUploadsDir(rootDir string) {
	goFileManager.CreateManyFolders(rootDir+"/static/uploads/imgs")
	goFileManager.CreateManyFolders(rootDir+"/static/uploads/docs")
}


func (golang *GolangGenerator) printInstructions(appName string) {
	fmt.Println(readme.InstrunctionsBeforeRun(appName))
}

func (golang *GolangGenerator) generateConfig(rootDir string, db string) {
	var config *golangsamples.ConfigGo
	file := "config.go"
	goFileManager.CreateSigleFile(rootDir, file)
	goFileManager.WriteFile(rootDir, file, goImport.ImportForConfig())
	goFileManager.WriteFile(rootDir, file, config.CreateConfig(db))
}


func (golang *GolangGenerator) generateReadme(rootDir string, db string, appType string) {
	file := "README.md"
	goFileManager.CreateSigleFile(rootDir, file)
	switch appType {
	case "mvc":
		goFileManager.WriteFile(rootDir, file, readme.ReadmeMVC(db))
	case "api":
		goFileManager.WriteFile(rootDir, file, readme.ReadmeAPI(db))
	}
}


func (golang *GolangGenerator) generateGitIgnore(rootDir string, appType string) {
	var ignore *helpers.GitIgnoreGo
	file := ".gitignore"
	goFileManager.CreateSigleFile(rootDir, file)
	switch appType {
	case "mvc":
		goFileManager.WriteFile(rootDir, file, ignore.GitIgnoreMvc())
	case "api":
		goFileManager.WriteFile(rootDir, file, ignore.GitIgnoreAPI())
	}
}


func (golang *GolangGenerator) generateMain(rootDir string, appType string) {
	file := "app.go"
	goFileManager.CreateSigleFile(rootDir, file)
	switch appType {
	case "mvc":
		goFileManager.WriteFile(rootDir, file, goImport.ImportForMvcApp())
	case  "api":
		goFileManager.WriteFile(rootDir, file, goImport.ImportForRestApi())
	}
	goFileManager.WriteFile(rootDir, file, goImport.AppMainCode())
}


func (golang *GolangGenerator) generateMySqlDB(rootDir string) {
	var mysql *dbsamples.MySqlDB
	dbDir := rootDir+"/database"
	file := "db_task.sql"
	goFileManager.CreateManyFolders(dbDir)
	goFileManager.CreateSigleFile(dbDir, file)
	goFileManager.WriteFile(dbDir, file, mysql.GetDatabaseScript())
}


func (golang *GolangGenerator) generatePostgresDB(rootDir string) {
	var postgres *dbsamples.PostgresDB
	dbDir := rootDir+"/database"
	file := "db_task.sql"
	goFileManager.CreateManyFolders(dbDir)
	goFileManager.CreateSigleFile(dbDir, file)
	goFileManager.WriteFile(dbDir, file, postgres.GetDatabaseScript())
}


func (golang *GolangGenerator) generateRequirements(rootDir string, db string) {
	var deps *scripts.GothonDeps
	file := "requirements.txt"
	goFileManager.CreateSigleFile(rootDir, file)
	goFileManager.WriteFile(rootDir, file, deps.Requirements(db))
}

func (golang *GolangGenerator) generatePostmanCollections(rootDir string) {
	var postman *collectionsamples.Postman
	collectionsFolder := rootDir+"/__postman_collections__"
	collection := "postman_collection.json"
	testRun := "postman_test_run.json"
	goFileManager.CreateManyFolders(collectionsFolder)
	goFileManager.CreateManyFiles(collectionsFolder, collection, testRun)
	goFileManager.WriteFile(collectionsFolder, collection, postman.PostmanCollection(rootDir))
	goFileManager.WriteFile(collectionsFolder, testRun, postman.PostmanTestRun(rootDir))
}


func (golang *GolangGenerator) generateHelpers(rootDir string) {
	var helper *golangsamples.Helper
	helpersFolder := rootDir+"/helpers"
	httpFile := "http_code.go"
	constFile := "constants.go"
	uploaderFile := "file_uploader.go"
	passwordHandlerFile := "password_handler.go"
	goFileManager.CreateManyFolders(helpersFolder)
	goFileManager.CreateSigleFile(helpersFolder, httpFile)
	goFileManager.CreateSigleFile(helpersFolder, constFile)
	goFileManager.CreateSigleFile(helpersFolder, uploaderFile)
	goFileManager.CreateSigleFile(helpersFolder, passwordHandlerFile)
	goFileManager.WriteFile(helpersFolder, httpFile, helper.HttpCodes())
	goFileManager.WriteFile(helpersFolder, constFile, helper.Constants())
	goFileManager.WriteFile(helpersFolder, uploaderFile, helper.FileUploader())
	goFileManager.WriteFile(helpersFolder, passwordHandlerFile, helper.PasswordHandler())
}


func (golang *GolangGenerator) generateModels(rootDir string) {
	var model *golangsamples.Model
	modelsFolder := rootDir+"/models"
	roleFile := "role.go"
	userFile := "user.go"
	taskFile := "task.go"
	goFileManager.CreateManyFolders(modelsFolder)
	goFileManager.CreateSigleFile(modelsFolder, roleFile)
	goFileManager.CreateSigleFile(modelsFolder, taskFile)
	goFileManager.CreateSigleFile(modelsFolder, userFile)
	goFileManager.WriteFile(modelsFolder, roleFile, model.RoleModel())
	goFileManager.WriteFile(modelsFolder, userFile, model.UserModel())
	goFileManager.WriteFile(modelsFolder, taskFile, model.TaskModel())
}


func (golang *GolangGenerator) generateStaticFiles(rootDir string) {
	var staticFile *libs.StaticFile
	var bootstrap *libs.BootstrapLib
	var jquery *libs.JqueryLib

	staticCss := rootDir+"/static/css"
	staticJs := rootDir+"/static/js"
	staticImgs := rootDir+"/static/images"
	staticLibBootstrapCss := rootDir+"/static/lib/bootstrap/css"
	staticLibBootstrapJs := rootDir+"/static/lib/bootstrap/css"
	staticLibJquery := rootDir+"/static/lib/jquery"
	
	jsFile := "script.js"
	cssFile := "style.css"
	bootstrapCssFile := "bootstrap.min.css"
	bootstrapJsFile := "bootstrap.min.js"
	jqueryJsFile := "jquery.min.js"

	goFileManager.CreateManyFolders(staticCss)
	goFileManager.CreateManyFolders(staticJs)
	goFileManager.CreateManyFolders(staticImgs)
	goFileManager.CreateManyFolders(staticLibBootstrapCss)
	goFileManager.CreateManyFolders(staticLibBootstrapJs)
	goFileManager.CreateManyFolders(staticLibJquery)
	goFileManager.CreateSigleFile(staticCss, cssFile)
	goFileManager.CreateSigleFile(staticJs, jsFile)
	goFileManager.CreateSigleFile(staticLibBootstrapCss, bootstrapCssFile)
	goFileManager.CreateSigleFile(staticLibBootstrapJs, bootstrapJsFile)
	goFileManager.CreateSigleFile(staticLibJquery, jqueryJsFile)
	goFileManager.WriteFile(staticCss, cssFile, staticFile.CssContent())
	goFileManager.WriteFile(staticJs, jsFile, staticFile.JsContent())
	goFileManager.WriteFile(staticLibBootstrapCss, bootstrapCssFile, bootstrap.BootstrapMinCss())
	goFileManager.WriteFile(staticLibBootstrapJs, bootstrapJsFile, bootstrap.BootstrapMinJs())
	goFileManager.WriteFile(staticLibJquery, jqueryJsFile, jquery.JqueryMinJs())
}


func (golang *GolangGenerator) generateMvcControllers(rootDir string) {
	var mvcController *golangsamples.MvcController
	controllersFolder := rootDir+"/controllers"
	roleFile := "role_controller.go"
	userFile := "user_controller.go"
	taskFile := "task_controller.go"
	authFile := "auth_controller.go"
	frontFile := "front_controller.go"
	goFileManager.CreateManyFolders(controllersFolder)
	goFileManager.CreateSigleFile(controllersFolder, roleFile)
	goFileManager.CreateSigleFile(controllersFolder, taskFile)
	goFileManager.CreateSigleFile(controllersFolder, userFile)
	goFileManager.CreateSigleFile(controllersFolder, authFile)
	goFileManager.CreateSigleFile(controllersFolder, frontFile)
	goFileManager.WriteFile(controllersFolder, roleFile, mvcController.RoleController())
	goFileManager.WriteFile(controllersFolder, userFile, mvcController.UserController())
	goFileManager.WriteFile(controllersFolder, taskFile, mvcController.TaskController())
	goFileManager.WriteFile(controllersFolder, authFile, mvcController.AuthController())
	goFileManager.WriteFile(controllersFolder, frontFile, mvcController.FrontController())
}


func (golang *GolangGenerator) generateApiControllers(rootDir string) {
	var apiController *golangsamples.ApiController
	apiControllersFolder := rootDir+"/api_controllers"
	roleFile := "role_api.go"
	userFile := "user_api.go"
	taskFile := "task_api.go"
	authFile := "auth_api.go"
	rootFile := "root_api.go"
	goFileManager.CreateManyFolders(apiControllersFolder)
	goFileManager.CreateSigleFile(apiControllersFolder, roleFile)
	goFileManager.CreateSigleFile(apiControllersFolder, taskFile)
	goFileManager.CreateSigleFile(apiControllersFolder, userFile)
	goFileManager.CreateSigleFile(apiControllersFolder, authFile)
	goFileManager.CreateSigleFile(apiControllersFolder, rootFile)
	goFileManager.WriteFile(apiControllersFolder, roleFile, apiController.RoleApiController())
	goFileManager.WriteFile(apiControllersFolder, userFile, apiController.UserApiController())
	goFileManager.WriteFile(apiControllersFolder, taskFile, apiController.TaskApiController())
	goFileManager.WriteFile(apiControllersFolder, authFile, apiController.AuthApiController())
	goFileManager.WriteFile(apiControllersFolder, rootFile, apiController.RootApiController())
}


func (golang *GolangGenerator) generateViews(rootDir string) {

	var layout *golangsamples.Layout
	var perr *golangsamples.PageError
	var front *golangsamples.FrontTemplate
	var auth *golangsamples.AuthTemplate
	var user *golangsamples.UserTemplate
	var task *golangsamples.TaskTemplate
	var role *golangsamples.RoleTemplate

	templatesFolder := rootDir+"/templates"
	layoutsFolder := templatesFolder+"/layouts"
	errorFolder := templatesFolder+"/error"
	frontFolder := templatesFolder+"/front"
	authFolder := templatesFolder+"/auth"
	userFolder := templatesFolder+"/user"
	taskFolder := templatesFolder+"/task"
	roleFolder := templatesFolder+"/role"
	
	backLayoutFile := "back-layout.html"
	frontLayoutFile := "front-layout.html"
	normalMenuFile := "normal-menu.html"
	adminMenuFile := "admin-menu.html"

	indexFile := "index.html"
	loginFile := "login.html"
	homeFile := "home.html"
	err404File := "404.html"
	err404MenuFile := "404-menu.html"
	
	userAddFile := "add.html"
	userEditFile := "edit.html"
	userDataFile := "user-data.html"
	userSearchFile := "search.html"
	userSearchResultsFile := "search-results.html"
	userShowFile := "show.html"
	userDetailsFile := "details.html"
	userUploadFile := "upload-image.html"

	taskAddFile := "add.html"
	taskEditFile := "edit.html"
	taskShowFile := "show.html"
	taskSearchFile := "search.html"
	taskSearchResultsFile := "search-results.html"
	taskDetailsFile := "details.html"
	userTasksFile := "user-tasks.html"

	roleAddFile := "add.html"
	roleShowFile := "show.html"
	roleDetailsFile := "details.html"

	goFileManager.CreateManyFolders(templatesFolder)
	goFileManager.CreateManyFolders(layoutsFolder)
	goFileManager.CreateManyFolders(frontFolder)
	goFileManager.CreateManyFolders(authFolder)
	goFileManager.CreateManyFolders(userFolder)
	goFileManager.CreateManyFolders(roleFolder)
	goFileManager.CreateManyFolders(taskFolder)
	goFileManager.CreateManyFolders(errorFolder)
	
	goFileManager.CreateManyFiles(errorFolder, err404File, err404MenuFile)
	goFileManager.CreateSigleFile(frontFolder, indexFile)
	goFileManager.CreateManyFiles(authFolder, loginFile, homeFile)
	goFileManager.CreateManyFiles(layoutsFolder, frontLayoutFile, backLayoutFile, normalMenuFile, adminMenuFile)
	goFileManager.CreateManyFiles(roleFolder, roleAddFile, roleShowFile, roleDetailsFile)
	goFileManager.CreateManyFiles(taskFolder, taskAddFile, taskEditFile, taskSearchFile, taskSearchResultsFile, taskShowFile, taskDetailsFile, userTasksFile)
	goFileManager.CreateManyFiles(userFolder, userAddFile, userEditFile, userDataFile, userSearchFile, userSearchResultsFile, userShowFile, userSearchFile, userDetailsFile, userUploadFile)

	goFileManager.WriteFile(layoutsFolder, frontLayoutFile, layout.FontLayout(rootDir))
	goFileManager.WriteFile(layoutsFolder, backLayoutFile, layout.BackLayout(rootDir))
	goFileManager.WriteFile(layoutsFolder, adminMenuFile, layout.AdminMenu())
	goFileManager.WriteFile(layoutsFolder, normalMenuFile, layout.NormalMenu())

	goFileManager.WriteFile(frontFolder, indexFile, front.IndexTemplate())
	goFileManager.WriteFile(authFolder, loginFile, auth.LoginTemplate(rootDir))
	goFileManager.WriteFile(authFolder, homeFile, auth.HomeTemplate(rootDir))
	goFileManager.WriteFile(errorFolder, err404File, perr.Error404())
	goFileManager.WriteFile(errorFolder, err404MenuFile, perr.Error404Menu())

	goFileManager.WriteFile(userFolder, userAddFile, user.AddTemplate())
	goFileManager.WriteFile(userFolder, userEditFile, user.EditTemplate())
	goFileManager.WriteFile(userFolder, userShowFile, user.ShowTemplate())
	goFileManager.WriteFile(userFolder, userDetailsFile, user.DetailsTemplate())
	goFileManager.WriteFile(userFolder, userSearchFile, user.SearchTemplate())
	goFileManager.WriteFile(userFolder, userSearchResultsFile, user.SearchResultsTemplate())
	goFileManager.WriteFile(userFolder, userDataFile, user.UserDataTemplate())
	goFileManager.WriteFile(userFolder, userUploadFile, user.UserUploadTemplate())

	goFileManager.WriteFile(taskFolder, taskAddFile, task.AddTemplate())
	goFileManager.WriteFile(taskFolder, taskEditFile, task.EditTemplate())
	goFileManager.WriteFile(taskFolder, taskShowFile, task.ShowTemplate())
	goFileManager.WriteFile(taskFolder, taskSearchFile, task.SearchTemplate())
	goFileManager.WriteFile(taskFolder, taskSearchResultsFile, task.SearchResultsTemplate())
	goFileManager.WriteFile(taskFolder, taskDetailsFile, task.DetailsTemplate())
	goFileManager.WriteFile(taskFolder, userTasksFile, task.UserTasksTemplate())

	goFileManager.WriteFile(roleFolder, roleAddFile, role.AddTemplate())
	goFileManager.WriteFile(taskFolder, roleDetailsFile, role.DetailsTemplate())
	goFileManager.WriteFile(taskFolder, roleShowFile, role.ShowTemplate())
}

*/