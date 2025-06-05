package seeder

import (
	"context"
	"final_project/internal/domain/category"
	"final_project/internal/domain/item"
	"final_project/internal/domain/post"
	rolepermission "final_project/internal/domain/role_permission"
	"final_project/internal/domain/user"
	"final_project/internal/pkg/enums"
	"final_project/internal/pkg/hash"
	"final_project/internal/pkg/helpers"
	"fmt"
	"os"
)

type Seeder struct {
	rolePerRepo  rolepermission.Repository
	itemRepo     item.Repository
	userRepo     user.Repository
	categoryRepo category.Repository
	postRepo     post.Repository
}

func NewSeeder(rolePerRepo rolepermission.Repository, itemRepo item.Repository, userRepo user.Repository, categoryRepo category.Repository, postRepo post.Repository) *Seeder {
	return &Seeder{
		rolePerRepo:  rolePerRepo,
		itemRepo:     itemRepo,
		userRepo:     userRepo,
		categoryRepo: categoryRepo,
		postRepo:     postRepo,
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
			Password:    "admin1234",
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
			Password:    "admin1234",
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
			Password:    "admin1234",
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
			Password:    "user1234",
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

	// Kiểm tra bảng user có rỗng không
	isEmpty, err := s.userRepo.IsTableEmpty(ctx)
	if err != nil {
		return err
	}
	if !isEmpty {
		fmt.Println("Post had data...")
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
			Password:    "superadmin",
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
			Password:    "contentmanager",
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
			Password:    "warehousemanager",
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
			Password:    "123456",
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
