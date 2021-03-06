package app

import "github.com/Sirupsen/logrus"

//Classe de logger padrão

//Logger define a interface do logger que vai ser exposta pelo RequestScope
type Logger interface {
	//SetField adiciona o campo que deve passar pelo log em cada mensagem
	SetField(name, value string)

	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
}

type logger struct {
	logger *logrus.Logger
	fields logrus.Fields
}

// NewLogger cria o logger com os campos que devem ser adicionados para cada mensagem.
func NewLogger(l *logrus.Logger, fields logrus.Fields) Logger {
	return &logger{
		logger: l,
		fields: fields,
	}
}

func (l *logger) SetField(name, value string) {
	l.fields[name] = value
}

func (l *logger) SetFieldInt(name string, value int) {
	l.fields[name] = value
}

func (l *logger) Debugf(format string, args ...interface{}) {
	l.tagged().Debugf(format, args...)
}

func (l *logger) Infof(format string, args ...interface{}) {
	l.tagged().Infof(format, args...)
}

func (l *logger) Warnf(format string, args ...interface{}) {
	l.tagged().Warnf(format, args...)
}

func (l *logger) Errorf(format string, args ...interface{}) {
	l.tagged().Errorf(format, args...)
}

func (l *logger) Debug(args ...interface{}) {
	l.tagged().Debug(args...)
}

func (l *logger) Info(args ...interface{}) {
	l.tagged().Info(args...)
}

func (l *logger) Warn(args ...interface{}) {
	l.tagged().Warn(args...)
}

func (l *logger) Error(args ...interface{}) {
	l.tagged().Error(args...)
}

func (l *logger) tagged() *logrus.Entry {
	return l.logger.WithFields(l.fields)
}
