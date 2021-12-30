package path

import (
	"os/exec"
	
	"github.com/ldez/gomoddirectives"
)

//GetModule 获取go.mod中的module值
func GetModule() (string, error) {
	mod, err := gomoddirectives.GetModuleFile()
	if err != nil {
		return "", err
	}

	return mod.Module.Mod.String(), nil
}

func GetPwdByModule() (string, error) {
	cmd := exec.Command("go", "list")
	raw, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(raw), nil
}