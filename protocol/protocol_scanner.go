package protocol

import (
	"github.com/cristian-ion-neagu/onionscan/config"
	"github.com/cristian-ion-neagu/onionscan/report"
)

type Scanner interface {
	ScanProtocol(hiddenService string, onionscanConfig *config.OnionScanConfig, report *report.OnionScanReport)
}
