package api

import (
	"encoding/xml"
	"github.com/ekalinin/gogalert/gmeta"
	"testing"
)

var (
	testXml = []byte(`<?xml version="1.0" encoding="ISO-8859-1" standalone="yes"?>
		<!DOCTYPE GANGLIA_XML [
		<!ELEMENT GANGLIA_XML (GRID|CLUSTER|HOST)*>
		<!ATTLIST GANGLIA_XML VERSION CDATA #REQUIRED>
		<!ATTLIST GANGLIA_XML SOURCE CDATA #REQUIRED>
		<!ELEMENT GRID (CLUSTER | GRID | HOSTS | METRICS)*>
		<!ATTLIST GRID NAME CDATA #REQUIRED>
		<!ATTLIST GRID AUTHORITY CDATA #REQUIRED>
		<!ATTLIST GRID LOCALTIME CDATA #IMPLIED>
		<!ELEMENT CLUSTER (HOST | HOSTS | METRICS)*>
		<!ATTLIST CLUSTER NAME CDATA #REQUIRED>
		<!ATTLIST CLUSTER OWNER CDATA #IMPLIED>
		<!ATTLIST CLUSTER LATLONG CDATA #IMPLIED>
		<!ATTLIST CLUSTER URL CDATA #IMPLIED>
		<!ATTLIST CLUSTER LOCALTIME CDATA #REQUIRED>
		<!ELEMENT HOST (METRIC)*>
		<!ATTLIST HOST NAME CDATA #REQUIRED>
		<!ATTLIST HOST IP CDATA #REQUIRED>
		<!ATTLIST HOST LOCATION CDATA #IMPLIED>
		<!ATTLIST HOST TAGS CDATA #IMPLIED>
		<!ATTLIST HOST REPORTED CDATA #REQUIRED>
		<!ATTLIST HOST TN CDATA #IMPLIED>
		<!ATTLIST HOST TMAX CDATA #IMPLIED>
		<!ATTLIST HOST DMAX CDATA #IMPLIED>
		<!ATTLIST HOST GMOND_STARTED CDATA #IMPLIED>
		<!ELEMENT METRIC (EXTRA_DATA*)>
		<!ATTLIST METRIC NAME CDATA #REQUIRED>
		<!ATTLIST METRIC VAL CDATA #REQUIRED>
		<!ATTLIST METRIC TYPE (string | int8 | uint8 | int16 | uint16 | int32 | uint32 | float | double | timestamp) #REQUIRED>
		<!ATTLIST METRIC UNITS CDATA #IMPLIED>
		<!ATTLIST METRIC TN CDATA #IMPLIED>
		<!ATTLIST METRIC TMAX CDATA #IMPLIED>
		<!ATTLIST METRIC DMAX CDATA #IMPLIED>
		<!ATTLIST METRIC SLOPE (zero | positive | negative | both | unspecified) #IMPLIED>
		<!ATTLIST METRIC SOURCE (gmond) 'gmond'>
		<!ELEMENT EXTRA_DATA (EXTRA_ELEMENT*)>
		<!ELEMENT EXTRA_ELEMENT EMPTY>
		<!ATTLIST EXTRA_ELEMENT NAME CDATA #REQUIRED>
		<!ATTLIST EXTRA_ELEMENT VAL CDATA #REQUIRED>
		<!ELEMENT HOSTS EMPTY>
		<!ATTLIST HOSTS UP CDATA #REQUIRED>
		<!ATTLIST HOSTS DOWN CDATA #REQUIRED>
		<!ATTLIST HOSTS SOURCE (gmond | gmetad) #REQUIRED>
		<!ELEMENT METRICS (EXTRA_DATA*)>
		<!ATTLIST METRICS NAME CDATA #REQUIRED>
		<!ATTLIST METRICS SUM CDATA #REQUIRED>
		<!ATTLIST METRICS NUM CDATA #REQUIRED>
		<!ATTLIST METRICS TYPE (string | int8 | uint8 | int16 | uint16 | int32 | uint32 | float | double | timestamp) #REQUIRED>
		<!ATTLIST METRICS UNITS CDATA #IMPLIED>
		<!ATTLIST METRICS SLOPE (zero | positive | negative | both | unspecified) #IMPLIED>
		<!ATTLIST METRICS SOURCE (gmond) 'gmond'>
		]>
		<GANGLIA_XML VERSION="3.6.0" SOURCE="gmetad">
		<GRID NAME="unspecified" AUTHORITY="http://monitor.example.com/ganglia/" LOCALTIME="1436989284">
		<CLUSTER NAME="cluster-example.com" LOCALTIME="1436989282" OWNER="example" LATLONG="100" URL="200">
		<HOST NAME="example.com" IP="127.0.0.1" REPORTED="1436989274" TN="9" TMAX="20" DMAX="0" LOCATION="unspecified" GMOND_STARTED="1411930720" TAGS="unspecified">
		<METRIC NAME="disk_free" VAL="1106.528" TYPE="double" UNITS="GB" TN="117" TMAX="180" DMAX="0" SLOPE="both" SOURCE="gmond">
		<EXTRA_DATA>
		<EXTRA_ELEMENT NAME="GROUP" VAL="disk"/>
		<EXTRA_ELEMENT NAME="DESC" VAL="Total free disk space"/>
		<EXTRA_ELEMENT NAME="TITLE" VAL="Disk Space Available"/>
		</EXTRA_DATA>
		</METRIC>
		<METRIC NAME="bytes_out" VAL="7167.18" TYPE="float" UNITS="bytes/sec" TN="277" TMAX="300" DMAX="0" SLOPE="both" SOURCE="gmond">
		<EXTRA_DATA>
		<EXTRA_ELEMENT NAME="GROUP" VAL="network"/>
		<EXTRA_ELEMENT NAME="DESC" VAL="Number of bytes out per second"/>
		<EXTRA_ELEMENT NAME="TITLE" VAL="Bytes Sent"/>
		</EXTRA_DATA>
		</METRIC>
		<METRIC NAME="disk_free" VAL="945.966" TYPE="double" UNITS="GB" TN="80" TMAX="180" DMAX="0" SLOPE="both" SOURCE="gmond">
		<EXTRA_DATA>
		<EXTRA_ELEMENT NAME="GROUP" VAL="disk"/>
		<EXTRA_ELEMENT NAME="DESC" VAL="Total free disk space"/>
		<EXTRA_ELEMENT NAME="TITLE" VAL="Disk Space Available"/>
		</EXTRA_DATA>
		</METRIC>
		<METRIC NAME="swap_total" VAL="2102460" TYPE="float" UNITS="KB" TN="21" TMAX="1200" DMAX="0" SLOPE="zero" SOURCE="gmond">
		<EXTRA_DATA>
		<EXTRA_ELEMENT NAME="GROUP" VAL="memory"/>
		<EXTRA_ELEMENT NAME="DESC" VAL="Total amount of swap space displayed in KBs"/>
		<EXTRA_ELEMENT NAME="TITLE" VAL="Swap Space Total"/>
		</EXTRA_DATA>
		</METRIC>
		<METRIC NAME="part_max_used" VAL="60.5" TYPE="float" UNITS="%" TN="80" TMAX="180" DMAX="0" SLOPE="both" SOURCE="gmond">
		<EXTRA_DATA>
		<EXTRA_ELEMENT NAME="GROUP" VAL="disk"/>
		<EXTRA_ELEMENT NAME="DESC" VAL="Maximum percent used for all partitions"/>
		<EXTRA_ELEMENT NAME="TITLE" VAL="Maximum Disk Space Used"/>
		</EXTRA_DATA>
		</METRIC>
		</HOST>
		</CLUSTER>
		</GRID>
		</GANGLIA_XML>`)
)

func Test_ParseXML(t *testing.T) {
	resXml := &GMetaWrapper{gmeta.Parse(testXml)}

	if resXml.Version != "3.6.0" {
		t.Error("Wrong Version")
	}
	if resXml.Source != "gmetad" {
		t.Error("Wrong source")
	}
	if len(resXml.Grids) != 1 {
		t.Error("Wrong number of grids in the cluster")
	}
	grid := resXml.Grids[0]
	if grid.Name != "unspecified" {
		t.Error("Wrong name of the grid")
	}
	if grid.Authority != "http://monitor.example.com/ganglia/" {
		t.Error("Wrong Authority of the grid")
	}
}

func Test_Parse(t *testing.T) {

}

func Test_ValidateCondition(t *testing.T) {
	if !ValidateCondition(5.0, "eq", 5.0) {
		t.Error("5 == 5")
	}
	if ValidateCondition(5.0, "eq", 5.1) {
		t.Error("5 != 5.1")
	}
	if !ValidateCondition(5.0, "gt", 4.0) {
		t.Error("5 > 4")
	}
	if ValidateCondition(5.0, "gt", 5.0) {
		t.Error("5 !> 5")
	}
	if !ValidateCondition(5.0, "ge", 5.0) {
		t.Error("5 >= 5")
	}
	if !ValidateCondition(5.1, "ge", 5.0) {
		t.Error("5.1 >= 5")
	}
	if ValidateCondition(5.0, "ge", 5.1) {
		t.Error("5 !>= 5.1")
	}
	if !ValidateCondition(5.0, "lt", 5.1) {
		t.Error("5.0 < 5.1")
	}
	if ValidateCondition(5.1, "lt", 5.0) {
		t.Error("5.1 !< 5")
	}
	if !ValidateCondition(5.0, "le", 5.1) {
		t.Error("5 <= 5.1")
	}
	if !ValidateCondition(5.0, "le", 5.0) {
		t.Error("5 <= 5")
	}
	if ValidateCondition(5.1, "le", 5.0) {
		t.Error("5.1 !<= 5")
	}
}

func Test_GMetaResponseFind(t *testing.T) {
	testXml := []byte(`
		<GANGLIA_XML VERSION="3.6.0" SOURCE="gmetad">
			<GRID NAME="unspecified" AUTHORITY="http://monitor.example.com/ganglia/" LOCALTIME="1436989284">
				<CLUSTER NAME="cluster-example.com" LOCALTIME="1436989282" OWNER="example" LATLONG="100" URL="200">
					<HOST NAME="example.com" IP="127.0.0.1" REPORTED="1436989274" TN="9" TMAX="20" DMAX="0" LOCATION="unspecified" GMOND_STARTED="1411930720" TAGS="unspecified">
						<METRIC NAME="disk_free" VAL="1106.528" TYPE="double" UNITS="GB" TN="117" TMAX="180" DMAX="0" SLOPE="both" SOURCE="gmond">
							<EXTRA_DATA>
								<EXTRA_ELEMENT NAME="GROUP" VAL="disk"/>
								<EXTRA_ELEMENT NAME="DESC" VAL="Total free disk space"/>
								<EXTRA_ELEMENT NAME="TITLE" VAL="Disk Space Available"/>
							</EXTRA_DATA>
						</METRIC>
						<METRIC NAME="bytes_out" VAL="7167.18" TYPE="float" UNITS="bytes/sec" TN="277" TMAX="300" DMAX="0" SLOPE="both" SOURCE="gmond">
							<EXTRA_DATA>
								<EXTRA_ELEMENT NAME="GROUP" VAL="network"/>
								<EXTRA_ELEMENT NAME="DESC" VAL="Number of bytes out per second"/>
								<EXTRA_ELEMENT NAME="TITLE" VAL="Bytes Sent"/>
							</EXTRA_DATA>
						</METRIC>
						<METRIC NAME="swap_total" VAL="2102460" TYPE="float" UNITS="KB" TN="21" TMAX="1200" DMAX="0" SLOPE="zero" SOURCE="gmond">
							<EXTRA_DATA>
								<EXTRA_ELEMENT NAME="GROUP" VAL="memory"/>
								<EXTRA_ELEMENT NAME="DESC" VAL="Total amount of swap space displayed in KBs"/>
								<EXTRA_ELEMENT NAME="TITLE" VAL="Swap Space Total"/>
							</EXTRA_DATA>
						</METRIC>
						<METRIC NAME="part_max_used" VAL="60.5" TYPE="float" UNITS="%" TN="80" TMAX="180" DMAX="0" SLOPE="both" SOURCE="gmond">
							<EXTRA_DATA>
								<EXTRA_ELEMENT NAME="GROUP" VAL="disk"/>
								<EXTRA_ELEMENT NAME="DESC" VAL="Maximum percent used for all partitions"/>
								<EXTRA_ELEMENT NAME="TITLE" VAL="Maximum Disk Space Used"/>
							</EXTRA_DATA>
						</METRIC>
					</HOST>
				</CLUSTER>
			</GRID>
		</GANGLIA_XML>`)

	var tmpXml gmeta.GMetaResponse
	xml.Unmarshal(testXml, &tmpXml)
	resXml := &GMetaWrapper{tmpXml}

	if resp := len(resXml.Find(&MetricFilter{"", "", "", "", 0})); resp != 4 {
		t.Error("Wrong response length: ", resp)
	}
	if resp := len(resXml.Find(&MetricFilter{"disk_free", "", "", "", 0})); resp != 1 {
		t.Error("Wrong response length: ", resp)
	}
	if resp := len(resXml.Find(&MetricFilter{"disk_free!!", "", "", "", 0})); resp != 0 {
		t.Error("Wrong response length: ", resp)
	}
	if resp := len(resXml.Find(&MetricFilter{"", "example.com", "", "", 0})); resp != 4 {
		t.Error("Wrong response length: ", resp)
	}
	if resp := len(resXml.Find(&MetricFilter{"", "example.com!", "", "", 0})); resp != 0 {
		t.Error("Wrong response length: ", resp)
	}
	if resp := len(resXml.Find(&MetricFilter{"disk_free", "example.com", "", "", 0})); resp != 1 {
		t.Error("Wrong response length: ", resp)
	}
	if resp := len(resXml.Find(&MetricFilter{"", "", "cluster-example.com", "", 0})); resp != 4 {
		t.Error("Wrong response length: ", resp)
	}
	if resp := len(resXml.Find(&MetricFilter{"", "", "cluster-example.com!", "", 0})); resp != 0 {
		t.Error("Wrong response length: ", resp)
	}
	if resp := len(resXml.Find(&MetricFilter{"swap_total", "", "cluster-example.com", "", 0})); resp != 1 {
		t.Error("Wrong response length: ", resp)
	}
	if resp := len(resXml.Find(&MetricFilter{"", "", "", "gt", 100})); resp != 3 {
		t.Error("Wrong response length: ", resp)
	}
	if resp := len(resXml.Find(&MetricFilter{"part_max_used", "", "", "gt", 80})); resp != 0 {
		t.Error("Wrong response length: ", resp)
	}
}
