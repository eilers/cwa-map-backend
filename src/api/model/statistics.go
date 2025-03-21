/*
 *   Corona-Warn-App / cwa-map-backend
 *
 *   (C) 2020, T-Systems International GmbH
 *
 *   Deutsche Telekom AG and all other contributors /
 *   copyright owners license this file to you under the Apache
 *   License, Version 2.0 (the "License"); you may not use this
 *   file except in compliance with the License.
 *   You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 *   Unless required by applicable law or agreed to in writing,
 *   software distributed under the License is distributed on an
 *   "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 *   KIND, either express or implied.  See the License for the
 *   specific language governing permissions and limitations
 *   under the License.
 */

package model

import "com.t-systems-mms.cwa/repositories"

type ReportStatisticsDTO struct {
	Subject     string `json:"subject"`
	ReportCount uint   `json:"report_count"`
}

func (dto ReportStatisticsDTO) FromModel(model repositories.ReportStatistics) *ReportStatisticsDTO {
	return &ReportStatisticsDTO{
		Subject:     model.Subject,
		ReportCount: model.Count,
	}
}
