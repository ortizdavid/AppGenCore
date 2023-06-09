package helpers

type GitIgnoreGo struct {}

func (ign *GitIgnoreGo) GitIgnoreAPI()  string {
return `Uploads/

*.pyc
__pycache__/
`
}

func (ign *GitIgnoreGo) GitIgnoreMvc()  string {
return `Uploads/

*.pyc
__pycache__/
`
}

