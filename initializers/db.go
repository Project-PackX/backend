package initializers

import (
	"fmt"
	"os"
	"time"

	"github.com/Project-PackX/backend/enums"
	"github.com/Project-PackX/backend/models"
	"github.com/Project-PackX/backend/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

// Defining the application databse structure with gorm
var DB *gorm.DB

// Connectiing to the database based on the environment variables
func ConnectToDatabase() {
	logger := utils.Logger

	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"), 5432)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Silent),
	})

	// Error handling
	if err != nil {
		logger.Error("Couldn't connect to the database")
	}
}

// FOR TESTING PURPOSES
func DropTables() {
	DB.Exec("DROP TABLE IF EXISTS public.users;")
	DB.Exec("DROP TABLE IF EXISTS public.packages;")
	DB.Exec("DROP TABLE IF EXISTS public.statuses;")
	DB.Exec("DROP TABLE IF EXISTS public.packagestatuses;")
	DB.Exec("DROP TABLE IF EXISTS public.couriers;")
	DB.Exec("DROP TABLE IF EXISTS public.lockers;")
	DB.Exec("DROP TABLE IF EXISTS public.lockergroups;")
	DB.Exec("DROP TABLE IF EXISTS public.packageslockers")
	DB.Exec("DROP TABLE IF EXISTS public.reset_password_code")
}

// Migrating the DB tables into Go models
func SyncDB() {
	DB.AutoMigrate(&models.Package{})
	DB.AutoMigrate(&models.Courier{})
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Status{})
	DB.AutoMigrate(&models.PackageStatus{})
	DB.AutoMigrate(&models.Locker{})
	DB.AutoMigrate(&models.LockerGroup{})
	DB.AutoMigrate(&models.PackageLocker{})
	DB.AutoMigrate(&models.ResetPasswordCode{})
}

// Generate test datas
func GenerateTestEntries() {

	// Users

	felh1 := models.User{
		Name:        "Kovács Bea",
		Address:     "Liliom utca 4.",
		Phone:       "+36201956673",
		Email:       "k.bea@mail.com",
		AccessLevel: enums.AccessLevel.Normal,
	}
	DB.Create(&felh1)

	felh2 := models.User{
		Name:        "Szalma Géza",
		Address:     "Egressy körút 58.",
		Phone:       "+36605385438",
		Email:       "szalmag@mail.com",
		AccessLevel: enums.AccessLevel.Normal,
	}
	DB.Create(&felh2)

	felh3 := models.User{
		Name:        "Veres Péter",
		Address:     "Malom út 12.",
		Phone:       "+36504098931",
		Email:       "vrsptr@mail.com",
		AccessLevel: enums.AccessLevel.Admin,
	}
	DB.Create(&felh3)

	// Packages

	csomag1 := models.Package{
		UserID:       2,
		Size:         enums.Sizes.Medium,
		Price:        37990,
		Note:         "Utánvét",
		CourierID:    1,
		DeliveryDate: time.Now().Add(time.Hour * 24 * 5),
	}
	DB.Create(&csomag1)

	csomag2 := models.Package{
		UserID:       1,
		Size:         enums.Sizes.Small,
		Price:        225000,
		Note:         "Javítás",
		CourierID:    1,
		DeliveryDate: time.Now().Add(time.Hour * 24 * 5),
	}
	DB.Create(&csomag2)

	csomag3 := models.Package{
		UserID:       2,
		Size:         enums.Sizes.Medium,
		Price:        17490,
		Note:         "-",
		CourierID:    2,
		DeliveryDate: time.Now().Add(time.Hour * 24 * 5),
	}
	DB.Create(&csomag3)

	csomag4 := models.Package{
		UserID:       3,
		Size:         enums.Sizes.Small,
		Price:        3989,
		Note:         "Cserekészülék",
		CourierID:    2,
		DeliveryDate: time.Now().Add(time.Hour * 24 * 5),
	}
	DB.Create(&csomag4)

	csomag5 := models.Package{
		UserID:       1,
		Size:         enums.Sizes.Large,
		Price:        55990,
		Note:         "-",
		CourierID:    1,
		DeliveryDate: time.Now().Add(time.Hour * 24 * 5),
	}
	DB.Create(&csomag5)

	csomag6 := models.Package{
		UserID:       1,
		Size:         enums.Sizes.Small,
		Price:        3490,
		Note:         "-",
		CourierID:    2,
		DeliveryDate: time.Now().Add(time.Hour * 24 * 5),
	}
	DB.Create(&csomag6)

	// Possible package statuses

	statusz1 := models.Status{
		Id:   1,
		Name: enums.Statuses.Dispatch,
	}
	DB.Create(&statusz1)

	statusz2 := models.Status{
		Id:   2,
		Name: enums.Statuses.Transit,
	}
	DB.Create(&statusz2)

	statusz3 := models.Status{
		Id:   3,
		Name: enums.Statuses.Warehouse,
	}
	DB.Create(&statusz3)

	statusz4 := models.Status{
		Id:   4,
		Name: enums.Statuses.Delivery,
	}
	DB.Create(&statusz4)

	statusz5 := models.Status{
		Id:   5,
		Name: enums.Statuses.Delivered,
	}
	DB.Create(&statusz5)

	statusz6 := models.Status{
		Id:   6,
		Name: enums.Statuses.Canceled,
	}
	DB.Create(&statusz6)

	// Package statuses

	csomagstatusz1 := models.PackageStatus{
		Package_id: 1,
		Status_id:  1,
	}
	DB.Create(&csomagstatusz1)

	csomagstatusz2 := models.PackageStatus{
		Package_id: 2,
		Status_id:  4,
	}
	DB.Create(&csomagstatusz2)

	csomagstatusz3 := models.PackageStatus{
		Package_id: 3,
		Status_id:  3,
	}
	DB.Create(&csomagstatusz3)

	csomagstatusz4 := models.PackageStatus{
		Package_id: 4,
		Status_id:  5,
	}
	DB.Create(&csomagstatusz4)

	csomagstatusz5 := models.PackageStatus{
		Package_id: 5,
		Status_id:  2,
	}
	DB.Create(&csomagstatusz5)

	// Couriers

	futar1 := models.Courier{
		Name:  "Kiss Bendegúz",
		Phone: "+36403437791",
	}
	DB.Create(&futar1)

	futar2 := models.Courier{
		Name:  "Némedi Emma",
		Phone: "+36301984673",
	}
	DB.Create(&futar2)

	// Lockers

	locker1 := models.Locker{
		City:      "Győr",
		Address:   "Szent István út 23.",
		Capacity:  7,
		Latitude:  47.683337112393474,
		Longitude: 17.623955422088322,
	}
	DB.Create(&locker1)

	locker2 := models.Locker{
		City:      "Győr",
		Address:   "Kiss Ernő utca 5.",
		Capacity:  5,
		Latitude:  47.68834622211481,
		Longitude: 17.623955422088322,
	}
	DB.Create(&locker2)

	locker3 := models.Locker{
		City:      "Győr",
		Address:   "Lomnic utca 30.",
		Capacity:  5,
		Latitude:  47.67035183513254,
		Longitude: 17.63988174907635,
	}
	DB.Create(&locker3)

	locker4 := models.Locker{
		City:      "Szombathely",
		Address:   "Paragvári utca 74.",
		Capacity:  5,
		Latitude:  47.24350217822487,
		Longitude: 17.623955422088322,
	}
	DB.Create(&locker4)

	locker5 := models.Locker{
		City:      "Szombathely",
		Address:   "Gömör utca 3.",
		Capacity:  5,
		Latitude:  47.22994517025941,
		Longitude: 16.60908613742963,
	}
	DB.Create(&locker5)

	locker6 := models.Locker{
		City:      "Szombathely",
		Address:   "Éhen Gyula tér 3.",
		Capacity:  10,
		Latitude:  47.23677632684086,
		Longitude: 16.631628691405695,
	}
	DB.Create(&locker6)

	locker7 := models.Locker{
		City:      "Szombathely",
		Address:   "Sziget utca 7.",
		Capacity:  15,
		Latitude:  47.23858035784418,
		Longitude: 16.64677093558233,
	}
	DB.Create(&locker7)

	// Lockergroups

	lgroup1 := models.LockerGroup{
		ID:   1,
		City: "Győr",
	}
	DB.Create(&lgroup1)

	lgroup2 := models.LockerGroup{
		ID:   2,
		City: "Szombathely",
	}
	DB.Create(&lgroup2)

	// PackagesLockers

	pl1 := models.PackageLocker{
		Package_id: 1,
		Locker_id:  2,
	}
	DB.Create(&pl1)

	pl2 := models.PackageLocker{
		Package_id: 2,
		Locker_id:  1,
	}
	DB.Create(&pl2)

	pl3 := models.PackageLocker{
		Package_id: 3,
		Locker_id:  6,
	}
	DB.Create(&pl3)

	pl4 := models.PackageLocker{
		Package_id: 4,
		Locker_id:  1,
	}
	DB.Create(&pl4)

	pl5 := models.PackageLocker{
		Package_id: 5,
		Locker_id:  6,
	}
	DB.Create(&pl5)
}
