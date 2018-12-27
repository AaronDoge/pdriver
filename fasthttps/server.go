package fasthttps

import (
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
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
