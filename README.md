# mitail(みている)
You should not use the `watch` command for monitoring.
The `watch` command can see only now.
`mitail` is monitroing tool for me.
`mitail` is very simple commands, `mitail` has feature that is only print execution output to stdio.

```console
▶ ./mitail -t  'dig @8.8.8.8 google.co.jp | grep -e "Query time" -e "status: " | tr -d "\n" | cut -d" " -f5,6,9,10,11'
2021-07-28T22:40:10+09:00	status: NOERROR, Query time: 13
2021-07-28T22:40:11+09:00	status: NOERROR, Query time: 9
2021-07-28T22:40:12+09:00	status: NOERROR, Query time: 9
2021-07-28T22:40:13+09:00	status: NOERROR, Query time: 11
2021-07-28T22:40:14+09:00	status: NOERROR, Query time: 9
2021-07-28T22:40:15+09:00	status: NOERROR, Query time: 9
2021-07-28T22:40:16+09:00	status: NOERROR, Query time: 9
2021-07-28T22:40:17+09:00	status: NOERROR, Query time: 11
```
