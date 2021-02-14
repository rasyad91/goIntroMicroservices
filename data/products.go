package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

//Product defintes the structure for an API product
// swagger:model
type Product struct {
	// the id for this user
	// required: false
	// min: 1
	ID int `json:"id"`
	// the name for this user
	// required: true
	Name string `json:"name" validate:"required"`
	// the description of the product
	Description string `json:"description"`
	// the price of the product
	// required: true
	// more than 0
	Price float32 `json:"price" validate:"required,gt=0"`
	// Product SKU (Stock Keeping Unit)
	// required: true
	// example : "abc-def-ghi"
	SKU        string `json:"sku" validate:"required,sku"`
	CreatedOn  string `json:"-"`
	UpdatedOn  string `json:"-"`
	DeletednOn string `json:"-"`
}

// Products slice of Product
type Products []*Product

// Validate validates Product Struct
func (p *Product) Validate() error {
	v := validator.New()
	v.RegisterValidation("sku", validateSKU)
	return v.Struct(p)
}

func validateSKU(fl validator.FieldLevel) bool {
	//sku format abc-abcd-abcde
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := re.FindAllString(fl.Field().String(), -1)
	if len(matches) != 1 {
		return false
	}
	return true
}

// ToJSON serializes the contents of the collection to JSON
// NewEncoder provides better performance than json.Unmarshal as it does not
// have to buffer the output into an in memory slice of bytes
// this reduces allocations and the overheads of the service
//
// https://golang.org/pkg/encoding/json/#NewEncoder
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// FromJSON decodes single Product from JSON
func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

// GetProducts method for main to get Products from data, returns a list of products
func GetProducts() Products {
	return productList
}

// AddProduct appends new product to database
func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

// UpdateProduct updates existing data with respect to ID
func UpdateProduct(id int, p *Product) error {
	pos, err := findProduct(id)
	if err != nil {
		return err
	}
	p.ID = id
	productList[pos] = p
	return nil
}

// DeleteProduct updates existing data with respect to ID
func DeleteProduct(id int) error {
	i, err := findProduct(id)
	if err != nil {
		return err
	}
	productList = append(productList[0:i], productList[i+1:]...)
	return nil
}

func GetProductByID(id int) (Product, error) {
	i, err := findProduct(id)
	if err != nil {
		return Product{}, err
	}
	return *productList[i], nil
}

// ErrProductNotFound - Error
var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (int, error) {
	for i, p := range productList {
		if p.ID == id {
			return i, nil
		}
	}
	return -1, ErrProductNotFound
}

func getNextID() int {
	lp := len(productList) + 1
	return lp
	// lp := productList[len(productList) - 1]
	// return lp.ID + 1
}

var productList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc001",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "def001",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
