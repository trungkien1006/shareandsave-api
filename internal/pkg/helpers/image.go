package helpers

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"mime"
	"os"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
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
	// Loại bỏ prefix
	if commaIdx := strings.Index(base64Str, ","); commaIdx != -1 {
		base64Str = base64Str[commaIdx+1:]
	}

	decoded, err := base64.RawStdEncoding.DecodeString(base64Str)
	if err != nil {
		return "", fmt.Errorf("Không thể giải mã base64: %w", err)
	}

	img, format, err := image.Decode(bytes.NewReader(decoded))
	if err != nil {
		return "", fmt.Errorf("Không thể giải mã hình: %w", err)
	}

	resized := imaging.Resize(img, width, height, imaging.Lanczos)

	buf := new(bytes.Buffer)
	switch format {
	case "jpeg", "jpg":
		err = jpeg.Encode(buf, resized, nil)
	case "png":
		err = png.Encode(buf, resized)
	default:
		return "", fmt.Errorf("Format ảnh không hỗ trợ: %s", format)
	}

	if err != nil {
		return "", fmt.Errorf("Không thể mã hóa ảnh đã resize: %w", err)
	}

	encoded := base64.StdEncoding.EncodeToString(buf.Bytes())

	return fmt.Sprintf("data:image/%s;base64,%s", format, encoded), nil
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

func ImageToBase64(imagePath string) (string, error) {
	// Đọc toàn bộ file ảnh
	data, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return "", fmt.Errorf("không thể đọc file ảnh: %w", err)
	}

	// Xác định loại MIME (vd: image/png, image/jpeg)
	ext := filepath.Ext(imagePath)
	mimeType := mime.TypeByExtension(ext)
	if mimeType == "" {
		mimeType = "application/octet-stream" // fallback
	}

	// Encode thành base64
	base64Str := base64.StdEncoding.EncodeToString(data)

	// Trả về dạng chuẩn `data:<mime>;base64,...`
	result := fmt.Sprintf("data:%s;base64,%s", mimeType, base64Str)

	return result, nil
}

// ProcessImageBase64 xử lý ảnh base64: resize, nén chất lượng, đổi định dạng
func ProcessImageBase64(inputBase64 string, width, height uint, quality int) (string, error) {
	// Nếu đã resize thì bỏ qua
	if strings.Contains(inputBase64, ";resized;") {
		return inputBase64, nil
	}

	// Tách prefix
	prefix := ""
	mimeType := ""
	if idx := strings.Index(inputBase64, ","); idx != -1 {
		prefix = inputBase64[:idx+1]
		inputBase64 = inputBase64[idx+1:]

		if strings.Contains(prefix, "image/jpeg") || strings.Contains(prefix, "image/jpg") {
			mimeType = "jpeg"
		} else if strings.Contains(prefix, "image/png") {
			mimeType = "png"
		} else {
			return "", errors.New("unsupported image format")
		}
	} else {
		return "", errors.New("invalid base64 image data")
	}

	// Decode base64
	imgData, err := base64.StdEncoding.DecodeString(inputBase64)
	if err != nil {
		return "", err
	}

	// Decode ảnh
	img, _, err := image.Decode(bytes.NewReader(imgData))
	if err != nil {
		return "", err
	}

	// Chuyển sang RGB để đảm bảo không bị lỗi encode JPEG
	rgbImg := ensureRGB(img)

	// Resize bằng imaging (Lanczos chất lượng cao)
	resizedImg := imaging.Resize(rgbImg, int(width), int(height), imaging.Lanczos)

	// Encode lại ảnh
	var buf bytes.Buffer
	switch mimeType {
	case "jpeg":
		err = jpeg.Encode(&buf, resizedImg, &jpeg.Options{Quality: quality})
	case "png":
		err = png.Encode(&buf, resizedImg)
	default:
		return "", errors.New("unsupported image format for encoding")
	}
	if err != nil {
		return "", err
	}

	// Encode base64
	encoded := base64.StdEncoding.EncodeToString(buf.Bytes())

	// Gắn lại prefix với dấu hiệu đã resize
	newPrefix := strings.Replace(prefix, "base64,", "resized;base64,", 1)
	return newPrefix + encoded, nil
}

func ensureRGB(img image.Image) image.Image {
	b := img.Bounds()
	rgbImg := image.NewRGBA(b)

	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			c := color.NRGBAModel.Convert(img.At(x, y)).(color.NRGBA)

			alpha := float64(c.A) / 255
			r := uint8(float64(c.R)*alpha + 255*(1-alpha))
			g := uint8(float64(c.G)*alpha + 255*(1-alpha))
			b := uint8(float64(c.B)*alpha + 255*(1-alpha))

			rgbImg.Set(x, y, color.RGBA{r, g, b, 255})
		}
	}
	return rgbImg
}
