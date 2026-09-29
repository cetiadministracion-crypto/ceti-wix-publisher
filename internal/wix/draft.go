package wix

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/cetiadministracion-crypto/ceti-wix-publisher/internal/ricos"
)

const draftPostsURL = "https://www.wixapis.com/blog/v3/draft-posts"

type DraftInput struct {
	Title       string
	Excerpt     string
	Slug        string
	Language    string
	MemberID    string
	AltText     string
	RichContent ricos.RichContent
	Hero        *MediaFile
}

type DraftResult struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

type createDraftResponse struct {
	DraftPost DraftResult `json:"draftPost"`
}

func (c *Client) CreateDraft(input DraftInput) (*DraftResult, error) {
	payload := map[string]any{
		"draftPost": map[string]any{
			"title":             input.Title,
			"excerpt":           input.Excerpt,
			"language":          input.Language,
			"memberId":          input.MemberID,
			"seoSlug":           input.Slug,
			"richContent":       input.RichContent,
			"commentingEnabled": true,
			"media": map[string]any{
				"displayed": true,
				"custom":    true,
				"wixMedia": map[string]any{
					"image": map[string]any{
						"id":       input.Hero.ID,
						"url":      input.Hero.URL,
						"filename": input.Hero.Filename,
					},
				},
			},
			"heroImage": map[string]any{
				"id":       input.Hero.ID,
				"url":      input.Hero.URL,
				"filename": input.Hero.Filename,
				"altText":  input.AltText,
			},
		},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("encode draft: %w", err)
	}

	req, err := http.NewRequest(
		http.MethodPost,
		draftPostsURL,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", c.APIKey)
	req.Header.Set("wix-site-id", c.SiteID)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("create Wix draft: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		responseBody, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf(
			"create Wix draft failed: HTTP %d: %s",
			resp.StatusCode,
			string(responseBody),
		)
	}

	var result createDraftResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("decode Wix draft response: %w", err)
	}

	if result.DraftPost.ID == "" {
		return nil, fmt.Errorf("Wix returned an empty draft ID")
	}

	return &result.DraftPost, nil
}
