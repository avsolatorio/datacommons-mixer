// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"encoding/json"
	"testing"

	pb "github.com/datacommonsorg/mixer/proto"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"
)

const covidJSON = `{
  "triples": [
    {
      "subjectId": "Covid19CumulativeCases",
      "predicate": "typeOf",
      "objectId": "StatisticalVariable",
      "objectName": "StatisticalVariable",
      "objectTypes": ["Class"],
      "provenanceId": "dc/5l5zxr1"
    },
    {
      "subjectId": "Covid19CumulativeCases",
      "predicate": "statType",
      "objectId": "measuredValue",
      "objectName": "measuredValue",
      "objectTypes": ["Property"],
      "provenanceId": "dc/5l5zxr1"
    },
    {
      "subjectId": "Covid19CumulativeCases",
      "predicate": "provenance",
      "objectId": "dc/5l5zxr1",
      "objectName": "http://schema.datacommons.org",
      "objectTypes": ["Provenance"],
      "provenanceId": "dc/5l5zxr1"
    },
    {
      "subjectId": "Covid19CumulativeCases",
      "predicate": "populationType",
      "objectId": "MedicalConditionIncident",
      "objectName": "MedicalConditionIncident",
      "objectTypes": ["Class"],
      "provenanceId": "dc/5l5zxr1"
    },
    {
      "subjectId": "Covid19CumulativeCases",
      "predicate": "medicalStatus",
      "objectId": "ConfirmedOrProbableCase",
      "objectName": "ConfirmedOrProbableCase",
      "objectTypes": ["MedicalStatusEnum"],
      "provenanceId": "dc/5l5zxr1"
    },
    {
      "subjectId": "Covid19CumulativeCases",
      "predicate": "measurementMethod",
      "objectId": "COVID19",
      "provenanceId": "dc/5l5zxr1"
    },
    {
      "subjectId": "Covid19CumulativeCases",
      "predicate": "measuredProperty",
      "objectId": "cumulativeCount",
      "objectName": "cumulativeCount",
      "objectTypes": ["Property"],
      "provenanceId": "dc/5l5zxr1"
    },
    {
      "subjectId": "Covid19CumulativeCases",
      "predicate": "incidentType",
      "objectId": "COVID_19",
      "objectName": "COVID_19",
      "objectTypes": ["MedicalConditionIncidentTypeEnum"],
      "provenanceId": "dc/5l5zxr1"
    },
    {
      "subjectId": "Covid19CumulativeCases",
      "predicate": "constraintProperties",
      "objectId": "medicalStatus",
      "objectName": "medicalStatus",
      "objectTypes": ["Property"],
      "provenanceId": "dc/5l5zxr1"
    },
    {
      "subjectId": "Covid19CumulativeCases",
      "predicate": "constraintProperties",
      "objectId": "incidentType",
      "objectName": "incidentType",
      "objectTypes": ["Property"],
      "provenanceId": "dc/5l5zxr1"
    }
  ]
}`

const populationJSON = `{
  "triples": [
    {
      "subjectId": "TotalPopulation",
      "predicate": "typeOf",
      "objectId": "StatisticalVariable",
      "objectName": "StatisticalVariable",
      "objectTypes": ["Class"],
      "provenanceId": "dc/5l5zxr1"
    },
    {
      "subjectId": "TotalPopulation",
      "predicate": "statType",
      "objectId": "measuredValue",
      "objectName": "measuredValue",
      "objectTypes": ["Property"],
      "provenanceId": "dc/5l5zxr1"
    },
    {
      "subjectId": "TotalPopulation",
      "predicate": "provenance",
      "objectId": "dc/5l5zxr1",
      "objectName": "http://schema.datacommons.org",
      "objectTypes": ["Provenance"],
      "provenanceId": "dc/5l5zxr1"
    },
    {
      "subjectId": "TotalPopulation",
      "predicate": "populationType",
      "objectId": "Person",
      "objectName": "Person",
      "objectTypes": ["Class"],
      "provenanceId": "dc/5l5zxr1"
    },
    {
      "subjectId": "TotalPopulation",
      "predicate": "measurementMethod",
      "objectId": "CensusACS5yrSurvey",
      "objectName": "CensusACS5yrSurvey",
      "objectTypes": ["CensusSurveyEnum"],
      "provenanceId": "dc/5l5zxr1"
    },
    {
      "subjectId": "TotalPopulation",
      "predicate": "measuredProperty",
      "objectId": "count",
      "objectName": "count",
      "objectTypes": ["Property"],
      "provenanceId": "dc/5l5zxr1"
		},
    {
      "subjectId": "dc/d20yq6r828zrc",
      "subjectTypes": [
        "StatisticalVariable"
      ],
      "predicate": "measurementDenominator",
      "objectId": "TotalPopulation",
      "objectTypes": [
        "StatisticalVariable"
      ],
      "provenanceId": "dc/cweckx1"
    }
  ]
}`

func TestTriplesToStatsVar(t *testing.T) {
	var covidStatsVarTriples TriplesCache
	var populationStatsVarTriples TriplesCache
	err := json.Unmarshal([]byte(covidJSON), &covidStatsVarTriples)
	if err != nil {
		t.Errorf("Unmarshal got err %v", err)
		return
	}
	err = json.Unmarshal([]byte(populationJSON), &populationStatsVarTriples)
	if err != nil {
		t.Errorf("Unmarshal got err %v", err)
		return
	}
	for _, c := range []struct {
		statsVar     string
		triples      *TriplesCache
		wantStatsVar *StatisticalVariable
		wantErr      bool
	}{
		{
			"Covid19CumulativeCases",
			&covidStatsVarTriples,
			&StatisticalVariable{
				PopType: "MedicalConditionIncident",
				PVs: map[string]string{
					"incidentType":  "COVID_19",
					"medicalStatus": "ConfirmedOrProbableCase",
				},
				MeasuredProp:      "cumulativeCount",
				MeasurementMethod: "COVID19",
				StatType:          "measured",
			},
			false,
		},
		{
			"TotalPopulation",
			&populationStatsVarTriples,
			&StatisticalVariable{
				PopType:           "Person",
				MeasuredProp:      "count",
				MeasurementMethod: "CensusACS5yrSurvey",
				StatType:          "measured",
			},
			false,
		},
	} {
		gotStatsVar, err := triplesToStatsVar(c.statsVar, c.triples)
		if c.wantErr {
			if err == nil {
				t.Errorf("triplesToStatsVar want error for triples %v", c.triples)
			}
			continue
		}
		if err != nil {
			t.Errorf("triplesToStatsVar(%v) = %v", c.triples, err)
			continue
		}
		if diff := cmp.Diff(gotStatsVar, c.wantStatsVar); diff != "" {
			t.Errorf("triplesToStatsVar() got diff %+v", diff)
		}
	}
}

func TestFilterAndRank(t *testing.T) {
	for _, c := range []struct {
		input   *pb.ObsTimeSeries
		mmethod string
		unit    string
		op      string
		want    *pb.ObsTimeSeries
	}{
		// Default ranking
		{
			&pb.ObsTimeSeries{
				SourceSeries: []*pb.ObsTimeSeries_SourceSeries{
					{
						Val:               map[string]float64{"2011": 100, "2012": 101},
						MeasurementMethod: "CensusPEPSurvey",
						ImportName:        "CensusPEP",
					},
					{
						Val:               map[string]float64{"2011": 101, "2012": 102, "2013": 103},
						MeasurementMethod: "CensusACS5yrSurvey",
						ImportName:        "CensusACS5YearSurvey",
					},
				},
			},
			"",
			"",
			"",
			&pb.ObsTimeSeries{
				Data:   map[string]float64{"2011": 100, "2012": 101},
				Source: "CensusPEP",
			},
		},
		// Filter by mmethod
		{
			&pb.ObsTimeSeries{
				SourceSeries: []*pb.ObsTimeSeries_SourceSeries{
					{
						Val:               map[string]float64{"2011": 100, "2012": 101},
						MeasurementMethod: "CensusPEPSurvey",
						ImportName:        "CensusPEP",
					},
					{
						Val:               map[string]float64{"2011": 101, "2012": 102, "2013": 103},
						MeasurementMethod: "CensusACS5yrSurvey",
						ImportName:        "CensusACS5YearSurvey",
					},
				},
			},
			"CensusACS5yrSurvey",
			"",
			"",
			&pb.ObsTimeSeries{
				Data:   map[string]float64{"2011": 101, "2012": 102, "2013": 103},
				Source: "CensusACS5YearSurvey",
			},
		},
		// Filter by observation period
		{
			&pb.ObsTimeSeries{
				SourceSeries: []*pb.ObsTimeSeries_SourceSeries{
					{
						Val:               map[string]float64{"2011": 100, "2012": 101},
						MeasurementMethod: "CensusPEPSurvey",
						ImportName:        "CensusPEP",
						ObservationPeriod: "P1Y",
					},
					{
						Val:               map[string]float64{"2017": 101},
						MeasurementMethod: "CensusPEPSurvey",
						ImportName:        "CensusPEP",
						ObservationPeriod: "P2Y",
					},
				},
			},
			"",
			"",
			"P2Y",
			&pb.ObsTimeSeries{
				Data:   map[string]float64{"2017": 101},
				Source: "CensusPEP",
			},
		},
		// No match
		{
			&pb.ObsTimeSeries{
				SourceSeries: []*pb.ObsTimeSeries_SourceSeries{
					{
						Val:               map[string]float64{"2011": 100, "2012": 101},
						MeasurementMethod: "CensusPEPSurvey",
						ImportName:        "CensusPEP",
						ObservationPeriod: "P1Y",
					},
					{
						Val:               map[string]float64{"2017": 101},
						MeasurementMethod: "CensusPEPSurvey",
						ImportName:        "CensusPEP",
						ObservationPeriod: "P2Y",
					},
				},
			},
			"",
			"",
			"P3Y",
			&pb.ObsTimeSeries{},
		},
	} {
		got := filterAndRank(c.input, c.mmethod, c.op, c.unit)
		if diff := cmp.Diff(got, c.want, protocmp.Transform()); diff != "" {
			t.Errorf("filterAndRank() got diff %+v", diff)
		}
	}
}
