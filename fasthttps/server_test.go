package fasthttps

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init()  {
	Routes()
}

func TestSendJson(t *testing.T)  {
	req, err := http.NewRequest(http.MethodGet, "/sendjson", nil)
	if err != nil {
		t.Fatal("创建请求失败！ERR:", err.Error())
	}

	rw := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(rw, req)

	log.Println("CODE:", rw.Code)
	log.Println("BODY:", rw.Body.String())
}
