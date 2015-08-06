package response

import (
	"encoding/xml"
	"testing"
)

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

	var resXml GMetaResponse
	xml.Unmarshal(testXml, &resXml)

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
