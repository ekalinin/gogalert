package api

import (
	"../response"
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strconv"
	"strings"
)

type DataSource struct {
	Path string
	Host string
	Port int
}

func (this *DataSource) Read() ([]byte, error) {

	if this.Path != "" {
		xmlFile, err := os.Open(this.Path)
		CheckError(err)
		defer xmlFile.Close()
		return ioutil.ReadAll(xmlFile)
	}

	if this.Host != "" {
		connectString := []string{this.Host, strconv.Itoa(this.Port)}
		conn, err := net.Dial("tcp", strings.Join(connectString, ":"))
		CheckError(err)
		defer conn.Close()
		return ioutil.ReadAll(conn)
	}

	return []byte{}, errors.New("Not a file nor a socket!")
}

// Parses XML fetched from source
func Parse(source *DataSource) response.GMetaResponse {
	xmlData, err := source.Read()
	CheckError(err)

	return (ParseXML(xmlData))
}

// Parse input xml text
func ParseXML(xmlbody []byte) response.GMetaResponse {
	var gResp response.GMetaResponse
	var xmlClear []byte

	// Exclude part of xml
	ganglia_xml_start_idx := bytes.Index(xmlbody, []byte("<GANGLIA_XML"))
	if ganglia_xml_start_idx == -1 {
		xmlClear = xmlbody
	} else {
		xmlClear = xmlbody[ganglia_xml_start_idx:]
	}
	xml.Unmarshal(xmlClear, &gResp)
	return gResp
}

// Check error, panic if it is
func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
