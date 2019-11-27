package handlers

import (
	"errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/ldsec/medco-connector/restapi/models"
	"github.com/ldsec/medco-connector/restapi/server/operations/genomic_annotations"
	utilserver "github.com/ldsec/medco-connector/util/server"
	"github.com/sirupsen/logrus"
	"strings"
)

var genomicAnnotationsTypes = []string{"variant_name", "protein_change", "hugo_gene_symbol"}

func MedCoGenomicAnnotationsGetValuesHandler(params genomic_annotations.GetValuesParams, principal *models.User) middleware.Responder {

	err := utilserver.DBConnection.Ping()
	if err != nil {
		logrus.Error("Impossible to connect to DB: " + err.Error())
		return genomic_annotations.NewGetValuesDefault(500).WithPayload(&genomic_annotations.GetValuesDefaultBody{
			Message: "Impossible to connect to DB " + err.Error(),
		})
	}

	var query string

	query, err = buildGetValuesQuery(params)
	if err != nil {
		logrus.Error("Query execution error " + err.Error())
		return genomic_annotations.NewGetValuesDefault(500).WithPayload(&genomic_annotations.GetValuesDefaultBody{
			Message: "Query execution error: " + err.Error(),
		})
	}

	//escaping * characters
	rows, err := utilserver.DBConnection.Query(query, strings.ReplaceAll(params.Value, "*", "\\*"), *params.Limit)
	if err != nil {
		logrus.Error("Query execution error: " + err.Error())
		return genomic_annotations.NewGetValuesDefault(500).WithPayload(&genomic_annotations.GetValuesDefaultBody{
			Message: "Query execution error: " + err.Error(),
		})
	}
	defer rows.Close()

	var annotations []string
	var annotation string

	for rows.Next() {
		err := rows.Scan(&annotation)
		if err != nil {
			logrus.Error("Query result reading error: " + err.Error())
			return genomic_annotations.NewGetValuesDefault(500).WithPayload(&genomic_annotations.GetValuesDefaultBody{
				Message: "Query result reading error: " + err.Error(),
			})
		}
		annotations = append(annotations, annotation)
	}

	if len(annotations) > 0 {
		return genomic_annotations.NewGetValuesOK().WithPayload(annotations)
	} else {
		return genomic_annotations.NewGetValuesNotFound()
	}
}

func MedCoGenomicAnnotationsGetVariantsHandler(params genomic_annotations.GetVariantsParams, principal *models.User) middleware.Responder {

	err := utilserver.DBConnection.Ping()
	if err != nil {
		logrus.Error("Impossible to connect to DB: " + err.Error())
		return genomic_annotations.NewGetVariantsDefault(500).WithPayload(&genomic_annotations.GetVariantsDefaultBody{
			Message: "Impossible to connect to DB: " + err.Error(),
		})
	}
	zygosity := ""
	if len(params.Zygosity) > 0 {
		zygosity = params.Zygosity[0]

		for i := 1; i < len(params.Zygosity); i++ {
			zygosity += "|" + params.Zygosity[i]
		}
	}

	var query string

	query, err = buildGetVariantsQuery(params)
	if err != nil {
		logrus.Error("Query execution error " + err.Error())
		return genomic_annotations.NewGetValuesDefault(500).WithPayload(&genomic_annotations.GetValuesDefaultBody{
			Message: "Query execution error: " + err.Error(),
		})
	}

	rows, err := utilserver.DBConnection.Query(query, params.Value, zygosity)
	if err != nil {
		logrus.Error("Query execution error" + err.Error())
		return genomic_annotations.NewGetVariantsDefault(500).WithPayload(&genomic_annotations.GetVariantsDefaultBody{
			Message: "Query execution error " + err.Error(),
		})
	}
	defer rows.Close()

	var variants []string
	var variant string

	for rows.Next() {
		err := rows.Scan(&variant)
		if err != nil {
			logrus.Error("Query result reading error" + err.Error())
			return genomic_annotations.NewGetVariantsDefault(500).WithPayload(&genomic_annotations.GetVariantsDefaultBody{
				Message: "Query result reading error " + err.Error(),
			})
		}
		variants = append(variants, variant)
	}

	if len(variants) > 0 {
		return genomic_annotations.NewGetVariantsOK().WithPayload(variants)
	} else {
		return genomic_annotations.NewGetVariantsNotFound()
	}
}

func buildGetValuesQuery(params genomic_annotations.GetValuesParams) (string, error) {

	if contains(genomicAnnotationsTypes, params.Annotation) {
		return "SELECT annotation_value FROM genomic_annotations." + params.Annotation + " WHERE annotation_value ~* $1 ORDER BY annotation_value LIMIT $2", nil
	} else {
		return "", errors.New("Requested invalid annotation type: " + params.Annotation)
	}

}

func buildGetVariantsQuery(params genomic_annotations.GetVariantsParams) (string, error) {

	if contains(genomicAnnotationsTypes, params.Annotation) {
		if *params.Encrypted {
			return "SELECT variant_id_enc FROM genomic_annotations.genomic_annotations WHERE lower(" + params.Annotation + ") = lower($1) AND annotations ~* $2 ORDER BY variant_id", nil
		} else {
			return "SELECT variant_id FROM genomic_annotations.genomic_annotations WHERE lower(" + params.Annotation + ") = lower($1) AND annotations ~* $2 ORDER BY variant_id", nil
		}
	} else {
		return "", errors.New("Requested invalid annotation type: " + params.Annotation)
	}

}

func contains(slice []string, value string) bool {
	for _, s := range slice {
		if s == value {
			return true
		}
	}
	return false
}
