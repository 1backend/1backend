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
)

var jsonHandler = slog.NewJSONHandler(log.Writer(), &slog.HandlerOptions{
	Level: slog.LevelDebug,
})

var Logger = slog.New(jsonHandler)

var Debug = Logger.Debug
var Info = Logger.Info
var Warn = Logger.Warn
var Error = Logger.Error
