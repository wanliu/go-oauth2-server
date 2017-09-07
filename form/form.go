package form

import (
	"fmt"
	"log"
	"net/url"
	"regexp"

	goform "github.com/go-playground/form"
	"github.com/go-playground/locales/zh_Hans"
	ut "github.com/go-playground/universal-translator"
	zh_Hans_translations "github.com/wanliu/zh_Hans"
	"gopkg.in/go-playground/validator.v9"
)

type Form struct {
	Errors     validator.ValidationErrors
	custErrors [][2]string
}

var (
	decoder  *goform.Decoder
	validate *validator.Validate
	trans    ut.Translator
	uni      *ut.UniversalTranslator
)
var validPhone = regexp.MustCompile(`^1[3|4|5|7|8][0-9]{9}$`)

func cnmobileValidFunc(fl validator.FieldLevel) bool {

	if validPhone.MatchString(fl.Field().String()) {
		return true
	}

	return false
}

func cnmobileRegisFunc(ut ut.Translator) (err error) {
	if err = ut.Add("cnmobile", "{0} 必须是个手机号码", false); err != nil {
		return
	}

	return

}

func ParseForm(v interface{}, values url.Values) error {
	return decoder.Decode(v, values)
}

func (f *Form) ValidForm(form interface{}) bool {
	var ok bool

	// f.SetTranslator("zh_Hans")
	err := validate.Struct(form)
	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return false
		}

		if f.Errors, ok = err.(validator.ValidationErrors); ok {
			if len(f.Errors) > 0 {
				return false
			}
		}
	}

	return true
}

func (f *Form) ErrorStatus(name string) string {
	err := f.Error(name)
	if err == nil {
		return ""
	}

	return "has-error"
}

func (f *Form) Error(name string) validator.FieldError {
	for _, err := range f.Errors {
		if err.Field() == name {
			return err
		}
	}
	return nil
}

func (f *Form) ErrorText(name string) string {
	err := f.Error(name)
	if err == nil {
		if msg := f.getCustError(name); len(msg) > 0 {
			return msg
		} else {
			return ""
		}
	}

	return err.Translate(trans)
}

func (f *Form) ErrorMsg(name string) string {
	var msg string
	err := f.Error(name)
	if err == nil {
		if msg = f.getCustError(name); len(msg) == 0 {
			return ""
		}
	} else {
		msg = f.ErrorText(name)
	}

	return fmt.Sprintf(`<span id="help-%s" class="help-block">%s</span>`, name, msg)
}

func (f *Form) AddError(name, msg string) {
	f.custErrors = append(f.custErrors, [2]string{name, msg})
}

func (f *Form) getCustError(name string) string {
	for _, errs := range f.custErrors {
		if errs[0] == name {
			return errs[1]
		}
	}

	return ""
}

func translateFunc(ut ut.Translator, fe validator.FieldError) string {

	t, err := ut.T(fe.Tag(), fe.Field())
	if err != nil {
		log.Printf("warning: error translating FieldError: %#v", fe)
		return fe.(error).Error()
	}

	return t
}

func init() {
	decoder = goform.NewDecoder()
	validate = validator.New()
	validate.RegisterValidation("cnmobile", cnmobileValidFunc)
	// en := en.New()
	zh := zh_Hans.New()
	uni = ut.New(zh, zh)
	trans, _ = uni.GetTranslator("zh_Hans")
	validate.RegisterTranslation("cnmobile", trans, cnmobileRegisFunc, translateFunc)
	zh_Hans_translations.RegisterDefaultTranslations(validate, trans)
}
