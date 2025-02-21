package version

import (
	"flag"
	"fmt"
	"runtime"

	"gitlab.atcatw.org/atca/community-edition/trojan-go/common"
	"gitlab.atcatw.org/atca/community-edition/trojan-go/constant"
	"gitlab.atcatw.org/atca/community-edition/trojan-go/option"
)

type versionOption struct {
	flag *bool
}

func (*versionOption) Name() string {
	return "version"
}

func (*versionOption) Priority() int {
	return 10
}

func (c *versionOption) Handle() error {
	if *c.flag {
		fmt.Println("Trojan-Go", constant.Version)
		fmt.Println("Go Version:", runtime.Version())
		fmt.Println("OS/Arch:", runtime.GOOS+"/"+runtime.GOARCH)
		fmt.Println("Git Commit:", constant.Commit)
		fmt.Println("")
		fmt.Println("Developed by trli (p4gefau1t and fregie)")
		fmt.Println("Licensed under GNU General Public License version 3")
		fmt.Println("GitHub Repository:\thttps://gitlab.atcatw.org/atca/community-edition/trojan-go")
		fmt.Println("Trojan-Go Documents:\thttps://p4gefau1t.github.io/trojan-go/")
		return nil
	}
	return common.NewError("not set")
}

func init() {
	option.RegisterHandler(&versionOption{
		flag: flag.Bool("version", false, "Display version and help info"),
	})
}
