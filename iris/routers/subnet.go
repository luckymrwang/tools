package routers

import (
	"tools/iris/controllers"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
)

func SubnetHub(party iris.Party) {
	party.Get("/subnets/clusters/cidr", hero.Handler(new(controllers.SubnetController).GetClusterCIDR))
}
