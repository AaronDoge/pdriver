package testRedis

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	if val := Get("zhangliang"); val == "malatang" {
		t.Logf("get value = %s", val)
	} else {
		t.Logf("FAIL")
	}
}

func TestSet(t *testing.T) {
	key, val := "Monday", "dongbeicai"

	if err := Set(key, val); err != nil {
		t.Logf("set %s: %s error, %s", key, val, err.Error())
	} else {
		t.Logf("PASS")
	}
}

var clusterClient = InitCodisProxyCluster()

func TestClusterGet(t *testing.T) {
	val := ClusterGet("huoguo", clusterClient)
	fmt.Println("huoguo: ", val)
}
func TestClusterSet(t *testing.T) {
	if err := ClusterSet("coffee", "nestle", clusterClient); err != nil {
		fmt.Println(err.Error())
		return
	}
}