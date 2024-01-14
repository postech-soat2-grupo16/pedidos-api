package main

import (
	"log"
	"net/http"
	"time"
)

import (
	_ "github.com/lib/pq"
	"github.com/postech-soat2-grupo16/pedidos-api/api"
)

//	@title			Fast Food API
//	@version		1.0
//	@description	Here you will find everything you need to have the best possible integration with our APIs.
//	@termsOfService	http://fastfood.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.fastfood.io/support
//	@contact.email	support@fastfood.io

// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	db := api.SetupDB()
	r := api.SetupRouter(db)

	server := &http.Server{
		Addr:              ":8001",
		ReadHeaderTimeout: 3 & time.Second,
		Handler:           r,
	}
	log.Println(server.ListenAndServe())
}
