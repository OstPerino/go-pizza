package models

type pizzaSize struct {
}

type Pizza struct {
	// TODO: если оставить gorm.Model и сделать миграцию то миграция не отработает
	ID    uint `gorm:"primary_key"`
	Name  string
	Size  string
	Price uint32
}
