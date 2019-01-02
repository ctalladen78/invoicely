package domain

// Item model
type Item struct {
	ID          string `json:"id"`
	Status      string `json:"status"`
	Ticket      string `json:"ticket"`
	Description string `json:"description"`

	// Duration of time spent on this item, in minutes
	Duration uint32 `json:"duration"`
	Sprint   string `json:"sprint"`
}

// ItemRepository interfaces with items table
type ItemRepository struct {
	datastore Datastore
}

// NewItemRepository returns a new instance of item repository
func NewItemRepository(datastore Datastore) *ItemRepository {
	return &ItemRepository{datastore: datastore}
}

// Get a single item
func (r *ItemRepository) Get(id string) (*Item, error) {
	var item *Item
	if err := r.datastore.Get(id, &item); err != nil {
		return nil, err
	}
	return item, nil
}

// List all items
func (r *ItemRepository) List() (*[]Item, error) {
	var items *[]Item
	if err := r.datastore.List(&items); err != nil {
		return nil, err
	}
	return items, nil
}

// Store a new item
func (r *ItemRepository) Store(item *Item) error {
	return r.datastore.Store(item)
}
