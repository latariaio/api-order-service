package service_order

type ServiceOrderService struct {
	repo *ServiceOrderRepository
}

func NewServiceOrderService(repo *ServiceOrderRepository) *ServiceOrderService {
	return &ServiceOrderService{repo: repo}
}

func (s *ServiceOrderService) Create(serviceOrder *ServiceOrder) error {
	return s.repo.Insert(serviceOrder)
}

func (s *ServiceOrderService) GetAll() ([]ServiceOrder, error) {
	return s.repo.FindAll()
}

func (s *ServiceOrderService) FindByID(id string) (*ServiceOrder, error) {
	return s.repo.FindByID(id)
}

func (s *ServiceOrderService) Update(id string, serviceOrder *ServiceOrder) error {
	return s.repo.Update(id, serviceOrder)
}

func (s *ServiceOrderService) Delete(id string) error {
	return s.repo.Delete(id)
}
