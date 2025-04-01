package domain

type CustomerRepository interface {
	// CreateCustomer creates a new customer
	CreateCustomer(customer Customer) (Customer, error)
	// GetCustomer retrieves a customer by ID
	GetCustomer(id string) (Customer, error)
	// UpdateCustomer updates an existing customer
	UpdateCustomer(customer Customer) (Customer, error)
	// DeleteCustomer deletes a customer by ID
	DeleteCustomer(id string) error
	// ListCustomers retrieves a list of customers with pagination
	ListCustomers(page, pageSize int) ([]Customer, error)
}
