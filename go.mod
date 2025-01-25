module go-crud-app

go 1.16

require (
	// Adicione suas dependências aqui, por exemplo:
	github.com/gin-gonic/gin v1.7.7
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/jinzhu/gorm v1.9.16
	gorm.io/driver/sqlite v1.5.7 // indirect
	gorm.io/gorm v1.25.7-0.20240204074919-46816ad31dde // indirect
)

replace github.com/jinzhu/gorm/dialects/sqlite => github.com/go-gorm/sqlite v1.0.6
