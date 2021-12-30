package path

import "testing"

func TestGetModule(t *testing.T) {
	got, err := GetModule()
	if  err != nil  {
		t.Errorf("GetModule() error = %v ", err )
		return
	}
	t.Logf("module:%s",got)
}

func TestGetPwdByModule(t *testing.T) {
	got, err := GetPwdByModule()
	if err != nil {
		t.Errorf("GetPwdByModule() error = %v", err)
		return
	}
	t.Logf("path:%s",got)
}
