package seeder

import (
	"context"
	"final_project/internal/domain/category"
	importinvoice "final_project/internal/domain/import_invoice"
	"final_project/internal/domain/item"
	"final_project/internal/domain/post"
	rolepermission "final_project/internal/domain/role_permission"
	"final_project/internal/domain/user"
	"final_project/internal/domain/warehouse"
	"final_project/internal/pkg/enums"
	"final_project/internal/pkg/hash"
	"final_project/internal/pkg/helpers"
	"fmt"
	"log"
	"os"
	"time"
)

type Seeder struct {
	rolePerRepo      rolepermission.Repository
	itemRepo         item.Repository
	userRepo         user.Repository
	categoryRepo     category.Repository
	postRepo         post.Repository
	postService      *post.PostService
	imInvoiceRepo    importinvoice.Repository
	imInvoiceService *importinvoice.ImportInvoiceService
}

func NewSeeder(rolePerRepo rolepermission.Repository, itemRepo item.Repository, userRepo user.Repository, categoryRepo category.Repository, postRepo post.Repository, postService *post.PostService, imInvoiceRepo importinvoice.Repository, imInvoiceService *importinvoice.ImportInvoiceService) *Seeder {
	return &Seeder{
		rolePerRepo:      rolePerRepo,
		itemRepo:         itemRepo,
		userRepo:         userRepo,
		categoryRepo:     categoryRepo,
		postRepo:         postRepo,
		postService:      postService,
		imInvoiceRepo:    imInvoiceRepo,
		imInvoiceService: imInvoiceService,
	}
}

func (s *Seeder) Seed() error {
	if err := s.seedPermission(); err != nil {
		return err
	}

	if err := s.seedRole(); err != nil {
		return err
	}

	if err := s.seedRolePer(); err != nil {
		return err
	}

	if err := s.seedCategory(); err != nil {
		return err
	}

	if err := s.seedItems(); err != nil {
		return err
	}

	if err := s.seedUsers(); err != nil {
		return err
	}

	if err := s.seedPosts(); err != nil {
		return err
	}

	if err := s.seedImportInvoice(); err != nil {
		return err
	}

	return nil
}

func (s *Seeder) seedPermission() error {
	ctx := context.Background()

	fmt.Println("Start seed permission...")

	isEmpty, err := s.rolePerRepo.IsPermissionTableEmpty(ctx)
	if err != nil {
		return err
	}
	if !isEmpty {
		fmt.Println("Permission had data...")
		return nil // Đã có dữ liệu, không seed nữa
	}

	var permissions = []rolepermission.Permission{
		//user management permissions
		{Name: "Create User", Code: "create_user"},
		{Name: "Read User", Code: "read_user"},
		{Name: "Update User", Code: "update_user"},
		{Name: "Delete User", Code: "delete_user"},

		//item management permissions
		{Name: "Create Item", Code: "create_item"},
		{Name: "Read Item", Code: "read_item"},
		{Name: "Update Item", Code: "update_item"},
		{Name: "Delete Item", Code: "delete_item"},

		//post management permissions
		{Name: "Create Post", Code: "create_post"},
		{Name: "Read Post", Code: "read_post"},
		{Name: "Update Post", Code: "update_post"},
		{Name: "Delete Post", Code: "delete_post"},

		//admin management permissions
		{Name: "Create Admin", Code: "create_admin"},
		{Name: "Read Admin", Code: "read_admin"},
		{Name: "Update Admin", Code: "update_admin"},
		{Name: "Delete Admin", Code: "delete_admin"},

		//notification management permissions
		{Name: "Create Notification", Code: "create_notification"},

		//import_invoice management permissions
		{Name: "Read Import Invoice", Code: "read_import_invoice"},
		{Name: "Create Import Invoice", Code: "create_import_invoice"},
		{Name: "Update Import Invoice", Code: "update_import_invoice"},
		{Name: "Lock Import Invoice", Code: "lock_import_invoice"},
		{Name: "Delete Import Invoice", Code: "delete_import_invoice"},

		//export_invoice management permissions
		{Name: "Read Export Invoice", Code: "read_export_invoice"},
		{Name: "Create Export Invoice", Code: "create_export_invoice"},
		{Name: "Update Export Invoice", Code: "update_export_invoice"},
		{Name: "Lock Export Invoice", Code: "lock_export_invoice"},
		{Name: "Delete Export Invoice", Code: "delete_export_invoice"},

		//warehouse management permissions
		{Name: "Read Warehouse", Code: "read_warehouse"},

		//item_warehouse management permissions
		{Name: "Read Item Warehouse", Code: "read_item_warehouse"},
	}

	if err := s.rolePerRepo.SavePermission(&permissions); err != nil {
		return err
	}

	fmt.Println("Finish seed permission...")

	return nil
}

func (s *Seeder) seedRole() error {
	ctx := context.Background()

	fmt.Println("Start seed role...")

	isEmpty, err := s.rolePerRepo.IsRoleTableEmpty(ctx)
	if err != nil {
		return err
	}
	if !isEmpty {
		fmt.Println("Role had data...")
		return nil
	}

	var roles = []rolepermission.Role{
		{Name: "Client"},
		{Name: "Super Admin"},
		{Name: "Content Manager"},
		{Name: "Warehouse Manager"},
		// {Name: "Human Resources Manager"},
		// {Name: "Client Manager"},
	}

	if err := s.rolePerRepo.SaveRole(&roles); err != nil {
		return err
	}

	fmt.Println("Finish seed role...")

	return nil
}

func (s *Seeder) seedRolePer() error {
	ctx := context.Background()

	fmt.Println("Start seed role permission...")

	isEmpty, err := s.rolePerRepo.IsRolePermissionTableEmpty(ctx)
	if err != nil {
		return err
	}
	if !isEmpty {
		fmt.Println("Role permission had data...")
		return nil
	}

	var rolePermissionConfig = map[string][]string{
		"Client":            {""},
		"Super Admin":       {"*"},
		"Content Manager":   {"create_post", "read_post", "update_post", "delete_post", "read_notification"},
		"Warehouse Manager": {"read_warehouse", "read_item_warehouse", "read_import_invoice", "create_import_invoice", "update_import_invoice", "lock_import_invoice", "delete_import_invoice", "read_export_invoice", "create_export_invoice", "update_export_invoice", "lock_export_invoice", "delete_export_invoice", "read_notification"},
		// "Human Resources Manager": {"read_admin", "update_admin", "delete_admin", "create_admin", "read_notification"},
		// "Client Manager":          {"read_user", "update_user", "delete_user", "read_notification", "read_request", "reply_request", "delete_request"},
	}

	var roles []rolepermission.Role
	var permissions []rolepermission.Permission
	var rolePerms []rolepermission.RolePermission

	if err := s.rolePerRepo.GetAllRoles(&roles); err != nil {
		return err
	}

	if err := s.rolePerRepo.GetAllPermission(&permissions); err != nil {
		return err
	}

	permCodeToID := make(map[string]uint)

	for _, p := range permissions {
		permCodeToID[p.Code] = p.ID
	}

	for _, role := range roles {
		codes, exists := rolePermissionConfig[role.Name]

		if !exists {
			continue // nếu role đó không có cấu hình thì bỏ qua
		}

		if len(codes) == 1 && codes[0] == "*" {
			for _, p := range permissions {
				rolePerms = append(rolePerms, rolepermission.RolePermission{
					RoleID:       role.ID,
					PermissionID: p.ID,
				})
			}
		} else {
			for _, code := range codes {
				if permID, ok := permCodeToID[code]; ok {
					rolePerms = append(rolePerms, rolepermission.RolePermission{
						RoleID:       role.ID,
						PermissionID: permID,
					})
				}
			}
		}
	}

	if err := s.rolePerRepo.SaveRolePermission(&rolePerms); err != nil {
		return err
	}

	fmt.Println("Finish seed role permission...")

	return nil
}

func (s *Seeder) seedCategory() error {
	ctx := context.Background()

	fmt.Println("Start seed category...")

	isEmpty, err := s.categoryRepo.IsTableEmpty(ctx)
	if err != nil {
		return err
	}

	if !isEmpty {
		fmt.Println("Category had data...")
		return nil
	}

	categorys := []category.Category{
		{Name: "Khác"},
		{Name: "Sách"},
		{Name: "Thời trang"},
		{Name: "Dụng cụ cơ khí"},
		{Name: "Vật dụng cá nhân"},
		{Name: "Giấy tờ tùy thân"},
		{Name: "Tài liệu học tập"},
	}

	for i := range categorys {
		if err := s.categoryRepo.Save(ctx, &categorys[i]); err != nil {
			return err
		}
	}

	fmt.Println("Finish seed category...")

	return nil
}

func (s *Seeder) seedItems() error {
	ctx := context.Background()

	fmt.Println("Start seed items...")

	isEmpty, err := s.itemRepo.IsTableEmpty(ctx)
	if err != nil {
		return err
	}
	if !isEmpty {
		fmt.Println("Item had data...")
		return nil
	}

	var categories []category.Category

	// Lấy tất cả categories
	err = s.categoryRepo.GetAllCategories(ctx, &categories)

	fmt.Println("Số category lấy được:", len(categories))

	if err != nil {
		return err
	}

	// Map tên category -> ID
	categoryMap := make(map[string]uint)
	for _, c := range categories {
		categoryMap[c.Name] = c.ID
	}

	// Seed item với category cụ thể
	items := []item.Item{
		{Name: "Giáo trình Chính trị 1", Description: "Giáo trình học phần chính trị cơ bản cho sinh viên năm nhất.", CategoryID: categoryMap["Sách"]},
		{Name: "Giáo trình Thiết kế Web", Description: "Giáo trình dạy về HTML, CSS, và JavaScript cơ bản.", CategoryID: categoryMap["Tài liệu học tập"]},
		{Name: "Sách Toán Cao Cấp", Description: "Tài liệu học Toán cao cấp cho khối kỹ thuật.", CategoryID: categoryMap["Sách"]},
		{Name: "Giáo trình Cấu trúc Dữ liệu và Giải thuật", Description: "Tài liệu học thuật về các thuật toán và cấu trúc dữ liệu cơ bản.", CategoryID: categoryMap["Tài liệu học tập"]},
		{Name: "Máy tính Casio fx-570VN Plus", Description: "Máy tính cầm tay cần thiết cho các môn học tự nhiên và kỹ thuật.", CategoryID: categoryMap["Vật dụng cá nhân"]},
		{Name: "Balo Laptop Chống Sốc", Description: "Balo chống sốc bảo vệ laptop và tài liệu học tập.", CategoryID: categoryMap["Vật dụng cá nhân"]},
		{Name: "Giáo trình Kinh tế Vi mô", Description: "Tài liệu học môn Kinh tế Vi mô cho sinh viên Kinh tế.", CategoryID: categoryMap["Sách"]},
		{Name: "Bút dạ quang", Description: "Bộ 5 cây bút dạ quang nhiều màu để ghi chú nội dung quan trọng.", CategoryID: categoryMap["Vật dụng cá nhân"]},
		{Name: "Giáo trình Nguyên lý Kế toán", Description: "Sách học phần cơ bản cho ngành kế toán.", CategoryID: categoryMap["Sách"]},
		{Name: "Tập vở ô ly A4", Description: "Tập vở cũ còn dư dùng để ghi chép bài giảng.", CategoryID: categoryMap["Tài liệu học tập"]},
		{Name: "Sách Hóa học Đại cương", Description: "Tài liệu học Hóa học cơ bản dành cho sinh viên khoa tự nhiên.", CategoryID: categoryMap["Sách"]},
		{Name: "Thước kẻ 30cm", Description: "Dụng cụ học tập thường dùng trong các bài toán vẽ hình hoặc kỹ thuật.", CategoryID: categoryMap["Dụng cụ cơ khí"]},
		{Name: "Giáo trình Lập trình C C++", Description: "Sách học về ngôn ngữ lập trình C/C++ từ cơ bản đến nâng cao.", CategoryID: categoryMap["Sách"]},
		{Name: "Sách Giáo dục Quốc phòng", Description: "Tài liệu học quốc phòng an ninh cho sinh viên năm nhất.", CategoryID: categoryMap["Sách"]},
		{Name: "Laptop cũ Dell Vostro", Description: "Máy tính cũ dành cho sinh viên học tập và làm bài tập nhóm.", CategoryID: categoryMap["Vật dụng cá nhân"]},
		{Name: "Tai nghe chụp tai", Description: "Tai nghe phục vụ học online, nghe giảng, và làm việc nhóm.", CategoryID: categoryMap["Vật dụng cá nhân"]},
		{Name: "Giáo trình Tâm lý học đại cương", Description: "Sách học phần Tâm lý học cho khối xã hội.", CategoryID: categoryMap["Sách"]},
		{Name: "Chuột không dây Logitech", Description: "Thiết bị hỗ trợ học tập và làm việc với máy tính.", CategoryID: categoryMap["Vật dụng cá nhân"]},
		{Name: "Giáo trình Xác suất Thống kê", Description: "Tài liệu dành cho các ngành kỹ thuật và kinh tế.", CategoryID: categoryMap["Tài liệu học tập"]},
		{Name: "Sách Nhập môn Công nghệ Thông tin", Description: "Tài liệu học cơ sở ngành CNTT cho sinh viên năm nhất.", CategoryID: categoryMap["Sách"]},
		{Name: "Đèn học LED để bàn", Description: "Đèn bàn dùng học bài, tiết kiệm điện và dịu mắt.", CategoryID: categoryMap["Vật dụng cá nhân"]},
		{Name: "Ổ cứng di động 500GB", Description: "Dùng để lưu trữ tài liệu học tập, phần mềm, bài tập nhóm.", CategoryID: categoryMap["Vật dụng cá nhân"]},
		{Name: "Giáo trình Pháp luật đại cương", Description: "Tài liệu học bắt buộc cho tất cả các ngành.", CategoryID: categoryMap["Sách"]},
		{Name: "Thẻ nhớ SD 64GB", Description: "Thiết bị lưu trữ dành cho sinh viên học media, CNTT.", CategoryID: categoryMap["Vật dụng cá nhân"]},
		{Name: "Giáo trình Xây dựng Đảng", Description: "Tài liệu học chính trị dành cho sinh viên năm cuối.", CategoryID: categoryMap["Sách"]},
		{Name: "Bút bi Thiên Long", Description: "Bộ bút bi dư dùng để ghi chép bài vở.", CategoryID: categoryMap["Vật dụng cá nhân"]},
		{Name: "Bìa hồ sơ và túi clear A4", Description: "Dùng để nộp bài, tài liệu, báo cáo môn học.", CategoryID: categoryMap["Tài liệu học tập"]},
		{Name: "Máy in cũ Canon", Description: "Máy in cũ hỗ trợ in báo cáo, tài liệu học tập tại nhà.", CategoryID: categoryMap["Vật dụng cá nhân"]},
		{Name: "Ghế gấp học bài", Description: "Ghế gấp gọn, dùng khi ngồi học trong phòng trọ nhỏ.", CategoryID: categoryMap["Vật dụng cá nhân"]},
		{Name: "Bảng trắng mini", Description: "Bảng trắng dùng ghi chú, trình bày khi học nhóm.", CategoryID: categoryMap["Tài liệu học tập"]},
	}

	for i := range items {
		strBase64Image, err := helpers.ResizeImageFromFileToBase64(os.Getenv("IMAGE_PATH")+"/item.png", enums.ItemImageWidth, enums.ItemImageHeight)
		if err != nil {
			return err
		}

		items[i].Image = strBase64Image

		if err := s.itemRepo.Save(ctx, &items[i]); err != nil {
			return err
		}
	}

	fmt.Println("Finish seed items...")

	return nil
}

func (s *Seeder) seedUsers() error {
	ctx := context.Background()

	fmt.Println("Start seed users...")

	// Kiểm tra bảng user có rỗng không
	isEmpty, err := s.userRepo.IsTableEmpty(ctx)
	if err != nil {
		return err
	}
	if !isEmpty {
		fmt.Println("User had data...")
		return nil
	}

	// Lấy tất cả role từ DB
	var roles []rolepermission.Role
	if err := s.rolePerRepo.GetAllRoles(&roles); err != nil {
		return err
	}

	// Map role name -> ID để dễ gán
	roleMap := make(map[string]uint)
	for _, r := range roles {
		roleMap[r.Name] = r.ID
	}

	// 3 admin user với 3 role admin khác nhau
	adminUsers := []user.User{
		{
			Email:       "superadmin@example.com",
			Password:    "Admin1234",
			FullName:    "Super Admin",
			Status:      1,
			RoleID:      roleMap["Super Admin"],
			PhoneNumber: "0900000001",
			Address:     "Hà Nội",
			GoodPoint:   100,
			Major:       "Quản trị",
			Active:      true,
		},
		{
			Email:       "content@example.com",
			Password:    "Admin1234",
			FullName:    "Content Manager",
			Status:      1,
			RoleID:      roleMap["Content Manager"],
			PhoneNumber: "0900000002",
			Address:     "Hồ Chí Minh",
			GoodPoint:   100,
			Major:       "Quản trị",
			Active:      true,
		},
		{
			Email:       "warehouse@example.com",
			Password:    "Admin1234",
			FullName:    "Warehouse Manager",
			Status:      1,
			RoleID:      roleMap["Warehouse Manager"],
			PhoneNumber: "0900000003",
			Address:     "Đà Nẵng",
			GoodPoint:   100,
			Major:       "Quản trị",
			Active:      true,
		},
	}

	// 27 client user
	clientUsers := make([]user.User, 27)
	for i := 0; i < 27; i++ {
		clientUsers[i] = user.User{
			Email:       fmt.Sprintf("client%02d@example.com", i+1),
			Password:    "User1234",
			FullName:    fmt.Sprintf("Client User %02d", i+1),
			Status:      1,
			RoleID:      roleMap["Client"],
			PhoneNumber: fmt.Sprintf("090000%04d", i+4),
			Address:     "Khách hàng",
			GoodPoint:   i % 10,
			Major:       "Khách hàng",
			Active:      true,
		}
	}

	// Gộp admin và client lại
	users := append(adminUsers, clientUsers...)

	// Lưu tất cả user vào DB
	for _, u := range users {
		hashedPassword, err := hash.HashPassword(u.Password)

		if err != nil {
			return err
		}

		strBase64Image, err := helpers.ResizeImageFromFileToBase64(os.Getenv("IMAGE_PATH")+"/user.png", enums.UserImageWidth, enums.UserImageHeight)

		if err != nil {
			return err
		}

		u.Password = hashedPassword
		u.Avatar = strBase64Image

		if err := s.userRepo.Save(ctx, &u); err != nil {
			return err
		}
	}

	fmt.Println("Finish seed users...")

	return nil
}

func (s *Seeder) seedPosts() error {
	ctx := context.Background()

	fmt.Println("Start seed posts...")

	// Kiểm tra bảng post có rỗng không
	isEmpty, err := s.postRepo.IsTableEmpty(ctx)
	if err != nil {
		return err
	}
	if !isEmpty {
		fmt.Println("Post had data...")
		return nil
	}

	// Generate post default base64 image
	base64, err := helpers.ImageToBase64(os.Getenv("IMAGE_PATH") + "/post.png")
	if err != nil {
		return err
	}

	postDefaultImage, err := helpers.ProcessImageBase64(base64, uint(enums.ItemImageWidth), uint(enums.ItemImageHeight), 75, helpers.FormatJPEG)
	if err != nil {
		return err
	}

	// Generate item default base64 image
	base64, err = helpers.ImageToBase64(os.Getenv("IMAGE_PATH") + "/item.png")
	if err != nil {
		return err
	}

	itemDefaultImage, err := helpers.ProcessImageBase64(base64, uint(enums.ItemImageWidth), uint(enums.ItemImageHeight), 75, helpers.FormatJPEG)
	if err != nil {
		return err
	}

	authorMap := map[string]uint{
		"superadmin": 1,
		"client1":    4,
		"client2":    5,
		"client3":    6,
	}

	oldItemInPost := map[int][]post.OldItemsInPost{
		0: {
			{
				ItemID:   1,
				Image:    itemDefaultImage,
				Quantity: 1,
			},
		},
		1: {
			{
				ItemID:   4,
				Image:    itemDefaultImage,
				Quantity: 2,
			},
			{
				ItemID:   5,
				Image:    itemDefaultImage,
				Quantity: 1,
			},
			{
				ItemID:   6,
				Image:    itemDefaultImage,
				Quantity: 3,
			},
		},
		2: {
			{
				ItemID:   7,
				Image:    itemDefaultImage,
				Quantity: 2,
			},
			{
				ItemID:   8,
				Image:    itemDefaultImage,
				Quantity: 1,
			},
			{
				ItemID:   9,
				Image:    itemDefaultImage,
				Quantity: 3,
			},
		},
	}

	posts := []post.CreatePost{

		{
			AuthorID:    authorMap["superadmin"],
			AuthorName:  "Superadmin",
			Type:        int(enums.PostTypeGiveAwayOldItem),
			Slug:        s.postService.GenerateSlug("Bài viết số 1"),
			Title:       "Bài viết số 1",
			Content:     "[]",
			Info:        `{}`,
			Description: "Mô tả về bài đăng plaplaploplo...",
			Status:      int8(enums.PostStatusPending),
			Images:      []string{postDefaultImage},
			Tag: []string{
				"Vật dụng cá nhân",
			},
			Items:    []item.Item{},
			OldItems: oldItemInPost[0],
			NewItems: []post.NewItemsInPost{},
		},

		{
			AuthorID:    authorMap["client3"],
			AuthorName:  "Client3",
			Type:        int(enums.PostTypeGiveAwayOldItem),
			Slug:        s.postService.GenerateSlug("Bài viết số 2"),
			Title:       "Bài viết số 2",
			Content:     "[]",
			Info:        `{}`,
			Description: "Mô tả về bài đăng plaplaploplo...",
			Status:      int8(enums.PostStatusPending),
			Images:      []string{postDefaultImage},
			Tag: []string{
				"Đồ dùng học tập",
			},
			Items:    []item.Item{},
			OldItems: oldItemInPost[1],
			NewItems: []post.NewItemsInPost{},
		},

		{
			AuthorID:   authorMap["superadmin"],
			AuthorName: "Superadmin",
			Type:       int(enums.PostTypeSeekLoseItem),
			Slug:       s.postService.GenerateSlug("Bài viết số 3"),
			Title:      "Bài viết số 3",
			Content:    "",
			Info: `{
				"lostDate": "2025-06-05 12:21:09.055600",
				"lostLocation": "Địa điểm 3",
				"reward": "100",
				"category": "Ví"
			}`,
			Description: "Mô tả về bài đăng plaplaploplo...",
			Status:      int8(enums.PostStatusPending),
			Images:      []string{postDefaultImage},
			Tag: []string{
				"Ví",
			},
			Items:    []item.Item{},
			OldItems: []post.OldItemsInPost{},
			NewItems: []post.NewItemsInPost{},
		},

		{
			AuthorID:   authorMap["superadmin"],
			AuthorName: "Superadmin",
			Type:       int(enums.PostTypeSeekLoseItem),
			Slug:       s.postService.GenerateSlug("Bài viết số 4"),
			Title:      "Bài viết số 4",
			Content:    "",
			Info: `{
				"lostDate": "2025-06-05 12:21:09.055636",
				"lostLocation": "Địa điểm 4",
				"reward": "100",
				"category": "Sách"
			}`,
			Description: "Mô tả về bài đăng plaplaploplo...",
			Status:      int8(enums.PostStatusPending),
			Images:      []string{postDefaultImage},
			Tag: []string{
				"Sách",
				"Tài liệu học tập",
			},
			Items:    []item.Item{},
			OldItems: []post.OldItemsInPost{},
			NewItems: []post.NewItemsInPost{},
		},

		{
			AuthorID:    authorMap["client3"],
			AuthorName:  "Client3",
			Type:        int(enums.PostTypeGiveAwayOldItem),
			Slug:        s.postService.GenerateSlug("Bài viết số 5"),
			Title:       "Bài viết số 5",
			Content:     "[]",
			Info:        `{}`,
			Description: "Mô tả về bài đăng plaplaploplo...",
			Status:      int8(enums.PostStatusPending),
			Images:      []string{postDefaultImage},
			Tag: []string{
				"Đồ dùng học tập",
			},
			Items:    []item.Item{},
			OldItems: oldItemInPost[1],
			NewItems: []post.NewItemsInPost{},
		},

		{
			AuthorID:   authorMap["client2"],
			AuthorName: "Client2",
			Type:       int(enums.PostTypeSeekLoseItem),
			Slug:       s.postService.GenerateSlug("Bài viết số 6"),
			Title:      "Bài viết số 6",
			Content:    "",
			Info: `{
				"lostDate": "2025-06-05 12:21:09.055672",
				"lostLocation": "Địa điểm 6",
				"reward": "100",
				"category": "Thất lạc"
			}`,
			Description: "Mô tả về bài đăng plaplaploplo...",
			Status:      int8(enums.PostStatusPending),
			Images:      []string{postDefaultImage},
			Tag: []string{
				"Thất lạc",
			},
			Items:    []item.Item{},
			OldItems: []post.OldItemsInPost{},
			NewItems: []post.NewItemsInPost{},
		},

		{
			AuthorID:   authorMap["client2"],
			AuthorName: "Client2",
			Type:       int(enums.PostTypeFoundItem),
			Slug:       s.postService.GenerateSlug("Bài viết số 7"),
			Title:      "Bài viết số 7",
			Content:    "",
			Info: `{
				"foundDate": "2025-06-05 12:21:09.055690",
				"foundLocation": "Địa điểm 7"
			}`,
			Description: "Mô tả về bài đăng plaplaploplo...",
			Status:      int8(enums.PostStatusPending),
			Images:      []string{postDefaultImage},
			Tag: []string{
				"Balo",
			},
			Items:    []item.Item{},
			OldItems: oldItemInPost[0],
			NewItems: []post.NewItemsInPost{},
		},

		{
			AuthorID:    authorMap["superadmin"],
			AuthorName:  "Superadmin",
			Type:        int(enums.PostTypeGiveAwayOldItem),
			Slug:        s.postService.GenerateSlug("Bài viết số 8"),
			Title:       "Bài viết số 8",
			Content:     "[]",
			Info:        `{}`,
			Description: "Mô tả về bài đăng plaplaploplo...",
			Status:      int8(enums.PostStatusPending),
			Images:      []string{postDefaultImage},
			Tag: []string{
				"Balo",
			},
			Items:    []item.Item{},
			OldItems: oldItemInPost[1],
			NewItems: []post.NewItemsInPost{},
		},

		{
			AuthorID:   authorMap["superadmin"],
			AuthorName: "Superadmin",
			Type:       int(enums.PostTypeSeekLoseItem),
			Slug:       s.postService.GenerateSlug("Bài viết số 9"),
			Title:      "Bài viết số 9",
			Content:    "",
			Info: `{
				"lostDate": "2025-06-05 12:21:09.055731",
				"lostLocation": "Địa điểm 9",
				"reward": "100",
				"category": "Sách"
			}`,
			Description: "Mô tả về bài đăng plaplaploplo...",
			Status:      int8(enums.PostStatusPending),
			Images:      []string{postDefaultImage},
			Tag: []string{
				"Sách",
				"Tài liệu học tập",
			},
			Items:    []item.Item{},
			OldItems: []post.OldItemsInPost{},
			NewItems: []post.NewItemsInPost{},
		},

		{
			AuthorID:   authorMap["superadmin"],
			AuthorName: "Superadmin",
			Type:       int(enums.PostTypeFoundItem),
			Slug:       s.postService.GenerateSlug("Bài viết số 10"),
			Title:      "Bài viết số 10",
			Content:    "",
			Info: `{
				"foundDate": "2025-06-05 12:21:09.055756",
				"foundLocation": "Địa điểm 10"
			}`,
			Description: "Mô tả về bài đăng plaplaploplo...",
			Status:      int8(enums.PostStatusPending),
			Images:      []string{postDefaultImage},
			Tag: []string{
				"Đồ dùng học tập",
			},
			Items:    []item.Item{},
			OldItems: oldItemInPost[0],
			NewItems: []post.NewItemsInPost{},
		},

		{
			AuthorID:   authorMap["client1"],
			AuthorName: "Client1",
			Type:       int(enums.PostTypeSeekLoseItem),
			Slug:       s.postService.GenerateSlug("Bài viết số 11"),
			Title:      "Bài viết số 11",
			Content:    "",
			Info: `{
				"lostDate": "2025-06-05 12:21:09.055776",
				"lostLocation": "Địa điểm 11",
				"reward": "100",
				"category": "Túi xách"
			}`,
			Description: "Mô tả về bài đăng plaplaploplo...",
			Status:      int8(enums.PostStatusPending),
			Images:      []string{postDefaultImage},
			Tag: []string{
				"Túi xách",
			},
			Items:    []item.Item{},
			OldItems: []post.OldItemsInPost{},
			NewItems: []post.NewItemsInPost{},
		},

		{
			AuthorID:   authorMap["client2"],
			AuthorName: "Client2",
			Type:       int(enums.PostTypeSeekLoseItem),
			Slug:       s.postService.GenerateSlug("Bài viết số 12"),
			Title:      "Bài viết số 12",
			Content:    "",
			Info: `{
				"lostDate": "2025-06-05 12:21:09.055800",
				"lostLocation": "Địa điểm 12",
				"reward": "100",
				"category": "Điện thoại"
			}`,
			Description: "Mô tả về bài đăng plaplaploplo...",
			Status:      int8(enums.PostStatusPending),
			Images:      []string{postDefaultImage},
			Tag: []string{
				"Điện thoại",
			},
			Items:    []item.Item{},
			OldItems: []post.OldItemsInPost{},
			NewItems: []post.NewItemsInPost{},
		},

		{
			AuthorID:    authorMap["client2"],
			AuthorName:  "Client2",
			Type:        int(enums.PostTypeGiveAwayOldItem),
			Slug:        s.postService.GenerateSlug("Bài viết số 13"),
			Title:       "Bài viết số 13",
			Content:     "[]",
			Info:        `{}`,
			Description: "Mô tả về bài đăng plaplaploplo...",
			Status:      int8(enums.PostStatusPending),
			Images:      []string{postDefaultImage},
			Tag: []string{
				"Áo khoác",
			},
			Items:    []item.Item{},
			OldItems: oldItemInPost[0],
			NewItems: []post.NewItemsInPost{},
		},

		{
			AuthorID:    authorMap["superadmin"],
			AuthorName:  "Superadmin",
			Type:        int(enums.PostTypeGiveAwayOldItem),
			Slug:        s.postService.GenerateSlug("Bài viết số 14"),
			Title:       "Bài viết số 14",
			Content:     "[]",
			Info:        `{}`,
			Description: "Mô tả về bài đăng plaplaploplo...",
			Status:      int8(enums.PostStatusPending),
			Images:      []string{postDefaultImage},
			Tag: []string{
				"Vật dụng cá nhân",
			},
			Items:    []item.Item{},
			OldItems: oldItemInPost[1],
			NewItems: []post.NewItemsInPost{},
		},

		{
			AuthorID:   authorMap["client2"],
			AuthorName: "Client2",
			Type:       int(enums.PostTypeSeekLoseItem),
			Slug:       s.postService.GenerateSlug("Bài viết số 15"),
			Title:      "Bài viết số 15",
			Content:    "",
			Info: `{
				"lostDate": "2025-06-05 12:21:09.055857",
				"lostLocation": "Địa điểm 15",
				"reward": "100",
				"category": "Tai nghe"
			}`,
			Description: "Mô tả về bài đăng plaplaploplo...",
			Status:      int8(enums.PostStatusPending),
			Images:      []string{postDefaultImage},
			Tag: []string{
				"Tai nghe",
			},
			Items:    []item.Item{},
			OldItems: []post.OldItemsInPost{},
			NewItems: []post.NewItemsInPost{},
		},

		{
			AuthorID:    authorMap["client1"],
			AuthorName:  "Client1",
			Type:        int(enums.PostTypeGiveAwayOldItem),
			Slug:        s.postService.GenerateSlug("Bài viết số 16"),
			Title:       "Bài viết số 16",
			Content:     "[]",
			Info:        `{}`,
			Description: "Mô tả về bài đăng plaplaploplo...",
			Status:      int8(enums.PostStatusPending),
			Images:      []string{postDefaultImage},
			Tag: []string{
				"Tai nghe",
			},
			Items:    []item.Item{},
			OldItems: oldItemInPost[0],
			NewItems: []post.NewItemsInPost{},
		},

		{
			AuthorID:    authorMap["client3"],
			AuthorName:  "Client3",
			Type:        int(enums.PostTypeGiveAwayOldItem),
			Slug:        s.postService.GenerateSlug("Bài viết số 17"),
			Title:       "Bài viết số 17",
			Content:     "[]",
			Info:        `{}`,
			Description: "Mô tả về bài đăng plaplaploplo...",
			Status:      int8(enums.PostStatusPending),
			Images:      []string{postDefaultImage},
			Tag: []string{
				"Balo",
			},
			Items:    []item.Item{},
			OldItems: oldItemInPost[1],
			NewItems: []post.NewItemsInPost{},
		},

		{
			AuthorID:   authorMap["client1"],
			AuthorName: "Client1",
			Type:       int(enums.PostTypeSeekLoseItem),
			Slug:       s.postService.GenerateSlug("Bài viết số 18"),
			Title:      "Bài viết số 18",
			Content:    "",
			Info: `{
				"lostDate": "2025-06-05 12:21:09.055914",
				"lostLocation": "Địa điểm 18",
				"reward": "100",
				"category": "Đồ dùng học tập"
			}`,
			Description: "Mô tả về bài đăng plaplaploplo...",
			Status:      int8(enums.PostStatusPending),
			Images:      []string{postDefaultImage},
			Tag: []string{
				"Đồ dùng học tập",
			},
			Items:    []item.Item{},
			OldItems: []post.OldItemsInPost{},
			NewItems: []post.NewItemsInPost{},
		},

		{
			AuthorID:   authorMap["superadmin"],
			AuthorName: "Superadmin",
			Type:       int(enums.PostTypeSeekLoseItem),
			Slug:       s.postService.GenerateSlug("Bài viết số 19"),
			Title:      "Bài viết số 19",
			Content:    "",
			Info: `{
				"lostDate": "2025-06-05 12:21:09.055939",
				"lostLocation": "Địa điểm 19",
				"reward": "100",
				"category": "Ví"
			}`,
			Description: "Mô tả về bài đăng plaplaploplo...",
			Status:      int8(enums.PostStatusPending),
			Images:      []string{postDefaultImage},
			Tag: []string{
				"Ví",
			},
			Items:    []item.Item{},
			OldItems: []post.OldItemsInPost{},
			NewItems: []post.NewItemsInPost{},
		},

		{
			AuthorID:   authorMap["client2"],
			AuthorName: "Client2",
			Type:       int(enums.PostTypeFoundItem),
			Slug:       s.postService.GenerateSlug("Bài viết số 20"),
			Title:      "Bài viết số 20",
			Content:    "",
			Info: `{
				"foundDate": "2025-06-05 12:21:09.055963",
				"foundLocation": "Địa điểm 20"
			}`,
			Description: "Mô tả về bài đăng plaplaploplo...",
			Status:      int8(enums.PostStatusPending),
			Images:      []string{postDefaultImage},
			Tag: []string{
				"Vật dụng cá nhân",
			},
			Items:    []item.Item{},
			OldItems: oldItemInPost[1],
			NewItems: []post.NewItemsInPost{},
		},

		{
			AuthorID:   authorMap["client1"],
			AuthorName: "Client1",
			Type:       int(enums.PostTypeFoundItem),
			Slug:       s.postService.GenerateSlug("Bài viết số 21"),
			Title:      "Bài viết số 21",
			Content:    "",
			Info: `{
				"foundDate": "2025-06-05 12:21:09.055986",
				"foundLocation": "Địa điểm 21"
			}`,
			Description: "Mô tả về bài đăng plaplaploplo...",
			Status:      int8(enums.PostStatusPending),
			Images:      []string{postDefaultImage},
			Tag: []string{
				"Ví",
			},
			Items:    []item.Item{},
			OldItems: oldItemInPost[2],
			NewItems: []post.NewItemsInPost{},
		},

		{
			AuthorID:    authorMap["client2"],
			AuthorName:  "Client2",
			Type:        int(enums.PostTypeGiveAwayOldItem),
			Slug:        s.postService.GenerateSlug("Bài viết số 22"),
			Title:       "Bài viết số 22",
			Content:     "[]",
			Info:        `{}`,
			Description: "Mô tả về bài đăng plaplaploplo...",
			Status:      int8(enums.PostStatusPending),
			Images:      []string{postDefaultImage},
			Tag: []string{
				"Áo khoác",
			},
			Items:    []item.Item{},
			OldItems: oldItemInPost[0],
			NewItems: []post.NewItemsInPost{},
		},

		{
			AuthorID:   authorMap["client3"],
			AuthorName: "Client3",
			Type:       int(enums.PostTypeSeekLoseItem),
			Slug:       s.postService.GenerateSlug("Bài viết số 23"),
			Title:      "Bài viết số 23",
			Content:    "",
			Info: `{
				"lostDate": "2025-06-05 12:21:09.056026",
				"lostLocation": "Địa điểm 23",
				"reward": "100",
				"category": "Đồ dùng học tập"
			}`,
			Description: "Mô tả về bài đăng plaplaploplo...",
			Status:      int8(enums.PostStatusPending),
			Images:      []string{postDefaultImage},
			Tag: []string{
				"Đồ dùng học tập",
			},
			Items:    []item.Item{},
			OldItems: []post.OldItemsInPost{},
			NewItems: []post.NewItemsInPost{},
		},

		{
			AuthorID:   authorMap["client3"],
			AuthorName: "Client3",
			Type:       int(enums.PostTypeSeekLoseItem),
			Slug:       s.postService.GenerateSlug("Bài viết số 24"),
			Title:      "Bài viết số 24",
			Content:    "",
			Info: `{
				"lostDate": "2025-06-05 12:21:09.056052",
				"lostLocation": "Địa điểm 24",
				"reward": "100",
				"category": "Đồ dùng học tập"
			}`,
			Description: "Mô tả về bài đăng plaplaploplo...",
			Status:      int8(enums.PostStatusPending),
			Images:      []string{postDefaultImage},
			Tag: []string{
				"Đồ dùng học tập",
			},
			Items:    []item.Item{},
			OldItems: []post.OldItemsInPost{},
			NewItems: []post.NewItemsInPost{},
		},

		{
			AuthorID:    authorMap["client1"],
			AuthorName:  "Client1",
			Type:        int(enums.PostTypeGiveAwayOldItem),
			Slug:        s.postService.GenerateSlug("Bài viết số 25"),
			Title:       "Bài viết số 25",
			Content:     "[]",
			Info:        `{}`,
			Description: "Mô tả về bài đăng plaplaploplo...",
			Status:      int8(enums.PostStatusPending),
			Images:      []string{postDefaultImage},
			Tag: []string{
				"Áo khoác",
			},
			Items:    []item.Item{},
			OldItems: oldItemInPost[0],
			NewItems: []post.NewItemsInPost{},
		},

		{
			AuthorID:   authorMap["superadmin"],
			AuthorName: "Superadmin",
			Type:       int(enums.PostTypeSeekLoseItem),
			Slug:       s.postService.GenerateSlug("Bài viết số 26"),
			Title:      "Bài viết số 26",
			Content:    "",
			Info: `{
				"lostDate": "2025-06-05 12:21:09.056096",
				"lostLocation": "Địa điểm 26",
				"reward": "100",
				"category": "Vật dụng cá nhân"
			}`,
			Description: "Mô tả về bài đăng plaplaploplo...",
			Status:      int8(enums.PostStatusPending),
			Images:      []string{postDefaultImage},
			Tag: []string{
				"Vật dụng cá nhân",
			},
			Items:    []item.Item{},
			OldItems: []post.OldItemsInPost{},
			NewItems: []post.NewItemsInPost{},
		},

		{
			AuthorID:   authorMap["client1"],
			AuthorName: "Client 1",
			Type:       int(enums.PostTypeOther),
			Slug:       s.postService.GenerateSlug("Bài viết khác"),
			Title:      "Bài viết khác",
			Content:    "",
			Info: `{
			}`,
			Description: "Mô tả về bài đăng plaplaploplo...",
			Status:      int8(enums.PostStatusPending),
			Images:      []string{postDefaultImage},
			Tag:         []string{},
			Items:       []item.Item{},
			OldItems:    []post.OldItemsInPost{},
			NewItems:    []post.NewItemsInPost{},
		},
	}

	for key, value := range posts {
		if value.Info != "{}" {
			content, _ := s.postService.GenerateContent(value.Info)

			posts[key].Content = content
		} else {
			posts[key].Content = "[]"
		}

		err := s.postRepo.Save(ctx, &posts[key])
		if err != nil {
			fmt.Println("Seed post error...")
		}
	}

	fmt.Println("Finish seed post...")

	return nil
}

func (s *Seeder) seedImportInvoice() error {
	ctx := context.Background()
	log.Println("Start seeding import invoices...")

	// Kiểm tra bảng ImportInvoice có trống không
	isEmpty, err := s.imInvoiceRepo.IsTableEmpty(ctx)
	if err != nil {
		return fmt.Errorf("failed to check if import invoice table is empty: %w", err)
	}

	if !isEmpty {
		log.Println("Import invoice table already has data, skipping seed...")
		return nil
	}

	// Dữ liệu mẫu cho ImportInvoice
	importInvoices := []importinvoice.ImportInvoice{
		{
			SenderID:     1,
			SenderName:   "Nguyen Van A",
			ReceiverID:   2,
			ReceiverName: "Tran Thi B",
			Classify:     1,
			Description:  "Phiếu nhập sách giáo khoa",
			ItemImportInvoice: []importinvoice.ItemImportInvoice{
				{
					ItemID:      1,
					ItemName:    "",
					Quantity:    10,
					Description: "Sách giáo khoa Toán lớp 12",
				},
				{
					ItemID:      2,
					ItemName:    "",
					Quantity:    5,
					Description: "Sách giáo khoa Văn lớp 12",
				},
			},
			ItemCount: 15,
			CreatedAt: time.Now(),
		},
		{
			SenderID:     3,
			SenderName:   "Le Van C",
			ReceiverID:   4,
			ReceiverName: "Pham Thi D",
			Classify:     1,
			Description:  "Phiếu nhập quần áo mùa đông",
			ItemImportInvoice: []importinvoice.ItemImportInvoice{
				{
					ItemID:      3,
					ItemName:    "",
					Quantity:    20,
					Description: "Áo khoác nam size M",
				},
			},
			ItemCount: 20,
			CreatedAt: time.Now().Add(-1 * time.Hour),
		},
		{
			SenderID:     5,
			SenderName:   "Hoang Van E",
			ReceiverID:   6,
			ReceiverName: "Nguyen Thi F",
			Classify:     1,
			Description:  "Phiếu nhập dụng cụ sửa chữa",
			ItemImportInvoice: []importinvoice.ItemImportInvoice{
				{
					ItemID:      4,
					ItemName:    "",
					Quantity:    8,
					Description: "Bộ cờ lê đa năng",
				},
				{
					ItemID:      5,
					ItemName:    "",
					Quantity:    12,
					Description: "Máy khoan cầm tay",
				},
			},
			ItemCount: 20,
			CreatedAt: time.Now().Add(-2 * time.Hour),
		},
		{
			SenderID:     7,
			SenderName:   "Tran Van G",
			ReceiverID:   8,
			ReceiverName: "Le Thi H",
			Classify:     1,
			Description:  "Phiếu nhập đồ dùng cá nhân",
			ItemImportInvoice: []importinvoice.ItemImportInvoice{
				{
					ItemID:      6,
					ItemName:    "",
					Quantity:    15,
					Description: "Bàn chải đánh răng",
				},
			},
			ItemCount: 15,
			CreatedAt: time.Now().Add(-3 * time.Hour),
		},
		{
			SenderID:     9,
			SenderName:   "Pham Van I",
			ReceiverID:   10,
			ReceiverName: "Hoang Thi J",
			Classify:     1,
			Description:  "Phiếu nhập giấy tờ",
			ItemImportInvoice: []importinvoice.ItemImportInvoice{
				{
					ItemID:      7,
					ItemName:    "",
					Quantity:    5,
					Description: "Hộ chiếu",
				},
				{
					ItemID:      8,
					ItemName:    "",
					Quantity:    3,
					Description: "Chứng minh nhân dân",
				},
			},
			ItemCount: 8,
			CreatedAt: time.Now().Add(-4 * time.Hour),
		},
		{
			SenderID:     11,
			SenderName:   "Nguyen Van K",
			ReceiverID:   12,
			ReceiverName: "Tran Thi L",
			Classify:     1,
			Description:  "Phiếu nhập tài liệu học tập",
			ItemImportInvoice: []importinvoice.ItemImportInvoice{
				{
					ItemID:      9,
					ItemName:    "",
					Quantity:    25,
					Description: "Sổ tay ghi chú",
				},
			},
			ItemCount: 25,
			CreatedAt: time.Now().Add(-5 * time.Hour),
		},
		{
			SenderID:     13,
			SenderName:   "Le Van M",
			ReceiverID:   14,
			ReceiverName: "Pham Thi N",
			Classify:     1,
			Description:  "Phiếu nhập sách tham khảo",
			ItemImportInvoice: []importinvoice.ItemImportInvoice{
				{
					ItemID:      10,
					ItemName:    "",
					Quantity:    7,
					Description: "Sách luyện thi IELTS",
				},
				{
					ItemID:      11,
					ItemName:    "",
					Quantity:    6,
					Description: "Sách luyện thi TOEIC",
				},
			},
			ItemCount: 13,
			CreatedAt: time.Now().Add(-6 * time.Hour),
		},
		{
			SenderID:     15,
			SenderName:   "Hoang Van O",
			ReceiverID:   16,
			ReceiverName: "Nguyen Thi P",
			Classify:     1,
			Description:  "Phiếu nhập quần áo mùa hè",
			ItemImportInvoice: []importinvoice.ItemImportInvoice{
				{
					ItemID:      12,
					ItemName:    "",
					Quantity:    18,
					Description: "Áo thun nam size L",
				},
			},
			ItemCount: 18,
			CreatedAt: time.Now().Add(-7 * time.Hour),
		},
		{
			SenderID:     17,
			SenderName:   "Tran Van Q",
			ReceiverID:   18,
			ReceiverName: "Le Thi R",
			Classify:     1,
			Description:  "Phiếu nhập thiết bị cơ khí",
			ItemImportInvoice: []importinvoice.ItemImportInvoice{
				{
					ItemID:      13,
					ItemName:    "",
					Quantity:    10,
					Description: "Máy mài góc",
				},
				{
					ItemID:      14,
					ItemName:    "",
					Quantity:    5,
					Description: "Bộ tua vít đa năng",
				},
			},
			ItemCount: 15,
			CreatedAt: time.Now().Add(-8 * time.Hour),
		},
		{
			SenderID:     19,
			SenderName:   "Pham Van S",
			ReceiverID:   20,
			ReceiverName: "Hoang Thi T",
			Classify:     1,
			Description:  "Phiếu nhập đồ vệ sinh cá nhân",
			ItemImportInvoice: []importinvoice.ItemImportInvoice{
				{
					ItemID:      15,
					ItemName:    "",
					Quantity:    30,
					Description: "Kem đánh răng",
				},
			},
			ItemCount: 30,
			CreatedAt: time.Now().Add(-9 * time.Hour),
		},
		{
			SenderID:     21,
			SenderName:   "Nguyen Van U",
			ReceiverID:   22,
			ReceiverName: "Tran Thi V",
			Classify:     1,
			Description:  "Phiếu nhập giấy tờ hành chính",
			ItemImportInvoice: []importinvoice.ItemImportInvoice{
				{
					ItemID:      16,
					ItemName:    "",
					Quantity:    4,
					Description: "Sổ hộ khẩu",
				},
			},
			ItemCount: 4,
			CreatedAt: time.Now().Add(-10 * time.Hour),
		},
		{
			SenderID:     23,
			SenderName:   "Le Van W",
			ReceiverID:   24,
			ReceiverName: "Pham Thi X",
			Classify:     1,
			Description:  "Phiếu nhập tài liệu ôn thi",
			ItemImportInvoice: []importinvoice.ItemImportInvoice{
				{
					ItemID:      17,
					ItemName:    "",
					Quantity:    12,
					Description: "Tài liệu ôn thi đại học",
				},
				{
					ItemID:      18,
					ItemName:    "",
					Quantity:    8,
					Description: "Sách hướng dẫn lập trình",
				},
			},
			ItemCount: 20,
			CreatedAt: time.Now().Add(-11 * time.Hour),
		},
		{
			SenderID:     5,
			SenderName:   "Hoang Van E",
			ReceiverID:   6,
			ReceiverName: "Nguyen Thi F",
			Classify:     1,
			Description:  "Phiếu nhập sách văn học",
			ItemImportInvoice: []importinvoice.ItemImportInvoice{
				{
					ItemID:      19,
					ItemName:    "",
					Quantity:    9,
					Description: "Tiểu thuyết lãng mạn",
				},
			},
			ItemCount: 9,
			CreatedAt: time.Now().Add(-12 * time.Hour),
		},
		{
			SenderID:     7,
			SenderName:   "Tran Van G",
			ReceiverID:   8,
			ReceiverName: "Le Thi H",
			Classify:     1,
			Description:  "Phiếu nhập phụ kiện thời trang",
			ItemImportInvoice: []importinvoice.ItemImportInvoice{
				{
					ItemID:      20,
					ItemName:    "",
					Quantity:    25,
					Description: "Mũ thời trang unisex",
				},
			},
			ItemCount: 25,
			CreatedAt: time.Now().Add(-13 * time.Hour),
		},
	}

	for i, inv := range importInvoices {
		// Lấy số hóa đơn
		invoiceNum, err := s.imInvoiceRepo.GetImportInvoiceNum(ctx)
		if err != nil {
			return fmt.Errorf("failed to get invoice number for invoice %d: %w", i+1, err)
		}
		importInvoices[i].InvoiceNum = invoiceNum
		importInvoices[i].IsLock = false

		// Kiểm tra SenderID tồn tại
		senderExisted, err := s.userRepo.IsExist(ctx, inv.SenderID)
		if err != nil {
			return fmt.Errorf("failed to check sender existence for invoice %d: %w", i+1, err)
		}
		if !senderExisted {
			return fmt.Errorf("sender ID %d does not exist for invoice %d", inv.SenderID, i+1)
		}

		// Kiểm tra ItemID và gán ItemName
		for j, itemInv := range inv.ItemImportInvoice {
			var item item.Item
			err := s.itemRepo.GetByID(ctx, &item, itemInv.ItemID)
			if err != nil {

				return fmt.Errorf("failed to get item ID %d for invoice %d: %w", itemInv.ItemID, i+1, err)
			}
			if item.ID == 0 {

				return fmt.Errorf("item ID %d does not exist for invoice %d", itemInv.ItemID, i+1)
			}
			importInvoices[i].ItemImportInvoice[j].ItemName = item.Name
		}

		// Gom nhóm các món đồ thành Warehouse và sinh ItemWareHouse
		warehouses := make(map[uint]warehouse.Warehouse)
		for _, itemInv := range inv.ItemImportInvoice {
			if wh, ok := warehouses[itemInv.ItemID]; ok {
				wh.Quantity += int(itemInv.Quantity)
				warehouses[itemInv.ItemID] = wh
			} else {
				wh := warehouse.Warehouse{
					ItemID:      itemInv.ItemID,
					ItemName:    itemInv.ItemName,
					SKU:         s.imInvoiceService.GenerateSKU(int(itemInv.ItemID)),
					Classify:    inv.Classify,
					Description: "",
					Quantity:    int(itemInv.Quantity),
					StockPlace:  "",
				}
				warehouses[itemInv.ItemID] = wh
			}

			// Sinh ItemWareHouse cho mỗi món đồ
			wh := warehouses[itemInv.ItemID]
			var itemWHs []warehouse.ItemWareHouse
			for k := 0; k < int(itemInv.Quantity); k++ {
				itemCode, err := s.imInvoiceService.GenerateUniqueDigitString(9)
				if err != nil {

					return fmt.Errorf("failed to generate code for item %d in invoice %d: %w", itemInv.ItemID, i+1, err)
				}
				itemWHs = append(itemWHs, warehouse.ItemWareHouse{
					ItemID:      itemInv.ItemID,
					ItemName:    itemInv.ItemName,
					Code:        itemCode,
					Description: itemInv.Description,
				})
			}
			wh.ItemWareHouse = itemWHs
			warehouses[itemInv.ItemID] = wh
		}

		// Chuyển map warehouses thành slice
		var warehouseSlice []warehouse.Warehouse
		for _, wh := range warehouses {
			warehouseSlice = append(warehouseSlice, wh)
		}

		// Lưu ImportInvoice và Warehouse
		if err := s.imInvoiceRepo.CreateImportInvoice(ctx, importInvoices[i], &warehouseSlice); err != nil {
			return fmt.Errorf("failed to create import invoice %d: %w", i+1, err)
		}
	}

	log.Printf("Successfully seeded %d import invoices", len(importInvoices))
	return nil
}
