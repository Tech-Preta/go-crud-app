module go-crud-app

go 1.16

require (
	github.com/gorilla/mux v1.8.1
	gorm.io/driver/sqlite v1.5.7
	gorm.io/gorm v1.25.7-0.20240204074919-46816ad31dde
)

replace github.com/jinzhu/gorm/dialects/sqlite => github.com/go-gorm/sqlite v1.0.6
