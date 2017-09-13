package api

type IndexApiHandle struct {
	BaseApiHandle
}

func (this *IndexApiHandle) Get() {
	this.Ctx.WriteString("我是API")
}
