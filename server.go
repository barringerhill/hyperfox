package main;

import (
	"strconv"
	
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/iris-contrib/middleware/cors"
)

func server() {
	// db
	allblue := Allblue{};

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	});
	
	// app
	app := iris.New();

	// v1
	v1 := app.Party("/", crs).AllowMethods(iris.MethodOptions)

	{
		v1.Get("/", func (ctx context.Context) {
			ctx.JSON(iris.Map{"result": "Hello, world!"});
		});

		v1.Get("/read", func (ctx context.Context) {
			value := ctx.URLParam("page");
			if value == "" { value = "0" }
			
			page, err := strconv.Atoi(value);
			assert(err);

			ctx.JSON(allblue.read(page));
		});

		v1.Get("/search", func (ctx context.Context) {
			text := ctx.URLParam("text");
			value := ctx.URLParam("page");
			if value == "" { value = "0" }

			page, err := strconv.Atoi(value);
			assert(err);

			ctx.JSON(allblue.search(text, page));
		});

	}
	
	app.Run(iris.Addr(":1439"));
}
