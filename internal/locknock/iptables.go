package locknock

import (
	"html/template"
	"strings"
)

type IPTablesParams struct {
	// TargetPort is a port to which we hide
	TargetPort int
	// KnockPorts are ports needed to be knocked to access TargetPort
	KnockPorts []int
	// TargetReapTimeoutSecs is how long recent module of iptables
	// keeps the temporary rule to the target port after knocking was complete
	TargetReapTimeoutSecs int
	// InternalReapTimeoutSecs is how long recent module of iptables in internal knock rules
	// essentialy its now fast the knock packets should be sent to be registered
	InternalReapTimeoutSecs int
}

type IPTablesRulesRenderer struct {
	Params IPTablesParams
}

func (m *IPTablesRulesRenderer) Render() string {
	tmpl, err := template.New("locknock").Funcs(*templateHelpers).Parse(dedent(`
	{{- $length := len .KnockPorts -}}
	{{- $internalReapTimeoutSecs := .InternalReapTimeoutSecs -}}
	iptables -N LOCKNOCK
	iptables -F LOCKNOCK
	iptables -A LOCKNOCK -p tcp -m state --state RELATED,ESTABLISHED -j ACCEPT
	iptables -A LOCKNOCK -m recent --rcheck --seconds {{.TargetReapTimeoutSecs}} --reap --name knock{{sum $length -1}} --rsource -j ACCEPT
	{{- range $index, $knockPort := .KnockPorts }}
	{{- if eq $index 0 }}
	iptables -A LOCKNOCK -p udp -m udp --dport {{$knockPort}} -m recent --set --name knock{{$index}} --rsource -j DROP
	{{- else }}
	iptables -A LOCKNOCK -p udp -m recent --rcheck --seconds {{$internalReapTimeoutSecs}} --reap --name knock{{sum $index -1}} --rsource -m udp --dport {{$knockPort}} -m recent --set --name knock{{$index}} --rsource -j DROP
	{{- end }}
	{{- end }}
	iptables -A LOCKNOCK -p tcp -m tcp --dport {{.TargetPort}} -j DROP
	`))
	if err != nil {
		panic(err)
	}
	var sb strings.Builder
	err = tmpl.Execute(&sb, m.Params)
	if err != nil {
		panic(err)
	}
	return sb.String()
}
