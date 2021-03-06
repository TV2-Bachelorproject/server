package model

import (
	"time"

	"github.com/TV2-Bachelorproject/server/model/public"

	"github.com/TV2-Bachelorproject/server/model/private"
	"github.com/TV2-Bachelorproject/server/model/user"
	"github.com/TV2-Bachelorproject/server/pkg/db"
)

var tables = []interface{}{
	&private.Person{},
	&private.Service{},
	&public.Category{},
	&public.Genre{},
	&public.Credit{},
	&public.CreditGroup{},
	&public.Production{},
	&public.Program{},
	&public.Season{},
	&public.Serie{},
	&user.User{},
}

func Migrate() {
	time.Sleep(100 * time.Millisecond)
	db.Migrate(tables...)

	//Add foreignKeys for Program
	db.Model(&public.Program{}).
		AddForeignKey("production_id", "productions(id)", "RESTRICT", "RESTRICT").
		AddForeignKey("season_id", "seasons(id)", "RESTRICT", "RESTRICT").
		AddForeignKey("category_id", "categories(id)", "RESTRICT", "RESTRICT").
		AddForeignKey("serie_id", "series(id)", "RESTRICT", "RESTRICT")

}

func Reset() {
	time.Sleep(100 * time.Millisecond)
	db.Reset(tables...)
}
