package logger

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	goLog "log"
	"runtime"
	"strings"
)

var logger *goLog.Logger

const (
	programCounter = 3
	callersToSkip  = 3
	RequestIdTag   = "request_id"
)

type Tag struct {
	Name  string
	Value any
}

func WithTags(values ...string) []Tag {
	tags := make([]Tag, 0, len(values)/2)
	i := 0
	for i < len(values)-1 {
		tags = append(tags, Tag{Name: values[i], Value: values[i+1]})
		i += 2
	}

	return tags
}

func (t Tag) String() string {
	return fmt.Sprintf("[%s:%v] ", t.Name, t.Value)
}

func Trace(ctx context.Context, tags []Tag, msg string) {
	log(ctx, "TRACE", tags, msg)
}

func Info(ctx context.Context, tags []Tag, msg string) {
	log(ctx, "INFO", tags, msg)
}

func Warning(ctx context.Context, tags []Tag, msg string) {
	log(ctx, "WARNING", tags, msg)
}

func Error(ctx context.Context, tags []Tag, err error, msg string) {
	log(ctx, "ERROR", append(tags, Tag{"error", err.Error()}), msg)
}

func Panic(ctx context.Context, tags []Tag, msg string) {
	log(ctx, "PANIC", tags, msg)
}

func log(ctx context.Context, level string, customTags []Tag, msg string) {
	var fixedTags []Tag
	fixedTags = append(fixedTags, getRequestIdTag(ctx))
	fixedTags = append(fixedTags, customTags...)
	s := ""
	for _, t := range fixedTags {
		s += t.String()
	}

	logger.Printf("level:%s %s: %s| %s\n", level, getFileName(), s, msg)
}

func getRequestIdTag(ctx context.Context) Tag {
	return Tag{RequestIdTag, ctx.Value(echo.HeaderXRequestID)}
}

func getFileName() string {
	// Ask runtime.Callers for up to 10 pcs, including runtime.Callers itself.
	pc := make([]uintptr, programCounter)
	n := runtime.Callers(callersToSkip, pc)
	if n == 0 {
		// No pcs available. Stop now.
		// This can happen if the first argument to runtime.Callers is large.
		return ""
	}

	pc = pc[:n] // pass only valid pcs to runtime.CallersFrames
	frames := runtime.CallersFrames(pc)

	for {
		frame, _ := frames.Next()
		if strings.Contains(frame.File, "runtime/") || strings.Contains(frame.File, "logger/") {
			continue
		}
		// Get file name
		file := frame.File
		for i := len(frame.File) - 1; i > 0; i-- {
			if frame.File[i] == '/' {
				file = frame.File[i+1:]
				break
			}
		}

		return fmt.Sprintf("%s:%d", file, frame.Line)
	}
}

func init() {
	logger = goLog.Default()
	logger.SetFlags(goLog.LstdFlags | goLog.LUTC)
}
