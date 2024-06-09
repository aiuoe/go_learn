module github.com/aiuoe/go_learn

go 1.22.4

replace github.com/aiuoe/gomodule => ./module

require (
	github.com/aiuoe/gomodule v0.0.0-00010101000000-000000000000
	gorm.io/driver/sqlite v1.5.5
	gorm.io/gorm v1.25.10
)

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.22 // indirect
)
