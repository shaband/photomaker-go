package main

import (
	"github.com/shaband/photomaker-go/pkgs/infrastucture/database"
	"github.com/shaband/photomaker-go/pkgs/modules/clients"
	"github.com/shaband/photomaker-go/pkgs/modules/settings"
	"gorm.io/gorm"
)

func main() {

	database.Init()
	db := database.GetConnection()

	// settingsSeeder(db)
	createClient(db)
}

func settingsSeeder(db *gorm.DB) {

	settings := []*settings.Setting{
		&settings.Setting{
			Label: "من نحن",
			Value: `رواد في الابداع الفني الرقمي لصنع الصورة الاحترافية.

			نعمل في صانع الصورة بفرق متخصصة ونسخر كافة الوسائل التقنية الحديثة لتعزيز مكانة عملائنا.
			
			الإبداع مزيج بين المنطق والخيال ... هكذا نحن`,
			Page: "about",
			Type: "text",
		},

		&settings.Setting{
			Label: "بروفايل الشركة",
			Value: "http://training.aljazeera.net/mritems/Documents/2016/2/16/e782075b14c84729a88e703e0776f66a_100.pdf",
			Page:  "about",
			Type:  "file",
		},

		&settings.Setting{
			Label: "فيسبوك",
			Value: "http://training.aljazeera.net/mritems/Documents/2016/2/16/e782075b14c84729a88e703e0776f66a_100.pdf",
			Page:  "about",
			Type:  "text",
		},

		&settings.Setting{
			Label: "twitter",
			Value: "http://training.aljazeera.net/mritems/Documents/2016/2/16/e782075b14c84729a88e703e0776f66a_100.pdf",
			Page:  "footer",
			Type:  "text",
		},
		&settings.Setting{
			Label: "snapchat",
			Value: "http://training.aljazeera.net/mritems/Documents/2016/2/16/e782075b14c84729a88e703e0776f66a_100.pdf",
			Page:  "footer",
			Type:  "text",
		},

		&settings.Setting{
			Label: "instagram",
			Value: "http://training.aljazeera.net/mritems/Documents/2016/2/16/e782075b14c84729a88e703e0776f66a_100.pdf",
			Page:  "footer",
			Type:  "text",
		},

		&settings.Setting{
			Label: "youtube",
			Value: "http://training.aljazeera.net/mritems/Documents/2016/2/16/e782075b14c84729a88e703e0776f66a_100.pdf",
			Page:  "footer",
			Type:  "text",
		},

		&settings.Setting{
			Label: "email",
			Value: "photo@phont.com",
			Page:  "contact",
			Type:  "text",
		},

		&settings.Setting{
			Label: "phone",
			Value: "01214141414",
			Page:  "contact",
			Type:  "text",
		},
	}

	for _, value := range settings {
		db.Create(value)
	}
}

type Fakable interface {
	Fake() interface{}
}

func CreateFakeEntity(entity Fakable) interface{} {

	return entity.Fake()

}

func createClient(db *gorm.DB) clients.Client {

	client := clients.Client{}

	CreateFakeEntity(&client)
	return client
}
