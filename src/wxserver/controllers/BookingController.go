package controllers

import (
	"time"

	"strconv"

	"wxserver/model"

	"fmt"

	"github.com/astaxie/beego"
)

type BookingController struct {
	beego.Controller
}

func (this *BookingController) Post() {
	name := this.GetString("name")
	phone := this.GetString("phone")
	room, _ := strconv.Atoi(this.Ctx.Input.Param(":roomId"))
	date, _ := time.Parse("2006-01-02", this.GetString("date"))
	newBook := model.Booking{Customer: name, Phone: phone, DateBooked: date, Room: room}

	created, exitcode := newBook.Insert()
	if created {
		if exitcode == 0 {
			this.Ctx.Redirect(302, "/status/success")
		} else {
			this.Ctx.Redirect(302, "/status/fail")
		}
	} else {
		this.Ctx.Redirect(302, "/status/fail")
	}

}

func (this *BookingController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplName = "booking.tpl"

	date := this.Ctx.Input.Param(":date")
	var dateBooked time.Time
	if len(date) > 0 {
		dateBooked, _ = time.Parse("2006-01-02", date)
	} else {
		dateBooked = time.Now()
	}

	booking := model.Booking{DateBooked: dateBooked}
	resuls := booking.Query()

	totalRooms, _ := strconv.Atoi(beego.AppConfig.String("total_rooms"))
	Rooms := make([]int, totalRooms)

	for _, book := range *resuls {
		Rooms[book.Room-1] = 1
	}

	fmt.Println(Rooms)

}
