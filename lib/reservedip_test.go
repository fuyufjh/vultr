package lib

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ReservedIP_ListReservedIP_Fail(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusNotAcceptable, ``)
	defer server.Close()

	_, err := client.ListReservedIP()
	if err == nil {
		t.Error(err)
	}
}

func Test_ReservedIP_ListReservedIP_Ok_Empty(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusNotAcceptable, `{}`)
	defer server.Close()

	list, err := client.ListReservedIP()
	if err == nil {
		t.Error(err)
	}
	assert.Equal(t, len(list), 0)
}

func Test_ReservedIP_ListReservedIP_Ok(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK,
		`{
      "4":{"SUBID":4,"DCID":5,"ip_type":"v4","subnet":"subnet1",
           "subnet_size":8,"label":"label","attached_SUBID":false},
      "9":{"SUBID":9,"DCID":5,"ip_type":"v6","subnet":"subnet2",
           "subnet_size":16,"label":"label","attached_SUBID":123}
      }`)
	defer server.Close()

	ips, err := client.ListReservedIP()
	if err != nil {
		t.Error(err)
	}
	if assert.NotNil(t, ips) {
		assert.Equal(t, 2, len(ips))
		// keys could be in random order
		for _, ip := range ips {
			switch ip.ID {
			case "4":
				assert.Equal(t, ip.RegionID, 5)
				assert.Equal(t, ip.IPType, "v4")
				assert.Equal(t, ip.Subnet, "subnet1")
				assert.Equal(t, ip.SubnetSize, 8)
				assert.Equal(t, ip.Label, "label")
				assert.Equal(t, ip.AttachedTo, "")
			case "9":
				assert.Equal(t, ip.RegionID, 5)
				assert.Equal(t, ip.IPType, "v6")
				assert.Equal(t, ip.Subnet, "subnet2")
				assert.Equal(t, ip.SubnetSize, 16)
				assert.Equal(t, ip.Label, "label")
				assert.Equal(t, ip.AttachedTo, "123")
			default:
				t.Error("Unknown ReservedIP")
			}
		}
	}
}

func Test_ReservedIP_CreateReservedIP_Fail(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusNotAcceptable, ``)
	defer server.Close()

	_, err := client.CreateReservedIP(1, "ip", "")
	if err == nil {
		t.Error(err)
	}
}

func Test_ReservedIP_CreateReservedIP_OK(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, `{"SUBID":4711}`)
	defer server.Close()

	id, err := client.CreateReservedIP(1, "ip", "")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, id, "4711")
}

func Test_ReservedIP_DestroyReservedIP_Fail(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusNotAcceptable, ``)
	defer server.Close()

	err := client.DestroyReservedIP("subid")
	if err == nil {
		t.Error(err)
	}
}

func Test_ReservedIP_DestroyReservedIP_OK(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, ``)
	defer server.Close()

	err := client.DestroyReservedIP("subid")
	if err != nil {
		t.Error(err)
	}
}

func Test_ReservedIP_AttachReservedIP_Fail(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusNotAcceptable, ``)
	defer server.Close()

	err := client.AttachReservedIP("ip", "subid")
	if err == nil {
		t.Error(err)
	}
}

func Test_ReservedIP_AttachReservedIP_OK(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, ``)
	defer server.Close()

	err := client.AttachReservedIP("subid", "ip")
	if err != nil {
		t.Error(err)
	}
}

func Test_ReservedIP_ConvertReservedIP_Fail(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusNotAcceptable, ``)
	defer server.Close()

	_, err := client.ConvertReservedIP("subid", "ip")
	if err == nil {
		t.Error(err)
	}
}

func Test_ReservedIP_ConvertReservedIP_OK(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, `{"SUBID":4711}`)
	defer server.Close()

	id, err := client.ConvertReservedIP("subid", "ip")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, id, "4711")
}

func Test_ReservedIP_DetachReservedIP_Fail(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusNotAcceptable, ``)
	defer server.Close()

	err := client.DetachReservedIP("subid", "ip")
	if err == nil {
		t.Error(err)
	}
}

func Test_ReservedIP_DetachReservedIP_OK(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, ``)
	defer server.Close()

	err := client.DetachReservedIP("subid", "ip")
	if err != nil {
		t.Error(err)
	}
}
