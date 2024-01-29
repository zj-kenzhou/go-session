package config

var _config SessionConfig

type SessionConfig struct {

	/** token的长久有效期(单位:秒) -1代表永久 */
	Timeout int64

	/*token临时有效期 [指定时间内无操作就视为token过期] (单位: 秒), -1代表不限制 */
	ActivityTimeout int64

	/*持久化的key 的前缀*/
	KeyPrefix string

	/** redis Host */
	Host []string

	Username string

	/** redis password */
	Password string

	SentinelUsername string

	/** redis password */
	SentinelPassword string

	/** redis db */
	Db int

	// 主节点名称（集群需要）
	MasterName string

	// 连接名称（集群需要）
	ClientName string

	// 连接池最大数量
	PoolSize int

	// 等待连接超时时间，单位分钟
	PoolTimeout int

	// 最小空闲连接数量
	MinIdleConns int

	// 最大空闲连接数量
	MaxIdleConns int

	// 空闲连接存活时间 ，单位分钟
	ConnMaxIdleTime int

	// 连接最大存活时间 ，单位分钟
	ConnMaxLifetime int
}

func SetConfig(config SessionConfig) {
	_config = config
}

func GetConfig() SessionConfig {
	return _config
}
