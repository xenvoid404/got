package got

import (
	"fmt"
	"reflect"
)

// oke, jadi gini. untuk awal ini keknya bikin gimana caranya
// supaya si got bisa mendeteksi tag 'validate' daris struct yang dikirim.
//
// misalnya ada struct seperti ini:
// type User struct {
//   Name string `validate:"required"`
//   Email string `validate:"required"`
// }

// harusnya sih nanti butuh struct utama seperti ini,
// sekarang belum tau buat nyimpen apaan.
type Validator struct{}

// nah ini buat inisiasi instance got, kek library pada umumnya lah,
// pakenya kan berarti gini:
//
//	got.New()
func New() *Validator {
	return &Validator{}
}

func (v *Validator) Struct(s any) error {
	// nah disini pertama kali ambil tipe
	// dari struct yang dikirim (maksudnya s),
	t := reflect.TypeOf(s)

	// terus pastiin yang dikirim bener bener struct
	// kalo bukan struct tolak dengan error
	if t.Kind() != reflect.Struct {
		return fmt.Errorf("got: mengharapkan struct, ini malah %s", t.Kind())
	}

	// disini ngapain?
	// oiya, loop semua field dulu
	// terus ambil tag validate
	// misalnya disini pake rule required
	// kek: `validate:"required"`
	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		tag := sf.Tag.Get("validate")
		if tag == "required" {
			fmt.Println("field butuh required:", sf.Name)
		}
	}

	return nil
}
