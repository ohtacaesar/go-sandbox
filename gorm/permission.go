package gorm

func init() {
	models = append(models, &Permission{})
}

// Permission はフィールドレベルのPermissionをテストするための構造体です。
// https://gorm.io/docs/models.html#Field-Level-Permission
type Permission struct {
	ID uint   `gorm:"primarykey;->"`
	V1 string `gorm:"<-:create"`
	V2 string `gorm:"<-:update"`
	V3 string `gorm:"->:false"`
}
