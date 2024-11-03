package nonna

const (
	BASE_PATH              = "/nonna"
	SYSTEM_UPSTREAM_FILE   = "/nonna_system/upstream"
	SYSTEM_DOWNSTREAM_FILE = "/nonna_system/downstream"
)

var (
    QUEUE *ExtraQueue
)

func init() {
    go func() {
		QUEUE = NewExtraQueue()
	}()
}
