package Models

import (
	"genity/Config"
	"github.com/gin-gonic/gin"
	"time"
	"github.com/satori/go.uuid"
)

type Data struct {
	Title string `json:"Title";sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UUID4 uuid.UUID `json:"UUID4";gorm:"primary_key"`
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

// BeforeCreate will set a UUID rather than numeric ID.
func (data *Data) BeforeCreate(scope *gin.Context)  {
	uuid := uuid.NewV4()
	scope.Set("ID", uuid)
}