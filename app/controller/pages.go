package controller

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

func StartPage(writer http.ResponseWriter, request *http.Request, params httprouter.Params){
	content := `<!DOCTYPE html>
				<html lang="en">
				<head>
					<meta charset="UTF-8">
					<title>Home</title>
				</head>
				<body>
					<h1>Welcome to the home page of this site!</h1>
					<ul>
						<li>The first element of the bulleted list</li>
						<li>The second element of the bulleted list</li>
						<li>The third element of the bulleted list</li>
					</ul>
				</body>
				</html>`
	//create html-template
	html, err := template.New("example").Parse(content)
	if err != nil {
		http.Error(writer, err.Error(), 400)
		return
	}

	//output the template to the client in the browser
	err = html.Execute(writer, content)
}
