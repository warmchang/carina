package volume

import "carina/pkg/devicemanager/types"

const (
	THIN     = "thin-"
	SNAP     = "snap-"
	LVVolume = "volume-"
)

// 本接口负责对外提供方法
// 处理业务逻辑并调用lvm接口
type LocalVolume interface {
	CreateVolume(lvName, vgName string, size, ratio uint64) error
	DeleteVolume(lvName, vgName string) error
	ResizeVolume(lvName, vgName string, size, ratio uint64) error
	VolumeList(lvName, vgName string) error

	CreateSnapshot(snapName, lvName, vgName string) error
	DeleteSnapshot(snapName, vgName string) error
	RestoreSnapshot(snapName, vgName string) error
	SnapshotList(lvName, vgName string) error

	CloneVolume(lvName, vgName, newLvName string) error

	// 额外的方法
	GetCurrentVgStruct() ([]*types.VgGroup, error)
	AddNewDiskToVg(disk, vgName string) error
	RemoveDiskInVg(disk, vgName string) error
}
