#!/bin/sh

sudo ipset restore <<EOF
  -! create china hash:net family inet hashsize 2048 maxelem 65536
  {{range $i, $ip := .Ips}}add china {{$ip.Ip}}/{{$ip.Cidr}}
  {{end}}
EOF
