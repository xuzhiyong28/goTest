package settings

type App struct {
	JwtSecret string
	PageSize  int
	PrefixUrl string
	RuntimeRootPath string
	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string
	ExportSavePath string
	QrCodeSavePath string
	FontSavePath   string
	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}
