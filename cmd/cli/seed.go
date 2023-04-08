package main

import (
	"github.com/shaband/photomaker-go/pkgs/infrastucture/database"
	"github.com/shaband/photomaker-go/pkgs/modules/categories"
	"github.com/shaband/photomaker-go/pkgs/modules/clients"
	"github.com/shaband/photomaker-go/pkgs/modules/contacts"
	"github.com/shaband/photomaker-go/pkgs/modules/services"
	"github.com/shaband/photomaker-go/pkgs/modules/settings"
	"github.com/shaband/photomaker-go/pkgs/modules/sliders"
	"gorm.io/gorm"
)

func main() {

	database.Init()
	db := database.GetConnection()
	database.MakeMigration(db)
	settingsSeeder(db)
	for i := 0; i < 10; i++ {
		// StoreEntityFakeDataForEntity(db, &categories.Category{})
		StoreDataForEntities(db)
	}
}

func settingsSeeder(db *gorm.DB) {

	settingsData := []*settings.Setting{
		{
			Slug:  "about",
			Label: "من نحن",
			Value: `رواد في الابداع الفني الرقمي لصنع الصورة الاحترافية.

			نعمل في صانع الصورة بفرق متخصصة ونسخر كافة الوسائل التقنية الحديثة لتعزيز مكانة عملائنا.
			
			الإبداع مزيج بين المنطق والخيال ... هكذا نحن`,
			Page: "about",
			Type: "text",
		},

		{
			Slug:  "profile",
			Label: "بروفايل الشركة",
			Value: "http://training.aljazeera.net/mritems/Documents/2016/2/16/e782075b14c84729a88e703e0776f66a_100.pdf",
			Page:  "about",
			Type:  "file",
		},

		{
			Slug:  "facebook",
			Label: "فيسبوك",
			Value: "http://training.aljazeera.net/mritems/Documents/2016/2/16/e782075b14c84729a88e703e0776f66a_100.pdf",
			Page:  "about",
			Type:  "text",
		},

		{
			Slug:  "twitter",
			Label: "twitter",
			Value: "http://training.aljazeera.net/mritems/Documents/2016/2/16/e782075b14c84729a88e703e0776f66a_100.pdf",
			Page:  "footer",
			Type:  "text",
		},
		{
			Slug:  "snapchat",
			Label: "snapchat",
			Value: "http://training.aljazeera.net/mritems/Documents/2016/2/16/e782075b14c84729a88e703e0776f66a_100.pdf",
			Page:  "footer",
			Type:  "text",
		},

		{
			Slug:  "instagram",
			Label: "instagram",
			Value: "http://training.aljazeera.net/mritems/Documents/2016/2/16/e782075b14c84729a88e703e0776f66a_100.pdf",
			Page:  "footer",
			Type:  "text",
		},

		{
			Slug:  "youtube",
			Label: "youtube",
			Value: "http://training.aljazeera.net/mritems/Documents/2016/2/16/e782075b14c84729a88e703e0776f66a_100.pdf",
			Page:  "footer",
			Type:  "text",
		},

		{
			Slug:  "email",
			Label: "email",
			Value: "photo@phont.com",
			Page:  "contact",
			Type:  "text",
		},

		{
			Slug:  "phone",
			Label: "phone",
			Value: "01214141414",
			Page:  "contact",
			Type:  "text",
		},

		{
			Slug:  "google",
			Label: "google",
			Value: "https://google.com",
			Page:  "contact",
			Type:  "text",
		},

		{
			Slug:  "pinterest",
			Label: "pinterest",
			Value: "https://pinterest.com",
			Page:  "contact",
			Type:  "text",
		},
		{
			Slug:  "behance",
			Label: "behance",
			Value: "https://behance.com",
			Page:  "contact",
			Type:  "text",
		},
		{
			Slug:  "vimeo",
			Label: "vimeo",
			Value: "https://vimeo.com",
			Page:  "contact",
			Type:  "text",
		},
	}

	for _, value := range settingsData {
		db.Where(settings.Setting{Slug: value.Slug}).FirstOrCreate(value)
	}
}

type Fakable interface {
	Fake() interface{}
}

func CreateFakeEntity(entity Fakable) interface{} {

	return entity.Fake()

}

// func createClient(db *gorm.DB) {

// 	client := clients.Client{}

// 	CreateFakeEntity(&client)

// 	db.Create(&client)

// }

func StoreEntityFakeDataForEntity(db *gorm.DB, entity Fakable) {

	CreateFakeEntity(entity)

	db.Create(entity)

}
func StoreDataForEntities(db *gorm.DB) {

	var entites []Fakable = []Fakable{
		&clients.Client{},
		&categories.CategoryImage{},
		&services.Service{},
		&sliders.Slider{},
		&contacts.ServiceTypeItem{},
	}

	for _, entity := range entites {
		CreateFakeEntity(entity)

		db.Create(entity)
	}

}
