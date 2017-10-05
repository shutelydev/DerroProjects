package model

import (
	"time"

	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

type Booking struct {
	Id         int
	Customer   string
	Phone      string
	DateBooked time.Time `orm:"column(dateBooked);type(date)"`
	Room       int
	TimeStamp  time.Time `orm:"auto_now;type(timeStamp);column(timeStamp)"`
}

var o orm.Ormer

func init() {
	orm.RegisterDataBase("default", "postgres", "postgres://postgres:19840526@localhost/Fishing?sslmode=disable")
	// register model
	orm.RegisterModel(new(Booking))
	o = orm.NewOrm()
	o.Using("Fishing")
}

func (b Booking) Insert() (bSuccess bool, bExitCode int) {
	if created, _, err := o.ReadOrCreate(&b, "room", "dateBooked"); err == nil {
		if created {
			return true, 0
		} else {
			return true, 1
		}
	} else {
		fmt.Println(err)
		return false, -1

	}
}

func (b Booking) Delete() (dels int64) {
	o := orm.NewOrm()

	seter := o.QueryTable("Booking")
	if b.Id > 0 {
		seter = seter.Filter("id", b.Id)
	}

	if len(b.Customer) > 0 {
		seter = seter.Filter("customer", b.Customer)
	}

	if len(b.Phone) > 0 {
		seter = seter.Filter("phone", b.Phone)
	}

	if b.DateBooked.IsZero() == false {
		seter = seter.Filter("dateBooked", b.DateBooked)
	}

	if b.Room > 0 {
		seter = seter.Filter("room", b.Room)
	}

	if num, err := seter.Delete(); err == nil {
		return num
	} else {
		fmt.Println(err)
		return -1
	}
}

func (b Booking) Query() *[]Booking {
	var books []Booking

	seter := o.QueryTable("Booking")
	if b.Id > 0 {
		seter = seter.Filter("id", b.Id)
	}

	if len(b.Customer) > 0 {
		seter = seter.Filter("customer", b.Customer)
	}

	if len(b.Phone) > 0 {
		seter = seter.Filter("phone", b.Phone)
	}

	if b.DateBooked.IsZero() == false {
		seter = seter.Filter("dateBooked", b.DateBooked)
	}

	if b.Room > 0 {
		seter = seter.Filter("room", b.Room)
	}

	seter.All(&books)

	return &books
}

func (b Booking) Update(params orm.Params) (updates int64) {
	o := orm.NewOrm()

	seter := o.QueryTable("Booking")
	if b.Id > 0 {
		seter = seter.Filter("id", b.Id)
	}

	if len(b.Customer) > 0 {
		seter = seter.Filter("customer", b.Customer)
	}

	if len(b.Phone) > 0 {
		seter = seter.Filter("phone", b.Phone)
	}

	if b.DateBooked.IsZero() == false {
		seter = seter.Filter("dateBooked", b.DateBooked)
	}

	if b.Room > 0 {
		seter = seter.Filter("room", b.Room)
	}

	if num, err := seter.Update(params); err == nil {
		return num
	} else {
		fmt.Println(err)
		return -1
	}
}
