package main

import (
	"fmt"
	"strconv"

	zsend "github.com/blacked/go-zabbix"
)

func makePrefix(prefix, key string) string {
	return fmt.Sprintf(
		"%s.%s", prefix, key,
	)
}

func createMetrics(
	hostname string,
	stats *Stats,
	metrics []*zsend.Metric,
	prefix string,
) []*zsend.Metric {

	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"status",
			),
			stats.Status,
		),
	)
	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"recv",
			),
			strconv.Itoa(int(stats.SMS.Recv)),
		),
	)
	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"recv.queued",
			),
			strconv.Itoa(int(stats.SMS.RecvQueued)),
		),
	)
	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"sent",
			),
			strconv.Itoa(int(stats.SMS.Sent)),
		),
	)
	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"sent.queued",
			),
			strconv.Itoa(int(stats.SMS.SentQueued)),
		),
	)
	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"store",
			),
			strconv.Itoa(int(stats.SMS.StoreSize)),
		),
	)

	for _, provider := range stats.SMSC {

		metrics = append(
			metrics,
			zsend.NewMetric(
				hostname,
				makePrefix(
					prefix,
					fmt.Sprintf("status.provider.[%s]", provider.ID),
				),
				provider.Status,
			),
		)
		metrics = append(
			metrics,
			zsend.NewMetric(
				hostname,
				makePrefix(
					prefix,
					fmt.Sprintf("uptime.provider.[%s]", provider.ID),
				),
				fmt.Sprintf("%f", provider.Uptime),
			),
		)
		metrics = append(
			metrics,
			zsend.NewMetric(
				hostname,
				makePrefix(
					prefix,
					fmt.Sprintf("recv.provider.[%s]", provider.ID),
				),
				strconv.Itoa(int(provider.Recv)),
			),
		)
		metrics = append(
			metrics,
			zsend.NewMetric(
				hostname,
				makePrefix(
					prefix,
					fmt.Sprintf("sent.provider.[%s]", provider.ID),
				),
				strconv.Itoa(int(provider.Sent)),
			),
		)
		metrics = append(
			metrics,
			zsend.NewMetric(
				hostname,
				makePrefix(
					prefix,
					fmt.Sprintf("failed.provider.[%s]", provider.ID),
				),
				strconv.Itoa(int(provider.Failed)),
			),
		)
		metrics = append(
			metrics,
			zsend.NewMetric(
				hostname,
				makePrefix(
					prefix,
					fmt.Sprintf("queued.provider.[%s]", provider.ID),
				),
				strconv.Itoa(int(provider.Queued)),
			),
		)

	}

	return metrics
}
