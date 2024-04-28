package main

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

var sqlTemplate = `
CREATE DICTIONARY {{ .name }}
(
    {{ .structure }}
)
PRIMARY KEY {{ .primaryKey }}
SOURCE({{.source }})
LAYOUT({{.layout }})
LIFETIME({{.lifetime }})`

var tpl = template.Must(template.New("sql").Parse(sqlTemplate))

func (d Dictionary) ToSQL() (string, error) {
	name := d.Name
	primaryKey := ""
	if d.Structure.ID != nil {
		primaryKey = d.Structure.ID.Name
	} else {
		keys := []string{}
		for _, attribute := range d.Structure.Key.Attributes {
			keys = append(keys, attribute.Name)
		}

		primaryKey = strings.Join(keys, ",")
	}

	buff := bytes.NewBufferString("")

	err := tpl.Execute(buff, map[string]string{
		"name":       name,
		"primaryKey": primaryKey,
		"source":     d.getSourceSQLString(),
		"layout":     d.getLayoutSQLString(),
		"lifetime":   d.getLifetimeSQLString(),
		"structure":  d.getStructureSQLString(),
	})
	if err != nil {
		return "", err
	}

	return buff.String(), nil
}

func (d Dictionary) getSourceSQLString() string {
	if d.Source.File != nil {
		return fmt.Sprintf("FILE(PATH '%s' FORMAT '%s')", d.Source.File.Path, d.Source.File.Format)
	} else if d.Source.HTTP != nil {
		return fmt.Sprintf("HTTP(URL '%s' FORMAT '%s')", d.Source.HTTP.URL, d.Source.HTTP.Format)
	} else if d.Source.Clickhouse != nil {
		sourceConfig := ""
		if d.Source.Clickhouse.Host == "localhost" {
			sourceConfig = fmt.Sprintf("DB '%s' TABLE '%s'", d.Source.Clickhouse.DB, d.Source.Clickhouse.Table)
		} else {
			sourceConfig = fmt.Sprintf("HOST '%s' PORT %d USER '%s' PASSWORD '%s' DB '%s' TABLE '%s'", d.Source.Clickhouse.Host, d.Source.Clickhouse.Port, d.Source.Clickhouse.User, d.Source.Clickhouse.Password, d.Source.Clickhouse.DB, d.Source.Clickhouse.Table)
		}

		if d.Source.Clickhouse.Where != nil {
			sourceConfig += fmt.Sprintf(" WHERE '%s'", *d.Source.Clickhouse.Where)
		}
		return fmt.Sprintf("CLICKHOUSE(%s)", sourceConfig)
	}
	return "INVALID()"
}

func (d Dictionary) getLayoutSQLString() string {
	if d.Layout.Hashed != nil {
		return "HASHED()"
	} else if d.Layout.ComplexKeyHashed != nil {
		return "COMPLEX_KEY_HASHED()"
	} else {
		return "INVALID()"
	}
}

func (d Dictionary) getLifetimeSQLString() string {
	return fmt.Sprintf("MIN %d MAX %d", d.Lifetime.Min, d.Lifetime.Max)
}

func (d Dictionary) getStructureSQLString() string {
	attrs := []string{}
	for _, attribute := range d.Structure.Attributes {
		attr := fmt.Sprintf("%s %s", attribute.Name, attribute.Type)
		if attribute.NullValue != nil {
			attr += fmt.Sprintf(" DEFAULT '%s'", *attribute.NullValue)
		}
		attrs = append(attrs, attr)
	}
	return strings.Join(attrs, ",\n")
}
