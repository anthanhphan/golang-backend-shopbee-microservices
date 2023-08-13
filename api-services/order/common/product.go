package common

import (
	driver "database/sql/driver"
	"encoding/json"
	"fmt"
)

type Product struct {
	Name     string  `json:"name" gorm:"-"`
	Price    float64 `json:"price" gorm:"-"`
	Quantity int     `json:"quantity" gorm:"-"`
	Image    string  `json:"image" gorm:"-"`
}

func (Product) TableName() string {
	return "products"
}

func (img *Product) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf(fmt.Sprint("Failed to unmarshal JSONB value", value))
	}

	var product Product
	if err := json.Unmarshal(bytes, &product); err != nil {
		return err
	}

	// syntax: *pointer_variable = value | when we change value of a pointer
	*img = product
	return nil
}

// write down DB, must be implement method Value
func (j *Product) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}

	return json.Marshal(j)
}

type Products []Product

func (j *Products) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf(fmt.Sprintf("Failed to unmarshal JSONB %s", value))
	}

	var product []Product
	if err := json.Unmarshal(bytes, &product); err != nil {
		return err
	}

	*j = product
	return nil
}

func (j *Products) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}

	return json.Marshal(j)
}
