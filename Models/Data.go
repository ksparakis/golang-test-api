package Models

import (
	"genity/Config"
	"time"
	"github.com/satori/go.uuid"
)

type Data struct {
	Title string `json:"Title";gorm:"unique"`
	UUID4 uuid.UUID `json:"UUID4";gorm:"primaryKey;unique";sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Timestamp time.Time `json:"Timestamp";sql:"DEFAULT:current_timestamp"`
}

func (d *Data) TableName() string{
	return "Data"
}

//GetUserByID ... Fetch only one user by Id
func GetDataByTitle(data *Data, title string) (err error) {
 	if err = Config.DB.Where("title = ?", title).First(data).Error;

	err != nil {
		return err
	}

	return nil
}

//CreateUser ... Insert New data
func CreateData(data *Data) (err error) {
	if err = Config.DB.Create(data).Error; err != nil {
		return err
	}
	return nil
}

