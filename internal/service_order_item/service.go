package service_order_item

type ServiceOrderItemService struct {
	repo *ServiceOrderItemRepository
}

func NewServiceOrderItemService(repo *ServiceOrderItemRepository) *ServiceOrderItemService {
	return &ServiceOrderItemService{
		repo: repo,
	}
}

func (s *ServiceOrderItemService) Create(item *ServiceOrderItem) error {
	return s.repo.Insert(item)
}

func (s *ServiceOrderItemService) Update(item *ServiceOrderItem) error {
	return s.repo.Update(item)
}

func (s *ServiceOrderItemService) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s *ServiceOrderItemService) Get(id string) (*ServiceOrderItem, error) {
	return s.repo.FindById(id)
}
