package api

import (
	"../response"
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"strconv"
	"strings"
)

// Read xml file & parse it
func ParseFile(path string) response.GMetaResponse {

	xmlFile, err := os.Open(path)
	CheckError(err)
	defer xmlFile.Close()

	return (readAndParse(xmlFile))
}

// Read xml from socket & parse it
func ParseSocket(host string, port int) response.GMetaResponse {
	connectString := []string{host, strconv.Itoa(port)}

	conn, err := net.Dial("tcp", strings.Join(connectString, ":"))
	CheckError(err)
	defer conn.Close()

	return (readAndParse(conn))
}

func readAndParse(r io.Reader) response.GMetaResponse {
	xmlData, err := ioutil.ReadAll(r)
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
