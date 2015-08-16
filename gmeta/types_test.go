package gmeta

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
	if len(host.Metrics) != 2 {
		t.Error("Wrong number of metrics", len(host.Metrics))
	}
	metric := host.Metrics[0]
	if metric.String() != "Metric: Name=disk_free Val=1106.528 Type=double Units=GB Tn=117 Tmax=180 Dmax=0 Slope=both Source=gmond" {
		t.Error("Wrong metric", metric)
	}
	elements := metric.Extra.Elements
	if len(elements) != 3 {
		t.Error("Wrong number of extra elements", len(elements))
	}
	if elements[0].Name != "GROUP" && elements[0].Val != "disk" {
		t.Error("Wrong Element", elements[0])
	}
}
