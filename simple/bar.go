package simple

type BarRepository struct {
}

func NewBarRepository() *BarRepository {
	return &BarRepository{}
}

type BarService struct {
	*BarRepository
}

func NewBarService(repo *BarRepository) *BarService {
	return &BarService{repo}
}
