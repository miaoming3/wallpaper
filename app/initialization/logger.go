package initialization

import (
	"github.com/gin-gonic/gin"
	"github.com/miaoming3/wallpaper/app/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"path/filepath"
	"time"
)

func InitLoggers() {

	logger := zap.New(zapcore.NewCore(getEncoder(), encode(), zapcore.DebugLevel), zap.AddCaller())
	global.SugarLogger = logger.Sugar()
}

func encode() zapcore.WriteSyncer {
	path := filepath.Join(global.SysConfig.Director, "log.txt")
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		log.Panicln("创建日志文件失败")
	}
	file, _ := os.Create(path)
	return zapcore.AddSync(file)
}

func getEncoder() zapcore.Encoder {
	zapcon := zap.NewProductionEncoderConfig()
	if gin.Mode() == "dev" {
		zapcon = zap.NewDevelopmentEncoderConfig()
	}
	zapcon.EncodeTime = DateTimeTimeEncoder
	zapcon.EncodeLevel = zapcore.CapitalLevelEncoder

	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

func encodeTimeLayout(t time.Time, layout string, enc zapcore.PrimitiveArrayEncoder) {
	type appendTimeEncoder interface {
		AppendTimeLayout(time.Time, string)
	}

	if enc, ok := enc.(appendTimeEncoder); ok {
		enc.AppendTimeLayout(t, layout)
		return
	}

	enc.AppendString(t.Format(layout))
}

func DateTimeTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	encodeTimeLayout(t, time.DateTime, enc)
}
