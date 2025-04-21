package plusgorm

import (
	"bytes"
	"context"
	"os"
	"path"
	"strings"
	"unicode"

	"github.com/spf13/cast"
	"gorm.io/gorm"
)

// StrToCamel underline to camel
func StrToCamel(str string) string {
	var newStr string
	strSplit := strings.Split(str, "_")
	for _, p := range strSplit {
		switch len(p) {
		case 0:
		case 1:
			newStr += strings.ToUpper(p[0:1])
		default:
			newStr += strings.ToUpper(p[0:1]) + p[1:]
		}
	}
	return newStr
}

// StrToUnderline camel to underline eg:AaBb ==> aa_bb
func StrToUnderline(str string) string {
	var buf bytes.Buffer
	for i, r := range str {
		if unicode.IsUpper(r) {
			if i > 0 && unicode.IsLower(rune(str[i-1])) {
				buf.WriteRune('_')
			}
			buf.WriteRune(unicode.ToLower(r))
		} else {
			buf.WriteRune(r)
		}
	}
	return buf.String()
}

// CloneDB deep copy db (db.Statement)
func CloneDB(db *gorm.DB) *gorm.DB {
	return db.WithContext(context.Background())
}

// NewDB new db
func NewDB(db *gorm.DB) *gorm.DB {
	return db.Session(&gorm.Session{NewDB: true, Context: context.Background()})
	// return db.Session(&gorm.Session{NewDB: true, Context: db.Statement.Context})
}

// CleanDB clean db
func CleanDB(db *gorm.DB) *gorm.DB {
	// return db.Session(&gorm.Session{NewDB: true, Context: context.Background()})
	return db.Session(&gorm.Session{NewDB: true, Context: db.Statement.Context})
}
func writeToFile(file string, data interface{}, append bool) error {
	var err error
	err = mkdir(file, true)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0777)
	if append {
		f, err = os.OpenFile(file, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	}
	if err != nil {
		return err
	}

	_, err = f.WriteString(cast.ToString(data))
	_ = f.Close()
	return err

}
func mkdir(file string, cascade bool) error {
	exist, err := isExistFile(file)
	if err != nil || !exist {
		f := path.Dir(file)
		if cascade {
			err = os.MkdirAll(f, 0755)
			if err != nil {
				return err
			}
			return nil
		}
		err = os.Mkdir(f, 0755)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}
func isExistFile(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
