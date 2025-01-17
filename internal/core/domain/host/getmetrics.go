package host

import (
	"errors"
	"fmt"
	"strings"

	"golang.org/x/sync/errgroup"

	"github.com/bartosian/suimon/internal/core/domain/enums"
	"github.com/bartosian/suimon/internal/core/ports"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// rpcMethodToMetric maps an RPC method to a metric type.
	rpcMethodToMetric = map[enums.RPCMethod]enums.MetricType{
		enums.RPCMethodGetTotalTransactionBlocks:         enums.MetricTypeTotalTransactionBlocks,
		enums.RPCMethodGetLatestCheckpointSequenceNumber: enums.MetricTypeLatestCheckpoint,
		enums.RPCMethodGetSuiSystemState:                 enums.MetricTypeSuiSystemState,
		enums.RPCMethodGetValidatorsApy:                  enums.MetricTypeValidatorsApy,
		enums.RPCMethodGetEpochs:                         enums.MetricTypeEpochsHistory,
	}
	// rpcMethodToParams maps an RPC method to a params list.
	rpcMethodToParams = map[enums.RPCMethod][]any{
		enums.RPCMethodGetEpochs: {nil, 100, true},
	}
	// prometheusToMetric maps a Prometheus metric name to a metric type.
	prometheusToMetric = map[enums.PrometheusMetricName]enums.MetricType{
		enums.PrometheusMetricNameTotalTransactionCertificates: enums.MetricTypeTotalTransactionCertificates,
		enums.PrometheusMetricNameTotalTransactionEffects:      enums.MetricTypeTotalTransactionEffects,
		enums.PrometheusMetricNameHighestKnownCheckpoint:       enums.MetricTypeHighestKnownCheckpoint,
		enums.PrometheusMetricNameHighestSyncedCheckpoint:      enums.MetricTypeHighestSyncedCheckpoint,
		enums.PrometheusMetricNameLastExecutedCheckpoint:       enums.MetricTypeLastExecutedCheckpoint,
		enums.PrometheusMetricNameCurrentEpoch:                 enums.MetricTypeCurrentEpoch,
		enums.PrometheusMetricNameEpochTotalDuration:           enums.MetricTypeEpochTotalDuration,
		enums.PrometheusMetricNameCurrentRound:                 enums.MetricTypeCurrentRound,
		enums.PrometheusMetricNameHighestProcessedRound:        enums.MetricTypeHighestProcessedRound,
		enums.PrometheusMetricNameLastCommittedRound:           enums.MetricTypeLastCommittedRound,
		enums.PrometheusMetricNamePrimaryNetworkPeers:          enums.MetricTypePrimaryNetworkPeers,
		enums.PrometheusMetricNameWorkerNetworkPeers:           enums.MetricTypeWorkerNetworkPeers,
		enums.PrometheusMetricNameSuiNetworkPeers:              enums.MetricTypeSuiNetworkPeers,
		enums.PrometheusMetricNameSkippedConsensusTransactions: enums.MetricTypeSkippedConsensusTransactions,
		enums.PrometheusMetricNameTotalSignatureErrors:         enums.MetricTypeTotalSignatureErrors,
		enums.PrometheusMetricNameUptime:                       enums.MetricTypeUptime,
		enums.PrometheusMetricNameCertificatesCreated:          enums.MetricTypeCertificatesCreated,
		enums.PrometheusMetricNameNonConsensusLatencySum:       enums.MetricTypeNonConsensusLatencySum,
	}
	// tableToRpcMethods maps a table type to a list of RPC methods.
	tableToRpcMethods = map[enums.TableType][]enums.RPCMethod{
		enums.TableTypeNode: {
			enums.RPCMethodGetTotalTransactionBlocks,
			enums.RPCMethodGetLatestCheckpointSequenceNumber,
		},
		enums.TableTypeRPC: {
			enums.RPCMethodGetTotalTransactionBlocks,
			enums.RPCMethodGetLatestCheckpointSequenceNumber,
			enums.RPCMethodGetSuiSystemState,
			enums.RPCMethodGetValidatorsApy,
		},
		enums.TableTypeEpochsHistory: {
			enums.RPCMethodGetEpochs,
		},
	}
	// tablesToCallMetrics maps a table type to a boolean value indicating whether to call metrics for that table type.
	tablesToCallMetrics = map[enums.TableType]bool{
		enums.TableTypeNode:      true,
		enums.TableTypeValidator: true,
	}
)

// getPrometheusMetricsForTableType returns a list of Prometheus metrics that should be collected for a table type.
func getPrometheusMetricsForTableType(table enums.TableType) ports.Metrics {
	metrics := ports.Metrics{
		enums.PrometheusMetricNameTotalTransactionCertificates: {
			MetricType: enums.PrometheusMetricTypeCounter,
		},
		enums.PrometheusMetricNameTotalTransactionEffects: {
			MetricType: enums.PrometheusMetricTypeCounter,
		},
		enums.PrometheusMetricNameHighestKnownCheckpoint: {
			MetricType: enums.PrometheusMetricTypeGauge,
		},
		enums.PrometheusMetricNameHighestSyncedCheckpoint: {
			MetricType: enums.PrometheusMetricTypeGauge,
		},
		enums.PrometheusMetricNameLastExecutedCheckpoint: {
			MetricType: enums.PrometheusMetricTypeGauge,
		},
		enums.PrometheusMetricNameCurrentEpoch: {
			MetricType: enums.PrometheusMetricTypeGauge,
		},
		enums.PrometheusMetricNameEpochTotalDuration: {
			MetricType: enums.PrometheusMetricTypeGauge,
		},
		enums.PrometheusMetricNameSuiNetworkPeers: {
			MetricType: enums.PrometheusMetricTypeGauge,
		},
		enums.PrometheusMetricNameUptime: {
			MetricType: enums.PrometheusMetricTypeCounter,
		},
	}

	if table == enums.TableTypeValidator {
		metrics[enums.PrometheusMetricNameLastCommittedRound] = ports.MetricConfig{
			MetricType: enums.PrometheusMetricTypeGauge,
		}
		metrics[enums.PrometheusMetricNamePrimaryNetworkPeers] = ports.MetricConfig{
			MetricType: enums.PrometheusMetricTypeGauge,
		}
		metrics[enums.PrometheusMetricNameHighestProcessedRound] = ports.MetricConfig{
			MetricType: enums.PrometheusMetricTypeGauge,
			Labels: prometheus.Labels{
				"source": "own",
			},
		}
		metrics[enums.PrometheusMetricNameWorkerNetworkPeers] = ports.MetricConfig{
			MetricType: enums.PrometheusMetricTypeGauge,
		}
		metrics[enums.PrometheusMetricNameTotalSignatureErrors] = ports.MetricConfig{
			MetricType: enums.PrometheusMetricTypeCounter,
		}
		metrics[enums.PrometheusMetricNameSkippedConsensusTransactions] = ports.MetricConfig{
			MetricType: enums.PrometheusMetricTypeCounter,
		}
		metrics[enums.PrometheusMetricNameCurrentRound] = ports.MetricConfig{
			MetricType: enums.PrometheusMetricTypeGauge,
		}
		metrics[enums.PrometheusMetricNameCertificatesCreated] = ports.MetricConfig{
			MetricType: enums.PrometheusMetricTypeCounter,
		}
	}

	return metrics
}

// GetPrometheusMetrics fetches Prometheus metrics for the host and sends the metric values to the host's Metrics field.
// The function constructs a URL using the host's HostPort and Ports fields, sends an HTTP GET request to the URL, and parses the response as a Prometheus text format.
// Returns an error if the HTTP request fails or if the response cannot be parsed as Prometheus text format.
func (host *Host) GetPrometheusMetrics() error {
	metricsDef := getPrometheusMetricsForTableType(host.TableType)

	result, err := host.gateways.prometheus.CallFor(metricsDef)
	if err != nil {
		return err
	}

	if result == nil {
		return errors.New("failed to get metrics from Prometheus")
	}

	for metricName, metricValue := range result {
		metricType, ok := prometheusToMetric[metricName]
		if !ok {
			// ignore unused metric
			delete(result, metricName)
			continue
		}

		if err := host.Metrics.SetValue(metricType, metricValue.Value); err != nil {
			return err
		}

		// delete processed metric from result map
		delete(result, metricName)

		if metricType != enums.MetricTypeUptime {
			continue
		}

		if value, ok := metricValue.Labels["version"]; ok {
			versionInfo := strings.SplitN(value, "-", 2)

			if err := host.Metrics.SetValue(enums.MetricTypeVersion, versionInfo[0]); err != nil {
				return err
			}

			if len(versionInfo) == 2 {
				if err := host.Metrics.SetValue(enums.MetricTypeCommit, versionInfo[1]); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// GetMetrics fetches data from the host by calling three different methods asynchronously: GetTotalTransactionNumber, GetLatestCheckpoint, and GetPrometheusMetrics.
// The function waits for all three methods to complete before returning.
// Returns an error if any of the three methods fail or return an error.
func (host *Host) GetMetrics() error {
	var errGroup errgroup.Group

	rpcMethods := tableToRpcMethods[host.TableType]
	for _, method := range rpcMethods {
		method := method

		errGroup.Go(func() error {
			return host.GetDataByMetric(method)
		})
	}

	if ok := tablesToCallMetrics[host.TableType]; ok {
		errGroup.Go(func() error {
			return host.GetPrometheusMetrics()
		})
	}

	if err := errGroup.Wait(); err != nil {
		return fmt.Errorf("failed to get metrics for table %s, host: %s: %w", host.TableType, host.Endpoint.Address, err)
	}

	return nil
}

// GetDataByMetric is a method of the Host struct that retrieves data for a given RPC method
// and stores it as a metric in the Metrics struct. It takes an RPCMethod input parameter and
// returns an error if the method is not supported.
func (host *Host) GetDataByMetric(method enums.RPCMethod) error {
	metric, ok := rpcMethodToMetric[method]
	if !ok {
		return fmt.Errorf("unsupported RPC method: %v", method)
	}

	params := rpcMethodToParams[method]

	result, err := host.gateways.rpc.CallFor(method, params...)
	if err != nil {
		return err
	}

	return host.Metrics.SetValue(metric, result)
}
