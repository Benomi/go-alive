package notifiers

import s "github.com/Benomi/go-alive/strategies"

type Notifier interface {
	NotifySpecificPortHealthCheckResult(result s.SpecificPortHealthCheckResult) error
	NotifyHealthCheckResult(result s.HealthCheckResult) error
}
