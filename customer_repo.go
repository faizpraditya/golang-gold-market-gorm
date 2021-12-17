package main

import "fmt"

type CustomerRepo struct {
	BaseRepo
}

type ResultJoin struct {
	FirstName string
	Username  string
}

// type SelectAll struct {
// 	ID        string
// 	FirstName string
// 	LastName  string
// 	BirthDate string
// 	Address   string
// 	Status    string
// 	Username  string
// 	Password  string
// 	Email     string
// 	IsActive  string
// }

type SelectAll struct {
	Customer
	UserCredential
}

// Sebaiknya ngebalikin customernya, dengan catatan idnya auto increment atau pakai uuid, krn biasanya bakal dipakai di frontend
// func (cr *CustomerRepo) Insert(newCustomer Customer) error {
func (cr *CustomerRepo) Insert(newCustomer Customer) (Customer, error) {
	// Cek validasi manual
	fmt.Println("insert")
	result := cr.conn.Db.Create(&newCustomer)
	// Cek validasi manual
	fmt.Println(newCustomer.ID)
	return newCustomer, cr.HandleError(result)
}

func (cr *CustomerRepo) Delete(customer Customer) error {
	result := cr.conn.Db.Delete(&customer)
	return cr.HandleError(result)
}

func (cr *CustomerRepo) FindById(id string) Customer {
	var customer Customer
	// var userCredential UserCredential
	// var resultJoin ResultJoin
	var selectAll SelectAll
	// result := cr.conn.Db.Find(&customer, "id = ?", id)
	// result := cr.conn.Db.First(&customer, "id = ?", id)
	// UserCredential nama struct
	// Preload 2 kali select, bukan join
	// result := cr.conn.Db.Preload("UserCredential").First(&customer, "id = ?", id)
	// Model harus spesifik ke struct yang mana
	// result := cr.conn.Db.Model(&Customer{}).Select("customers.first_name,mst_user_credential.username").Where("customer_id = ?", id).Joins("join mst_user_credential on mst_user_credential.customer_id = customers.id").Scan(&resultJoin)
	// Bisa pakai select(*)
	// .Select("customers.*,mst_user_credential.*")
	// .Select("*")
	// Kalau mau semua ga perlu pakai select
	result := cr.conn.Db.Model(&Customer{}).Where("customer_id = ?", id).Joins("join mst_user_credential on mst_user_credential.customer_id = customers.id").Scan(&selectAll)
	// fmt.Println("Result join : ", resultJoin)
	fmt.Println("Result join : ", selectAll)
	err := cr.HandleError(result)
	if err != nil {
		return Customer{}
	}
	return customer
}

func (cr *CustomerRepo) ShowFirst() Customer {
	var customer Customer
	result := cr.conn.Db.First(&customer)
	err := cr.HandleError(result)
	if err != nil {
		return Customer{}
	}
	return customer
}

func (cr *CustomerRepo) ShowLast() Customer {
	var customer Customer
	result := cr.conn.Db.Last(&customer)
	err := cr.HandleError(result)
	if err != nil {
		return Customer{}
	}
	return customer
}

func (cr *CustomerRepo) FindAll() []Customer {
	var customer []Customer
	result := cr.conn.Db.Find(&customer)
	err := cr.HandleError(result)
	if err != nil {
		return nil
	}
	return customer
}

func (cr *CustomerRepo) ShowTok(offset int, limit int) []Customer {
	var customer []Customer
	result := cr.conn.Db.Offset(offset).Limit(limit).Find(&customer)
	err := cr.HandleError(result)
	if err != nil {
		return nil
	}
	return customer
}

func (cr *CustomerRepo) ShowPage(page int, limit int) []Customer {
	var customers []Customer
	result := cr.conn.Db.Offset((page - 1) * limit).Limit(limit).Find(&customers)
	err := cr.HandleError(result)
	if err != nil {
		return nil
	}
	return customers

}

func (cr *CustomerRepo) HardDelete(customer Customer) error {
	result := cr.conn.Db.Unscoped().Delete(&customer)
	return cr.HandleError(result)
}

func (cr *CustomerRepo) HardDeleteByName(customer Customer) error {
	result := cr.conn.Db.Unscoped().Delete(&customer, "first_name = ?", customer.FirstName)
	return cr.HandleError(result)
}

func (cr *CustomerRepo) Update(updateCustomerInfo Customer) (Customer, error) {
	// Bisa pakai where tapi harus by ID (kek hard delete)
	// Select("") untuk memilih spesifik kolom untuk update, misal .Select("first_name")
	// Jadi kalau client update zero value bakal update tapi yang zero value dianggap tidak ada perubahan (kalau ga ada select)
	// Kalau pakai select (*) semua bakal diupdate, dan kalau parameter input updatenya tidak ada, maka akan mengupdate dengan default value
	// result := cr.conn.Db.Model(&updateCustomerInfo).Select("*").Updates(updateCustomerInfo)
	// Updates, multiple column
	// Update, single column
	result := cr.conn.Db.Model(&updateCustomerInfo).Updates(updateCustomerInfo)
	err := cr.HandleError(result)
	if err != nil {
		return Customer{}, err
	}
	return updateCustomerInfo, nil
}

func (cr *CustomerRepo) UpdateFirstName(updateCustomerInfo Customer) (Customer, error) {
	result := cr.conn.Db.Model(&updateCustomerInfo).Select("first_name").Updates(updateCustomerInfo)
	err := cr.HandleError(result)
	if err != nil {
		return Customer{}, err
	}
	return updateCustomerInfo, nil
}

// Otomatis isi ke customer_has_product
func (cr *CustomerRepo) OpenProductForExistingCustomer(customerWithProduct Customer) error {
	result := cr.conn.Db.Model(&customerWithProduct).Updates(customerWithProduct)
	return cr.HandleError(result)
}

func NewCustomerRepo(conn *DBConn) *CustomerRepo {
	return &CustomerRepo{BaseRepo{conn: conn}}
}
