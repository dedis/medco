package survivalserver

import (
	"fmt"
	"strings"

	utilserver "github.com/ldsec/medco/connector/util/server"
	"github.com/sirupsen/logrus"
)

// getConceptCodes returns all concept codes for a given path and its descendants
func getConceptCodes(path string) ([]string, error) {

	tableCd, pathURI := extractTableAndURI(path)

	preparedPath := prepareLike(pathURI)

	tableName, err := getTableName(tableCd)
	if err != nil {
		return nil, err
	}

	description := fmt.Sprintf("getConceptCodes (table name: %s, path: %s), SQL: %s", tableName, preparedPath, conceptCodes)
	logrus.Debugf(" running: %s", description)

	rows, err := utilserver.I2B2DBConnection.Query(conceptCodes, tableName, preparedPath)
	if err != nil {
		err = fmt.Errorf("while selecting concept codes: %s, DB operation: %s", err.Error(), description)
		logrus.Error(err)
		return nil, err
	}

	resString := new(string)
	res := make([]string, 0)
	for rows.Next() {

		err = rows.Scan(resString)
		if err != nil {
			err = fmt.Errorf("while scanning SQL record: %s, DB operation: %s", err.Error(), description)
			logrus.Error(err)
			return nil, err
		}

		res = append(res, *resString)
	}
	logrus.Tracef("concept codes are %v, DB operation: %s", res, description)
	err = rows.Close()
	if err != nil {
		err = fmt.Errorf("while closing SQL record stream: %s, DB operation: %s", err.Error(), description)
		logrus.Error(err)
		return nil, err
	}
	logrus.Debugf("successfully retrieved %d concept codes, DB operation: %s", len(res), description)

	return res, nil

}

// getModifierCodes returns all modifier codes for a given path and its descendants, and exactly matching the appliedPath
func getModifierCodes(path string, appliedPath string) ([]string, error) {

	tableCD, path := extractTableAndURI(path)

	preparedPath := prepareLike(path)

	tableName, err := getTableName(tableCD)
	if err != nil {
		return nil, err
	}

	preparedAppliedPath := prepareEqual(appliedPath)
	description := fmt.Sprintf("getModifierCodes (table name: %s, path: %s, applied path: %s), SQL: %s", tableName, preparedPath, preparedAppliedPath, modifierCodes)
	logrus.Debugf("running: %s", description)
	rows, err := utilserver.I2B2DBConnection.Query(modifierCodes, tableName, preparedPath, preparedAppliedPath)
	if err != nil {
		err = fmt.Errorf("while selecting modifier codes: %s, DB operation: %s", err.Error(), description)
		return nil, err
	}

	resString := new(string)
	res := make([]string, 0)
	for rows.Next() {

		err = rows.Scan(resString)
		if err != nil {
			err = fmt.Errorf("while scanning SQL record: %s, DB operation: %s", err.Error(), description)
			return nil, err
		}

		res = append(res, *resString)
	}
	logrus.Tracef("modifier codes are %v, DB operation: %s", res, description)

	err = rows.Close()
	if err != nil {
		err = fmt.Errorf("while closing SQL record stream: %s, DB operation: %s", err.Error(), description)
		return nil, err
	}

	logrus.Debugf("successfully retrieved %d modifier codes, DB operation: %s", len(res), description)
	return res, nil
}

// getTableName get the ontology table name for a given table code (in I2B2, the first node ofa URI is the table CD)
// getTableName returns an error when no entry was found for the provided table code.
func getTableName(tableCD string) (string, error) {
	upper := strings.ToUpper(tableCD)
	description := fmt.Sprintf("getTableName (table code: %s), SQL : %s", upper, tableName)
	logrus.Debugf("runnig: %s", description)
	row := utilserver.I2B2DBConnection.QueryRow(tableName, upper)
	ret := new(string)
	err := row.Scan(ret)
	if err != nil {
		err = fmt.Errorf("while getting table name: %s, DB operation: %s", err.Error(), description)
		logrus.Error(err)
		return "", err
	}
	logrus.Debugf(`successfully selected table name "%s", DB operation: %s`, *ret, description)

	return strings.ToLower(*ret), nil
}

// extracts table name and URI
func extractTableAndURI(pathURI string) (tableName, pathWoTable string) {
	pathURI = strings.Trim(pathURI, "/")
	tokens := strings.Split(pathURI, "/")
	switch len(tokens) {
	case 0:
		return "", ""
	case 1:
		return tokens[0], ""
	default:
		return tokens[0], "/" + strings.Join(tokens[1:], "/")
	}
}

// prepareLike prepare path for LIKE operator
func prepareLike(pathURI string) string {
	path := strings.Replace(pathURI, "/", `\\`, -1)
	if strings.HasSuffix(path, "%") {
		return path
	}
	if strings.HasSuffix(path, `\\`) {
		return path + "%"
	}
	return path + `\\%`
}

// prepareEqual prepare path for LIKE operator
func prepareEqual(pathURI string) string {
	return strings.Replace(pathURI, "/", `\`, -1)
}

const conceptCodes = `
SELECT medco_ont.get_concept_codes($1,$2);
`

const modifierCodes = `
SELECT medco_ont.get_modifier_codes($1, $2,$3);
`

const tableName = `
SELECT c_table_name from medco_ont.table_access
WHERE c_table_cd = $1;
`