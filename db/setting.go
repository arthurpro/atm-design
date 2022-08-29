package db

type Setting struct {
	Key   string `gorm:"primaryKey"`
	Value float64
}

func (s *Setting) Balance() (float64, error) {
	tx := db.Where("key = 'balance'").First(&s)
	return s.Value, tx.Error
}

func (s *Setting) UpdateBalance(amount float64) error {
	s.Key = "balance"
	s.Value = amount
	tx := db.Save(s)
	return tx.Error
}

func (s *Setting) OverdraftFee() (float64, error) {
	tx := db.Where("key = 'overdraft_fee'").First(&s)
	return s.Value, tx.Error
}

func (s *Setting) BillValue() (float64, error) {
	tx := db.Where("key = 'bill_value'").First(&s)
	return s.Value, tx.Error
}
