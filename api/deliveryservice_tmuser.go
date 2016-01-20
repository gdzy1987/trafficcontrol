// Copyright 2015 Comcast Cable Communications Management, LLC

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file was initially generated by gen_to_start.go (add link), as a start
// of the Traffic Ops golang data model

package api

import (
	"encoding/json"
	_ "github.com/Comcast/traffic_control/traffic_ops/experimental/server/output_format" // needed for swagger
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

type DeliveryserviceTmuser struct {
	LastUpdated time.Time                  `db:"last_updated" json:"lastUpdated"`
	Links       DeliveryserviceTmuserLinks `json:"_links" db:-`
}

type DeliveryserviceTmuserLinks struct {
	Self                string              `db:"self" json:"_self"`
	DeliveryserviceLink DeliveryserviceLink `json:"deliveryservice" db:-`
	TmUserLink          TmUserLink          `json:"tm_user" db:-`
}

// @Title getDeliveryserviceTmuserById
// @Description retrieves the deliveryservice_tmuser information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    DeliveryserviceTmuser
// @Resource /api/2.0
// @Router /api/2.0/deliveryservice_tmuser/{id} [get]
func getDeliveryserviceTmuserById(id int, db *sqlx.DB) (interface{}, error) {
	ret := []DeliveryserviceTmuser{}
	arg := DeliveryserviceTmuser{}
	arg.Links.DeliveryserviceLink.ID = int64(id)
	queryStr := "select *, concat('" + API_PATH + "deliveryservice_tmuser/', id) as self "
	queryStr += ", concat('" + API_PATH + "deliveryservice/', deliveryservice) as deliveryservice_id_ref"
	queryStr += ", concat('" + API_PATH + "tm_user/', tm_user) as tm_user_id_ref"
	queryStr += " from deliveryservice_tmuser where Links.DeliveryserviceLink.ID=:Links.DeliveryserviceLink.ID"
	nstmt, err := db.PrepareNamed(queryStr)
	err = nstmt.Select(&ret, arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	nstmt.Close()
	return ret, nil
}

// @Title getDeliveryserviceTmusers
// @Description retrieves the deliveryservice_tmuser
// @Accept  application/json
// @Success 200 {array}    DeliveryserviceTmuser
// @Resource /api/2.0
// @Router /api/2.0/deliveryservice_tmuser [get]
func getDeliveryserviceTmusers(db *sqlx.DB) (interface{}, error) {
	ret := []DeliveryserviceTmuser{}
	queryStr := "select *, concat('" + API_PATH + "deliveryservice_tmuser/', id) as self "
	queryStr += ", concat('" + API_PATH + "deliveryservice/', deliveryservice) as deliveryservice_id_ref"
	queryStr += ", concat('" + API_PATH + "tm_user/', tm_user) as tm_user_id_ref"
	queryStr += " from deliveryservice_tmuser"
	err := db.Select(&ret, queryStr)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return ret, nil
}

// @Title postDeliveryserviceTmuser
// @Description enter a new deliveryservice_tmuser
// @Accept  application/json
// @Param                 Body body     DeliveryserviceTmuser   true "DeliveryserviceTmuser object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/deliveryservice_tmuser [post]
func postDeliveryserviceTmuser(payload []byte, db *sqlx.DB) (interface{}, error) {
	var v DeliveryserviceTmuser
	err := json.Unmarshal(payload, &v)
	if err != nil {
		log.Println(err)
	}
	sqlString := "INSERT INTO deliveryservice_tmuser("
	sqlString += "deliveryservice"
	sqlString += ",tm_user"
	sqlString += ") VALUES ("
	sqlString += ":deliveryservice"
	sqlString += ",:tm_user"
	sqlString += ")"
	result, err := db.NamedExec(sqlString, v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}

// @Title putDeliveryserviceTmuser
// @Description modify an existing deliveryservice_tmuserentry
// @Accept  application/json
// @Param   id              path    int     true        "The row id"
// @Param                 Body body     DeliveryserviceTmuser   true "DeliveryserviceTmuser object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/deliveryservice_tmuser/{id}  [put]
func putDeliveryserviceTmuser(id int, payload []byte, db *sqlx.DB) (interface{}, error) {
	var v DeliveryserviceTmuser
	err := json.Unmarshal(payload, &v)
	v.Links.DeliveryserviceLink.ID = int64(id) // overwrite the id in the payload
	if err != nil {
		log.Println(err)
		return nil, err
	}
	v.LastUpdated = time.Now()
	sqlString := "UPDATE deliveryservice_tmuser SET "
	sqlString += "deliveryservice = :deliveryservice"
	sqlString += ",tm_user = :tm_user"
	sqlString += ",last_updated = :last_updated"
	sqlString += " WHERE Links.DeliveryserviceLink.ID=:Links.DeliveryserviceLink.ID"
	result, err := db.NamedExec(sqlString, v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}

// @Title delDeliveryserviceTmuserById
// @Description deletes deliveryservice_tmuser information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    DeliveryserviceTmuser
// @Resource /api/2.0
// @Router /api/2.0/deliveryservice_tmuser/{id} [delete]
func delDeliveryserviceTmuser(id int, db *sqlx.DB) (interface{}, error) {
	arg := DeliveryserviceTmuser{}
	arg.Links.DeliveryserviceLink.ID = int64(id)
	result, err := db.NamedExec("DELETE FROM deliveryservice_tmuser WHERE id=:id", arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}
