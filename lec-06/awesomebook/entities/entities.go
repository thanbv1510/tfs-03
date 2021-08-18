package entities

type BookEntity struct {
	ID       int `gorm:"AUTO_INCREMENT;primary"`
	Name     string
	Author   string
	Category string
	Price    int
}

type CartBook struct {
	BookID int `gorm:"primaryKey;column:book_entity_id"`
	CartID int `gorm:"primaryKey;column:cart_entity_id"`
	Qty    int
}

type CartEntity struct {
	ID           int          `gorm:"AUTO_INCREMENT;primary"`
	UserID       int          `gorm:"column:user_id"`
	BookEntities []BookEntity `gorm:"many2many:cart_books"`
}

type UserEntity struct {
	ID            int           `gorm:"AUTO_INCREMENT;primary"`
	Username      string        `gorm:"unique;"`
	Passwd        string        `gorm:"column:password"`
	Email         string        `gorm:"unique"`
	OrderEntities []OrderEntity `gorm:"foreignKey:user_id"`
	CartEntity    CartEntity    `gorm:"foreignKey:user_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	PaymentEntity PaymentEntity `gorm:"foreignKey:user_id"`
}

type PaymentEntity struct {
	ID      int `gorm:"AUTO_INCREMENT;primary"`
	UserID  int `gorm:"column:user_id"`
	OrderID int `gorm:"column:order_id"`
	Amount  float64
	IsPaid  bool
}

type OrderEntity struct {
	ID     int `gorm:"AUTO_INCREMENT;primary"`
	UserID int `gorm:"column:user_id"`

	PaymentEntity PaymentEntity `gorm:"foreignKey:order_id"`
	BookEntities  []BookEntity  `gorm:"many2many:order_books"`
}

type OrderBook struct {
	BookID  int `gorm:"primaryKey;column:book_entity_id"`
	OrderID int `gorm:"primaryKey;column:order_entity_id"`
	Qty     int
}
