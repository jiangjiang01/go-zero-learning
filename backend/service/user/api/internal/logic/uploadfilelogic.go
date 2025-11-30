// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"go-zero-learning/common/errorx"
	"go-zero-learning/service/user/api/internal/svc"
	"go-zero-learning/service/user/api/internal/types"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
)

type UploadFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadFileLogic {
	return &UploadFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 文件上传逻辑
// 注意：文件上传使用 multipart/form-data， 需要直接处理 http.Request
func (l *UploadFileLogic) UploadFile(r *http.Request) (resp *types.UploadFileResp, err error) {
	// 1. 解析 multipart form
	err = r.ParseMultipartForm(l.svcCtx.Config.Upload.MaxSize)
	if err != nil {
		l.Errorf("解析表单失败：%v", err)
		return nil, errorx.NewBusinessError(errorx.CodeInvalidParam, "文件过大或格式错误")
	}

	// 2. 获取文件
	file, header, err := r.FormFile("file")
	if err != nil {
		l.Errorf("获取文件失败：%v", err)
		return nil, errorx.NewBusinessError(errorx.CodeInvalidParam, "请选择要上传的文件")
	}
	defer file.Close()

	// 3. 验证文件大小
	if header.Size > l.svcCtx.Config.Upload.MaxSize {
		return nil, errorx.NewBusinessError(errorx.CodeInvalidParam,
			fmt.Sprintf("文件大小超过限制（最大 %dMB）", l.svcCtx.Config.Upload.MaxSize/1024/1024))
	}

	// 4. 验证文件类型
	contentType := header.Header.Get("Content-Type")
	if !l.isAllowedType(contentType) {
		return nil, errorx.NewBusinessError(errorx.CodeInvalidParam,
			fmt.Sprintf("不支持的文件类型：%s", contentType))
	}

	// 5. 获取文件分类（可选）
	category := strings.TrimSpace(r.FormValue("category"))
	if category == "" {
		category = "default" // 默认分类
	}

	// 6. 生成文件名（UUID + 原始扩展名）
	ext := filepath.Ext(header.Filename)
	filename := fmt.Sprintf("%s%s", uuid.New().String(), ext)

	// 7. 创建目录（按分类和日期分目录）
	now := time.Now()
	dateDir := now.Format("2006/01/02")
	uploadDir := filepath.Join(l.svcCtx.Config.Upload.Path, category, dateDir)
	err = os.MkdirAll(uploadDir, 0755)
	if err != nil {
		l.Errorf("创建目录失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 8. 保存文件
	filePath := filepath.Join(uploadDir, filename)
	dst, err := os.Create(filePath)
	if err != nil {
		l.Errorf("创建文件失败：%v", err)
		return nil, errorx.ErrInternalError
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		l.Errorf("保存文件失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 9. 构建访问 URL
	url := fmt.Sprintf("%s/%s/%s/%s", l.svcCtx.Config.Upload.BaseURL, category, dateDir, filename)

	// 10. 构建响应
	resp = &types.UploadFileResp{
		URL:      url,
		Filename: filename,
		Size:     header.Size,
	}

	return resp, nil
}

// 检查文件类型是否允许
func (l *UploadFileLogic) isAllowedType(contentType string) bool {
	for _, allowedType := range l.svcCtx.Config.Upload.AllowedTypes {
		if contentType == allowedType {
			return true
		}
	}
	return false
}
