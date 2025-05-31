package helpers

import (
	"bytes"
	"encoding/base64"
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/nfnt/resize"
)

// ✅ Hàm resize ảnh từ file ảnh
func ResizeImageFromFileToBase64(inputPath string, width int, height int) (string, error) {
	// Mở và decode ảnh từ file
	img, err := imaging.Open(inputPath)
	if err != nil {
		return "", errors.New("Không thể mở ảnh: " + err.Error())
	}

	// Resize ảnh
	resized := imaging.Resize(img, width, height, imaging.Lanczos)

	// Encode ảnh đã resize vào buffer
	buf := new(bytes.Buffer)
	ext := strings.ToLower(filepath.Ext(inputPath))

	var mimeType string
	switch ext {
	case ".jpg", ".jpeg":
		err = jpeg.Encode(buf, resized, nil)
		mimeType = "image/jpeg"
	case ".png":
		err = png.Encode(buf, resized)
		mimeType = "image/png"
	default:
		return "", errors.New("Định dạng ảnh không hỗ trợ: " + ext)
	}

	if err != nil {
		return "", errors.New("Không thể mã hóa ảnh đã resize: " + err.Error())
	}

	// Encode sang base64
	base64Str := base64.StdEncoding.EncodeToString(buf.Bytes())

	// Thêm prefix MIME để dùng trực tiếp
	return "data:" + mimeType + ";base64," + base64Str, nil
}

// ✅ Hàm resize ảnh từ base64
func ResizeImageFromBase64(base64Str string, width int, height int) (string, error) {
	// Loại bỏ prefix nếu có (data:image/png;base64,...)
	commaIdx := strings.Index(base64Str, ",")
	if commaIdx != -1 {
		base64Str = base64Str[commaIdx+1:]
	}

	decoded, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return "", errors.New("Không thể giải mã base64: " + err.Error())
	}

	reader := bytes.NewReader(decoded)

	img, format, err := image.Decode(reader)
	if err != nil {
		return "", errors.New("Không thể giải mã hình: " + err.Error())
	}

	resized := imaging.Resize(img, width, height, imaging.Lanczos)

	buf := new(bytes.Buffer)
	switch format {
	case "jpeg":
		err = jpeg.Encode(buf, resized, nil)
	case "png":
		err = png.Encode(buf, resized)
	default:
		return "", errors.New("Format ảnh không hỗ trợ: " + format)
	}

	if err != nil {
		return "", errors.New("Không thể mã hóa ảnh đã resize: " + err.Error())
	}

	// Encode lại sang base64
	base64Encoded := base64.StdEncoding.EncodeToString(buf.Bytes())

	// Thêm prefix cho base64 result
	var mimeType string
	if format == "jpeg" {
		mimeType = "image/jpeg"
	} else if format == "png" {
		mimeType = "image/png"
	}
	finalResult := "data:" + mimeType + ";base64," + base64Encoded

	return finalResult, nil
}

func EncodeImageToBase64(imagePath string) (string, error) {
	// Mở file ảnh
	file, err := os.Open(imagePath)
	if err != nil {
		return "", errors.New("Không thể mở file ảnh: " + err.Error())
	}
	defer file.Close()

	// Decode ảnh để biết định dạng
	img, format, err := image.Decode(file)
	if err != nil {
		return "", errors.New("Không thể decode ảnh: " + err.Error())
	}

	// Encode lại ảnh vào buffer
	buf := new(bytes.Buffer)
	switch strings.ToLower(format) {
	case "jpeg", "jpg":
		err = jpeg.Encode(buf, img, nil)
	case "png":
		err = png.Encode(buf, img)
	default:
		return "", errors.New("Định dạng ảnh không hỗ trợ: " + format)
	}

	if err != nil {
		return "", errors.New("Không thể encode ảnh: " + err.Error())
	}

	// Chuyển sang chuỗi base64
	base64Str := base64.StdEncoding.EncodeToString(buf.Bytes())

	// Xác định đúng MIME type
	var mimeType string
	switch strings.ToLower(format) {
	case "jpeg", "jpg":
		mimeType = "image/jpeg"
	case "png":
		mimeType = "image/png"
	}

	// Thêm prefix
	return "data:" + mimeType + ";base64," + base64Str, nil
}

type ImageFormat string

const (
	FormatJPEG ImageFormat = "jpeg"
	FormatPNG  ImageFormat = "png"
)

// ProcessImageBase64 xử lý ảnh base64: resize, nén chất lượng, đổi định dạng
func ProcessImageBase64(inputBase64 string, width, height uint, quality int, outputFormat ImageFormat) (string, error) {
	// Loại bỏ tiền tố nếu có
	if idx := strings.Index(inputBase64, ","); idx != -1 {
		inputBase64 = inputBase64[idx+1:]
	}

	// Decode base64 -> []byte
	imgData, err := base64.StdEncoding.DecodeString(inputBase64)
	if err != nil {
		return "", err
	}

	// Decode ảnh
	img, _, err := image.Decode(bytes.NewReader(imgData))
	if err != nil {
		return "", err
	}

	// Resize ảnh
	resizedImg := resize.Resize(width, height, img, resize.Lanczos3)

	// Encode lại ảnh đã resize với format + quality
	var buf bytes.Buffer
	switch outputFormat {
	case FormatJPEG:
		err = jpeg.Encode(&buf, resizedImg, &jpeg.Options{Quality: quality})
	case FormatPNG:
		err = png.Encode(&buf, resizedImg)
	default:
		return "", errors.New("unsupported output format")
	}
	if err != nil {
		return "", err
	}

	// Encode lại thành base64
	outputBase64 := base64.StdEncoding.EncodeToString(buf.Bytes())

	// Thêm tiền tố nếu muốn
	prefix := "data:image/" + string(outputFormat) + ";base64,"
	return prefix + outputBase64, nil
}
