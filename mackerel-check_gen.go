// Code generated by "tools/gen_mackerel_check.pl"; DO NOT EDIT
package main

import (
	"fmt"

	"github.com/mackerelio/go-check-plugins/check-aws-sqs-queue-size/lib"
	"github.com/mackerelio/go-check-plugins/check-cert-file/lib"
	"github.com/mackerelio/go-check-plugins/check-elasticsearch/lib"
	"github.com/mackerelio/go-check-plugins/check-file-age/lib"
	"github.com/mackerelio/go-check-plugins/check-file-size/lib"
	"github.com/mackerelio/go-check-plugins/check-http/lib"
	"github.com/mackerelio/go-check-plugins/check-jmx-jolokia/lib"
	"github.com/mackerelio/go-check-plugins/check-load/lib"
	"github.com/mackerelio/go-check-plugins/check-log/lib"
	"github.com/mackerelio/go-check-plugins/check-mailq/lib"
	"github.com/mackerelio/go-check-plugins/check-masterha/lib"
	"github.com/mackerelio/go-check-plugins/check-memcached/lib"
	"github.com/mackerelio/go-check-plugins/check-mysql/lib"
	"github.com/mackerelio/go-check-plugins/check-ntpoffset/lib"
	"github.com/mackerelio/go-check-plugins/check-postgresql/lib"
	"github.com/mackerelio/go-check-plugins/check-procs/lib"
	"github.com/mackerelio/go-check-plugins/check-redis/lib"
	"github.com/mackerelio/go-check-plugins/check-solr/lib"
	"github.com/mackerelio/go-check-plugins/check-ssh/lib"
	"github.com/mackerelio/go-check-plugins/check-tcp/lib"
	"github.com/mackerelio/go-check-plugins/check-uptime/lib"
)

func runPlugin(plug string) error {
	switch plug {
	case "aws-sqs-queue-size":
		checkawssqsqueuesize.Do()
	case "cert-file":
		checkcertfile.Do()
	case "elasticsearch":
		checkelasticsearch.Do()
	case "file-age":
		checkfileage.Do()
	case "file-size":
		checkfilesize.Do()
	case "http":
		checkhttp.Do()
	case "jmx-jolokia":
		checkjmxjolokia.Do()
	case "load":
		checkload.Do()
	case "log":
		checklog.Do()
	case "mailq":
		checkmailq.Do()
	case "masterha":
		checkmasterha.Do()
	case "memcached":
		checkmemcached.Do()
	case "mysql":
		checkmysql.Do()
	case "ntpoffset":
		checkntpoffset.Do()
	case "postgresql":
		checkpostgresql.Do()
	case "procs":
		checkprocs.Do()
	case "redis":
		checkredis.Do()
	case "solr":
		checksolr.Do()
	case "ssh":
		checkssh.Do()
	case "tcp":
		checktcp.Do()
	case "uptime":
		checkuptime.Do()
	default:
		return fmt.Errorf("unknown plugin: %q", plug)
	}
	return nil
}

var plugins = []string{
	"aws-sqs-queue-size",
	"cert-file",
	"elasticsearch",
	"file-age",
	"file-size",
	"http",
	"jmx-jolokia",
	"load",
	"log",
	"mailq",
	"masterha",
	"memcached",
	"mysql",
	"ntpoffset",
	"postgresql",
	"procs",
	"redis",
	"solr",
	"ssh",
	"tcp",
	"uptime",
}
