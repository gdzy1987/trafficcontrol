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

type FederationTmuser struct {
	LastUpdated time.Time             `db:"last_updated" json:"lastUpdated"`
	Links       FederationTmuserLinks `json:"_links" db:-`
}

type FederationTmuserLinks struct {
	Self           string         `db:"self" json:"_self"`
	FederationLink FederationLink `json:"federation" db:-`
	TmUserLink     TmUserLink     `json:"tm_user" db:-`
	RoleLink       RoleLink       `json:"role" db:-`
}

// @Title getFederationTmuserById
// @Description retrieves the federation_tmuser information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    FederationTmuser
// @Resource /api/2.0
// @Router /api/2.0/federation_tmuser/{id} [get]
func getFederationTmuserById(id int, db *sqlx.DB) (interface{}, error) {
	ret := []FederationTmuser{}
	arg := FederationTmuser{}
	arg.Links.FederationLink.ID = int64(id)
	queryStr := "select *, concat('" + API_PATH + "federation_tmuser/', id) as self "
	queryStr += ", concat('" + API_PATH + "federation/', federation) as federation_id_ref"
	queryStr += ", concat('" + API_PATH + "tm_user/', tm_user) as tm_user_id_ref"
	queryStr += ", concat('" + API_PATH + "role/', role) as role_id_ref"
	queryStr += " from federation_tmuser where Links.FederationLink.ID=:Links.FederationLink.ID"
	nstmt, err := db.PrepareNamed(queryStr)
	err = nstmt.Select(&ret, arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	nstmt.Close()
	return ret, nil
}

// @Title getFederationTmusers
// @Description retrieves the federation_tmuser
// @Accept  application/json
// @Success 200 {array}    FederationTmuser
// @Resource /api/2.0
// @Router /api/2.0/federation_tmuser [get]
func getFederationTmusers(db *sqlx.DB) (interface{}, error) {
	ret := []FederationTmuser{}
	queryStr := "select *, concat('" + API_PATH + "federation_tmuser/', id) as self "
	queryStr += ", concat('" + API_PATH + "federation/', federation) as federation_id_ref"
	queryStr += ", concat('" + API_PATH + "tm_user/', tm_user) as tm_user_id_ref"
	queryStr += ", concat('" + API_PATH + "role/', role) as role_id_ref"
	queryStr += " from federation_tmuser"
	err := db.Select(&ret, queryStr)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return ret, nil
}

// @Title postFederationTmuser
// @Description enter a new federation_tmuser
// @Accept  application/json
// @Param                 Body body     FederationTmuser   true "FederationTmuser object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/federation_tmuser [post]
func postFederationTmuser(payload []byte, db *sqlx.DB) (interface{}, error) {
	var v FederationTmuser
	err := json.Unmarshal(payload, &v)
	if err != nil {
		log.Println(err)
	}
	sqlString := "INSERT INTO federation_tmuser("
	sqlString += "federation"
	sqlString += ",tm_user"
	sqlString += ",role"
	sqlString += ") VALUES ("
	sqlString += ":federation"
	sqlString += ",:tm_user"
	sqlString += ",:role"
	sqlString += ")"
	result, err := db.NamedExec(sqlString, v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}

// @Title putFederationTmuser
// @Description modify an existing federation_tmuserentry
// @Accept  application/json
// @Param   id              path    int     true        "The row id"
// @Param                 Body body     FederationTmuser   true "FederationTmuser object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/federation_tmuser/{id}  [put]
func putFederationTmuser(id int, payload []byte, db *sqlx.DB) (interface{}, error) {
	var v FederationTmuser
	err := json.Unmarshal(payload, &v)
	v.Links.FederationLink.ID = int64(id) // overwrite the id in the payload
	if err != nil {
		log.Println(err)
		return nil, err
	}
	v.LastUpdated = time.Now()
	sqlString := "UPDATE federation_tmuser SET "
	sqlString += "federation = :federation"
	sqlString += ",tm_user = :tm_user"
	sqlString += ",role = :role"
	sqlString += ",last_updated = :last_updated"
	sqlString += " WHERE Links.FederationLink.ID=:Links.FederationLink.ID"
	result, err := db.NamedExec(sqlString, v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}

// @Title delFederationTmuserById
// @Description deletes federation_tmuser information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    FederationTmuser
// @Resource /api/2.0
// @Router /api/2.0/federation_tmuser/{id} [delete]
func delFederationTmuser(id int, db *sqlx.DB) (interface{}, error) {
	arg := FederationTmuser{}
	arg.Links.FederationLink.ID = int64(id)
	result, err := db.NamedExec("DELETE FROM federation_tmuser WHERE id=:id", arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}
