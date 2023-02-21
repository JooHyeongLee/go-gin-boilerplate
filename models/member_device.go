package models

import (
	"time"
)

const TableNameMemberDevice = "member_device"

// MemberDevice mapped from table <member_device>
type MemberDevice struct {
	ID                int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	MemberID          int64     `gorm:"column:member_id" json:"member_id"`
	DeviceCompany     string    `gorm:"column:device_company" json:"device_company"`           // 단말제조사
	DeviceModel       string    `gorm:"column:device_model" json:"device_model"`               // 단말모델
	MobileOs          string    `gorm:"column:mobile_os;not null" json:"mobile_os"`            // 모바일 OS : A - Android, I - iOS, W - Windows Moible
	MobileVersion     string    `gorm:"column:mobile_version" json:"mobile_version"`           // 모바일버전
	MobileVersionCode string    `gorm:"column:mobile_version_code" json:"mobile_version_code"` // 모바일버전코드
	MobileToken       string    `gorm:"column:mobile_token" json:"mobile_token"`
	UseYn             bool      `gorm:"column:use_yn" json:"use_yn"`
	CreatedDt         time.Time `gorm:"column:created_dt;not null" json:"created_dt"`
	UpdatedDt         time.Time `gorm:"column:updated_dt;not null" json:"updated_dt"`
}

// TableName MemberDevice's table name
func (*MemberDevice) TableName() string {
	return TableNameMemberDevice
}
