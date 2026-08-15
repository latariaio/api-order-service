package service

type Services struct {
	repo *ServiceRepository
}

func NewServices(repo *ServiceRepository) *Services {
	return &Services{repo: repo}
}

func (s *Services) Create(service *Service) error {
	return s.repo.Insert(service)
}

func (s *Services) GetAll() ([]Service, error) {
	services, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return services, nil
}

func (s *Services) GetByID(id string) (*Service, error) {
	service, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return service, nil
}

func (s *Services) Update(id string, service *Service) error {
	err := s.repo.Update(id, service)
	if err != nil {
		return err
	}
	return nil
}

func (s *Services) Delete(id string) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
