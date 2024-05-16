package locknock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIPTablesRulesRenderer(t *testing.T) {
	params := IPTablesParams{
		HiddenPort:              222,
		Knocks:                  []uint32{11111, 22222, 33333, 44444},
		TargetReapTimeoutSecs:   33,
		InternalReapTimeoutSecs: 10,
	}
	rules := IPTablesRulesRenderer{params}
	rendered := rules.Render()
	expected := dedent(`
	iptables -N LOCKNOCK
	iptables -F LOCKNOCK
	iptables -A LOCKNOCK -p tcp -m state --state RELATED,ESTABLISHED -j ACCEPT
	iptables -A LOCKNOCK -p tcp -m tcp --dport 222 -m recent --rcheck --seconds 33 --reap --name knock3 --rsource -j ACCEPT
	iptables -A LOCKNOCK -p udp --dport 0 -m u32 --u32 "0>>22&0x3c@8=0x2b67" -m recent --set --name knock0 --rsource -j RETURN
	iptables -A LOCKNOCK -p udp --dport 0 -m recent --rcheck --seconds 10 --reap --name knock0 --rsource -m u32 --u32 "0>>22&0x3c@8=0x56ce" -m recent --set --name knock1 --rsource -j RETURN
	iptables -A LOCKNOCK -p udp --dport 0 -m recent --rcheck --seconds 10 --reap --name knock1 --rsource -m u32 --u32 "0>>22&0x3c@8=0x8235" -m recent --set --name knock2 --rsource -j RETURN
	iptables -A LOCKNOCK -p udp --dport 0 -m recent --rcheck --seconds 10 --reap --name knock2 --rsource -m u32 --u32 "0>>22&0x3c@8=0xad9c" -m recent --set --name knock3 --rsource -j RETURN
	iptables -A LOCKNOCK -p tcp -m tcp --dport 222 -j DROP
	`)
	assert.Equal(t, expected, rendered)
}
