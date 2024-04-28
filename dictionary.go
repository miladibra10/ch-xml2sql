package main

import "encoding/xml"

type Dictionaries struct {
	XMLName      xml.Name     `xml:"dictionaries"`
	Dictionaries []Dictionary `xml:"dictionary"`
}

type Dictionary struct {
	XMLName   xml.Name  `xml:"dictionary"`
	Name      string    `xml:"name"`
	Source    Source    `xml:"source"`
	Lifetime  Lifetime  `xml:"lifetime"`
	Layout    Layout    `xml:"layout"`
	Structure Structure `xml:"structure"`
}

type Source struct {
	XMLName    xml.Name `xml:"source"`
	Clickhouse *struct {
		XMLName  xml.Name `xml:"clickhouse"`
		Host     string   `xml:"host"`
		Port     int      `xml:"port"`
		User     string   `xml:"user"`
		Password string   `xml:"password"`
		DB       string   `xml:"db"`
		Table    string   `xml:"table"`
		Where    *string  `xml:"where"`
	} `xml:"clickhouse"`
	HTTP *struct {
		XMLName xml.Name `xml:"http"`
		URL     string   `xml:"url"`
		Format  string   `xml:"format"`
	} `xml:"http"`
	File *struct {
		XMLName xml.Name `xml:"file"`
		Path    string   `xml:"path"`
		Format  string   `xml:"format"`
	} `xml:"file"`
}

type Lifetime struct {
	XMLName xml.Name `xml:"lifetime"`
	Min     int      `xml:"min"`
	Max     int      `xml:"max"`
}

type Layout struct {
	Hashed *struct {
		XMLName xml.Name `xml:"hashed"`
	} `xml:"hashed"`

	ComplexKeyHashed *struct {
		XMLName xml.Name `xml:"complex_key_hashed"`
	} `xml:"complex_key_hashed"`
}

type Structure struct {
	ID *struct {
		XMLName xml.Name `xml:"id"`
		Name    string   `xml:"name"`
	} `xml:"id"`
	Key *struct {
		XMLName    xml.Name    `xml:"key"`
		Attributes []Attribute `xml:"attribute"`
	} `xml:"key"`
	Attributes []Attribute `xml:"attribute"`
}

type Attribute struct {
	XMLName   xml.Name `xml:"attribute"`
	Name      string   `xml:"name"`
	Type      string   `xml:"type"`
	NullValue *string  `xml:"null_value"`
}
