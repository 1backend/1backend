/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package logger

import (
	"context"
	"log"
	"log/slog"
	"runtime"
	"strconv"
	"strings"
)

func NewLogger(fields ...slog.Attr) *slog.Logger {
	args := make([]any, len(fields))
	for i, f := range fields {
		args[i] = f
	}
	return base.With(args...)
}

var jsonHandler = slog.NewJSONHandler(log.Writer(), &slog.HandlerOptions{
	Level:     slog.LevelDebug,
	AddSource: false,
})

type fileOnlyHandler struct{ slog.Handler }

func (h fileOnlyHandler) Enabled(ctx context.Context, lvl slog.Level) bool {
	return h.Handler.Enabled(ctx, lvl)
}

func (h fileOnlyHandler) Handle(ctx context.Context, r slog.Record) error {
	nr := slog.NewRecord(r.Time, r.Level, r.Message, r.PC)
	r.Attrs(func(a slog.Attr) bool {
		if a.Key != slog.SourceKey {
			nr.AddAttrs(a)
		}
		return true
	})

	// Find first caller outside slog/runtime/this logger package.
	pcs := make([]uintptr, 32)
	n := runtime.Callers(3, pcs) // skip Callers + Handle + slog internals entry
	frames := runtime.CallersFrames(pcs[:n])
	for {
		fr, ok := frames.Next()
		if !ok {
			break
		}
		fn, f := fr.Function, fr.File
		if strings.HasPrefix(fn, "runtime.") ||
			strings.HasPrefix(fn, "log/slog.") ||
			strings.HasPrefix(fn, "log.") ||
			strings.Contains(fn, "logger.") {
			continue
		}
		nr.AddAttrs(slog.String("file", f+":"+strconv.Itoa(fr.Line))) // full path + line
		break
	}

	return h.Handler.Handle(ctx, nr)
}

func (h fileOnlyHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return fileOnlyHandler{h.Handler.WithAttrs(attrs)}
}

func (h fileOnlyHandler) WithGroup(name string) slog.Handler {
	return fileOnlyHandler{h.Handler.WithGroup(name)}
}

var base = slog.New(fileOnlyHandler{jsonHandler})

var Debug = base.Debug
var Info = base.Info
var Warn = base.Warn
var Error = base.Error

func init() {
	//
	// Some libraries used like autocert does not use slog directly,
	// so we need to redirect the stdlib log to our custom slog logger.
	//

	// Redirect stdlib log to our custom writer
	log.SetOutput(&logProxyWriter{log: base})
}

// logProxyWriter inspects the message and routes to slog with a guessed level
type logProxyWriter struct {
	log *slog.Logger
}

func (w *logProxyWriter) Write(p []byte) (n int, err error) {
	msg := strings.TrimSpace(string(p))

	// Trim stdlib log timestamp prefix: "YYYY/MM/DD HH:MM:SS "
	if len(msg) > 20 && msg[4] == '/' && msg[7] == '/' && msg[10] == ' ' {
		msg = msg[20:] // Drop the first 20 characters
	}

	lower := strings.ToLower(msg)

	// Heuristics: look for keywords
	switch {
	case strings.Contains(lower, "error"), strings.Contains(lower, "fail"):
		w.log.Error(msg)
	default:
		w.log.Info(msg)
	}

	return len(p), nil
}
