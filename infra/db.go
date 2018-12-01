package infra

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	fmt.Println("db.go init...")

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:1@tcp(127.0.0.1:3306)/zhongzhen?charset=utf8", 30)

	// register model
	// orm.RegisterModel(new(User))
}

func ExistDtu(regPack []byte) bool {
	s := fmt.Sprintf("%02x", regPack)
	o := orm.NewOrm()
	var maps []orm.Params
	fmt.Printf("SELECT * FROM DTU WHERE DTU_reg_pack = %s\n", s)
	num, _ := o.Raw("SELECT * FROM DTU WHERE DTU_reg_pack = ?", s).Values(&maps)
	fmt.Println(maps)
	if num > 0 {
		return true
	} else {
		return false
	}
}

func ShowAllTables() {
	o := orm.NewOrm()
	var lists []orm.ParamsList
	o.Raw("SHOW TABLES").ValuesList(&lists)
	for i := range lists {
		fmt.Printf("%d: %s\t", i, lists[i])
	}
	fmt.Printf("\n")
}
