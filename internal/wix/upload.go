package wix

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"
)

const generateUploadURL = "https://www.wixapis.com/site-media/v1/files/generate-upload-url"

type MediaFile struct {
	ID       string `json:"id"`
	URL      string `json:"url"`
	Filename string `json:"displayName"`
}

type generateUploadResponse struct {
	UploadURL string `json:"uploadUrl"`
}

type uploadResponse struct {
	File MediaFile `json:"file"`
}

func (c *Client) UploadImage(path string) (*MediaFile, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read image: %w", err)
	}

	filename := filepath.Base(path)
	mimeType := mime.TypeByExtension(filepath.Ext(path))
	if mimeType == "" {
		mimeType = http.DetectContentType(data)
	}

	body, err := json.Marshal(map[string]interface{}{
		"mimeType": mimeType,
		"fileName": filename,
		"private":  false,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, generateUploadURL, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", c.APIKey)
	req.Header.Set("wix-site-id", c.SiteID)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("generate upload URL: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		responseBody, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf(
			"generate upload URL failed: HTTP %d: %s",
			resp.StatusCode,
			string(responseBody),
		)
	}

	var generated generateUploadResponse
	if err := json.NewDecoder(resp.Body).Decode(&generated); err != nil {
		return nil, fmt.Errorf("decode upload URL: %w", err)
	}

	uploadReq, err := http.NewRequest(
		http.MethodPut,
		generated.UploadURL,
		bytes.NewReader(data),
	)
	if err != nil {
		return nil, err
	}

	uploadReq.Header.Set("Content-Type", mimeType)

	uploadResp, err := c.HTTPClient.Do(uploadReq)
	if err != nil {
		return nil, fmt.Errorf("upload image: %w", err)
	}
	defer uploadResp.Body.Close()

	if uploadResp.StatusCode < 200 || uploadResp.StatusCode >= 300 {
		responseBody, _ := io.ReadAll(uploadResp.Body)
		return nil, fmt.Errorf(
			"image upload failed: HTTP %d: %s",
			uploadResp.StatusCode,
			string(responseBody),
		)
	}

	var result uploadResponse
	if err := json.NewDecoder(uploadResp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("decode Wix upload response: %w", err)
	}

	if result.File.ID == "" {
		return nil, fmt.Errorf("Wix returned an empty media file ID")
	}

	return &result.File, nil
}
