package main

import (
	_ "backend/webapi/design"

	"github.com/shogo82148/goa-v1/design"
	"github.com/shogo82148/goa-v1/goagen/codegen"
	genapp "github.com/shogo82148/goa-v1/goagen/gen_app"
	gencontroller "github.com/shogo82148/goa-v1/goagen/gen_controller"
	genswagger "github.com/shogo82148/goa-v1/goagen/gen_swagger"
)

func main() {
	codegen.ParseDSL()
	codegen.Run(
		genapp.NewGenerator(
			genapp.API(design.Design),
			genapp.OutDir("app"),
			genapp.Target("app"),
		),
		gencontroller.NewGenerator(
			gencontroller.API(design.Design),
			gencontroller.OutDir("."),
			gencontroller.Pkg("api"),
		),
		genswagger.NewGenerator(
			genswagger.API(design.Design),
		),
	)
}
