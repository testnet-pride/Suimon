package checker

import (
	"bufio"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ybbus/jsonrpc/v3"

	"github.com/bartosian/sui_helpers/suimon/cmd/checker/enums"
)

func (host *Host) GetMetrics(httpClient *http.Client) {
	metricsURL := host.getUrl(requestTypeMetrics, false)

	result, err := httpClient.Get(metricsURL)
	if err != nil {
		return
	}

	defer result.Body.Close()

	reader := bufio.NewReader(result.Body)
	for {
		line, err := reader.ReadString('\n')
		if len(line) == 0 && err != nil {
			break
		}

		if strings.HasPrefix(line, "#") {
			continue
		}

		metric := strings.Split(line, " ")
		if len(metric) != 2 {
			continue
		}

		key, value := strings.TrimSpace(metric[0]), strings.TrimSpace(metric[1])

		metricName, err := enums.MetricTypeFromString(key)
		if err != nil {
			continue
		}

		if metricName == enums.MetricTypeUptime {
			versionMetric := versionRegex.FindStringSubmatch(key)
			version := strings.Split(versionMetric[1], "=")

			uptimeSeconds, err := strconv.Atoi(value)
			if err != nil {
				continue
			}

			value = fmt.Sprintf("%.1f days", float64(uptimeSeconds)/(60*60*24))

			versionInfo := strings.Split(version[1], "-")
			if len(versionInfo) != 2 {
				continue
			}

			host.Metrics.SetValue(enums.MetricTypeVersion, versionInfo[0])
			host.Metrics.SetValue(enums.MetricTypeCommit, versionInfo[1])
		}

		host.Metrics.SetValue(metricName, value)
	}
}

func (host *Host) GetTotalTransactionNumber() {
	if result := getRequestAttempt(host.rpcHttpClient, enums.RPCMethodGetTotalTransactionNumber); result != nil {
		host.Metrics.SetValue(enums.MetricTypeTotalTransactionsNumber, *result)

		return
	}

	if result := getRequestAttempt(host.rpcHttpsClient, enums.RPCMethodGetLatestCheckpointSequenceNumber); result != nil {
		host.Metrics.SetValue(enums.MetricTypeTotalTransactionsNumber, *result)
	}
}

func (host *Host) GetLatestCheckpoint() {
	if result := getRequestAttempt(host.rpcHttpClient, enums.RPCMethodGetLatestCheckpointSequenceNumber); result != nil {
		host.Metrics.SetValue(enums.MetricTypeLatestCheckpoint, *result)

		return
	}

	if result := getRequestAttempt(host.rpcHttpsClient, enums.RPCMethodGetLatestCheckpointSequenceNumber); result != nil {
		host.Metrics.SetValue(enums.MetricTypeLatestCheckpoint, *result)
	}
}

func getRequestAttempt(client jsonrpc.RPCClient, method enums.RPCMethod) *string {
	if result := getFromRPC(client, method); result != nil {
		result := strconv.Itoa(*result)

		return &result
	}

	return nil
}

func getFromRPC(rpcClient jsonrpc.RPCClient, method enums.RPCMethod) *int {
	respChan := make(chan *int)
	timeout := time.After(rpcClientTimeout)

	go func() {
		var response *int

		if err := rpcClient.CallFor(context.Background(), &response, method.String()); err != nil {
			return
		}

		respChan <- response
	}()

	select {
	case response := <-respChan:
		return response
	case <-timeout:
		return nil
	}
}

func (host *Host) getUrl(request requestType, secure bool) string {
	var (
		address  string
		port     string
		protocol = "http"
	)

	if secure {
		protocol = protocol + "s"
	}

	hostPort := host.HostPort

	if hostPort.IP != nil {
		address = *hostPort.IP
	} else if hostPort.Host != nil {
		address = *hostPort.GetHostWithPath()

		if hostPort.Path != nil {
			return fmt.Sprintf("%s://%s", protocol, address)
		}
	}

	switch request {
	case requestTypeRPC:
		port = rpcPortDefault
		if portRPC, ok := host.Ports[enums.PortTypeRPC]; ok {
			port = portRPC
		}

		return fmt.Sprintf("%s://%s:%s", protocol, address, port)
	case requestTypeMetrics:
		fallthrough
	default:
		port = metricsPortDefault
		if portMetrics, ok := host.Ports[enums.PortTypeMetrics]; ok {
			port = portMetrics
		}

		return fmt.Sprintf("%s://%s:%s/metrics", protocol, address, port)
	}
}