package deployment

import (
	mysql "app/backend/common/util/mysql"
	localtime "app/backend/common/util/time"
	// "encoding/json"
	// "fmt"
	// "log
	"log"
)

const (
	DEPLOYMENT_SELECT = "SELECT id, name, actionType, actionVerb, actionUrl, actionAt, actionOp, dcList, success, reason, json, comment FROM deployment where id=?"
	DEPLOYMENT_BYNAME = "SELECT id, name, actionType, actionVerb, actionUrl, actionAt, actionOp, dcList, success, reason, json, comment FROM deployment where name=? ORDER BY id DESC"
	DEPLOYMENT_INSERT = "INSERT INTO deployment(name, actionType, actionVerb, actionUrl, actionAt, actionOp, dcList, sucdes, reason, json, comment) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	VALID = 1
	INVALID = 0
)

type Deployment struct {
	Id         int32  `json:"id"`
	Name       string `json:"name"`
	ActionType int32  `json:"actionType"`
	ActionVerb string `json:"actionVerb"`
	ActionUrl  string `json:"actionUrl"`
	ActionAt   string `json:"actionAt"`
	ActionOp   int32  `json:"actionOp"`
	DcList     string `json:"dcList"`
	Success    int32  `json:"success"`
	Reason     string `json:"reason",omitempty`
	Json       string `json:"json"`
	Comment    string `json:"comment,omitempty"`
}

func NewDeployment(name, actionVerb, actionUrl, dcList, reason, json, comment string, actionType, actionOp, success int32) *Deployment {
	return &Deployment{
		Name: name,
		ActionType: actionType,
		ActionVerb: actionVerb,
		ActionUrl: actionUrl,
		ActionAt: localtime.NewLocalTime().String(),
		ActionOp: actionOp,
 		DcList: dcList,
		Success: success,
		Reason: reason,
		Json: json,
		Comment: comment,
	}
}

func (d *Deployment) QueryDeploymentById(id int32) {
	db := mysql.MysqlInstance().Conn()

	//Prepare select-statement
	stmt, err := db.Prepare(DEPLOYMENT_SELECT)
	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}
	defer stmt.Close()

	// Query user by id
	err = stmt.QueryRow(id).Scan(&d.Id, &d.Name, &d.ActionType, &d.ActionVerb, &d.ActionVerb, &d.ActionUrl, &d.ActionAt, &d.ActionOp, &d.DcList, &d.Success, &d.Reason, &d.Json, &d.Comment)
	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}
}

func (d *Deployment) InsertDeployment() {
	db := mysql.MysqlInstance().Conn()

	// Prepare insert-statement
	stmt, err := db.Prepare(DEPLOYMENT_INSERT)
	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}
	defer stmt.Close()

	// Update ActionAt
	d.ActionAt = localtime.NewLocalTime().String()

	// Insert a deployment
	_, err = stmt.Exec(d.Name, d.ActionType, d.ActionVerb, d.ActionUrl, d.ActionAt, d.ActionOp, d.DcList, d.Success, d.Reason, d.Json, d.Comment)
	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}

}

func QueryDeploymentByAppName(name string) *[]*Deployment {
	// New deployment point array
	deployments := new([]*Deployment)

	db := mysql.MysqlInstance().Conn()

	// Prepare select-by-name-statement
	stmt, err := db.Prepare(DEPLOYMENT_BYNAME)
	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(name)
	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		d := new(Deployment)
		err = rows.Scan(&d.Id, &d.Name, &d.ActionType, &d.ActionVerb, &d.ActionUrl, &d.ActionAt, &d.ActionOp, &d.DcList, &d.Success, &d.Reason, &d.Json, &d.Comment)
		if err != nil {
			log.Fatal(err)
			panic(err.Error())
		}
		*deployments = append(*deployments, d)
	}

	return deployments
}
