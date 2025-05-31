package seeder

import (
	"context"
	"final_project/internal/domain/item"
	rolepermission "final_project/internal/domain/role_permission"
	"final_project/internal/domain/user"
	"fmt"
)

type Seeder struct {
	rolePerRepo rolepermission.Repository
	itemRepo    item.Repository
	userRepo    user.Repository
}

func NewSeeder(rolePerRepo rolepermission.Repository, itemRepo item.Repository, userRepo user.Repository) *Seeder {
	return &Seeder{
		rolePerRepo: rolePerRepo,
		itemRepo:    itemRepo,
		userRepo:    userRepo,
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

func (s *Seeder) seedItems() error {
	ctx := context.Background()

	fmt.Println("Start seed items...")

	isEmpty, err := s.itemRepo.IsTableEmpty(ctx)
	if err != nil {
		return err
	}
	if !isEmpty {
		return nil
	}

	items := []item.Item{
		{Name: "Giáo trình Chính trị 1", Description: "Giáo trình học phần chính trị cơ bản cho sinh viên năm nhất.", Image: ""},
		{Name: "Giáo trình Thiết kế Web", Description: "Giáo trình dạy về HTML, CSS, và JavaScript cơ bản.", Image: ""},
		{Name: "Sách Toán Cao Cấp", Description: "Tài liệu học Toán cao cấp cho khối kỹ thuật.", Image: ""},
		{Name: "Giáo trình Cấu trúc Dữ liệu và Giải thuật", Description: "Tài liệu học thuật về các thuật toán và cấu trúc dữ liệu cơ bản.", Image: ""},
		{Name: "Máy tính Casio fx-570VN Plus", Description: "Máy tính cầm tay cần thiết cho các môn học tự nhiên và kỹ thuật.", Image: ""},
		{Name: "Balo Laptop Chống Sốc", Description: "Balo chống sốc bảo vệ laptop và tài liệu học tập.", Image: ""},
		{Name: "Giáo trình Kinh tế Vi mô", Description: "Tài liệu học môn Kinh tế Vi mô cho sinh viên Kinh tế.", Image: ""},
		{Name: "Bút dạ quang", Description: "Bộ 5 cây bút dạ quang nhiều màu để ghi chú nội dung quan trọng.", Image: ""},
		{Name: "Giáo trình Nguyên lý Kế toán", Description: "Sách học phần cơ bản cho ngành kế toán.", Image: ""},
		{Name: "Tập vở ô ly A4", Description: "Tập vở cũ còn dư dùng để ghi chép bài giảng.", Image: ""},
		{Name: "Sách Hóa học Đại cương", Description: "Tài liệu học Hóa học cơ bản dành cho sinh viên khoa tự nhiên.", Image: ""},
		{Name: "Thước kẻ 30cm", Description: "Dụng cụ học tập thường dùng trong các bài toán vẽ hình hoặc kỹ thuật.", Image: ""},
		{Name: "Giáo trình Lập trình C C++", Description: "Sách học về ngôn ngữ lập trình C/C++ từ cơ bản đến nâng cao.", Image: ""},
		{Name: "Sách Giáo dục Quốc phòng", Description: "Tài liệu học quốc phòng an ninh cho sinh viên năm nhất.", Image: ""},
		{Name: "Laptop cũ Dell Vostro", Description: "Máy tính cũ dành cho sinh viên học tập và làm bài tập nhóm.", Image: ""},
		{Name: "Tai nghe chụp tai", Description: "Tai nghe phục vụ học online, nghe giảng, và làm việc nhóm.", Image: ""},
		{Name: "Giáo trình Tâm lý học đại cương", Description: "Sách học phần Tâm lý học cho khối xã hội.", Image: ""},
		{Name: "Chuột không dây Logitech", Description: "Thiết bị hỗ trợ học tập và làm việc với máy tính.", Image: ""},
		{Name: "Giáo trình Xác suất Thống kê", Description: "Tài liệu dành cho các ngành kỹ thuật và kinh tế.", Image: ""},
		{Name: "Sách Nhập môn Công nghệ Thông tin", Description: "Tài liệu học cơ sở ngành CNTT cho sinh viên năm nhất.", Image: ""},
		{Name: "Đèn học LED để bàn", Description: "Đèn bàn dùng học bài, tiết kiệm điện và dịu mắt.", Image: ""},
		{Name: "Ổ cứng di động 500GB", Description: "Dùng để lưu trữ tài liệu học tập, phần mềm, bài tập nhóm.", Image: ""},
		{Name: "Giáo trình Pháp luật đại cương", Description: "Tài liệu học bắt buộc cho tất cả các ngành.", Image: ""},
		{Name: "Thẻ nhớ SD 64GB", Description: "Thiết bị lưu trữ dành cho sinh viên học media, CNTT.", Image: ""},
		{Name: "Giáo trình Xây dựng Đảng", Description: "Tài liệu học chính trị dành cho sinh viên năm cuối.", Image: ""},
		{Name: "Bút bi Thiên Long", Description: "Bộ bút bi dư dùng để ghi chép bài vở.", Image: ""},
		{Name: "Bìa hồ sơ và túi clear A4", Description: "Dùng để nộp bài, tài liệu, báo cáo môn học.", Image: ""},
		{Name: "Máy in cũ Canon", Description: "Máy in cũ hỗ trợ in báo cáo, tài liệu học tập tại nhà.", Image: ""},
		{Name: "Ghế gấp học bài", Description: "Ghế gấp gọn, dùng khi ngồi học trong phòng trọ nhỏ.", Image: ""},
		{Name: "Bảng trắng mini", Description: "Bảng trắng dùng ghi chú, trình bày khi học nhóm.", Image: ""},
	}

	for i := range items {
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
		fmt.Println("Users table not empty, skipping seed users.")
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
		if err := s.userRepo.Save(ctx, &u); err != nil {
			return err
		}
	}

	fmt.Println("Finish seed users...")

	return nil
}
