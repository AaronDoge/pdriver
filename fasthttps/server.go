package fasthttps

import (
	"encoding/json"
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"net/http"
)

func httpHadle(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/html")

	fmt.Println(ctx.Request.Body(), "hello fasthttp...")
}

func Start() {
	//	fs := &fasthttp.FS{
	//		Root: 	"./",
	//		GenerateIndexPages: true,
	//		Compress: true,
	//	}
	//
	//	handler := fs.NewRequestHandler()
	//
	//	if err := fasthttp.ListenAndServe(":8015", handler); err != nil {
	//		fmt.Println("error in ListenAndServe: ", err.Error())
	//		os.Exit(1)
	//	} else {
	//		fmt.Println("server started, listening prot: 8015...")
	//	}
	//}
	router := fasthttprouter.New()

	router.GET("/", httpHadle)
	router.GET("/find", func(ctx *fasthttp.RequestCtx) {
		fmt.Println("find the result...")
	})

	fasthttp.ListenAndServe(":8090", router.Handler)


}

func Routes() {
	http.HandleFunc("/sendjson", func(rw http.ResponseWriter, req *http.Request) {
		param := struct {
			Name 	string
		}{
			Name: "aaron.chen",
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode(param)
	})
}
