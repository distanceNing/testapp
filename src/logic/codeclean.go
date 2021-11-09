package logic

import "fmt"

// DataGateway 数据层接口
type DataGateway interface {
	Load() error
}

func MakeDataGateway() DataGateway {
	return &DataMapper{}
}

// DataMapper 数据层
type DataMapper struct {
}

func (d *DataMapper) Load() error {
	// panic("implement me")
	println("load data mapper")
	return nil
}

type ServiceInterface interface {
	Do() error
}

func MakeService() ServiceInterface {
	return &ServiceImpl{}
}

// ServiceImpl 逻辑层
type ServiceImpl struct {
	dataLoader DataGateway
}

func (s *ServiceImpl) Do() error {
	return nil
}

type Request struct {
}

type Response struct {
}

// ServiceController 控制层
type ServiceController struct {
	impl ServiceInterface
}

func (ctl *ServiceController) Proc(req *Request, rsp *Response) error {
	err := ctl.impl.Do()
	if err != nil {
		return err
	}
	return nil
}

func NewServiceController() *ServiceController {
	return &ServiceController{impl: MakeService()}
}

// View 数据展示层
type View interface {
	Print(response *Response) error
}

type WebView struct {
}

func (w *WebView) Print(response *Response) error {
	fmt.Printf("rsp:%v", response)
	return nil
}
