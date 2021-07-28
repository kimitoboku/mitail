# mitail(みている)
You should not use the `watch` command for monitoring.
The `watch` command can see only now.
`mitail` is monitroing tool for me.
`mitail` is very simple commands, `mitail` has feature that is only print execution output to stdio.

```console
./mitail -t  'dig @8.8.8.8 google.co.jp | grep -e "Query time" -e "status: " | tr -d "\n" | cut -d" " -f5,6,9,10,11'
2021-07-28 11:50:35.578988134 +0900 JST m=+0.027181683  status: NOERROR, Query time: 7
2021-07-28 11:50:36.607576553 +0900 JST m=+1.055770090  status: NOERROR, Query time: 7
2021-07-28 11:50:37.634456445 +0900 JST m=+2.082650014  status: NOERROR, Query time: 7
2021-07-28 11:50:38.666121483 +0900 JST m=+3.114315012  status: NOERROR, Query time: 7
2021-07-28 11:50:39.692187767 +0900 JST m=+4.140381298  status: NOERROR, Query time: 7
```