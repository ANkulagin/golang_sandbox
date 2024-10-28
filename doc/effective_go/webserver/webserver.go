package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
)

// Флаг для установки адреса сервиса HTTP
var addr = flag.String("addr", ":1718", "http service address") // Q=17, R=18

// Определение HTML-шаблона, который будет использоваться для отображения страницы
var templ = template.Must(template.New("qr").Parse(templateStr))

func main() {
	flag.Parse() // Разбор флагов командной строки
	// Назначение функции QR для обработки корневого маршрута
	http.Handle("/", http.HandlerFunc(QR))
	// Запуск сервера на указанном адресе
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err) // Вывод ошибки при неудачном запуске сервера
	}
}

// Обработчик QR - генерирует HTML с QR-кодом на основе данных формы
func QR(w http.ResponseWriter, req *http.Request) {
	templ.Execute(w, req.FormValue("s"))
}

// HTML-шаблон, который будет отображаться на веб-странице
const templateStr = `
<html>
<head>
<title>QR Link Generator</title>
</head>
<body>
{{if .}}
<img src="https://api.qrserver.com/v1/create-qr-code/?size=300x300&data={{.}}" />
<br>
{{.}}
<br>
<br>
{{end}}
<form action="/" name=f method="GET">
    <input maxLength=1024 size=70 name=s value="" title="Text to QR Encode">
    <input type=submit value="Show QR" name=qr>
</form>
</body>
</html>
`

/*
http://localhost:1718/

*/
