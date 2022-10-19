package renderers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)
var functions = template.FuncMap{

}
func RenderTemplate(w http.ResponseWriter, tmpl string){

	wd,error := os.Getwd()
if error != nil {
    panic(error)
}
root := filepath.Clean(filepath.Join(wd, "../../"))
fmt.Println(root)
parent := filepath.Dir(wd)
fmt.Println(parent)
	 tc,err := CreateTemplateCache(w)
	if err!= nil {
        log.Fatal("Error ceating cahce " ,err)
    }
	t :=tc[tmpl]
	
	parsedTemplate ,_ := template.ParseFiles(root+"/templates/"+tmpl)
	err = parsedTemplate.Execute(w,nil)

	if err != nil {
		fmt.Println("error parsing template" , err)
		return
	}

}

func CreateTemplateCache(w http.ResponseWriter) (map[string]*template.Template,error){
	wd,error := os.Getwd()
	if error != nil {
		panic(error)
	}
	root := filepath.Clean(filepath.Join(wd, "../../"))
	myCache := map[string]*template.Template{};
	pages , error := filepath.Glob(root+"/templates/*page.tmpl")
	if error !=nil {
		return myCache,error;
	}
	for _, page := range pages {
		name := filepath.Base(page);
		fmt.Println("page is currently " ,name)

		ts,err := template.New(name).Funcs(functions).ParseFiles(page)
		if (err!=nil){
			return myCache,err;
		}
		matches , err :=filepath.Glob(root+"/templates/*page.tmpl")
		if err !=nil {
			return myCache,err
		}
		if len(matches) > 0 {
			ts,err = template.ParseGlob(root+"/templates/*page.tmpl")
			if err !=nil {
				return myCache,err
			}
		}
		myCache[name]=ts
	}
	return nil,nil;
}