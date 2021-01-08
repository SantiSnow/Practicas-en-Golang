package main

import (
	"net/http"
)

func main() {

	//rutas
	http.HandleFunc("/", index)
	http.HandleFunc("/contacto", contacto)
	http.HandleFunc("/sobre-nosotros", aboutUs)

	//inicio del servidor
	http.ListenAndServe(":3000", nil)

}

func index(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Bienvenido al index de la aplicación"))
}

func contacto(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Esta es la pagina de contacto"))
}

func aboutUs(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Esta es la pagina sobre nosotros"))
}
