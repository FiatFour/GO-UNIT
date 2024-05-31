package core

//* Secondary Port (order_repository.go)

type OrderRepository interface { // Spec
	Save(order Order) error // Port
}
