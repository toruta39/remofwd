package remofwd

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestFetchDevices(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.nature.global/1/devices",
		httpmock.NewStringResponder(200, `[{"name":"リビング","id":"e9d9e78e-58cb-43b4-80de-86c36fad3d02","created_at":"2017-11-10T14:09:19Z","updated_at":"2020-05-01T17:25:20Z","mac_address":"a0:20:a6:1f:f6:b5","serial_number":"01000000002531","firmware_version":"Remo/1.0.62-gabbf5bd","temperature_offset":0,"humidity_offset":0,"users":[{"id":"2cb96f65-9f0b-441f-bf69-a3723a2815a1","nickname":"toruta39","superuser":true}],"newest_events":{"hu":{"val":51,"created_at":"2020-05-02T04:43:13Z"},"il":{"val":202,"created_at":"2020-05-02T05:22:37Z"},"mo":{"val":1,"created_at":"2020-05-02T05:29:25Z"},"te":{"val":26.6,"created_at":"2020-05-02T03:57:18Z"}}}]`))

	client := NewClient()
	devices, err := client.FetchDevices()

	assert.NoError(t, err)
	assert.Equal(t, 26.6, devices[0].Temperature().Value())

	info := httpmock.GetCallCountInfo()
	assert.Equal(t, 1, info["GET https://api.nature.global/1/devices"])
}
