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
	null "gopkg.in/guregu/null.v3"
	"log"
	"time"
)

type JobResult struct {
	Id          int64          `db:"id" json:"id"`
	Result      string         `db:"result" json:"result"`
	Description null.String    `db:"description" json:"description"`
	LastUpdated time.Time      `db:"last_updated" json:"lastUpdated"`
	Links       JobResultLinks `json:"_links" db:-`
}

type JobResultLinks struct {
	Self         string       `db:"self" json:"_self"`
	JobLink      JobLink      `json:"job" db:-`
	JobAgentLink JobAgentLink `json:"job_agent" db:-`
}

// @Title getJobResultById
// @Description retrieves the job_result information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    JobResult
// @Resource /api/2.0
// @Router /api/2.0/job_result/{id} [get]
func getJobResultById(id int, db *sqlx.DB) (interface{}, error) {
	ret := []JobResult{}
	arg := JobResult{}
	arg.Id = int64(id)
	queryStr := "select *, concat('" + API_PATH + "job_result/', id) as self "
	queryStr += ", concat('" + API_PATH + "job/', job) as job_id_ref"
	queryStr += ", concat('" + API_PATH + "job_agent/', job_agent) as job_agent_id_ref"
	queryStr += " from job_result where id=:id"
	nstmt, err := db.PrepareNamed(queryStr)
	err = nstmt.Select(&ret, arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	nstmt.Close()
	return ret, nil
}

// @Title getJobResults
// @Description retrieves the job_result
// @Accept  application/json
// @Success 200 {array}    JobResult
// @Resource /api/2.0
// @Router /api/2.0/job_result [get]
func getJobResults(db *sqlx.DB) (interface{}, error) {
	ret := []JobResult{}
	queryStr := "select *, concat('" + API_PATH + "job_result/', id) as self "
	queryStr += ", concat('" + API_PATH + "job/', job) as job_id_ref"
	queryStr += ", concat('" + API_PATH + "job_agent/', job_agent) as job_agent_id_ref"
	queryStr += " from job_result"
	err := db.Select(&ret, queryStr)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return ret, nil
}

// @Title postJobResult
// @Description enter a new job_result
// @Accept  application/json
// @Param                 Body body     JobResult   true "JobResult object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/job_result [post]
func postJobResult(payload []byte, db *sqlx.DB) (interface{}, error) {
	var v JobResult
	err := json.Unmarshal(payload, &v)
	if err != nil {
		log.Println(err)
	}
	sqlString := "INSERT INTO job_result("
	sqlString += "job"
	sqlString += ",job_agent"
	sqlString += ",result"
	sqlString += ",description"
	sqlString += ") VALUES ("
	sqlString += ":job"
	sqlString += ",:job_agent"
	sqlString += ",:result"
	sqlString += ",:description"
	sqlString += ")"
	result, err := db.NamedExec(sqlString, v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}

// @Title putJobResult
// @Description modify an existing job_resultentry
// @Accept  application/json
// @Param   id              path    int     true        "The row id"
// @Param                 Body body     JobResult   true "JobResult object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/job_result/{id}  [put]
func putJobResult(id int, payload []byte, db *sqlx.DB) (interface{}, error) {
	var v JobResult
	err := json.Unmarshal(payload, &v)
	v.Id = int64(id) // overwrite the id in the payload
	if err != nil {
		log.Println(err)
		return nil, err
	}
	v.LastUpdated = time.Now()
	sqlString := "UPDATE job_result SET "
	sqlString += "job = :job"
	sqlString += ",job_agent = :job_agent"
	sqlString += ",result = :result"
	sqlString += ",description = :description"
	sqlString += ",last_updated = :last_updated"
	sqlString += " WHERE id=:id"
	result, err := db.NamedExec(sqlString, v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}

// @Title delJobResultById
// @Description deletes job_result information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    JobResult
// @Resource /api/2.0
// @Router /api/2.0/job_result/{id} [delete]
func delJobResult(id int, db *sqlx.DB) (interface{}, error) {
	arg := JobResult{}
	arg.Id = int64(id)
	result, err := db.NamedExec("DELETE FROM job_result WHERE id=:id", arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}
