package camera_brand

type Service interface {
	CreateCameraBrand(name string) (*CameraBrand, error)
	GetAllCameraBrands() ([]CameraBrand, error)
	GetCameraBrandByID(id uint) (*CameraBrand, error)
	UpdateCameraBrand(id uint, name string) error
	DeleteCameraBrand(id uint) error
}

type service struct {
	repo CameraBrandRepository
}

func NewService(repo CameraBrandRepository) Service {
	return &service{repo: repo}
}

func (s *service) CreateCameraBrand(name string) (*CameraBrand, error) {
	brand := &CameraBrand{Name: name}
	err := s.repo.Create(brand)
	return brand, err
}

func (s *service) GetAllCameraBrands() ([]CameraBrand, error) {
	return s.repo.GetAll()
}

func (s *service) GetCameraBrandByID(id uint) (*CameraBrand, error) {
	return s.repo.GetByID(id)
}

func (s *service) UpdateCameraBrand(id uint, name string) error {
	brand, err := s.repo.GetByID(id)
	if err != nil {
		return err	
	}
	brand.Name = name
	return s.repo.Update(brand)
}

func (s *service) DeleteCameraBrand(id uint) error {
	return s.repo.Delete(id)
}
