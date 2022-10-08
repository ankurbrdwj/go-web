package renderers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

func RenderTemplate(w http.ResponseWriter, tmpl string){

	wd,error := os.Getwd()
if error != nil {
    panic(error)
}
root := filepath.Clean(filepath.Join(wd, "../../"))
fmt.Println(root)
parent := filepath.Dir(wd)
fmt.Println(parent)

	parsedTemplate ,_ := template.ParseFiles(root+"/templates/"+tmpl)
	err := parsedTemplate.Execute(w,nil)

	if err != nil {
		fmt.Println("error parsing template" , err)
		return
	}

}