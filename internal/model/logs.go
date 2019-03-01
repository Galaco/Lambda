package model

const (
	LogTypeApplication = 0
	LogTypeProject = 1
)

type Log struct {
	application []string
	project []string
}

func (log *Log) GetLogs(logType int) []string {
	switch logType {
	case LogTypeApplication:
		return log.application
	case LogTypeProject:
		return log.project
	default:
		return nil
	}
}

func (log *Log) AddLog(logType int, message string) {
	switch logType {
	case LogTypeApplication:
		log.application = append(log.application, message)
	case LogTypeProject:
		log.project = append(log.project, message)
	default:
		return
	}
}
