package routers

import (
	resource "tools/iris/controllers"

	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/hero"
)

func ResourceHubv2(party router.Party) {
	party.Get("/{resources}", hero.Handler(new(resource.ResourceController).List))
	party.Get("/{resources}/{name}", hero.Handler(new(resource.ResourceController).GetResouce))
	party.Get("/namespaces/{namespace}/{resources}", hero.Handler(new(resource.ResourceController).List))
	party.Get("/namespaces/{namespace}/{resources}/{name}", hero.Handler(new(resource.ResourceController).GetResouce))
	party.Delete("/namespaces/{namespace}/{resources}/{name}", hero.Handler(new(resource.ResourceController).DeleteResource))
}
