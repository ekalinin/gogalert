package response

import "testing"
import "encoding/xml"
import "fmt"

func Test_Types(t *testing.T) {

	testXml := []byte(`<GANGLIA_XML VERSION="3.6.0" SOURCE="gmetad">
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
	var resXml GMetaResponse
	xml.Unmarshal(testXml, &resXml)
	fmt.Println(resXml)

	if resXml.String() != "Resp: Version=3.6.0 Source=gmetad" {
		t.Error("Wrong resp: ", resXml)
	}
	grid := resXml.Grids[0]
	if grid.String() != "Grid: Name=unspecified Auth=http://monitor.example.com/ganglia/ Time=1436989284" {
		t.Error("Wrong grid", grid)
	}
	cluster := grid.Clusters[0]
	if cluster.String() != "Cluster: Name=cluster-example.com Time=1436989282 Owner=example LatLong=100 Url=200" {
		t.Error("Wrong cluster", cluster)
	}
	host := cluster.Hosts[0]
	if host.String() != "Host: Name=example.com IP=127.0.0.1 Reported=1436989274 Tn=9 Tmax=20 Dmax=0 Location=unspecified GMonStarted=1411930720 Tags=unspecified" {
		t.Error("Wrong cluster", host)
	}
}
