package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

type uploadResponse struct {
	Success   bool   `json:"success"`
	Code      string `json:"code"`
	Message   string `json:"message"`
	Data      data   `json:"data"`
	Images    string `json:"images"`
	RequestId string `json:"RequestId"`
}

type data struct {
	FileId    int    `json:"file_id"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	Filename  string `json:"filename"`
	Storename string `json:"storename"`
	Size      int    `json:"size"`
	Path      string `json:"path"`
	Hash      string `json:"hash"`
	Url       string `json:"url"`
	Delete    string `json:"delete"`
	Page      string `json:"page"`
}

var (
	uploadUrl = "https://smms.app/api/v2/upload"
	token     = "your token"
)

func uploadLocalFile(filePath string) (string, error) {
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// 创建multipart请求体
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("smfile", file.Name())
	if err != nil {
		return "", err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return "", err
	}
	err = writer.Close()
	if err != nil {
		return "", err
	}

	// 创建请求
	req, err := http.NewRequest("POST", uploadUrl, body)
	if err != nil {
		return "", err
	}

	// 设置请求头
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", token)

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 读取响应体
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	log.Println("UploadResponse:", string(responseBody))

	var res uploadResponse
	if err = json.Unmarshal(responseBody, &res); err != nil {
		return "", err
	}

	if res.Data.Url == "" {
		if res.Images != "" {
			return res.Images, nil
		}
		return "", errors.New(res.Message)
	}

	return res.Data.Url, nil
}

func uploadBase64(base64Img string) (string, error) {
	// 将Base64字符串解码为字节数组
	imgBytes, err := base64.StdEncoding.DecodeString(strings.Split(base64Img, ",")[1])
	if err != nil {
		return "", err
	}

	// 创建multipart请求体
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	// 创建表单文件部分
	part, err := writer.CreateFormFile("smfile", "image.png")
	if err != nil {
		return "", err
	}

	// 将解码后的图片字节写入表单文件部分
	_, err = part.Write(imgBytes)
	if err != nil {
		return "", err
	}

	// 关闭multipart writer
	err = writer.Close()
	if err != nil {
		return "", err
	}

	// 创建请求
	req, err := http.NewRequest("POST", uploadUrl, body)
	if err != nil {
		return "", err
	}

	// 设置请求头
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", token)

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 读取响应体
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	log.Println("UploadResponse:", string(responseBody))

	var res uploadResponse
	if err = json.Unmarshal(responseBody, &res); err != nil {
		return "", err
	}

	if res.Data.Url == "" {
		if res.Images != "" {
			return res.Images, nil
		}
		return "", errors.New(res.Message)
	}

	return res.Data.Url, nil
}

func UploadFile(imgBytes []byte, fileName string) (string, error) {
	// 创建multipart请求体
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("smfile", fileName)
	if err != nil {
		return "", err
	}

	// 将文件字节写入表单文件部分
	_, err = part.Write(imgBytes)
	if err != nil {
		return "", err
	}

	// 关闭multipart writer
	err = writer.Close()
	if err != nil {
		return "", err
	}

	// 创建请求
	req, err := http.NewRequest("POST", uploadUrl, body)
	if err != nil {
		return "", err
	}

	// 设置请求头
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", token)

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 读取响应体
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	log.Println("UploadResponse:", string(responseBody))

	var res uploadResponse
	if err = json.Unmarshal(responseBody, &res); err != nil {
		return "", err
	}

	if res.Data.Url == "" {
		if res.Images != "" {
			return res.Images, nil
		}
		return "", errors.New(res.Message)
	}
	return res.Data.Url, nil
}
