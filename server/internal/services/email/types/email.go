/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package email_svc

import "time"

type ErrorResponse struct {
	Error string `json:"error"`
}

type Event struct {
	Name string `json:"name"`
	Data any    `json:"data"`
}

type Email struct {
	InternalId string `json:"internalId,omitempty" swagger:"ignore"`

	// Unique identifier
	Id string `json:"id" binding:"required"`

	// List of recipient email addresses
	To []string `json:"to" binding:"required"`

	// List of CC recipient email addresses (optional)
	CC []string `json:"cc,omitempty"`

	// List of BCC recipient email addresses (optional)
	BCC []string `json:"bcc,omitempty"`

	// Email subject line
	Subject string `json:"subject" binding:"required"`

	// Email body content (plain text or HTML)
	Body string `json:"body" binding:"required"`

	// Content type: "text/plain" or "text/html"
	ContentType string `json:"contentType" binding:"required"`

	// Timestamp of email creation
	CreatedAt time.Time `json:"createdAt" binding:"required"`
}

type SendEmailRequest struct {
	// Unique identifier
	Id string `json:"id,omitempty"`

	// List of recipient email addresses
	To []string `json:"to" binding:"required"`

	// List of CC recipient email addresses (optional)
	CC []string `json:"cc,omitempty"`

	// List of BCC recipient email addresses (optional)
	BCC []string `json:"bcc,omitempty"`

	// Email subject line
	Subject string `json:"subject" binding:"required"`

	// Email body content (plain text or HTML)
	Body string `json:"body" binding:"required"`

	// Content type: "text/plain" or "text/html"
	ContentType string `json:"contentType,omitempty"`

	// List of file attachments (optional)
	Attachments []Attachment `json:"attachments"`
}

func (e *Email) GetId() string {
	return e.Id
}

type Attachment struct {
	// A File Svc file ID. Requires the file to be uploaded separately.
	// Recommended for mid to large-sized files.
	// If this field is specified, all other fields are optional.
	FileId string `json:"fileId,omitempty"`

	// File name for the attachment.
	// Required for inline attachments (i.e., those not using File Svc, see FileId).
	Filename string `json:"filename" binding:"required"`

	// MIME type of the file (e.g., "application/pdf", "image/png")
	// Required for inline attachments (i.e., those not using File Svc, see FileId).
	ContentType string `json:"contentType" binding:"required"`

	// Base64-encoded file content. Use this for small files.
	// Required for inline attachments (i.e., those not using File Svc, see FileId).
	Content string `json:"content,omitempty"`
}

type SendEmailResponse struct {
	EmailId string `json:"emailId"` // Unique identifier for the sent email
	Status  string `json:"status"`  // Status of the email send operation ("sent", "queued", etc.)
}
