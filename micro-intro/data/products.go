package data

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	protos "kokoichi206/go-expart/currency/protos/currency"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/hashicorp/go-hclog"
)

// swagger:model
type Product struct {
	// the id for the product
	//
	// required: false
	// min: 1
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func (p *Product) Validate() error {
	validate := validator.New()
	// add custom validation
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(p)
}

// Custom validation
func validateSKU(fl validator.FieldLevel) bool {
	// xxx-yyyy-zzzzz
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := re.FindAllString(fl.Field().String(), -1)

	if len(matches) != 1 {
		return false
	}

	return true
}

type Products []*Product

type ProductDB struct {
	currency protos.CurrencyClient
	log      hclog.Logger
}

func NewProductDB(c protos.CurrencyClient, l hclog.Logger) *ProductDB {
	return &ProductDB{
		currency: c,
		log:      l,
	}
}

// Cleaner way because it's kind of abstruction
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *ProductDB) GetProducts(currency string) (Products, error) {
	if currency == "" {
		return productList, nil
	}

	rate, err := p.getRate(currency)
	if err != nil {
		p.log.Error("Unable to get rate", "currency", currency, "error", err)
		return nil, err
	}
	// apply rate
	pr := Products{}
	for _, p := range productList {
		// make a copy
		np := *p
		np.Price = np.Price * rate
		pr = append(pr, &np)
	}

	return pr, nil
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}

	p.ID = id
	productList[pos] = p

	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func (p *ProductDB) getRate(destination string) (float64, error) {
	// get exchange rate
	rr := &protos.RateRequest{
		Base:        protos.Currencies(protos.Currencies_value["EUR"]),
		Destination: protos.Currencies(protos.Currencies_value[destination]),
	}
	resp, err := p.currency.GetRate(context.Background(), rr)
	return resp.Rate, err
}

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrProductNotFound
}

func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

var productList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "milky",
		Price:       2.44,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "strong",
		Price:       1.99,
		SKU:         "fjd32",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
