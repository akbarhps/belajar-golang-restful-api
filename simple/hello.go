package simple

type SayHello interface {
	SayHello(name string) string
}

type HelloService struct {
	SayHello
}

type SayHelloImpl struct {
}

func (s *SayHelloImpl) SayHello(name string) string {
	return "Hello " + name
}

func NewHelloService(sayHello SayHello) *HelloService {
	return &HelloService{sayHello}
}

func NewSayHelloImpl() *SayHelloImpl {
	return &SayHelloImpl{}
}
