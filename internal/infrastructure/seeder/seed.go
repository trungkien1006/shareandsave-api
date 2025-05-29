package seeder

import (
	"context"
	"final_project/internal/domain/admin"
	"final_project/internal/domain/item"
	rolepermission "final_project/internal/domain/role_permission"
	"final_project/internal/domain/user"
	"final_project/internal/pkg/enums"
	"final_project/internal/pkg/hash"
	"final_project/internal/pkg/helpers"
	"fmt"
	"os"
)

type Seeder struct {
	rolePerRepo rolepermission.Repository
	adminRepo   admin.Repository
	itemRepo    item.Repository
	userRepo    user.Repository
}

func NewSeeder(rolePerRepo rolepermission.Repository, adminRepo admin.Repository, itemRepo item.Repository, userRepo user.Repository) *Seeder {
	return &Seeder{
		rolePerRepo: rolePerRepo,
		adminRepo:   adminRepo,
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

	if err := s.seedAdmin(); err != nil {
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

		//request management permissions
		{Name: "Read Request", Code: "read_request"},
		{Name: "Reply Request", Code: "reply_request"},
		{Name: "Delete Request", Code: "delete_request"},

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
		"Super Admin":       {"*"},
		"Content Manager":   {"create_post", "read_post", "update_post", "delete_post", "read_notification"},
		"Warehouse Manager": {"read_warehouse", "read_item_warehouse", "read_import_invoice", "create_import_invoice", "update_import_invoice", "lock_import_invoice", "delete_import_invoice", "read_export_invoice", "create_export_invoice", "update_export_invoice", "lock_export_invoice", "delete_export_invoice", "read_notification", "read_request", "reply_request", "delete_request"},
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

func (s *Seeder) seedAdmin() error {
	ctx := context.Background()

	fmt.Println("Start seed admin...")

	isEmpty, err := s.adminRepo.IsTableEmpty(ctx)
	if err != nil {
		return err
	}
	if !isEmpty {
		return nil
	}

	var roles []rolepermission.Role

	if err := s.rolePerRepo.GetAllRoles(&roles); err != nil {
		return err
	}

	roleMap := make(map[string]uint)
	for _, r := range roles {
		roleMap[r.Name] = r.ID
	}

	admins := []admin.Admin{
		{Email: "superadmin@example.com", Password: "superadmin", FullName: "Super Admin", Status: 1, RoleID: roleMap["Super Admin"]},
		{Email: "content@example.com", Password: "contentmanager", FullName: "Content Manager", Status: 1, RoleID: roleMap["Content Manager"]},
		{Email: "warehouse@example.com", Password: "warehousemanager", FullName: "Warehouse Manager", Status: 1, RoleID: roleMap["Warehouse Manager"]},
		// {Email: "hr@example.com", Password: "hrmanager", FullName: "HR Manager", Status: 1, RoleID: roleMap["Human Resources Manager"]},
		// {Email: "client@example.com", Password: "clientmanager", FullName: "Client Manager", Status: 1, RoleID: roleMap["Client Manager"]},
	}

	for i := range admins {
		if _, err := s.adminRepo.Save(ctx, admins[i]); err != nil {
			return err
		}
	}

	fmt.Println("Finish seed admin...")

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

	fmt.Println("Check user table...")

	isEmpty, err := s.userRepo.IsTableEmpty(ctx)

	if err != nil {
		return err
	}

	if !isEmpty {
		return nil
	}

	fmt.Println("Seeding users...")

	users := []user.User{
		{
			Email:       "kien@gmail.com",
			Password:    "123456",
			Avatar:      "",
			FullName:    "Nguyễn Văn Kiên",
			PhoneNumber: "0123456789",
			Address:     "Hồ Chí Minh",
			Status:      1,
			GoodPoint:   0,
			Major:       "Công nghệ thông tin",
		},
		{
			Email:       "vinh@gmail.com",
			Password:    "123456",
			Avatar:      "",
			FullName:    "Nguyễn Văn Vinh",
			PhoneNumber: "0123456790",
			Address:     "Hồ Chí Minh",
			Status:      1,
			GoodPoint:   0,
			Major:       "Cơ khí",
		},
		{
			Email:       "khoa@gmail.com",
			Password:    "123456",
			Avatar:      "",
			FullName:    "Nguyễn Thị Khoa",
			PhoneNumber: "0123456791",
			Address:     "Hồ Chí Minh",
			Status:      1,
			GoodPoint:   0,
			Major:       "Kế toán doanh nghiệp",
		},
		{
			Email:       "hoang@gmail.com",
			Password:    "123456",
			Avatar:      "",
			FullName:    "Nguyễn Thị Hoàng",
			PhoneNumber: "0123456792",
			Address:     "Hồ Chí Minh",
			Status:      1,
			GoodPoint:   0,
			Major:       "Ô tô",
		},
		{
			Email:       "hien.tran@gmail.com",
			Password:    "123456",
			Avatar:      "",
			FullName:    "Trần Thị Hiền",
			PhoneNumber: "0981234561",
			Address:     "Hà Nội",
			Status:      1,
			GoodPoint:   0,
			Major:       "Sư phạm Toán",
		},
		{
			Email:       "minh.nguyen@gmail.com",
			Password:    "123456",
			Avatar:      "",
			FullName:    "Nguyễn Minh",
			PhoneNumber: "0934567890",
			Address:     "Đà Nẵng",
			Status:      1,
			GoodPoint:   0,
			Major:       "Thiết kế đồ họa",
		},
		{
			Email:       "thu.ha@gmail.com",
			Password:    "123456",
			Avatar:      "",
			FullName:    "Hà Thu",
			PhoneNumber: "0912345678",
			Address:     "Hải Phòng",
			Status:      1,
			GoodPoint:   0,
			Major:       "Ngôn ngữ Anh",
		},
		{
			Email:       "bao.tran@gmail.com",
			Password:    "123456",
			Avatar:      "",
			FullName:    "Trần Bảo",
			PhoneNumber: "0901234567",
			Address:     "Cần Thơ",
			Status:      1,
			GoodPoint:   0,
			Major:       "Công nghệ thực phẩm",
		},
		{
			Email:       "yen.pham@gmail.com",
			Password:    "123456",
			Avatar:      "",
			FullName:    "Phạm Thị Yến",
			PhoneNumber: "0987654321",
			Address:     "Hà Tĩnh",
			Status:      1,
			GoodPoint:   0,
			Major:       "Kỹ thuật phần mềm",
		},
		{
			Email:       "quang.le@gmail.com",
			Password:    "123456",
			Avatar:      "",
			FullName:    "Lê Văn Quang",
			PhoneNumber: "0976543210",
			Address:     "Nha Trang",
			Status:      1,
			GoodPoint:   0,
			Major:       "Điện - điện tử",
		},
		{
			Email:       "anh.do@gmail.com",
			Password:    "123456",
			Avatar:      "",
			FullName:    "Đỗ Thị Anh",
			PhoneNumber: "0965432109",
			Address:     "Huế",
			Status:      1,
			GoodPoint:   0,
			Major:       "Du lịch",
		},
		{
			Email:       "long.tran@gmail.com",
			Password:    "123456",
			Avatar:      "",
			FullName:    "Trần Văn Long",
			PhoneNumber: "0954321098",
			Address:     "Bình Dương",
			Status:      1,
			GoodPoint:   0,
			Major:       "Quản trị kinh doanh",
		},
		{
			Email:       "tuan.pham@gmail.com",
			Password:    "123456",
			Avatar:      "",
			FullName:    "Phạm Tuấn",
			PhoneNumber: "0943210987",
			Address:     "Hà Nam",
			Status:      1,
			GoodPoint:   0,
			Major:       "Marketing",
		},
		{
			Email:       "nhung.vo@gmail.com",
			Password:    "123456",
			Avatar:      "",
			FullName:    "Võ Thị Nhung",
			PhoneNumber: "0932109876",
			Address:     "Quảng Ngãi",
			Status:      1,
			GoodPoint:   0,
			Major:       "Tài chính ngân hàng",
		},
		{
			Email:       "loc.nguyen@gmail.com",
			Password:    "123456",
			Avatar:      "",
			FullName:    "Nguyễn Hữu Lộc",
			PhoneNumber: "0921098765",
			Address:     "Vũng Tàu",
			Status:      1,
			GoodPoint:   0,
			Major:       "Xây dựng",
		},
		{
			Email:       "my.tran@gmail.com",
			Password:    "123456",
			Avatar:      "",
			FullName:    "Trần Mỹ",
			PhoneNumber: "0910987654",
			Address:     "Đồng Nai",
			Status:      1,
			GoodPoint:   0,
			Major:       "Thiết kế nội thất",
		},
		{
			Email:       "dat.nguyen@gmail.com",
			Password:    "123456",
			Avatar:      "",
			FullName:    "Nguyễn Văn Đạt",
			PhoneNumber: "0909876543",
			Address:     "Gia Lai",
			Status:      1,
			GoodPoint:   0,
			Major:       "Công nghệ thông tin",
		},
		{
			Email:       "linh.le@gmail.com",
			Password:    "123456",
			Avatar:      "",
			FullName:    "Lê Ngọc Linh",
			PhoneNumber: "0898765432",
			Address:     "Lâm Đồng",
			Status:      1,
			GoodPoint:   0,
			Major:       "Thiết kế thời trang",
		},
		{
			Email:       "dung.tran@gmail.com",
			Password:    "123456",
			Avatar:      "",
			FullName:    "Trần Văn Dũng",
			PhoneNumber: "0887654321",
			Address:     "Phú Yên",
			Status:      1,
			GoodPoint:   0,
			Major:       "Kỹ thuật ô tô",
		},
		{
			Email:       "hanh.nguyen@gmail.com",
			Password:    "123456",
			Avatar:      "",
			FullName:    "Nguyễn Thị Hạnh",
			PhoneNumber: "0876543210",
			Address:     "Bạc Liêu",
			Status:      1,
			GoodPoint:   0,
			Major:       "Sư phạm Ngữ Văn",
		},
		{
			Email:       "bao.le@gmail.com",
			Password:    "123456",
			Avatar:      "",
			FullName:    "Lê Văn Bảo",
			PhoneNumber: "0865432109",
			Address:     "Kiên Giang",
			Status:      1,
			GoodPoint:   0,
			Major:       "Lập trình web",
		},
		{
			Email:       "trang.pham@gmail.com",
			Password:    "123456",
			Avatar:      "",
			FullName:    "Phạm Thị Trang",
			PhoneNumber: "0854321098",
			Address:     "Tiền Giang",
			Status:      1,
			GoodPoint:   0,
			Major:       "Chăm sóc sức khỏe",
		},
		{
			Email:       "duy.nguyen@gmail.com",
			Password:    "123456",
			Avatar:      "",
			FullName:    "Nguyễn Văn Duy",
			PhoneNumber: "0843210987",
			Address:     "Cà Mau",
			Status:      1,
			GoodPoint:   0,
			Major:       "Trí tuệ nhân tạo",
		},
		{
			Email:       "thu.nguyen@gmail.com",
			Password:    "123456",
			Avatar:      "",
			FullName:    "Nguyễn Thị Thu",
			PhoneNumber: "0832109876",
			Address:     "Sóc Trăng",
			Status:      1,
			GoodPoint:   0,
			Major:       "Quản trị nhà hàng - khách sạn",
		},
		{
			Email:       "khanh.vo@gmail.com",
			Password:    "123456",
			Avatar:      "",
			FullName:    "Võ Minh Khánh",
			PhoneNumber: "0821098765",
			Address:     "Tây Ninh",
			Status:      1,
			GoodPoint:   0,
			Major:       "Thương mại điện tử",
		},
		{
			Email:       "nhan.tran@gmail.com",
			Password:    "123456",
			Avatar:      "",
			FullName:    "Trần Hữu Nhân",
			PhoneNumber: "0810987654",
			Address:     "An Giang",
			Status:      1,
			GoodPoint:   0,
			Major:       "Mạng máy tính",
		},
		{
			Email:       "nga.nguyen@gmail.com",
			Password:    "123456",
			Avatar:      "",
			FullName:    "Nguyễn Thị Nga",
			PhoneNumber: "0809876543",
			Address:     "Đắk Lắk",
			Status:      1,
			GoodPoint:   0,
			Major:       "Điều dưỡng",
		},
	}

	for i := range users {
		strBase64Image, err := helpers.ResizeImageFromFileToBase64(os.Getenv("IMAGE_PATH")+"/user.png", enums.UserImageWidth, enums.UserImageHeight)
		if err != nil {
			return fmt.Errorf("error resizing image: %w", err)
		}

		users[i].Avatar = strBase64Image
		users[i].Password, err = hash.HashPassword(users[i].Password)

		if err != nil {
			return fmt.Errorf("error hash user password: %w", err)
		}

		if err := s.userRepo.Save(ctx, &users[i]); err != nil {
			return err
		}
	}

	fmt.Println("Seeding users success...")

	return nil
}
