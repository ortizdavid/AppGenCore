package generators

import (
	"github.com/ortizdavid/go-nopain/filemanager"
	//"github.com/ortizdavid/AppGenCore/helpers"
	//dbsamples "github.com/ortizdavid/AppGenCore/samples/db"
	phpsamples "github.com/ortizdavid/AppGenCore/samples/php"
	//"github.com/ortizdavid/AppGenCore/samples/scripts"
)

type PhpGenerator struct {}

var phpFileManager *filemanager.FileManager
var phpImport *phpsamples.AppImport


func (php *PhpGenerator) GenerateApp(appName string, appType string, db string) {
	/*
	switch appType {
	case "mvc":
		php.generateViews(appName)
		php.generateStaticFiles(appName)
		php.generateMvcControllers(appName)
	case "api":
		php.generateApiControllers(appName)
		php.generatePostmanCollections(appName)
	}
	switch db {
	case "mysql":
		php.generateMySqlDB(appName)
	case "postgres":
		php.generatePostgresDB(appName)
	}
	php.generateModels(appName)
	php.generateHelpers(appName)
	php.generateConfig(appName, db)
	php.generateRequirements(appName, db)
	php.generateMain(appName, appType)
	php.generateReadme(appName, db, appType)
	php.generateGitIgnore(appName, appType)
	php.generateUploadsDir(appName)
	php.printInstructions(appName)
	*/
}

/*

func (php *PhpGenerator) generateUploadsDir(rootDir string) {
	goFileManager.CreateManyFolders(rootDir+"/static/uploads/imgs")
	goFileManager.CreateManyFolders(rootDir+"/static/uploads/docs")
}


func (php *PhpGenerator) printInstructions(appName string) {
	fmt.Println(readme.InstrunctionsBeforeRun(appName))
}

func (php *PhpGenerator) generateConfig(rootDir string, db string) {
	var config *phpsamples.ConfigGo
	file := "config.php"
	goFileManager.CreateSingleFile(rootDir, file)
	goFileManager.WriteFile(rootDir, file, goImport.ImportForConfig())
	goFileManager.WriteFile(rootDir, file, config.CreateConfig(db))
}


func (php *PhpGenerator) generateReadme(rootDir string, db string, appType string) {
	file := "README.md"
	goFileManager.CreateSingleFile(rootDir, file)
	switch appType {
	case "mvc":
		goFileManager.WriteFile(rootDir, file, readme.ReadmeMVC(db))
	case "api":
		goFileManager.WriteFile(rootDir, file, readme.ReadmeAPI(db))
	}
}


func (php *PhpGenerator) generateGitIgnore(rootDir string, appType string) {
	var ignore *helpers.GitIgnoreGo
	file := ".gitignore"
	goFileManager.CreateSingleFile(rootDir, file)
	switch appType {
	case "mvc":
		goFileManager.WriteFile(rootDir, file, ignore.GitIgnoreMvc())
	case "api":
		goFileManager.WriteFile(rootDir, file, ignore.GitIgnoreAPI())
	}
}


func (php *PhpGenerator) generateMain(rootDir string, appType string) {
	file := "app.php"
	goFileManager.CreateSingleFile(rootDir, file)
	switch appType {
	case "mvc":
		goFileManager.WriteFile(rootDir, file, goImport.ImportForMvcApp())
	case  "api":
		goFileManager.WriteFile(rootDir, file, goImport.ImportForRestApi())
	}
	goFileManager.WriteFile(rootDir, file, goImport.AppMainCode())
}


func (php *PhpGenerator) generateMySqlDB(rootDir string) {
	var mysql *dbsamples.MySqlDB
	dbDir := rootDir+"/database"
	file := "db_task.sql"
	goFileManager.CreateManyFolders(dbDir)
	goFileManager.CreateSingleFile(dbDir, file)
	goFileManager.WriteFile(dbDir, file, mysql.GetDatabaseScript())
}


func (php *PhpGenerator) generatePostgresDB(rootDir string) {
	var postgres *dbsamples.PostgresDB
	dbDir := rootDir+"/database"
	file := "db_task.sql"
	goFileManager.CreateManyFolders(dbDir)
	goFileManager.CreateSingleFile(dbDir, file)
	goFileManager.WriteFile(dbDir, file, postgres.GetDatabaseScript())
}


func (php *PhpGenerator) generateRequirements(rootDir string, db string) {
	var deps *scripts.GothonDeps
	file := "requirements.txt"
	goFileManager.CreateSingleFile(rootDir, file)
	goFileManager.WriteFile(rootDir, file, deps.Requirements(db))
}

func (php *PhpGenerator) generatePostmanCollections(rootDir string) {
	var postman *collectionsamples.Postman
	collectionsFolder := rootDir+"/__postman_collections__"
	collection := "postman_collection.json"
	testRun := "postman_test_run.json"
	goFileManager.CreateManyFolders(collectionsFolder)
	goFileManager.CreateManyFiles(collectionsFolder, collection, testRun)
	goFileManager.WriteFile(collectionsFolder, collection, postman.PostmanCollection(rootDir))
	goFileManager.WriteFile(collectionsFolder, testRun, postman.PostmanTestRun(rootDir))
}


func (php *PhpGenerator) generateHelpers(rootDir string) {
	var helper *phpsamples.Helper
	helpersFolder := rootDir+"/helpers"
	httpFile := "http_code.php"
	constFile := "constants.php"
	uploaderFile := "file_uploader.php"
	passwordHandlerFile := "password_handler.php"
	goFileManager.CreateManyFolders(helpersFolder)
	goFileManager.CreateSingleFile(helpersFolder, httpFile)
	goFileManager.CreateSingleFile(helpersFolder, constFile)
	goFileManager.CreateSingleFile(helpersFolder, uploaderFile)
	goFileManager.CreateSingleFile(helpersFolder, passwordHandlerFile)
	goFileManager.WriteFile(helpersFolder, httpFile, helper.HttpCodes())
	goFileManager.WriteFile(helpersFolder, constFile, helper.Constants())
	goFileManager.WriteFile(helpersFolder, uploaderFile, helper.FileUploader())
	goFileManager.WriteFile(helpersFolder, passwordHandlerFile, helper.PasswordHandler())
}


func (php *PhpGenerator) generateModels(rootDir string) {
	var model *phpsamples.Model
	modelsFolder := rootDir+"/models"
	roleFile := "role.php"
	userFile := "user.php"
	taskFile := "task.php"
	goFileManager.CreateManyFolders(modelsFolder)
	goFileManager.CreateSingleFile(modelsFolder, roleFile)
	goFileManager.CreateSingleFile(modelsFolder, taskFile)
	goFileManager.CreateSingleFile(modelsFolder, userFile)
	goFileManager.WriteFile(modelsFolder, roleFile, model.RoleModel())
	goFileManager.WriteFile(modelsFolder, userFile, model.UserModel())
	goFileManager.WriteFile(modelsFolder, taskFile, model.TaskModel())
}


func (php *PhpGenerator) generateStaticFiles(rootDir string) {
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
	goFileManager.CreateSingleFile(staticCss, cssFile)
	goFileManager.CreateSingleFile(staticJs, jsFile)
	goFileManager.CreateSingleFile(staticLibBootstrapCss, bootstrapCssFile)
	goFileManager.CreateSingleFile(staticLibBootstrapJs, bootstrapJsFile)
	goFileManager.CreateSingleFile(staticLibJquery, jqueryJsFile)
	goFileManager.WriteFile(staticCss, cssFile, staticFile.CssContent())
	goFileManager.WriteFile(staticJs, jsFile, staticFile.JsContent())
	goFileManager.WriteFile(staticLibBootstrapCss, bootstrapCssFile, bootstrap.BootstrapMinCss())
	goFileManager.WriteFile(staticLibBootstrapJs, bootstrapJsFile, bootstrap.BootstrapMinJs())
	goFileManager.WriteFile(staticLibJquery, jqueryJsFile, jquery.JqueryMinJs())
}


func (php *PhpGenerator) generateMvcControllers(rootDir string) {
	var mvcController *phpsamples.MvcController
	controllersFolder := rootDir+"/controllers"
	roleFile := "role_controller.php"
	userFile := "user_controller.php"
	taskFile := "task_controller.php"
	authFile := "auth_controller.php"
	frontFile := "front_controller.php"
	goFileManager.CreateManyFolders(controllersFolder)
	goFileManager.CreateSingleFile(controllersFolder, roleFile)
	goFileManager.CreateSingleFile(controllersFolder, taskFile)
	goFileManager.CreateSingleFile(controllersFolder, userFile)
	goFileManager.CreateSingleFile(controllersFolder, authFile)
	goFileManager.CreateSingleFile(controllersFolder, frontFile)
	goFileManager.WriteFile(controllersFolder, roleFile, mvcController.RoleController())
	goFileManager.WriteFile(controllersFolder, userFile, mvcController.UserController())
	goFileManager.WriteFile(controllersFolder, taskFile, mvcController.TaskController())
	goFileManager.WriteFile(controllersFolder, authFile, mvcController.AuthController())
	goFileManager.WriteFile(controllersFolder, frontFile, mvcController.FrontController())
}


func (php *PhpGenerator) generateApiControllers(rootDir string) {
	var apiController *phpsamples.ApiController
	apiControllersFolder := rootDir+"/api_controllers"
	roleFile := "role_api.php"
	userFile := "user_api.php"
	taskFile := "task_api.php"
	authFile := "auth_api.php"
	rootFile := "root_api.php"
	goFileManager.CreateManyFolders(apiControllersFolder)
	goFileManager.CreateSingleFile(apiControllersFolder, roleFile)
	goFileManager.CreateSingleFile(apiControllersFolder, taskFile)
	goFileManager.CreateSingleFile(apiControllersFolder, userFile)
	goFileManager.CreateSingleFile(apiControllersFolder, authFile)
	goFileManager.CreateSingleFile(apiControllersFolder, rootFile)
	goFileManager.WriteFile(apiControllersFolder, roleFile, apiController.RoleApiController())
	goFileManager.WriteFile(apiControllersFolder, userFile, apiController.UserApiController())
	goFileManager.WriteFile(apiControllersFolder, taskFile, apiController.TaskApiController())
	goFileManager.WriteFile(apiControllersFolder, authFile, apiController.AuthApiController())
	goFileManager.WriteFile(apiControllersFolder, rootFile, apiController.RootApiController())
}


func (php *PhpGenerator) generateViews(rootDir string) {

	var layout *phpsamples.Layout
	var perr *phpsamples.PageError
	var front *phpsamples.FrontTemplate
	var auth *phpsamples.AuthTemplate
	var user *phpsamples.UserTemplate
	var task *phpsamples.TaskTemplate
	var role *phpsamples.RoleTemplate

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
	goFileManager.CreateSingleFile(frontFolder, indexFile)
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