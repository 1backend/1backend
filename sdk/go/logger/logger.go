/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package logger

import (
	"log"
	"log/slog"
	"strings"
)

var jsonHandler = slog.NewJSONHandler(log.Writer(), &slog.HandlerOptions{
	Level: slog.LevelDebug,
})

var Logger = slog.New(jsonHandler)

var Debug = Logger.Debug
var Info = Logger.Info
var Warn = Logger.Warn
var Error = Logger.Error

func init() {
	//
	// Some libraries used like autocert does not use slog directly,
	// so we need to redirect the stdlib log to our custom slog logger.
	//

	// Redirect stdlib log to your custom writer
	log.SetOutput(&logProxyWriter{})
}

// logProxyWriter inspects the message and routes to slog with a guessed level
type logProxyWriter struct{}

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
		Logger.Error(msg)
	default:
		Logger.Info(msg)
	}

	return len(p), nil
}
