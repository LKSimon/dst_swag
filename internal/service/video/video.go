package video

type Service interface {
	i()

	// 定义具体的service方法
	TODO()
}

type service struct{}

func (s *service) i() {}

func (s *service) TODO() {

}

func New() Service {
	return &service{}
}
