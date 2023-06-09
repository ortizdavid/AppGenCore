package helpers

type ReadMeGo struct {}

func (rdm *ReadMeGo) ReadmeAPI(db string)  string {
return `# Go REST API with `+StrDatabase(db)+`
This API manage user tasks
Users can Add, Remove, List and Edit daily tasks
`+rdm.Instrunctions(db)+``
}


func (rdm *ReadMeGo) ReadmeMVC(db string)  string {
return `# Go MVC APP with `+StrDatabase(db)+`
This API manage user tasks
Users can Add, Remove, List and Edit daily tasks
`+rdm.Instrunctions(db)+``
}


func (rdm *ReadMeGo) Instrunctions(db string) string {
return`## Pre requisites:
- `+StrDatabase(db)+`
- Go: version 3.9 

## Steps for run application: 
- Go to project directory: cd ProjectName
- Copy database script [db_task.sql](database/db_task.sql) to database management system
- Install virtual environment: pip install virtualenv
- Create virtualenvironment: virtualenv venv
- Activate virtual environment: venv/Scripts/activate
- Install all dependencies in [requirements.txt](requirements.txt): pip install -r requirements.txt
- Change database connection variables on: [config.go](config.go)
- Run application: go run main.go
- Open browser: http://localhost:5000
- Users for tests: admin01, admin02, user1, user2, 
- Passwords for all users: 12345678`
}

func (rdm *ReadMeGo) InstrunctionsBeforeRun(appName string) string {
return `
Instructions for run `+appName+`:
- Go project directory: cd `+appName+`
- Copy database script 'database/db_task.sql' to database management system
- Install virtual environment: pip install virtualenv
- Create virtualenvironment: virtualenv venv
- Activate virtual environment: venv/Scripts/activate
- Install all dependencies in 'requirements.txt': pip install -r requirements.txt
- Change database connection variables on: 'config.go'
- Run application: go run main.go
- Open browser: http://localhost:5000
- Users for tests: admin01, admin02, user1, user2, ...
- Passwords for all users: 12345678
`	
}