package controllers

import (
	"fmt"
	"github.com/lpisces/ashe/app"
	"github.com/revel/revel"
)

type Admin struct {
	*revel.Controller
}

func (c Admin) Index() revel.Result {
	return c.Render()
}

func (c Admin) Config() revel.Result {
	//rows, _ := app.Engine.Query("select * from config")
	exist, _ := app.Engine.IsTableExist("config")

	return c.RenderText(fmt.Sprintf("%+v", exist))
}
