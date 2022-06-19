package main

import (
	"bytes"
	"flag"
	"fmt"
	"os/exec"
	"regexp"
	"time"

	"github.com/vishvananda/netlink"
	"k8s.io/klog"
	utilexec "k8s.io/utils/exec"

	iptables_exec "tools/iris/iptables-exec"
	"tools/iris/routers"

	"github.com/kataras/iris/v12"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @host
// @BasePath /api/v1
func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("main recover:", r)
		}
	}()
	klog.InitFlags(nil)
	//Init the command-line flags.
	flag.Parse()
	defer klog.Flush()

	app := iris.New()

	routers.Init(app)

	//ticker()

	//smallest, err := findHostMTU(MTUIfacePattern)
	//if err == nil {
	//	fmt.Printf("the smallest is %v\n", smallest)
	//} else {
	//	fmt.Printf("err : %v\n", err)
	//}
	app.Run(iris.Addr(":9090"))
}

func p() string {
	time.Sleep(3 * time.Second)
	panic("xxxddd")
}

func ticker() {
	d := time.Duration(time.Second * 30)

	t := time.NewTicker(d)
	defer t.Stop()
	for {
		execer := utilexec.New()
		runner := iptables_exec.New(execer, iptables_exec.ProtocolIpv4)
		exists, err := runner.EnsureRule(iptables_exec.RulePosition("restored"), iptables_exec.TableNAT, iptables_exec.ChainPrerouting)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			return
		}
		fmt.Printf("exists: %v\n", exists)

		doexec()

		<-t.C
	}
}

func doexec() {
	cmd := exec.Command("iptables", "-t", "nat", "-S")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("failed to call cmd.Run(): %v\n", err)
		return
	}
	fmt.Printf("execout:\n%s\nexecerr:\n%s", stdout.String(), stderr.String())
}

// Configures MTU auto-detection.
var MTUIfacePattern = regexp.MustCompile(`^((en|wl|ww|sl|ib)[opsx].*|(eth|wlan|wwan).*)`)

// findHostMTU auto-detects the smallest host interface MTU.
func findHostMTU(matchRegex *regexp.Regexp) (int, error) {
	// Find all the interfaces on the host.
	links, err := netlink.LinkList()
	if err != nil {
		fmt.Printf("Failed to list interfaces. Unable to auto-detect MTU\n")
		return 0, err
	}

	// Iterate through them, keeping track of the lowest MTU.
	smallest := 0
	for _, l := range links {
		// Skip links that we know are not external interfaces.
		if matchRegex == nil || !matchRegex.MatchString(l.Attrs().Name) {
			fmt.Printf("Skipping interface for MTU detection mtu: %v, name: %v\n", l.Attrs().MTU, l.Attrs().Name)
			continue
		}
		fmt.Printf("Examining link for MTU calculation\n")
		if l.Attrs().MTU < smallest || smallest == 0 {
			smallest = l.Attrs().MTU
			fmt.Printf("The smallest is mtu: %v, name: %v\n", l.Attrs().MTU, l.Attrs().Name)
		}
	}

	if smallest == 0 {
		// We failed to find a usable interface. Default the MTU of the host
		// to 1460 - the smallest among common cloud providers.
		fmt.Printf("Failed to auto-detect host MTU - no interfaces matched the MTU interface pattern. To use auto-MTU, set mtuIfacePattern to match your host's interfaces\n")
		return 1460, nil
	}
	return smallest, nil
}
