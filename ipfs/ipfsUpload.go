package ipfs

import (
	"blockchain_services/config"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"mime/multipart"
	"net/http"
)

type IPFSResp struct {
	Hash string `json:"hash"`
	Name string `json:"name"`
	Size string `json:"size"`
}

func UploadToIPFS(content string) (string, error) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("file", content)
	if err := writer.Close(); err != nil {
		logrus.Errorf("error at write form-data input: %v", err)
		return "", err
	}

	req, _ := http.NewRequest("POST", config.IpfsUrl, payload)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Errorf("fail upload ipfs with content: %v, error: %v", content, err)
		return "", err
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorf("error at io read ipfs resp: %v", err)
		return "", err
	}

	fmt.Println(string(respBody))

	var respData IPFSResp
	if err = json.Unmarshal(respBody, &respData); err != nil {
		logrus.Errorf("error at unmarshal ipfs resp: %v", err)
		return "", err
	}

	return respData.Hash, nil
}
