package main;

import (
	"strconv"
	
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func server() {
	// db
	allblue := Allblue{};
	
	
	// app
	app := iris.New();
	app.Get("/", func (ctx context.Context) {
	 	ctx.JSON(iris.Map{"result": "Hello, world!"});
	});

	app.Get("/read", func (ctx context.Context) {
		value := ctx.URLParam("page")
		page, err := strconv.Atoi(value)
		assert(err);
	 	ctx.JSON(allblue.read(page));
	});

	app.Get("/search", func (ctx context.Context) {
		text := ctx.URLParam("text")
	 	ctx.JSON(allblue.search(text));
	});	
		
	app.Run(iris.Addr(":1439"));
}
