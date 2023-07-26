package filscan

import "testing"

func TestGetStatChainInfo(t *testing.T) {
	ret, err := GetStatChainInfo()
	if err != nil {
		t.Error(err)
	}
	t.Log(ret)

}

func TestGetStatChainInfoV2(t *testing.T) {
	ret, err := GetStatChainInfoV2()
	if err != nil {
		t.Error(err)
	}
	t.Log(ret)

}
