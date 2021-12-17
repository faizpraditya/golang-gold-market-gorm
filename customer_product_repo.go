package main

type CustomerProductRepo struct {
	BaseRepo
}

func (cr *CustomerProductRepo) Insert(newCustomerProduct CustomerProduct) error {
	result := cr.conn.Db.Create(&newCustomerProduct)
	return cr.HandleError(result)
}

func newCustomerProductRepo(conn *DBConn) *CustomerProductRepo {
	return &CustomerProductRepo{
		BaseRepo{
			conn: conn,
		},
	}
}
