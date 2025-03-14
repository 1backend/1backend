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

type SendEmailRequest struct {
	Email
	Attachments []File `json:"attachments"` // List of file attachments (optional)
}

type Email struct {
	Id          string    `json:"id"`                               // Unique identifier
	To          []string  `json:"to"            binding:"required"` // List of recipient email addresses
	CC          []string  `json:"cc,omitempty"`                     // List of CC recipient email addresses (optional)
	BCC         []string  `json:"bcc,omitempty"`                    // List of BCC recipient email addresses (optional)
	Subject     string    `json:"subject"       binding:"required"` // Email subject line
	Body        string    `json:"body"          binding:"required"` // Email body content (plain text or HTML)
	ContentType string    `json:"contentType"   binding:"required"` // Content type: "text/plain" or "text/html"
	CreatedAt   time.Time `json:"createdAt"     binding:"required"` // Timestamp of email creation
}

func (e *Email) GetId() string {
	return e.Id
}

type Attachment struct {
	Id      string `json:"id"      binding:"required"` // Unique identifier
	EmailId string `json:"emailId" binding:"required"` // Foreign key referencing Email
	File
}

func (a *Attachment) GetId() string {
	return a.Id
}

type File struct {
	Filename    string `json:"filename"    binding:"required"` // Name of the attached file
	Content     string `json:"content"     binding:"required"` // Base64-encoded content of the file
	ContentType string `json:"contentType" binding:"required"` // MIME type of the file (e.g., "application/pdf")
}

type SendEmailResponse struct {
	EmailId string `json:"emailId"` // Unique identifier for the sent email
	Status  string `json:"status"`  // Status of the email send operation ("sent", "queued", etc.)
}
