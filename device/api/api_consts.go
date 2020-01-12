package api

//global const
const (
	//全局上下文
	PATH_IOT = "/iot"
	//设备服务上下文
	PATH_DEVICE = "/device"
	//二级上下文
	PATH_SINGLE = "/single"
	PATH_BATCH  = "/batch"

	PATH_IOT_DEVICE_SINGLE = PATH_IOT + PATH_DEVICE + PATH_SINGLE
	PATH_IOT_DEVICE_BATCH  = PATH_IOT + PATH_DEVICE + PATH_BATCH

	//增删查改
	PATH_ADD    = "/add"
	PATH_UPDATE = "/update"
	PATH_DELETE = "/delete"
	PATH_QUERY  = "/query"
	PATH_EXEC   = "/exec"
)

const (
	PATH_DEVICETYPE_SINGLE       = "/devicetype"
	PATH_DEVICETYPE_SINGLE_QUERY = PATH_IOT_DEVICE_SINGLE + PATH_DEVICETYPE_SINGLE + PATH_QUERY
)
const (
	PATH_DEVICE_SINGLE_INFO   = "/device"
	PATH_DEVICE_SINGLE_ADD    = PATH_IOT_DEVICE_SINGLE + PATH_DEVICE_SINGLE_INFO + PATH_ADD
	PATH_DEVICE_SINGLE_DELETE = PATH_IOT_DEVICE_SINGLE + PATH_DEVICE_SINGLE_INFO + PATH_DELETE
	PATH_DEVICE_SINGLE_UPDATE = PATH_IOT_DEVICE_SINGLE + PATH_DEVICE_SINGLE_INFO + PATH_UPDATE
	PATH_DEVICE_SINGLE_QUERY  = PATH_IOT_DEVICE_SINGLE + PATH_DEVICE_SINGLE_INFO + PATH_QUERY
)
const (
	PATH_DEVICE_SINGLE_C        = "/device/c"
	PATH_DEVICE_SINGLE_C_QUERY  = PATH_IOT_DEVICE_SINGLE + PATH_DEVICE_SINGLE_C + PATH_QUERY
	PATH_DEVICE_SINGLE_C_UPDATE = PATH_IOT_DEVICE_SINGLE + PATH_DEVICE_SINGLE_C + PATH_UPDATE
)

const (
	PATH_DEVICE_SINGLE_F        = "/device/f"
	PATH_DEVICE_SINGLE_F_UPDATE = PATH_IOT_DEVICE_SINGLE + PATH_DEVICE_SINGLE_F + PATH_UPDATE
	PATH_DEVICE_SINGLE_F_EXEC   = PATH_IOT_DEVICE_SINGLE + PATH_DEVICE_SINGLE_F + PATH_EXEC
)

const (
	PATH_DEVICE_SINGLE_P        = "/device/p"
	PATH_DEVICE_SINGLE_P_UPDATE = PATH_IOT_DEVICE_SINGLE + PATH_DEVICE_SINGLE_P + PATH_UPDATE
	PATH_DEVICE_SINGLE_P_QUERY  = PATH_IOT_DEVICE_SINGLE + PATH_DEVICE_SINGLE_P + PATH_QUERY
)

const (
	PATH_DEVICE_SINGLE_R        = "/device/r"
	PATH_DEVICE_SINGLE_R_UPDATE = PATH_IOT_DEVICE_SINGLE + PATH_DEVICE_SINGLE_R + PATH_UPDATE
	PATH_DEVICE_SINGLE_R_QUERY  = PATH_IOT_DEVICE_SINGLE + PATH_DEVICE_SINGLE_R + PATH_QUERY
)

//batch api

const (
	PATH_DEVICETYPE_BATCH       = "/devicetype"
	PATH_DEVICETYPE_BATCH_QUERY = PATH_IOT_DEVICE_BATCH + PATH_DEVICETYPE_BATCH + PATH_QUERY
)

const (
	PATH_DEVICE_BATCH_INFO   = "/device"
	PATH_DEVICE_BATCH_ADD    = PATH_IOT_DEVICE_BATCH + PATH_DEVICE_BATCH_INFO + PATH_ADD
	PATH_DEVICE_BATCH_DELETE = PATH_IOT_DEVICE_BATCH + PATH_DEVICE_BATCH_INFO + PATH_DELETE
	PATH_DEVICE_BATCH_UPDATE = PATH_IOT_DEVICE_BATCH + PATH_DEVICE_BATCH_INFO + PATH_UPDATE
	PATH_DEVICE_BATCH_QUERY  = PATH_IOT_DEVICE_BATCH + PATH_DEVICE_BATCH_INFO + PATH_QUERY
)
const (
	PATH_DEVICE_BATCH_C       = "/device/c"
	PATH_DEVICE_BATCH_C_QUERY = PATH_IOT_DEVICE_BATCH + PATH_DEVICE_BATCH_C + PATH_QUERY
)
const (
	PATH_DEVICE_BATCH_F        = "/device/f"
	PATH_DEVICE_BATCH_F_UPDATE = PATH_IOT_DEVICE_BATCH + PATH_DEVICE_BATCH_F + PATH_UPDATE
	PATH_DEVICE_BATCH_F_EXEC   = PATH_IOT_DEVICE_BATCH + PATH_DEVICE_BATCH_F + PATH_EXEC
)
const (
	PATH_DEVICE_BATCH_P        = "/device/p"
	PATH_DEVICE_BATCH_P_UPDATE = PATH_IOT_DEVICE_BATCH + PATH_DEVICE_BATCH_P + PATH_UPDATE
	PATH_DEVICE_BATCH_P_QUERY  = PATH_IOT_DEVICE_BATCH + PATH_DEVICE_BATCH_P + PATH_QUERY
)

const (
	PATH_DEVICE_BATCH_R        = "/device/r"
	PATH_DEVICE_BATCH_R_UPDATE = PATH_IOT_DEVICE_BATCH + PATH_DEVICE_BATCH_R + PATH_UPDATE
	PATH_DEVICE_BATCH_R_QUERY  = PATH_IOT_DEVICE_BATCH + PATH_DEVICE_BATCH_R + PATH_QUERY
)
