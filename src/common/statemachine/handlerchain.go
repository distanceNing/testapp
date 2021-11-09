package statemachine

type Handler interface {
	Handle(req interface{}, rsp interface{}, ext map[string]interface{}) error
	// Skip(req interface{}, rsp interface{}, ext map[string]interface{}) bool
}

type HandlerChain struct {
	handlerList []Handler
}

func NewHandlerChain() *HandlerChain {
	return &HandlerChain{handlerList: make([]Handler, 0)}
}

func (hc *HandlerChain) addHandler(h Handler) {
	hc.handlerList = append(hc.handlerList, h)
}

func (hc *HandlerChain) Do(req interface{}, rsp interface{}, ext map[string]interface{}) error {
	for i := range hc.handlerList {
		err := hc.handlerList[i].Handle(req, rsp, ext)
		if err != nil {
			return err
		}
	}
	return nil
}
