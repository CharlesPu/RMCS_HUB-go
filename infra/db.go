/*******************************
@@Author     : Charles
@@Date       : 2018-12-20
@@Mail       : pu17rui@sina.com
@@Description:
* insert: 	user := new(User)
			user.MyName = "tom"
			user.MyAge = 10
			o.Insert(user)
* read:		DTU := DTU{CompanyID: 1, StationID: 1}
			o.Read(&DTU, "CompanyID", "StationID")
			fmt.Println(DTU.RegPack)
* update: 	o.QueryTable(new(User)).Filter("MyName", "jerry").Filter("MyAge", 10).Update(orm.Params{
				"MyAge": "11",
			})
* delete:	o.QueryTable("user_table").Filter("MyName", "jerry").Delete()
*******************************/
package infra

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const (
	ID_MAX_NUM = 32768
)

type DTU struct {
	RegPack   string `orm:"column(DTU_RegPack);pk"`
	CompanyID int    `orm:"column(DTU_CompanyID)"`
	StationID int    `orm:"column(DTU_StationID)"`
	IPC_DevID string `orm:"column(DTU_IPC_DevID)"`
}
type VibraPara struct {
	ID             int     `orm:"column(VibraPara_ID);pk"`
	CompanyID      int     `orm:"column(VibraPara_CompanyID)"`
	StationID      int     `orm:"column(VibraPara_StationID)"`
	BladeAm        float32 `orm:"column(VibraPara_BladeAm)"`
	BladeEffCnt    uint32  `orm:"column(VibraPara_BladeEffCnt)"`
	CylinderAm     float32 `orm:"column(VibraPara_CylinderAm)"`
	CylinderEffCnt uint32  `orm:"column(VibraPara_CylinderEffCnt)"`
	AlarmNumber    uint16  `orm:"column(VibraPara_AlarmNumber)"`
}

type BladeRTInfo struct {
	ID        int     `orm:"column(BladeRTInfo_ID);pk"`
	CompanyID int     `orm:"column(BladeRTInfo_CompanyID)"`
	StationID int     `orm:"column(BladeRTInfo_StationID)"`
	Position  float32 `orm:"column(BladeRTInfo_Position)"`
}

type CylinderRTInfo struct {
	ID        int     `orm:"column(CylinderRTInfo_ID);pk"`
	CompanyID int     `orm:"column(CylinderRTInfo_CompanyID)"`
	StationID int     `orm:"column(CylinderRTInfo_StationID)"`
	CyID      int     `orm:"column(CylinderRTInfo_CyID)"`
	Position  float32 `orm:"column(CylinderRTInfo_Position)"`
}

type VibraCtrl struct {
	ID            int     `orm:"column(VibraCtrl_ID);pk"`
	CompanyID     int     `orm:"column(VibraCtrl_CompanyID)"`
	StationID     int     `orm:"column(VibraCtrl_StationID)"`
	BladeAm       float32 `orm:"column(VibraCtrl_BladeAm)"`
	BladeVibraCnt uint32  `orm:"column(VibraCtrl_BladeVibraCnt)"`
	CylinderAm    float32 `orm:"column(VibraCtrl_CylinderAm)"`
	CylinderCycle uint32  `orm:"column(VibraCtrl_CylinderCycle)"`
}
type SingleCyCtrl struct {
	ID        int     `orm:"column(SingleCyCtrl_ID);pk"`
	CompanyID int     `orm:"column(SingleCyCtrl_CompanyID)"`
	StationID int     `orm:"column(SingleCyCtrl_StationID)"`
	CyID      int     `orm:"column(SingleCyCtrl_CyID)"`
	JogUp     int     `orm:"column(SingleCyCtrl_JogUp)"`
	JogDown   int     `orm:"column(SingleCyCtrl_JogDown)"`
	Reset     int     `orm:"column(SingleCyCtrl_Reset)"`
	Amplitude float32 `orm:"column(SingleCyCtrl_Amplitude)"`
	Cycle     uint32  `orm:"column(SingleCyCtrl_Cycle)"`
}

func init() {
	fmt.Println("db.go init...")
	orm.Debug = true

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:1@tcp(127.0.0.1:3306)/zhongzhen?charset=utf8", 30)

	orm.RegisterModel(new(DTU), new(VibraPara), new(BladeRTInfo), new(CylinderRTInfo), new(VibraCtrl), new(SingleCyCtrl))
}

// alter table name
func (*DTU) TableName() string            { return "DTU" }
func (*VibraPara) TableName() string      { return "VibrationPara" }
func (*BladeRTInfo) TableName() string    { return "BladeRTInfo" }
func (*CylinderRTInfo) TableName() string { return "CylinderRTInfo" }
func (*VibraCtrl) TableName() string      { return "VibrationControlPara" }
func (*SingleCyCtrl) TableName() string   { return "SingleCyControlPara" }

func ExistDtu(regPack []byte) bool {
	s := fmt.Sprintf("%02x", regPack)
	o := orm.NewOrm()

	exist := o.QueryTable(new(DTU)).Filter("RegPack", s).Exist()
	return exist
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

func TruncateTables(ts ...string) {
	o := orm.NewOrm()
	for _, t := range ts {
		//attetion: if use o.Raw("TRUNCATE TABLE ?", t).Exec()
		//orm will execute "TRUNCATE TABLE `table_name`" which
		//has extra ``
		sqlQuery := fmt.Sprintf("TRUNCATE TABLE %s", t)
		o.Raw(sqlQuery).Exec()
	}
}

func InsertVibrationParaT(vp *VibraPara) bool {
	o := orm.NewOrm()
	id, err := o.Insert(vp)
	/*create copy of old table if too full!*/
	if id > ID_MAX_NUM {
		year := time.Now().Year()
		month := time.Now().Month()
		day := time.Now().Day()
		newTableName := fmt.Sprintf("VibrationPara_%d%d%d", year, month, day)
		o.Raw("RENAME TABLE VibrationPara TO ?", newTableName).Exec()
		o.Raw("CREATE TABLE VibrationPara LIKE ?", newTableName).Exec()
		// or o.Raw("CREATE TABLE VibrationPara SELECT * FROM ? WHERE 1=2", newTableName).Exec()
	}
	if err == nil {
		return true
	} else {
		return false
	}
}

func InsertBladeRTInfoT(brti *BladeRTInfo) bool {
	o := orm.NewOrm()
	id, err := o.Insert(brti)
	/*create copy of old table if too full!*/
	if id > ID_MAX_NUM {
		year := time.Now().Year()
		month := time.Now().Month()
		day := time.Now().Day()
		newTableName := fmt.Sprintf("BladeRTInfo_%d%d%d", year, month, day)
		o.Raw("RENAME TABLE BladeRTInfo TO ?", newTableName).Exec()
		o.Raw("CREATE TABLE BladeRTInfo LIKE ?", newTableName).Exec()
		// or o.Raw("CREATE TABLE VibrationPara SELECT * FROM ? WHERE 1=2", newTableName).Exec()
	}
	if err == nil {
		return true
	} else {
		return false
	}
}

func InsertCylinderRTInfoT(crtis []CylinderRTInfo) bool {
	o := orm.NewOrm()
	var err error
	var id int64
	for _, val := range crtis {
		id, err = o.Insert(&val)
	}
	/*create copy of old table if too full!*/
	if id > ID_MAX_NUM {
		year := time.Now().Year()
		month := time.Now().Month()
		day := time.Now().Day()
		newTableName := fmt.Sprintf("CylinderRTInfo_%d%d%d", year, month, day)
		o.Raw("RENAME TABLE CylinderRTInfo TO ?", newTableName).Exec()
		o.Raw("CREATE TABLE CylinderRTInfo LIKE ?", newTableName).Exec()
		// or o.Raw("CREATE TABLE VibrationPara SELECT * FROM ? WHERE 1=2", newTableName).Exec()
	}
	if err == nil {
		return true
	} else {
		return false
	}
}

/*************************************************
@Description:
@Input:
@Output:
@Return: true: vc will be valuable, false: vc is not usable
@Others:
*************************************************/
func ReadVibrationControlPara() (vc VibraCtrl, ret bool) {
	// vc = new(VibraCtrl)
	o := orm.NewOrm()
	err := o.QueryTable(new(VibraCtrl)).One(&vc)
	// orm.ErrMultiRows  or  orm.ErrNoRows
	if err == nil { // exist
		// delete all
		o.Raw("TRUNCATE TABLE VibrationControlPara").Exec()
		return vc, true
	} else {
		return vc, false
	}
}

/*************************************************
@Description:
@Input:
@Output:
@Return:
@Others:
*************************************************/
func ReadSingleCyControlPara() (scc []SingleCyCtrl, n int64, ret bool) {
	o := orm.NewOrm()
	// scc = make([]SingleCyCtrl, 128)
	num, err := o.QueryTable(new(SingleCyCtrl)).Limit(-1).All(&scc)
	n = num
	if n == 0 || err != nil {
		return nil, 0, false
	} else {
		o.Raw("TRUNCATE TABLE SingleCyControlPara").Exec()
		return scc, n, true
	}
}
