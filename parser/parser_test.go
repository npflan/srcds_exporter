package parser

import (
	"reflect"
	"testing"

	"github.com/galexrt/srcds_exporter/models"
)

var parserTests = []struct {
	name     string
	request  string
	expected *models.Status
}{
	{
		"gmod",
		`hostname: Example server
version : 16.12.01/24 6729 secure
udp/ip  : 1.1.1.1:27015  (public ip: 1.1.1.1)
map     : rp_retribution_v2 at: 0 x, 0 y, 0 z
players : 1 (64 max)

# userid name                uniqueid            connected ping loss state  adr
#    218 "TestUser1"      STEAM_0:0:1015738 07:36       65    0 active 10.10.220.12:27005`,
		&models.Status{
			Hostname: "Example server",
			Version:  "16.12.01/24 6729 secure",
			Map:      "rp_retribution_v2",
			PlayerCount: *&models.PlayerCount{
				Max:     64,
				Current: 1,
			},
			Players: map[int]models.Player{
				218: *&models.Player{
					Username: "TestUser1",
					SteamID:  "STEAM_0:0:1015738",
					Ping:     65,
					Loss:     0,
					State:    "active",
					IP:       "10.10.220.12",
					ConnPort: 27005,
				},
			},
		},
	},
}

func TestParser(t *testing.T) {
	for _, tt := range parserTests {
		actual := Parse(tt.request)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("Parse(%s): expected %v, actual %v", tt.name, tt.expected, actual)
		}
	}
}