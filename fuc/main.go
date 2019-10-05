package main

import "fmt"

type (
	HelloRepo interface {
		CountNumber()
		ShowInternalNum()
	}
	Service interface {
		Count()
		Print()
	}
	HandlerA struct {
		sa Service
	}
	HandlerB struct {
		sb Service
	}
	OhioService struct {
		repo HelloRepo
	}
	HelloMongoDB struct {
		interNum int
	}
)

func NewHandlerA(service Service) *HandlerA {
	return &HandlerA{
		sa: service,
	}
}

func NewHandlerB(service Service) *HandlerB {
	return &HandlerB{
		sb: service,
	}
}

func (h *HandlerA) ShowNum() {
	fmt.Print(h.sa)
}

func (h *HandlerA) AddMore() {
	h.sa.Count()
}

func (h *HandlerB) ShowNum() {
	fmt.Print(h.sb)
}
func (h *HandlerB) AddMore() {
	h.sb.Count()
}

func NewOhioService(repo HelloRepo) *OhioService {
	return &OhioService{
		repo: repo,
	}
}

func NewHelloMongoDB() *HelloMongoDB {
	return &HelloMongoDB{
		interNum: 1,
	}
}

func (repo *HelloMongoDB) CountNumber() {
	repo.interNum++
}

func (repo *HelloMongoDB) ShowInternalNum() {
	fmt.Println(repo.interNum)
}

func (s *OhioService) Count() {
	s.repo.CountNumber()
}

func (s *OhioService) Print() {
	s.repo.ShowInternalNum()
}

func main() {
	//One time create
	repo := NewHelloMongoDB()
	serive := NewOhioService(repo)

	handlerA := NewHandlerA(serive)
	handlerB := NewHandlerB(serive)

	handlerA.ShowNum()
	handlerB.ShowNum()

}
