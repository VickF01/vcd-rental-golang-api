package vcd

type Service interface {
	Create(vcdRequest VCDRequest) (VCD, error)
	GetAllVCD() ([]VCD, error)
	GetOneVCD(ID int) (VCD, error)
	UpdateVCD(ID int, updateVCDRequest UpdateVCDRequest) (VCD, error)
	DeleteVCD(ID int) (VCD, error)
}

type service struct{ repository Repository }

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Create(vcdRequest VCDRequest) (VCD, error) {
	vcd := VCD{
		Title:       vcdRequest.Title,
		Price:       vcdRequest.Price,
		Stock:       vcdRequest.Stock,
		Description: vcdRequest.Description,
	}

	newVCD, err := s.repository.Create(vcd)
	return newVCD, err
}

func (s *service) GetAllVCD() ([]VCD, error) {
	vcds, err := s.repository.GetAllVCD()
	return vcds, err
}

func (s *service) GetOneVCD(ID int) (VCD, error) {
	vcd, err := s.repository.GetOneVCD(ID)
	return vcd, err
}

func (s *service) UpdateVCD(ID int, updateVCDRequest UpdateVCDRequest) (VCD, error) {
	vcd, err := s.repository.GetOneVCD(ID)
	if err != nil {
		return VCD{}, err
	}

	if updateVCDRequest.Title != "" {
		vcd.Title = updateVCDRequest.Title
	}
	if updateVCDRequest.Price != 0 {
		vcd.Price = updateVCDRequest.Price
	}
	if updateVCDRequest.Stock != 0 {
		vcd.Stock = updateVCDRequest.Stock
	}
	if updateVCDRequest.Description != "" {
		vcd.Description = updateVCDRequest.Description
	}

	updatedVCD, err := s.repository.UpdateVCD(vcd)
	if err != nil {
		return VCD{}, err
	}

	return updatedVCD, nil
}

func (s *service) DeleteVCD(ID int) (VCD, error) {
	vcd, _ := s.repository.GetOneVCD(ID)
	deletedVCD, err := s.repository.DeleteVCD(vcd)
	return deletedVCD, err
}
