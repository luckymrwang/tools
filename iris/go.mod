module tools/iris

go 1.14

require (
	common.dh.cn/test v0.0.1
	github.com/Joker/hpp v1.0.0 // indirect
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/emicklei/go-restful v2.12.0+incompatible
	github.com/evanphx/json-patch v4.2.0+incompatible // indirect
	github.com/gophercloud/gophercloud v0.1.0 // indirect
	github.com/gorilla/websocket v1.4.2
	github.com/henrylee2cn/mahonia v0.0.0-20150715080413-be6deb105fbc
	github.com/iris-contrib/middleware/cors v0.0.0-20200913183508-5d1bed0e6ea4
	github.com/iris-contrib/swagger/v12 v12.0.1
	github.com/k0kubun/colorstring v0.0.0-20150214042306-9440f1994b88 // indirect
	github.com/kataras/iris/v12 v12.2.0-alpha
	github.com/kr/pty v1.1.5 // indirect
	github.com/mattn/go-colorable v0.0.9 // indirect
	github.com/onsi/gomega v1.8.1 // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/sirupsen/logrus v1.7.0 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/swaggo/swag v1.6.5
	github.com/urfave/cli v1.22.2 // indirect
	github.com/yudai/pp v2.0.1+incompatible // indirect
	github.com/yuin/goldmark v1.1.32 // indirect
	gopkg.in/igm/sockjs-go.v2 v2.1.0
	gopkg.in/mgo.v2 v2.0.0-20180705113604-9856a29383ce
	gorm.io/driver/mysql v1.0.2
	gorm.io/gorm v1.20.2
	helm.sh/helm/v3 v3.3.4
	k8s.io/api v0.18.8
	k8s.io/client-go v0.18.8
	k8s.io/klog v1.0.0 // indirect
	rsc.io/letsencrypt v0.0.3 // indirect
	sigs.k8s.io/structured-merge-diff/v3 v3.0.0 // indirect
)

replace common.dh.cn/test v0.0.1 => ../../common.dh.cn/test
