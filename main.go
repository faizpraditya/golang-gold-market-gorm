package main

import "log"

func main() {
	// DB masuk ke native (ga pake gorm)
	// db, err := conn.DB()
	// db.Close()

	// Prefer dibungkus env variabel (contoh ada di connection.go)
	// conn.Debug().AutoMigrate(&Customer{})
	// conn.AutoMigrate(&Customer{})

	// NewDbConn()
	db := NewDbConn()
	// Bisa pass by value/reference
	// Dokumentasinya pakai pointer& (pass by reference)
	// db.Db.AutoMigrate(&Customer{})
	// db.Db.AutoMigrate(Customer{})
	defer db.Close()
	// db.Migration(Customer{})
	// db.Migration(&Customer{}, &UserCredential{})
	db.Migration(&Customer{}, &UserCredential{}, &Address{}, CustomerProduct{}, CustomerHasProduct{})

	// Many to many
	// customerProductRepo := newCustomerProductRepo(db)

	// customerProductRepo.Insert(CustomerProduct{
	// 	ID:          "CP001",
	// 	ProductName: "Deposito Rupiah",
	// })
	// customerProductRepo.Insert(CustomerProduct{
	// 	ID:          "CP002",
	// 	ProductName: "Deposito Dollar",
	// })
	// customerProductRepo.Insert(CustomerProduct{
	// 	ID:          "CP003",
	// 	ProductName: "Tabungan Emas",
	// })
	// customerProductRepo.Insert(CustomerProduct{
	// 	ID:          "CP004",
	// 	ProductName: "Tabungan",
	// })

	customerRepo := NewCustomerRepo(db)
	err := customerRepo.OpenProductForExistingCustomer(Customer{
		ID: "23ed591f-69fb-4f27-86f6-01f9f7fce784",
		Products: []CustomerProduct{
			{
				ID: "CP002",
			},
			{
				ID: "CP003",
			},
		},
	})
	if err != nil {
		log.Println(err.Error())
	}

	// 3 table
	// customerRepo := NewCustomerRepo(db)
	// _, err := customerRepo.Insert(Customer{
	// 	ID:        "C001",
	// 	FirstName: "Jution",
	// 	LastName:  "Chandra",
	// 	BirthDate: time.Time{},
	// 	Addresses: []Address{
	// 		{
	// 			ID:         "A003",
	// 			Streetname: "Jl.Jambu",
	// 			City:       "Lampung",
	// 			PostalCode: "111222",
	// 		},
	// 		{
	// 			ID:         "A004",
	// 			Streetname: "Jl.Mangga",
	// 			City:       "Lampung",
	// 			PostalCode: "111222",
	// 		},
	// 	},
	// 	Status: 1,
	// 	UserCredential: UserCredential{
	// 		ID:       "C001",
	// 		Username: "jution",
	// 		Password: "111222",
	// 		Email:    "jution@enigmacamp.com",
	// 		IsActive: true,
	// 	},
	// })
	// if err != nil {
	// 	log.Println(err.Error())
	// }

	// 2 table
	// customerRepo := NewCustomerRepo(db)
	// customer, err := customerRepo.Insert(Customer{
	// 	// ID:        "2e7b11f0-e02b-4a37-8a9d-a1f967ff0862",
	// 	FirstName: "Bruno",
	// 	LastName:  "Aguero",
	// 	BirthDate: time.Time{},
	// 	Address:   "Surabaya",
	// 	Status:    1,
	// 	UserCredential: UserCredential{
	// 		ID:       "001",
	// 		Username: "apabae",
	// 		Password: "1234",
	// 		Email:    "rezabukan@gmail.com",
	// 		IsActive: true,
	// 	},
	// })
	// customer := customerRepo.FindById("aa576d76-cdf5-4d62-afb2-0978c6a828e1")
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// fmt.Println(customer.ToString())

	// customerRepo := NewCustomerRepo(db)
	// Insert
	// err := customerRepo.Insert(Customer{
	// 	ID:        "C001",
	// 	FirstName: "Faiz",
	// 	LastName:  "Praditya",
	// 	BirthDate: time.Time{},
	// 	Address:   "Jakarta",
	// 	Status:    1,
	// })
	// customer, err := customerRepo.Insert(Customer{
	// 	FirstName: "Bruno",
	// 	LastName:  "Aguero",
	// 	BirthDate: time.Time{},
	// 	Address:   "Surabaya",
	// 	Status:    1,
	// })
	// Soft delete
	// err := customerRepo.Delete(Customer{ID: "C001"})
	// Hard delete
	// err := customerRepo.HardDelete(Customer{ID: "C001"})
	// Hard delete by name
	// err := customerRepo.HardDeleteByName(Customer{FirstName: "Faiz"})

	// Find by id
	// customer := customerRepo.FindById("87603f9d-2d52-41ce-a08b-b22128ea6005")
	// first := customerRepo.ShowFirst()
	// customers := customerRepo.FindAll()
	// customers := customerRepo.ShowPage(3, 1)

	// Update
	// customer, err := customerRepo.Update(Customer{
	// 	ID:        "df4f6194-0bf0-4ec5-a3f7-0637c893b6eb",
	// 	FirstName: "Bambang",
	// 	LastName:  "Rambutan",
	// 	BirthDate: time.Time{},
	// 	Address:   "Bandung",
	// 	Status:    1,
	// })

	// if err != nil {
	// 	log.Println(err.Error())
	// }

	// To string: ada keterangan nama field
	// fmt.Println(customer.ToString())
	// fmt.Println(customer)
	// fmt.Println(first.ToString())
	// fmt.Println(first)
	// fmt.Println(customers)

	// fmt.Println("New Customer : ", customer)
	// fmt.Println("Update Customer : ", customer)
}
