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

type FederationResolver struct {
	Id          int64                   `db:"id" json:"id"`
	IpAddress   string                  `db:"ip_address" json:"ipAddress"`
	LastUpdated time.Time               `db:"last_updated" json:"lastUpdated"`
	Links       FederationResolverLinks `json:"_links" db:-`
}

type FederationResolverLinks struct {
	Self     string   `db:"self" json:"_self"`
	TypeLink TypeLink `json:"type" db:-`
}

type FederationResolverLink struct {
	ID  int64  `db:"federation_resolver" json:"id"`
	Ref string `db:"federation_resolver_id_ref" json:"_ref"`
}

// @Title getFederationResolverById
// @Description retrieves the federation_resolver information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    FederationResolver
// @Resource /api/2.0
// @Router /api/2.0/federation_resolver/{id} [get]
func getFederationResolverById(id int, db *sqlx.DB) (interface{}, error) {
	ret := []FederationResolver{}
	arg := FederationResolver{}
	arg.Id = int64(id)
	queryStr := "select *, concat('" + API_PATH + "federation_resolver/', id) as self "
	queryStr += ", concat('" + API_PATH + "type/', type) as type_id_ref"
	queryStr += " from federation_resolver where id=:id"
	nstmt, err := db.PrepareNamed(queryStr)
	err = nstmt.Select(&ret, arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	nstmt.Close()
	return ret, nil
}

// @Title getFederationResolvers
// @Description retrieves the federation_resolver
// @Accept  application/json
// @Success 200 {array}    FederationResolver
// @Resource /api/2.0
// @Router /api/2.0/federation_resolver [get]
func getFederationResolvers(db *sqlx.DB) (interface{}, error) {
	ret := []FederationResolver{}
	queryStr := "select *, concat('" + API_PATH + "federation_resolver/', id) as self "
	queryStr += ", concat('" + API_PATH + "type/', type) as type_id_ref"
	queryStr += " from federation_resolver"
	err := db.Select(&ret, queryStr)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return ret, nil
}

// @Title postFederationResolver
// @Description enter a new federation_resolver
// @Accept  application/json
// @Param                 Body body     FederationResolver   true "FederationResolver object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/federation_resolver [post]
func postFederationResolver(payload []byte, db *sqlx.DB) (interface{}, error) {
	var v FederationResolver
	err := json.Unmarshal(payload, &v)
	if err != nil {
		log.Println(err)
	}
	sqlString := "INSERT INTO federation_resolver("
	sqlString += "ip_address"
	sqlString += ",type"
	sqlString += ") VALUES ("
	sqlString += ":ip_address"
	sqlString += ",:type"
	sqlString += ")"
	result, err := db.NamedExec(sqlString, v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}

// @Title putFederationResolver
// @Description modify an existing federation_resolverentry
// @Accept  application/json
// @Param   id              path    int     true        "The row id"
// @Param                 Body body     FederationResolver   true "FederationResolver object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/federation_resolver/{id}  [put]
func putFederationResolver(id int, payload []byte, db *sqlx.DB) (interface{}, error) {
	var v FederationResolver
	err := json.Unmarshal(payload, &v)
	v.Id = int64(id) // overwrite the id in the payload
	if err != nil {
		log.Println(err)
		return nil, err
	}
	v.LastUpdated = time.Now()
	sqlString := "UPDATE federation_resolver SET "
	sqlString += "ip_address = :ip_address"
	sqlString += ",type = :type"
	sqlString += ",last_updated = :last_updated"
	sqlString += " WHERE id=:id"
	result, err := db.NamedExec(sqlString, v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}

// @Title delFederationResolverById
// @Description deletes federation_resolver information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    FederationResolver
// @Resource /api/2.0
// @Router /api/2.0/federation_resolver/{id} [delete]
func delFederationResolver(id int, db *sqlx.DB) (interface{}, error) {
	arg := FederationResolver{}
	arg.Id = int64(id)
	result, err := db.NamedExec("DELETE FROM federation_resolver WHERE id=:id", arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}
