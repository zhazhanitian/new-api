package model

import (
	"errors"

	"github.com/QuantumNous/new-api/common"

	"gorm.io/gorm"
)

// Vendor 用于存储供应商信息，供模型引用
// Name 唯一，用于在模型中关联
// Icon 采用 @lobehub/icons 的图标名，前端可直接渲染
// Status 预留字段，1 表示启用
// 本表同样遵循 3NF 设计范式

type Vendor struct {
	Id          int            `json:"id"`
	Name        string         `json:"name" gorm:"size:128;not null;uniqueIndex:uk_vendor_name_delete_at,priority:1"`
	Description string         `json:"description,omitempty" gorm:"type:text"`
	Icon        string         `json:"icon,omitempty" gorm:"type:varchar(128)"`
	Status      int            `json:"status" gorm:"default:1"`
	CreatedTime int64          `json:"created_time" gorm:"bigint"`
	UpdatedTime int64          `json:"updated_time" gorm:"bigint"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index;uniqueIndex:uk_vendor_name_delete_at,priority:2"`
}

// Insert 创建新的供应商记录
func (v *Vendor) Insert() error {
	now := common.GetTimestamp()
	v.CreatedTime = now
	v.UpdatedTime = now
	return DB.Create(v).Error
}

// IsVendorNameDuplicated 检查供应商名称是否重复（排除自身 ID）
func IsVendorNameDuplicated(id int, name string) (bool, error) {
	if name == "" {
		return false, nil
	}
	var cnt int64
	err := DB.Model(&Vendor{}).Where("name = ? AND id <> ?", name, id).Count(&cnt).Error
	return cnt > 0, err
}

// Update 更新供应商记录
func (v *Vendor) Update() error {
	v.UpdatedTime = common.GetTimestamp()
	return DB.Save(v).Error
}

// Delete 软删除供应商
func (v *Vendor) Delete() error {
	return DB.Delete(v).Error
}

// GetVendorByID 根据 ID 获取供应商
func GetVendorByID(id int) (*Vendor, error) {
	var v Vendor
	err := DB.First(&v, id).Error
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// imageVendorCompany 把先前按产品名写入的供应商改成公司名。
// names 的第一项是要新建的公司名，后面是已经存在时优先沿用的别名。
type imageVendorCompany struct {
	product string
	names   []string
	desc    string
	icon    string
}

// EnsureImageModelVendors 补上图像模型还没有的公司供应商，并改掉按产品名插入的记录。
// 公司名已经存在时不新建，只把产品名那条并过去。
func EnsureImageModelVendors() {
	companies := []imageVendorCompany{
		{product: "FLUX", names: []string{"Black Forest Labs"}, desc: "Black Forest Labs 图像模型", icon: "Flux"},
		{product: "MAI", names: []string{"Microsoft", "微软"}, desc: "Microsoft 图像模型", icon: "Microsoft.Color"},
		{product: "Muse", names: []string{"Meta"}, desc: "Meta 图像模型", icon: "Meta.Color"},
		{product: "Qwen", names: []string{"阿里巴巴"}, desc: "阿里巴巴图像模型", icon: "Qwen.Color"},
		{product: "Hunyuan", names: []string{"腾讯"}, desc: "腾讯图像模型", icon: "Hunyuan.Color"},
	}
	for _, company := range companies {
		if err := ensureImageVendorCompany(company); err != nil {
			common.SysLog("ensure image vendor " + company.names[0] + ": " + err.Error())
		}
	}
}

func ensureImageVendorCompany(company imageVendorCompany) error {
	product, err := getVendorByName(company.product)
	if err != nil {
		return err
	}
	var existing *Vendor
	for _, name := range company.names {
		existing, err = getVendorByName(name)
		if err != nil {
			return err
		}
		if existing != nil {
			break
		}
	}
	if product != nil && existing != nil {
		if err = DB.Model(&Model{}).Where("vendor_id = ?", product.Id).Update("vendor_id", existing.Id).Error; err != nil {
			return err
		}
		return product.Delete()
	}
	if product != nil {
		product.Name = company.names[0]
		product.Description = company.desc
		product.Icon = company.icon
		product.Status = 1
		return product.Update()
	}
	if existing != nil {
		return nil
	}
	created := Vendor{
		Name:        company.names[0],
		Description: company.desc,
		Icon:        company.icon,
		Status:      1,
	}
	return created.Insert()
}

func getVendorByName(name string) (*Vendor, error) {
	var vendor Vendor
	err := DB.Where("name = ?", name).First(&vendor).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &vendor, nil
}

// GetAllVendors 获取全部供应商（分页）
func GetAllVendors(offset int, limit int) ([]*Vendor, error) {
	var vendors []*Vendor
	err := DB.Offset(offset).Limit(limit).Find(&vendors).Error
	return vendors, err
}

// SearchVendors 按关键字搜索供应商
func SearchVendors(keyword string, offset int, limit int) ([]*Vendor, int64, error) {
	db := DB.Model(&Vendor{})
	if keyword != "" {
		like := "%" + keyword + "%"
		db = db.Where("name LIKE ? OR description LIKE ?", like, like)
	}
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var vendors []*Vendor
	if err := db.Offset(offset).Limit(limit).Order("id DESC").Find(&vendors).Error; err != nil {
		return nil, 0, err
	}
	return vendors, total, nil
}
