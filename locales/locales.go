package locales

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
)

var supportedLanguages = []string{"en", "pt"}
var bundle *i18n.Bundle

func main() {
	name := map[string]string{"Name": "Tester"}
	msg_en, _ := Translate("test", "en", name)
	msg_pt, _ := Translate("test", "pt", name)

	println(msg_en)
	println(msg_pt)
}

func init() {
	bundle = i18n.NewBundle(language.Portuguese)
	bundle.RegisterUnmarshalFunc("yml", yaml.Unmarshal)

	for _, lang := range(supportedLanguages) {
		bundle.MustLoadMessageFile(lang + ".yml")
	}
}

// Receives the desired message, language, and contextual data (such as names and quantities)
func Translate(id string, language string, metadata map[string]string) (string, error) {
	loc := i18n.NewLocalizer(bundle, language)
	translation, err := loc.Localize(&i18n.LocalizeConfig{
		MessageID: id,
		TemplateData: metadata,
	})

	return translation, err
}