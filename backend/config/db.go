package config

import (
	"fmt"
	"project-se/entities"
	"project-se/entity"
	"time"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

// ฟังก์ชันคืนค่า Database Instance
func DB() *gorm.DB {
	return db
}

// ฟังก์ชันเชื่อมต่อฐานข้อมูล
func ConnectionDB() {
	database, err := gorm.Open(sqlite.Open("cabana.db?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connected to the database")
	db = database
}

// ฟังก์ชันตั้งค่าโครงสร้างฐานข้อมูลและเพิ่มข้อมูลเริ่มต้น
func SetupDatabase() {
	// AutoMigrate สำหรับสร้างตาราง
	db.AutoMigrate(
		&entity.Gender{},
		&entity.Roles{}, // เพิ่ม Role เข้าไปในระบบ
		&entity.DriverStatus{},
		&entity.VehicleStatus{},
		&entity.Position{},
		&entity.Employee{},
		&entity.Passenger{},
		&entity.Driver{},
		&entity.Message{},
		&entity.Booking{},
		&entity.Location{},
		&entity.VehicleType{},
		&entity.Vehicle{},
		&entity.StartLocation{},
		&entity.Destination{},
		&entity.DiscountType{},
		&entity.Promotion{},
		&entity.StatusPromotion{},
		&entity.Withdrawal{},
		&entity.Commission{},
		&entity.BankName{},
		&entity.TrainBook{},
		&entity.Trainers{},
		&entity.Rooms{},
		&entity.BankName{},
		&entity.RoomChat{},
		&entity.BookingStatus{},
		&entity.NametypeVechicle{},
		&entities.Payment{},
		&entities.Review{},
		&entities.Paid{},
		
	)

	GenderMale := entity.Gender{Gender: "Male"}
	GenderFemale := entity.Gender{Gender: "Female"}
	db.FirstOrCreate(&GenderMale, &entity.Gender{Gender: "Male"})
	db.FirstOrCreate(&GenderFemale, &entity.Gender{Gender: "Female"})

	// ข้อมูลตำแหน่งที่ต้องการเพิ่ม
	locations := []entity.Location{
		{Latitude: 14.989440874562565, Longitude: 102.09469233129263, Province: "นครราชสีมา", Place: "สถานีขนส่งผู้โดยสารจังหวัดนครราชสีมา แห่งที่2",DriverID:1},
		{Latitude: 14.97226216361242, Longitude: 102.07854163104108, Province: "นครราชสีมา", Place: "สถานีรถไฟโคราช",DriverID:2},
		{Latitude: 14.980969671175174, Longitude: 102.07643761780784, Province: "นครราชสีมา", Place: "เดอะมอลล์ โคราช",DriverID:3},
		{Latitude: 14.98183787602261, Longitude: 102.09010039126157, Province: "นครราชสีมา", Place: "เทอร์มินอล 21 โคราช",DriverID:4},
		{Latitude: 14.996281374785447, Longitude: 102.11693411904838, Province: "นครราชสีมา", Place: "เซ็นทรัล โคราช",DriverID:5},
		{Latitude: 14.901746803513126, Longitude: 102.00956884715538, Province: "นครราชสีมา", Place: "Café Amazon สาขา มทส. ประตู 4",DriverID:6},
		{Latitude: 14.978256144038262, Longitude: 102.09254730290546, Province: "นครราชสีมา", Place: "สถานีขนส่งนครราชสีมา"},
		{Latitude: 14.974824485355242, Longitude: 102.0981385474978, Province: "นครราชสีมา", Place: "Café Class ใกล้ลานย่าโม",},
		{Latitude: 14.986847325609906, Longitude: 102.09175265877519, Province: "นครราชสีมา", Place: "โรงเรียนอัสสัมชัญนครราชสีมา"},
		{Latitude: 13.745983305017283, Longitude: 100.5343802441482, Province: "กรุงเทพมหานคร", Place: "สยามพารากอน"},
		{Latitude: 13.98919288476311, Longitude: 100.61774675399516, Province: "กรุงเทพมหานคร", Place: "ฟิวเจอร์พาร์ค รังสิต แอน สเปลล์"},
		{Latitude: 13.813782036370695, Longitude: 100.54976354819318, Province: "กรุงเทพมหานคร", Place: "หน้าสถานีขนส่งหมอชิต 2"},
		{Latitude: 13.816038542388675, Longitude: 100.7251641441578, Province: "กรุงเทพมหานคร", Place: "ตลาดจตุจักร 2 (เมืองมีน)"},
	}

	// เพิ่มข้อมูล Location ลงในฐานข้อมูล
	for _, location := range locations {
		db.FirstOrCreate(&location, entity.Location{Latitude: location.Latitude, Longitude: location.Longitude, Place: location.Place})
	}


	statuses := []entity.DriverStatus{
		{
			Status: "Active",//ใช้บ่งบอกว่าสถานะของ Drivers
		},
		{
			Status: "Inactive",
		},
		{
			Status: "Accepting work",
		},
		{
			Status: "Finish work",
		},
	}
	
	for _, status := range statuses {
		db.Create(&status)
	}
	
	vehiclestatuses := []entity.VehicleStatus{
		{
			Status: "Active",
		},
		{
			Status: "Inactive",
		},
		{
			Status: "Maintenance",
		},
	}
	
	for _, status := range vehiclestatuses {
		db.Create(&status)
	}

	promotions := []entity.Promotion{
		{
			PromotionCode:        "DRIVE001",
			PromotionName:        "ส่งฟรี ไม่มีข้อแม้ !",
			PromotionDescription: "ฟรีสำหรับลูกค้าทุกท่าน",
			Discount:             100, // คิดเป็นส่วนลดเต็ม 100%
			EndDate:              time.Now().Add(30 * 24 * time.Hour),
			UseLimit:             88,
			UseCount:             0,
			DistancePromotion :   0.0,
			Photo:                "https://i.imgur.com/qsk8kle.gif",
			DiscountTypeID:       2, // Percent discount
			StatusPromotionID:    1, // ACTIVE
			DistanceCondition: "free",
		},
		{
			PromotionCode:        "DRIVE002",
			PromotionName:        "แค่ 5 กม. ก็ลดเลย!",
			PromotionDescription: "เดินทางในระยะทาง 5 กม. ขึ้นไป ลดทันที 50 บาท",
			Discount:             50.0,
			EndDate:              time.Now().Add(60 * 24 * time.Hour),
			UseLimit:             34,
			UseCount:             0,
			DistancePromotion :             5.0,
			Photo:                "https://i.imgur.com/PHZXXIY.gif",
			DiscountTypeID:       1, // Amount discount
			StatusPromotionID:             1, // ACTIVE
			DistanceCondition: "greater_equal",
		},
		{
			PromotionCode:        "DRIVE003",
			PromotionName:        "ระยะทางไกลก็ลดให้!",
			PromotionDescription: "รับส่วนลด 15% สำหรับการเดินทางในระยะทาง 20 กม. ขึ้นไป",
			Discount:             15.0,
			EndDate:              time.Now().Add(90 * 24 * time.Hour),
			UseLimit:             28,
			UseCount:             0,
			DistancePromotion :             20.0,
			Photo:                "https://i.imgur.com/Xp7h3Zn.gif",
			DiscountTypeID:       2, // Percent discount
			StatusPromotionID:              1, // ACTIVE
			DistanceCondition: "greater_equal",
		},
		{
			PromotionCode:        "DRIVE004",
			PromotionName:        "ยิ่งขยับ ยิ่งลด!",
			PromotionDescription: "รับส่วนลด 30 บาทเมื่อเดินทางในระยะทางเกิน 3 กม.",
			Discount:             30.0,
			EndDate:              time.Now().Add(120 * 24 * time.Hour),
			UseLimit:             72,
			UseCount:             0,
			DistancePromotion:             3.0,
			Photo:                "https://i.imgur.com/rkxFKdU.gif",
			DiscountTypeID:       1, // Amount discount
			StatusPromotionID:             1, // ACTIVE
			DistanceCondition: "greater",
		},
		{
			PromotionCode:        "DRIVE005",
			PromotionName:        "8 กม. ส่งฟรีจ้า",
			PromotionDescription: "รับบริการส่งฟรีเมื่อระยะทางไม่เกิน 8 กม.",
			Discount:             100.0, // คิดเป็นส่วนลดเต็ม 100%
			EndDate:              time.Now().Add(45 * 24 * time.Hour),
			UseLimit:             57,
			UseCount:             0,
			DistancePromotion :             8.0,
			Photo:                "https://i.imgur.com/7gbPAUZ.gif",
			DiscountTypeID:       2, // Percent discount
			StatusPromotionID:             1, // ACTIVE
			DistanceCondition: "less_equal",
		},
		{
			PromotionCode:        "DRIVE006",
			PromotionName:        "15 กม. ลดให้เลย 20%",
			PromotionDescription: "รับส่วนลด 20% สำหรับการเดินทางที่ระยะทางขั้นต่ำ 15 กม.",
			Discount:             20.0,
			EndDate:              time.Now().Add(180 * 24 * time.Hour),
			UseLimit:             52,
			UseCount:             0,
			DistancePromotion :             15.0,
			Photo:                "https://i.imgur.com/wNNnRDL.gif",
			DiscountTypeID:       2, // Percent discount
			StatusPromotionID:              2, // ACTIVE
			DistanceCondition: "greater_equal",
		},
	}
	// บันทึกข้อมูลโปรโมชั่นตัวอย่างลงในฐานข้อมูล
	for _, promo := range promotions {
		db.FirstOrCreate(
			&promo,
			entity.Promotion{
				PromotionCode:     promo.PromotionCode,
				DistanceCondition: promo.DistanceCondition, // เพิ่มเงื่อนไขนี้
			},
		)
	}

	// สร้างข้อมูลตัวอย่าง Status
	ActiveStatus := entity.StatusPromotion{StatusPromotion: "ACTIVE"}
	ExpiredStatus := entity.StatusPromotion{StatusPromotion: "EXPIRED"}
	db.FirstOrCreate(&ActiveStatus, &entity.StatusPromotion{StatusPromotion: "ACTIVE"})
	db.FirstOrCreate(&ExpiredStatus, &entity.StatusPromotion{StatusPromotion: "EXPIRED"})

	// สร้างข้อมูลตัวอ ย่าง DiscountType
	AmountDiscount := entity.DiscountType{DiscountType: "amount"}
	PercentDiscount := entity.DiscountType{DiscountType: "percent"}
	db.FirstOrCreate(&AmountDiscount, &entity.DiscountType{DiscountType: "amount"})
	db.FirstOrCreate(&PercentDiscount, &entity.DiscountType{DiscountType: "percent"})

	// สร้างข้อมูลตัวอย่าง BankName
	BankBangkok := entity.BankName{BankName: "ธนาคารกรุงเทพ"}
	BankKasikorn := entity.BankName{BankName: "ธนาคารกสิกรไทย"}
	BankSCB := entity.BankName{BankName: "ธนาคารไทยพาณิชย์"}
	BankKrungthai := entity.BankName{BankName: "ธนาคารกรุงไทย"}
	BankTMB := entity.BankName{BankName: "ธนาคารทหารไทย"}
	// ใช้ FirstOrCreate เพื่อป้องกันการสร้างข้อมูลซ้ำ
	db.FirstOrCreate(&BankBangkok, &entity.BankName{BankName: "ธนาคารกรุงเทพ"})
	db.FirstOrCreate(&BankKasikorn, &entity.BankName{BankName: "ธนาคารกสิกรไทย"})
	db.FirstOrCreate(&BankSCB, &entity.BankName{BankName: "ธนาคารไทยพาณิชย์"})
	db.FirstOrCreate(&BankKrungthai, &entity.BankName{BankName: "ธนาคารกรุงไทย"})
	db.FirstOrCreate(&BankTMB, &entity.BankName{BankName: "ธนาคารทหารไทย"})

	
	// สร้าง Position
	PositionOwner := entity.Position{Position: "Owner"}
	PositionEmployee := entity.Position{Position: "Employee"}
	PositionAdmin := entity.Position{Position: "Admin"}
	db.FirstOrCreate(&PositionOwner, &entity.Position{Position: "Owner"})
	db.FirstOrCreate(&PositionEmployee, &entity.Position{Position: "Employee"})
	db.FirstOrCreate(&PositionAdmin, &entity.Position{Position: "Admin"})

	// สร้าง Role
	rolePassenger := &entity.Roles{Role: "Passenger"}
	roleDriver := &entity.Roles{Role: "Driver"}
	roleEmployee := &entity.Roles{Role: "Employee"}
	roleAdmin := &entity.Roles{Role: "Admin"}
	db.FirstOrCreate(&rolePassenger, entity.Roles{Role: "Passenger"})
	db.FirstOrCreate(&roleDriver, entity.Roles{Role: "Driver"})
	db.FirstOrCreate(&roleEmployee, entity.Roles{Role: "Employee"})
	db.FirstOrCreate(&roleAdmin, entity.Roles{Role: "Admin"})

	// เข้ารหัสรหัสผ่าน
	hashedPassword, err := HashPassword("123456")
	if err != nil {
		panic("Failed to hash password for Driver")
	}

	// ข้อมูลพนักงานทั้งหมด
	employees := []*entity.Employee{
		{
			Firstname:   "Chanyeol",
			Lastname:    "Park",
			PhoneNumber: "0692345678",
			DateOfBirth: time.Date(1992, time.November, 27, 0, 0, 0, 0, time.UTC),
			StartDate:   time.Date(2013, time.December, 25, 0, 0, 0, 0, time.UTC),
			Salary:      100000.00,
			Email:       "Chanyeol@gmail.com",
			Password:    hashedPassword, // เก็บ Hash
			RolesID:     roleAdmin.ID,
			GenderID:    GenderMale.ID,
			PositionID:  PositionOwner.ID,
		},
		{
			Firstname:   "Seolhyun",
			Lastname:    "Kim",
			PhoneNumber: "0992345678",
			DateOfBirth: time.Date(1995, time.January, 3, 0, 0, 0, 0, time.UTC),
			StartDate:   time.Date(2019, time.December, 25, 0, 0, 0, 0, time.UTC),
			Salary:      40000.00,
			Email:       "Seolhyun@gmail.com",
			Password:    hashedPassword, // เก็บ Hash
			RolesID:     roleEmployee.ID,
			GenderID:    GenderFemale.ID,
			PositionID:  PositionEmployee.ID,
		},
		{
			Firstname:   "Jihoon",
			Lastname:    "Seo",
			PhoneNumber: "0892345678",
			DateOfBirth: time.Date(1997, time.April, 24, 0, 0, 0, 0, time.UTC),
			StartDate:   time.Date(2020, time.December, 25, 0, 0, 0, 0, time.UTC),
			Salary:      40000.00,
			Email:       "Jihoon@gmail.com",
			Password:    hashedPassword, // เก็บ Hash
			RolesID:     roleAdmin.ID,
			GenderID:    GenderMale.ID,
			PositionID:  PositionAdmin.ID,
		},
	}

	// วนลูปบันทึกข้อมูลพนักงานลงฐานข้อมูล
	for _, e := range employees {
		db.FirstOrCreate(e, entity.Employee{Email: e.Email})
	}

	// ข้อมูล Driver ทั้งหมด
	drivers := []*entity.Driver{
		{
			Firstname:                   "Somchai", // คนแรก (ตัวอย่างเดิม)
			Lastname:                    "Prasertsak",
			IdentificationNumber:        "1234567890123",
			DriverLicensenumber:         "48867890",
			DriverLicenseExpirationDate: time.Date(2027, time.December, 1, 0, 0, 0, 0, time.UTC),
			Email:                       "Somchai@gmail.com",
			PhoneNumber:                 "0812345678",
			Password:                    hashedPassword,
			Income:                      25000.50,
			DateOfBirth:                 time.Date(1985, time.December, 1, 0, 0, 0, 0, time.UTC),
			RoleID:                      roleDriver.ID,
			GenderID:                    GenderMale.ID,
			EmployeeID:                  2,
			VehicleID:					 1,
			LocationID: 6,
			DriverStatusID: 1,
		},
		{
			Firstname:                   "Somsak",
			Lastname:                    "Jantakan",
			IdentificationNumber:        "9876543210987",
			DriverLicensenumber:         "51234567",
			DriverLicenseExpirationDate: time.Date(2028, time.January, 15, 0, 0, 0, 0, time.UTC),
			Email:                       "SomsakJ@gmail.com",
			PhoneNumber:                 "0897654321",
			Password:                    hashedPassword,
			Income:                      27000.75,
			DateOfBirth:                 time.Date(1986, time.January, 15, 0, 0, 0, 0, time.UTC),
			RoleID:                      roleDriver.ID,
			GenderID:                    GenderMale.ID,
			EmployeeID:                  2,
			VehicleID:					 2,
			LocationID: 1,
			DriverStatusID: 1,
		},
		{
			Firstname:                   "Prasit",
			Lastname:                    "Thongchai",
			IdentificationNumber:        "4567891234567",
			DriverLicensenumber:         "67890123",
			DriverLicenseExpirationDate: time.Date(2026, time.June, 1, 0, 0, 0, 0, time.UTC),
			Email:                       "PrasitT@gmail.com",
			PhoneNumber:                 "0823456789",
			Password:                    hashedPassword,
			Income:                      28000.25,
			DateOfBirth:                 time.Date(1983, time.June, 15, 0, 0, 0, 0, time.UTC),
			RoleID:                      roleDriver.ID,
			GenderID:                    GenderMale.ID,
			EmployeeID:                  3,
			VehicleID:					 6,
			LocationID: 2,
			DriverStatusID: 1,
		},
		{
			Firstname:                   "Thannam",
			Lastname:                    "Suwan",
			IdentificationNumber:        "1231231231234",
			DriverLicensenumber:         "87654321",
			DriverLicenseExpirationDate: time.Date(2027, time.August, 20, 0, 0, 0, 0, time.UTC),
			Email:                       "Thannam@gmail.com",
			PhoneNumber:                 "0811112233",
			Password:                    hashedPassword,
			Income:                      29000.00,
			DateOfBirth:                 time.Date(1987, time.February, 10, 0, 0, 0, 0, time.UTC),
			RoleID:                      roleDriver.ID,
			GenderID:                    GenderFemale.ID,
			EmployeeID:                  3,
			VehicleID:					 4,
			LocationID: 3,
			DriverStatusID: 2,
		},
		{
			Firstname:                   "Anan",
			Lastname:                    "Phanwichai",
			IdentificationNumber:        "6543219876543",
			DriverLicensenumber:         "33445566",
			DriverLicenseExpirationDate: time.Date(2025, time.September, 10, 0, 0, 0, 0, time.UTC),
			Email:                       "Anan@gmail.com",
			PhoneNumber:                 "0888889999",
			Password:                    hashedPassword,
			Income:                      25000.00,
			DateOfBirth:                 time.Date(1988, time.September, 25, 0, 0, 0, 0, time.UTC),
			RoleID:                      roleDriver.ID,
			GenderID:                    GenderFemale.ID,
			EmployeeID:                  2,
			VehicleID:					 5,
			LocationID: 4,
			DriverStatusID: 2,
		},
		{
			Firstname:                   "Supa",
			Lastname:                    "Rungroj",
			IdentificationNumber:        "1112223334445",
			DriverLicensenumber:         "44332211",
			DriverLicenseExpirationDate: time.Date(2028, time.March, 30, 0, 0, 0, 0, time.UTC),
			Email:                       "Supa@gmail.com",
			PhoneNumber:                 "0801234567",
			Password:                    hashedPassword,
			Income:                      26000.75,
			DateOfBirth:                 time.Date(1984, time.April, 5, 0, 0, 0, 0, time.UTC),
			RoleID:                      roleDriver.ID,
			GenderID:                    GenderFemale.ID,
			EmployeeID:                  2,
			VehicleID:					 3,
			LocationID: 5,
			DriverStatusID: 1,
		},
	}

	// บันทึก Driver ทั้งหมดลงฐานข้อมูล
	for _, d := range drivers {
		db.FirstOrCreate(d, entity.Driver{DriverLicensenumber: d.DriverLicensenumber})
	}

	trainers := []*entity.Trainers{
		{
			FirstName: "Nontakarn",
			LastName: "saisok",
			Email: "nontakarn@gmail.com",
			BirthDay: time.Date(2003, time.December, 11, 17, 0, 0, 0, time.UTC),
			GenderID: 1,
			RoleID: 3,
		},
		{
			FirstName: "Somsak",
			LastName: "meemak",
			Email: "Somsak@gmail.com",
			BirthDay: time.Date(1983, time.September, 14, 17, 0, 0, 0, time.UTC),
			GenderID: 1,
			RoleID: 3,
		},
		{
			FirstName: "Chutima",
			LastName: "Ploysai",
			Email: "chutima@gmail.com",
			BirthDay: time.Date(1992, time.March, 11, 17, 0, 0, 0, time.UTC),
			GenderID: 2,
			RoleID: 3,
		},
		{
			FirstName: "Kanokwan",
			LastName: "Jantima",
			Email: "kanokwan@gmail.com",
			BirthDay: time.Date(1997, time.November, 29, 17, 0, 0, 0, time.UTC),
			GenderID: 2,
			RoleID: 3,
		},
		{
			FirstName: "Jirawat",
			LastName: "Somchai",
			Email: "jirawat@gmail.com",
			BirthDay: time.Date(1985, time.May, 9, 17, 0, 0, 0, time.UTC),
			GenderID: 1,
			RoleID: 3,
		},
		{
			FirstName: "Sutida",
			LastName: "Wongchai",
			Email: "sutida@gmail.com",
			BirthDay: time.Date(1990, time.February, 16, 17, 0, 0, 0, time.UTC),
			GenderID: 2,
			RoleID: 3,
		},
		{
			FirstName: "Natthapong",
			LastName: "Tangsiri",
			Email: "natthapong@gmail.com",
			BirthDay: time.Date(1988, time.August, 22, 17, 0, 0, 0, time.UTC),
			GenderID: 1,
			RoleID: 3,
		},
		{
			FirstName: "Waraporn",
			LastName: "Chanthai",
			Email: "waraporn@gmail.com",
			BirthDay: time.Date(1993, time.September, 6, 17, 0, 0, 0, time.UTC),
			GenderID: 2,
			RoleID: 3,
		},
		{
			FirstName: "Prasit",
			LastName: "Boonchai",
			Email: "prasit@gmail.com",
			BirthDay: time.Date(1983, time.June, 13, 17, 0, 0, 0, time.UTC),
			GenderID: 1,
			RoleID: 3,
		},
		{
			FirstName: "Siriporn",
			LastName: "Wansuk",
			Email: "siriporn@gmail.com",
			BirthDay: time.Date(1989, time.December, 4, 17, 0, 0, 0, time.UTC),
			GenderID: 2,
			RoleID: 3,
		},
	}
	
	// Add trainers to the database
	for _, e := range trainers {
		db.FirstOrCreate(e, entity.Trainers{Email: e.Email})
	}

	// สร้าง VehicleType
	vehicleType1 := entity.VehicleType{VehicleType: "Motorcycle"}
	vehicleType2 := entity.VehicleType{VehicleType: "Car"}

	// ใช้ db.FirstOrCreate เพื่อป้องกันการสร้างข้อมูลซ้ำ
	db.FirstOrCreate(&vehicleType1, entity.VehicleType{VehicleType: "Motorcycle"})
	db.FirstOrCreate(&vehicleType2, entity.VehicleType{VehicleType: "Car"})

	// สร้างข้อมูล NametypeVechicle
	nametypeVehicles := []entity.NametypeVechicle{
		{NameCar: "cabanabike", BaseFare: 20, PerKm: 5, Capacity: 1, VehicleTypeID: vehicleType1.ID},
		{NameCar: "cabanacar", BaseFare: 40, PerKm: 8, Capacity: 4, VehicleTypeID: vehicleType2.ID},
	}
	
	// เพิ่มข้อมูล NametypeVechicle โดยใช้ FirstOrCreate เพื่อตรวจสอบว่ามีข้อมูลอยู่แล้วหรือไม่
	for _, nametypeVehicles := range nametypeVehicles {
		db.FirstOrCreate(&nametypeVehicles, entity.NametypeVechicle{NameCar: nametypeVehicles.NameCar})
	}

	vehicles := []entity.Vehicle{
		// มอเตอร์ไซค์ 3 คัน
		{
			LicensePlate:               "กกก123",
			Brand:                      "Yamaha",
			VehicleModel:               "NMAX",
			Color:                      "Blue",
			DateOfPurchase:             time.Date(2021, time.January, 10, 0, 0, 0, 0, time.UTC),
			ExpirationDateOfVehicleAct: time.Date(2026, time.January, 10, 0, 0, 0, 0, time.UTC),
			Capacity:                   2,
			VehicleTypeID:              1, // มอเตอร์ไซค์
			EmployeeID:                 3,
		},
		{
			LicensePlate:               "ขข5678",
			Brand:                      "Honda",
			VehicleModel:               "PCX",
			Color:                      "Red",
			DateOfPurchase:             time.Date(2020, time.June, 15, 0, 0, 0, 0, time.UTC),
			ExpirationDateOfVehicleAct: time.Date(2025, time.June, 15, 0, 0, 0, 0, time.UTC),
			Capacity:                   2,
			VehicleTypeID:              1, // มอเตอร์ไซค์
			EmployeeID:                 3,
		},
		{
			LicensePlate:               "คค9012",
			Brand:                      "Kawasaki",
			VehicleModel:               "Z125",
			Color:                      "Green",
			DateOfPurchase:             time.Date(2022, time.February, 20, 0, 0, 0, 0, time.UTC),
			ExpirationDateOfVehicleAct: time.Date(2027, time.February, 20, 0, 0, 0, 0, time.UTC),
			Capacity:                   2,
			VehicleTypeID:              1, // มอเตอร์ไซค์
			EmployeeID:                 3,
		},

		// รถยนต์ 3 คัน
		{
			LicensePlate:               "งง3456",
			Brand:                      "Toyota",
			VehicleModel:               "Camry",
			Color:                      "White",
			DateOfPurchase:             time.Date(2019, time.March, 5, 0, 0, 0, 0, time.UTC),
			ExpirationDateOfVehicleAct: time.Date(2024, time.March, 5, 0, 0, 0, 0, time.UTC),
			Capacity:                   5,
			VehicleTypeID:              2, // รถยนต์
			EmployeeID:                 3,
		},
		{
			LicensePlate:               "จจ7890",
			Brand:                      "Honda",
			VehicleModel:               "Civic",
			Color:                      "Black",
			DateOfPurchase:             time.Date(2020, time.July, 20, 0, 0, 0, 0, time.UTC),
			ExpirationDateOfVehicleAct: time.Date(2025, time.July, 20, 0, 0, 0, 0, time.UTC),
			Capacity:                   5,
			VehicleTypeID:              2, // รถยนต์
			EmployeeID:                 3,
		},
		{
			LicensePlate:               "ฉฉ1234",
			Brand:                      "Mazda",
			VehicleModel:               "CX-5",
			Color:                      "Gray",
			DateOfPurchase:             time.Date(2021, time.September, 15, 0, 0, 0, 0, time.UTC),
			ExpirationDateOfVehicleAct: time.Date(2026, time.September, 15, 0, 0, 0, 0, time.UTC),
			Capacity:                   5,
			VehicleTypeID:              2, // รถยนต์
			EmployeeID:                 3,
		},
	}

	// บันทึกข้อมูล Vehicle ลงฐานข้อมูล
	for _, v := range vehicles {
		db.FirstOrCreate(&v, entity.Vehicle{LicensePlate: v.LicensePlate})
	}

	passengers := []entity.Passenger{
		{
			UserName:    "Anuwat",
			FirstName:   "Anuwat",
			LastName:    "Thongchai",
			PhoneNumber: "0811111111",
			Email:       "anuwat1@gmail.com",
			Password:    hashedPassword, // เก็บ Hash
			GenderID:    1,
			RoleID:      1,
		},
		{
			UserName:    "Chatchai",
			FirstName:   "Chatchai",
			LastName:    "Prasert",
			PhoneNumber: "0812222222",
			Email:       "chatchai2@gmail.com",
			Password:    hashedPassword,
			GenderID:    1,
			RoleID:      1,
		},
		{
			UserName:    "Kittipong",
			FirstName:   "Kittipong",
			LastName:    "Suwan",
			PhoneNumber: "0813333333",
			Email:       "kittipong3@gmail.com",
			Password:    hashedPassword,
			GenderID:    1,
			RoleID:      1,
		},
		{
			UserName:    "Nattapon",
			FirstName:   "Nattapon",
			LastName:    "Somchai",
			PhoneNumber: "0814444444",
			Email:       "nattapon4@gmail.com",
			Password:    hashedPassword,
			GenderID:    1,
			RoleID:      1,
		},
		{
			UserName:    "Siriporn",
			FirstName:   "Siriporn",
			LastName:    "Jantakan",
			PhoneNumber: "0815555555",
			Email:       "siriporn1@gmail.com",
			Password:    hashedPassword,
			GenderID:    2,
			RoleID:      1,
		},
		{
			UserName:    "Nanthicha",
			FirstName:   "Nanthicha",
			LastName:    "Phanwichai",
			PhoneNumber: "0816666666",
			Email:       "nanthicha2@gmail.com",
			Password:    hashedPassword,
			GenderID:    2,
			RoleID:      1,
		},
		{
			UserName:    "Chanidapa",
			FirstName:   "Chanidapa",
			LastName:    "Rungroj",
			PhoneNumber: "0817777777",
			Email:       "chanidapa3@gmail.com",
			Password:    hashedPassword,
			GenderID:    2,
			RoleID:      1,
		},
		{
			UserName:    "Supattra",
			FirstName:   "Supattra",
			LastName:    "Kraiwit",
			PhoneNumber: "0818888888",
			Email:       "supattra4@gmail.com",
			Password:    hashedPassword,
			GenderID:    GenderFemale.ID,
			RoleID:      1,
		},
	}

	// บันทึกข้อมูล Passenger ลงฐานข้อมูล
	for _, p := range passengers {
		db.FirstOrCreate(&p, entity.Passenger{Email: p.Email})
	}

	rooms := []entity.Rooms{
		{
			RoomName:        "A15",
			Capacity:        20,
			CurrentBookings: 0,
			TrainerID:       1,
			Detail:          "The application is designed to streamline the customer transportation experience, focusing on efficiency, safety, and convenience. As a driver, your role is crucial in ensuring a high-quality experience for all users. This section introduces the app’s main features, including navigation, booking management, earnings tracking, and communication tools. Drivers will learn how to set up their profiles, accept ride requests, and manage daily operations seamlessly.",
			Title:           "Introduction to the Driver Application",
		},
		{
			RoomName:        "D09",
			Capacity:        15,
			CurrentBookings: 0,
			TrainerID:       2,
			Detail:          "Building trust and maintaining professionalism with customers are key to success. This section emphasizes effective communication, maintaining a polite demeanor, and understanding diverse customer needs. Drivers will be guided on how to handle common scenarios, such as assisting passengers with luggage, addressing complaints constructively, and ensuring a comfortable journey for all.",
			Title:           "Professional Customer Interaction",
		},
		{
			RoomName:        "D00",
			Capacity:        25,
			CurrentBookings: 0,
			TrainerID:       3,
			Detail:          "A well-maintained vehicle is essential for providing safe and reliable services. This topic covers regular vehicle checkups, cleanliness standards, and handling minor issues before they escalate. Drivers will also learn about emergency protocols, such as responding to accidents or unexpected mechanical problems, ensuring both their safety and that of the passengers.",
			Title:           "Vehicle Maintenance and Safety",
		},
		{
			RoomName:        "B10",
			Capacity:        30,
			CurrentBookings: 0,
			TrainerID:       4,
			Detail:          "Efficient navigation is a key component of providing timely and reliable service. This section explains how to utilize the app’s GPS features to find optimal routes, avoid traffic, and adjust to real-time conditions. Drivers will also be trained on how to deal with unexpected detours and provide updates to passengers in case of delays.",
			Title:           "Using the Navigation System Effectively",
		},
		{
			RoomName:        "C14",
			Capacity:        10,
			CurrentBookings: 0,
			TrainerID:       5,
			Detail:          "Understanding the financial aspects of the app ensures transparency and satisfaction for both drivers and customers. This section teaches drivers how to process payments via the app, handle cash transactions when necessary, and access their earnings summary. Drivers will also be advised on managing disputes related to payments and maintaining accurate records.",
			Title:           "Handling Payments and Earnings Tracking",
		},
		{
			RoomName:        "D07",
			Capacity:        20,
			CurrentBookings: 0,
			TrainerID:       6,
			Detail:          "Drivers may encounter various unexpected scenarios during their service. This section prepares them to handle emergencies, such as accidents, medical situations, or confrontations, with a calm and professional approach. Drivers will also learn how to assist customers with special needs or provide extra support for vulnerable passengers.",
			Title:           "Emergency and Special Situations",
		},
		{
			RoomName:        "D11",
			Capacity:        15,
			CurrentBookings: 0,
			TrainerID:       7,
			Detail:          "Feedback is an essential part of maintaining quality service. This section discusses the importance of customer reviews and ratings, how to address constructive criticism, and ways to consistently improve performance. Drivers will also learn how to seek support from the app’s team if they face challenges or need assistance.",
			Title:           "Customer Feedback and Continuous Improvement",
		},
		{
			RoomName:        "B00",
			Capacity:        25,
			CurrentBookings: 0,
			TrainerID:       8,
			Detail:          "Operating as a driver involves adhering to laws and ethical standards. This topic outlines drivers’ responsibilities, including respecting traffic laws, ensuring customer privacy, and avoiding discrimination. Drivers will also be informed about the consequences of violating these standards and the importance of fostering a safe and inclusive environment.",
			Title:           "Legal and Ethical Responsibilities",
		},
	}

	// เพิ่มข้อมูล Room ลงในฐานข้อมูล
	for _, room := range rooms {
		db.FirstOrCreate(&room, entity.Rooms{RoomName: room.RoomName})
	}

	fmt.Println("Database setup and seeding completed")
}
