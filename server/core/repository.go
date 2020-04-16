package core

// Repository defines the API repository implementation should follow.
type Repository interface {
	Find(id string) (*WorkingDay, error)
	FindAll(selector map[string]interface{}) ([]*WorkingDay, error)
	Delete(kudo *WorkingDay) error
	Update(kudo *WorkingDay) error
	Create(kudo ...*WorkingDay) error
	Count() (int, error)
}
