/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package firehoseservice

import (
	"encoding/json"
	"log"
	"net/http"

	sdk "github.com/openorch/openorch/sdk/go"

	firehose "github.com/openorch/openorch/server/internal/services/firehose/types"
)

// @ID subscribeToEvents
// @Summary Subscribe to the Event Stream
// @Description Establish a subscription to the firehose events and accept a real time stream of them.
// @Tags Firehose Svc
// @Accept json
// @Produce text/event-stream
// @Success 200 {string} string "Event data"
// @Failure 401 {object} firehose.ErrorResponse "Unauthorized"
// @Failure 500 {object} firehose.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /firehose-svc/events/subscribe [get]
func (p *FirehoseService) Subscribe(
	w http.ResponseWriter,
	r *http.Request,
) {

	isAuthRsp, _, err := p.clientFactory.Client(sdk.WithTokenFromRequest(r)).
		UserSvcAPI.IsAuthorized(r.Context(), *firehose.PermissionFirehoseStream.Id).
		Execute()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if !isAuthRsp.GetAuthorized() {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`Unauthorized`))
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	eventsChannel := make(chan []*firehose.Event)
	subscriberID := p.subscribe(func(events []*firehose.Event) {
		eventsChannel <- events
	})
	defer p.unsubscribe(subscriberID)

	ctx := r.Context()

	go func() {
		defer func() {
			recover()
		}()
		<-ctx.Done()
		p.unsubscribe(subscriberID)
		close(eventsChannel)
	}()

	for {
		select {
		case events, ok := <-eventsChannel:
			if !ok {
				// The channel is closed, indicating the user has disconnected from the event stream.
				// Simply exit the loop and return.
				return
			}

			for _, event := range events {
				jsonResp, err := json.Marshal(event)
				if err != nil {
					log.Printf("Failed to marshal event: %v", err)
					continue
				}

				if _, err := w.Write([]byte("data: " + string(jsonResp) + "\n\n")); err != nil {
					log.Printf("Failed to write event to client: %v", err)
					return
				}

				if flusher, ok := w.(http.Flusher); ok {
					flusher.Flush()
				} else {
					log.Println("Warning: ResponseWriter does not support flushing, streaming might be delayed")
				}
			}
		}
	}
}
