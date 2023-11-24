package localize

import (
	"testing"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func TestArabicJson(t *testing.T) {
	bundle := i18n.NewBundle(language.Arabic)
	// bundle.RegisterUnmarshalFunc("json", i18n.UnmarshalJSONFile)
	bundle.MustLoadMessageFile("ar.json")

	localizer := i18n.NewLocalizer(bundle, "ar")

	keys := []string{"home", "about", "services", "gallery", "contact_us", "copyrights", "our Clients", "our Gallery", "service type", "name", "subject", "phone", "email", "attach file", "our work", "contact details", "subscribe", "download profile", "send"}

	for _, key := range keys {
		_, err := localizer.Localize(&i18n.LocalizeConfig{MessageID: key})
		if err != nil {
			t.Errorf("Failed to localize key '%s': %v", key, err)
		}
	}
}

func TestEnglishJson(t *testing.T) {
	bundle := i18n.NewBundle(language.English)
	// bundle.RegisterUnmarshalFunc("json", i18n.UnmarshalJSONFile)
	bundle.MustLoadMessageFile("en.json")

	localizer := i18n.NewLocalizer(bundle, "en")

	keys := []string{"home", "about", "services", "gallery", "contact_us", "copyrights", "our clients", "our Gallery", "service type", "name", "subject", "phone", "email", "attach file", "our work", "contact details", "subscribe", "download profile", "send"}

	for _, key := range keys {
		_, err := localizer.Localize(&i18n.LocalizeConfig{MessageID: key})
		if err != nil {
			t.Errorf("Failed to localize key '%s': %v", key, err)
		}
	}
}
