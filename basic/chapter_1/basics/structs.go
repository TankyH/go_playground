package basics

import "fmt"

// 接口的定义和实现

// Response interface
type Response interface {
	Data() []byte
}

// ResponseImpl 没有显式的 "继承" Response,
// 仅仅是实现了 Response Interface的方法, 就相当于实现了 Response
type ResponseImpl struct {
}

func (r *ResponseImpl) Data() []byte {
	return nil
}

// Service interface
type Service interface {
	Get() Response
	Post() Response
}

// ServiceImpl Implement Service
type ServiceImpl struct {
	method string
	body   []byte
}

func (s *ServiceImpl) Get() (Response, error) {
	// 这里用的 ResponseImpl return 的 Response interface, 而不是 ResponseImpl
	rsp := ResponseImpl{}
	return &rsp, nil
}

func (s *ServiceImpl) Post() (Response, error) {
	rsp := ResponseImpl{}
	return &rsp, nil
}

func Run() {
	s := ServiceImpl{}
	if response, err := s.Get(); err != nil {
		fmt.Printf("err: %v", err)
	} else {
		fmt.Println(response.Data())
	}
}
