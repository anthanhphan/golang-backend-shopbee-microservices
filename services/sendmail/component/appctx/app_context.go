package appctx

type AppContext interface {
}

type appCtx struct {
}

func NewAppContext() *appCtx {
	return &appCtx{}
}
