package i2b2

import (
	"github.com/ldsec/medco/connector/util/server"
	"testing"
)

func init() {
	utilserver.I2b2HiveURL = "http://localhost:8090/i2b2/services"
	utilserver.I2b2LoginDomain = "i2b2medcosrv0"
	utilserver.I2b2LoginProject = "MedCo"
	utilserver.I2b2LoginUser = "e2etest"
	utilserver.I2b2LoginPassword = "e2etest"
	utilserver.SetLogLevel("5")
}

// warning: all tests need the dev-local-3nodes medco deployment running locally, loaded with default data

// test ontology search query
func TestGetOntologyChildrenRoot(t *testing.T) {

	results, err := GetOntologyChildren("/")
	if err != nil {
		t.Fail()
	}
	t.Log(*results[0])
}

func TestGetOntologyChildrenNode(t *testing.T) {

	results, err := GetOntologyChildren("/E2ETEST/e2etest/")
	if err != nil {
		t.Fail()
	}
	t.Log(*results[0].MedcoEncryption)
}

func TestGetOntologyModifiers(t *testing.T) {

	results, err := GetOntologyModifiers("/E2ETEST/e2etest/1/")
	if err != nil {
		t.Fail()
	}
	t.Log(*results[0].MedcoEncryption)
}

func TestGetOntologyModifierChildren(t *testing.T) {

	results, err := GetOntologyModifierChildren("/E2ETEST/modifiers/", "/e2etest/%", "/E2ETEST/e2etest/1/")
	if err != nil {
		t.Fail()
	}
	t.Log(*results[0].MedcoEncryption)
}

func TestExecutePsmQuery(t *testing.T) {

	patientCount, patientSetID, err := ExecutePsmQuery(
		"testQuery",
		[][]string{{`\\SENSITIVE_TAGGED\medco\tagged\fa15afdd3ce192fffde16d4ed10690b206d7cc95bfce778797cc9a05c312a35d\`}},
		[]bool{false},
	)
	if err != nil {
		t.Fail()
	}
	t.Log("count:"+patientCount, "set ID:"+patientSetID)
}

func TestGetPatientSet(t *testing.T) {

	patientIDs, patientDummyFlags, err := GetPatientSet("9")
	if err != nil {
		t.Fail()
	}
	t.Log(patientIDs)
	t.Log(patientDummyFlags)
}