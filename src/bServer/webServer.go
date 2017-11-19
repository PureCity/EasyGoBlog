package bServer

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path"
)

const (
	TEMPLATES_DIR = "templates" //存放模板的文件夹
)

//网页服务类
//负责加载网页模板，启动网站服务
type WebServer struct {
	templates map[string]*template.Template
}

//读取缓存模板
func (web *WebServer) InitTemplates() (err error) {
	web.DisposeTemplates()

	fileInfoArr, err := ioutil.ReadDir(TEMPLATES_DIR)
	if err != nil {
		panic(err)
		return
	}

	var templateName, templatePath string
	for _, fileInfo := range fileInfoArr {
		templateName = fileInfo.Name()
		ext := path.Ext(templateName)
		//过滤非html文件
		if ext != "html" {
			continue
		}

		templatePath = TEMPLATES_DIR + "/" + templateName
		log.Println("读取模板：", templatePath)
		t := template.Must(template.ParseFiles(templatePath))
		//将模板存入缓存
		web.templates[templateName] = t
	}

	return
}

//初始化并释放已有的模板缓存
func (web *WebServer) DisposeTemplates() {
	web.templates = make(map[string]*template.Template)
}

//解析并返回模板页面至客户端
func (web *WebServer) renderHtml(w http.ResponseWriter, page string, locals map[string]interface{}) (err error) {
	err = web.templates[page].Execute(w, locals)
	return
}

func (web *WebServer) defaultHandle(w http.ResponseWriter, r *http.Request) {
	err := web.renderHtml(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//启动网站监听服务
func (web *WebServer) StartServer() (err error) {
	http.HandleFunc("/", web.defaultHandle)
	err = http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("启动服务器失败: [", err.Error(), "]")
	}
	return
}
