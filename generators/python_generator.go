package generators

import (
	"fmt"
	"github.com/ortizdavid/filemanager/core"
	"github.com/ortizdavid/AppGenCore/helpers"
	"github.com/ortizdavid/AppGenCore/samples/collections"
	dbsamples "github.com/ortizdavid/AppGenCore/samples/db"
	"github.com/ortizdavid/AppGenCore/samples/libs"
	pythonsamples "github.com/ortizdavid/AppGenCore/samples/python"
	"github.com/ortizdavid/AppGenCore/samples/scripts"
)

type PythonGenerator struct {}

var pyFileManager *core.FileManager
var pyImport *pythonsamples.AppImport
var readme *helpers.ReadMePy


func (python *PythonGenerator) GenerateApp(appName string, appType string, db string) {
	switch appType {
	case "mvc":
		python.generateViews(appName)
		python.generateStaticFiles(appName)
		python.generateMvcControllers(appName)
	case "api":
		python.generateApiControllers(appName)
		python.generatePostmanCollections(appName)
	}
	switch db {
	case "mysql":
		python.generateMySqlDB(appName)
	case "postgres":
		python.generatePostgresDB(appName)
	}
	python.generateModels(appName)
	python.generateHelpers(appName)
	python.generateConfig(appName, db)
	python.generateRequirements(appName, db)
	python.generateMain(appName, appType)
	python.generateReadme(appName, db, appType)
	python.generateGitIgnore(appName, appType)
	python.generateUploadsDir(appName)
	python.printInstructions(appName)
}


func (python *PythonGenerator) generateUploadsDir(rootDir string) {
	pyFileManager.CreateManyFolders(rootDir+"/static/uploads/imgs")
	pyFileManager.CreateManyFolders(rootDir+"/static/uploads/docs")
}


func (python *PythonGenerator) printInstructions(appName string) {
	fmt.Println(readme.InstrunctionsBeforeRun(appName))
}

func (python *PythonGenerator) generateConfig(rootDir string, db string) {
	var config *pythonsamples.ConfigPy
	file := "config.py"
	pyFileManager.CreateSigleFile(rootDir, file)
	pyFileManager.WriteFile(rootDir, file, pyImport.ImportForConfig())
	pyFileManager.WriteFile(rootDir, file, config.CreateConfig(db))
}


func (python *PythonGenerator) generateReadme(rootDir string, db string, appType string) {
	file := "README.md"
	pyFileManager.CreateSigleFile(rootDir, file)
	switch appType {
	case "mvc":
		pyFileManager.WriteFile(rootDir, file, readme.ReadmeMVC(db))
	case "api":
		pyFileManager.WriteFile(rootDir, file, readme.ReadmeAPI(db))
	}
}


func (python *PythonGenerator) generateGitIgnore(rootDir string, appType string) {
	var ignore *helpers.GitIgnorePy
	file := ".gitignore"
	pyFileManager.CreateSigleFile(rootDir, file)
	switch appType {
	case "mvc":
		pyFileManager.WriteFile(rootDir, file, ignore.GitIgnoreMvc())
	case "api":
		pyFileManager.WriteFile(rootDir, file, ignore.GitIgnoreAPI())
	}
}


func (python *PythonGenerator) generateMain(rootDir string, appType string) {
	file := "app.py"
	pyFileManager.CreateSigleFile(rootDir, file)
	switch appType {
	case "mvc":
		pyFileManager.WriteFile(rootDir, file, pyImport.ImportForMvcApp())
	case  "api":
		pyFileManager.WriteFile(rootDir, file, pyImport.ImportForRestApi())
	}
	pyFileManager.WriteFile(rootDir, file, pyImport.AppMainCode())
}


func (python *PythonGenerator) generateMySqlDB(rootDir string) {
	var mysql *dbsamples.MySqlDB
	dbDir := rootDir+"/database"
	file := "db_task.sql"
	pyFileManager.CreateManyFolders(dbDir)
	pyFileManager.CreateSigleFile(dbDir, file)
	pyFileManager.WriteFile(dbDir, file, mysql.GetDatabaseScript())
}


func (python *PythonGenerator) generatePostgresDB(rootDir string) {
	var postgres *dbsamples.PostgresDB
	dbDir := rootDir+"/database"
	file := "db_task.sql"
	pyFileManager.CreateManyFolders(dbDir)
	pyFileManager.CreateSigleFile(dbDir, file)
	pyFileManager.WriteFile(dbDir, file, postgres.GetDatabaseScript())
}


func (python *PythonGenerator) generateRequirements(rootDir string, db string) {
	var deps *scripts.PythonDeps
	file := "requirements.txt"
	pyFileManager.CreateSigleFile(rootDir, file)
	pyFileManager.WriteFile(rootDir, file, deps.Requirements(db))
}

func (python *PythonGenerator) generatePostmanCollections(rootDir string) {
	var postman *collectionsamples.Postman
	collectionsFolder := rootDir+"/__postman_collections__"
	collection := "postman_collection.json"
	testRun := "postman_test_run.json"
	pyFileManager.CreateManyFolders(collectionsFolder)
	pyFileManager.CreateManyFiles(collectionsFolder, collection, testRun)
	pyFileManager.WriteFile(collectionsFolder, collection, postman.PostmanCollection(rootDir))
	pyFileManager.WriteFile(collectionsFolder, testRun, postman.PostmanTestRun(rootDir))
}


func (python *PythonGenerator) generateHelpers(rootDir string) {
	var helper *pythonsamples.Helper
	helpersFolder := rootDir+"/helpers"
	httpFile := "http_code.py"
	constFile := "constants.py"
	uploaderFile := "file_uploader.py"
	passwordHandlerFile := "password_handler.py"
	pyFileManager.CreateManyFolders(helpersFolder)
	pyFileManager.CreateSigleFile(helpersFolder, httpFile)
	pyFileManager.CreateSigleFile(helpersFolder, constFile)
	pyFileManager.CreateSigleFile(helpersFolder, uploaderFile)
	pyFileManager.CreateSigleFile(helpersFolder, passwordHandlerFile)
	pyFileManager.WriteFile(helpersFolder, httpFile, helper.HttpCodes())
	pyFileManager.WriteFile(helpersFolder, constFile, helper.Constants())
	pyFileManager.WriteFile(helpersFolder, uploaderFile, helper.FileUploader())
	pyFileManager.WriteFile(helpersFolder, passwordHandlerFile, helper.PasswordHandler())
}


func (python *PythonGenerator) generateModels(rootDir string) {
	var model *pythonsamples.Model
	modelsFolder := rootDir+"/models"
	roleFile := "role.py"
	userFile := "user.py"
	taskFile := "task.py"
	pyFileManager.CreateManyFolders(modelsFolder)
	pyFileManager.CreateSigleFile(modelsFolder, roleFile)
	pyFileManager.CreateSigleFile(modelsFolder, taskFile)
	pyFileManager.CreateSigleFile(modelsFolder, userFile)
	pyFileManager.WriteFile(modelsFolder, roleFile, model.RoleModel())
	pyFileManager.WriteFile(modelsFolder, userFile, model.UserModel())
	pyFileManager.WriteFile(modelsFolder, taskFile, model.TaskModel())
}


func (python *PythonGenerator) generateStaticFiles(rootDir string) {
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

	pyFileManager.CreateManyFolders(staticCss)
	pyFileManager.CreateManyFolders(staticJs)
	pyFileManager.CreateManyFolders(staticImgs)
	pyFileManager.CreateManyFolders(staticLibBootstrapCss)
	pyFileManager.CreateManyFolders(staticLibBootstrapJs)
	pyFileManager.CreateManyFolders(staticLibJquery)
	pyFileManager.CreateSigleFile(staticCss, cssFile)
	pyFileManager.CreateSigleFile(staticJs, jsFile)
	pyFileManager.CreateSigleFile(staticLibBootstrapCss, bootstrapCssFile)
	pyFileManager.CreateSigleFile(staticLibBootstrapJs, bootstrapJsFile)
	pyFileManager.CreateSigleFile(staticLibJquery, jqueryJsFile)
	pyFileManager.WriteFile(staticCss, cssFile, staticFile.CssContent())
	pyFileManager.WriteFile(staticJs, jsFile, staticFile.JsContent())
	pyFileManager.WriteFile(staticLibBootstrapCss, bootstrapCssFile, bootstrap.BootstrapMinCss())
	pyFileManager.WriteFile(staticLibBootstrapJs, bootstrapJsFile, bootstrap.BootstrapMinJs())
	pyFileManager.WriteFile(staticLibJquery, jqueryJsFile, jquery.JqueryMinJs())
}


func (python *PythonGenerator) generateMvcControllers(rootDir string) {
	var mvcController *pythonsamples.MvcController
	controllersFolder := rootDir+"/controllers"
	roleFile := "role_controller.py"
	userFile := "user_controller.py"
	taskFile := "task_controller.py"
	authFile := "auth_controller.py"
	frontFile := "front_controller.py"
	pyFileManager.CreateManyFolders(controllersFolder)
	pyFileManager.CreateSigleFile(controllersFolder, roleFile)
	pyFileManager.CreateSigleFile(controllersFolder, taskFile)
	pyFileManager.CreateSigleFile(controllersFolder, userFile)
	pyFileManager.CreateSigleFile(controllersFolder, authFile)
	pyFileManager.CreateSigleFile(controllersFolder, frontFile)
	pyFileManager.WriteFile(controllersFolder, roleFile, mvcController.RoleController())
	pyFileManager.WriteFile(controllersFolder, userFile, mvcController.UserController())
	pyFileManager.WriteFile(controllersFolder, taskFile, mvcController.TaskController())
	pyFileManager.WriteFile(controllersFolder, authFile, mvcController.AuthController())
	pyFileManager.WriteFile(controllersFolder, frontFile, mvcController.FrontController())
}


func (python *PythonGenerator) generateApiControllers(rootDir string) {
	var apiController *pythonsamples.ApiController
	apiControllersFolder := rootDir+"/api_controllers"
	roleFile := "role_api.py"
	userFile := "user_api.py"
	taskFile := "task_api.py"
	authFile := "auth_api.py"
	rootFile := "root_api.py"
	pyFileManager.CreateManyFolders(apiControllersFolder)
	pyFileManager.CreateSigleFile(apiControllersFolder, roleFile)
	pyFileManager.CreateSigleFile(apiControllersFolder, taskFile)
	pyFileManager.CreateSigleFile(apiControllersFolder, userFile)
	pyFileManager.CreateSigleFile(apiControllersFolder, authFile)
	pyFileManager.CreateSigleFile(apiControllersFolder, rootFile)
	pyFileManager.WriteFile(apiControllersFolder, roleFile, apiController.RoleApiController())
	pyFileManager.WriteFile(apiControllersFolder, userFile, apiController.UserApiController())
	pyFileManager.WriteFile(apiControllersFolder, taskFile, apiController.TaskApiController())
	pyFileManager.WriteFile(apiControllersFolder, authFile, apiController.AuthApiController())
	pyFileManager.WriteFile(apiControllersFolder, rootFile, apiController.RootApiController())
}


func (python *PythonGenerator) generateViews(rootDir string) {

	var layout *pythonsamples.Layout
	var perr *pythonsamples.PageError
	var front *pythonsamples.FrontTemplate
	var auth *pythonsamples.AuthTemplate
	var user *pythonsamples.UserTemplate
	var task *pythonsamples.TaskTemplate
	var role *pythonsamples.RoleTemplate

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

	pyFileManager.CreateManyFolders(templatesFolder)
	pyFileManager.CreateManyFolders(layoutsFolder)
	pyFileManager.CreateManyFolders(frontFolder)
	pyFileManager.CreateManyFolders(authFolder)
	pyFileManager.CreateManyFolders(userFolder)
	pyFileManager.CreateManyFolders(roleFolder)
	pyFileManager.CreateManyFolders(taskFolder)
	pyFileManager.CreateManyFolders(errorFolder)
	
	pyFileManager.CreateManyFiles(errorFolder, err404File, err404MenuFile)
	pyFileManager.CreateSigleFile(frontFolder, indexFile)
	pyFileManager.CreateManyFiles(authFolder, loginFile, homeFile)
	pyFileManager.CreateManyFiles(layoutsFolder, frontLayoutFile, backLayoutFile, normalMenuFile, adminMenuFile)
	pyFileManager.CreateManyFiles(roleFolder, roleAddFile, roleShowFile, roleDetailsFile)
	pyFileManager.CreateManyFiles(taskFolder, taskAddFile, taskEditFile, taskSearchFile, taskSearchResultsFile, taskShowFile, taskDetailsFile, userTasksFile)
	pyFileManager.CreateManyFiles(userFolder, userAddFile, userEditFile, userDataFile, userSearchFile, userSearchResultsFile, userShowFile, userSearchFile, userDetailsFile, userUploadFile)

	pyFileManager.WriteFile(layoutsFolder, frontLayoutFile, layout.FontLayout(rootDir))
	pyFileManager.WriteFile(layoutsFolder, backLayoutFile, layout.BackLayout(rootDir))
	pyFileManager.WriteFile(layoutsFolder, adminMenuFile, layout.AdminMenu())
	pyFileManager.WriteFile(layoutsFolder, normalMenuFile, layout.NormalMenu())

	pyFileManager.WriteFile(frontFolder, indexFile, front.IndexTemplate())
	pyFileManager.WriteFile(authFolder, loginFile, auth.LoginTemplate(rootDir))
	pyFileManager.WriteFile(authFolder, homeFile, auth.HomeTemplate(rootDir))
	pyFileManager.WriteFile(errorFolder, err404File, perr.Error404())
	pyFileManager.WriteFile(errorFolder, err404MenuFile, perr.Error404Menu())

	pyFileManager.WriteFile(userFolder, userAddFile, user.AddTemplate())
	pyFileManager.WriteFile(userFolder, userEditFile, user.EditTemplate())
	pyFileManager.WriteFile(userFolder, userShowFile, user.ShowTemplate())
	pyFileManager.WriteFile(userFolder, userDetailsFile, user.DetailsTemplate())
	pyFileManager.WriteFile(userFolder, userSearchFile, user.SearchTemplate())
	pyFileManager.WriteFile(userFolder, userSearchResultsFile, user.SearchResultsTemplate())
	pyFileManager.WriteFile(userFolder, userDataFile, user.UserDataTemplate())
	pyFileManager.WriteFile(userFolder, userUploadFile, user.UserUploadTemplate())

	pyFileManager.WriteFile(taskFolder, taskAddFile, task.AddTemplate())
	pyFileManager.WriteFile(taskFolder, taskEditFile, task.EditTemplate())
	pyFileManager.WriteFile(taskFolder, taskShowFile, task.ShowTemplate())
	pyFileManager.WriteFile(taskFolder, taskSearchFile, task.SearchTemplate())
	pyFileManager.WriteFile(taskFolder, taskSearchResultsFile, task.SearchResultsTemplate())
	pyFileManager.WriteFile(taskFolder, taskDetailsFile, task.DetailsTemplate())
	pyFileManager.WriteFile(taskFolder, userTasksFile, task.UserTasksTemplate())

	pyFileManager.WriteFile(roleFolder, roleAddFile, role.AddTemplate())
	pyFileManager.WriteFile(taskFolder, roleDetailsFile, role.DetailsTemplate())
	pyFileManager.WriteFile(taskFolder, roleShowFile, role.ShowTemplate())
}