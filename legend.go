package legend

import (
	"net/http"

	"github.com/gernest/legend/orchestra"
	"github.com/gernest/legend/registry"
	"github.com/gernest/legend/service/config"
	"github.com/gernest/legend/service/mdown"
	"github.com/gernest/legend/service/store"
	"github.com/gernest/legend/service/view"
)

var coreServices = struct {
	Config string
	Store  string
	Mdown  string
	View   string
}{
	"config",
	"store",
	"mdown",
	"view",
}

func Boot() {

	// Rgister all core services
	registry.Register(coreServices.Config, &config.Service{})
	registry.Register(coreServices.Store, &store.DB{})
	registry.Register(coreServices.Mdown, &mdown.Markdown{})
	registry.Register(coreServices.View, &view.Renderer{})

	// Make sure at least one instance of config is running
	// before anything else.
	orchestra.Start(coreServices.Config)
	orchestra.Start(coreServices.Mdown)
	orchestra.Start(coreServices.Store)
	orchestra.Start(coreServices.View)

}

func HomePage(w http.ResponseWriter, r *http.Request) {
}

func NewUser(w http.ResponseWriter, r *http.Request) {

}

func Tweet(w http.ResponseWriter, r *http.Request) {

}

func Users(w http.ResponseWriter, r *http.Request) {

}
