package csdn_example

const (
	INFO = iota + 1
	DEBUG
	ERROR
)

type Logger interface {
	LogMessage(level int, message string, logger Logger)
	write(message string)
}

type AbstractLogger struct {
	level      int
	nextLogger Logger
}

func (s *AbstractLogger) write(message string) {

}

func (s *AbstractLogger) SetNextLogger(logger Logger) {
	s.nextLogger = logger
}

func (s *AbstractLogger) LogMessage(level int, message string, logger Logger) {
	if s.level <= level {
		logger.write(message)
	} else {
		if s.nextLogger != nil {
			s.nextLogger.LogMessage(level, message, s.nextLogger)
		}
	}
}

func testChainofResponsibilityPattern() Logger{
	errorLogger := new(ErrorLogger)
	errorLogger.abstractLogger.level = ERROR

	fileLogger := new(FileLogger)
	fileLogger.abstractLogger.level = DEBUG

	consoleLogger := new(ConsoleLogger)
	consoleLogger.abstractLogger.level = INFO

	errorLogger.abstractLogger.SetNextLogger(fileLogger)
	fileLogger.abstractLogger.SetNextLogger(consoleLogger)

	return errorLogger
}