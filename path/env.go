package path

import (
	"os/exec"
	"strings"
)

//GetModule 获取go.mod中的module值 TODO存在拒绝服务漏洞
//func GetModule() (string, error) {
//	mod, err := gomoddirectives.GetModuleFile()
//	if err != nil {
//		return "", err
//	}
//
//	return mod.Module.Mod.String(), nil
//}

func GetPwdByModule() (string, error) {
	cmd := exec.Command("go", "list")
	raw, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return strings.Replace(string(raw), "\n", "", -1), nil
}
