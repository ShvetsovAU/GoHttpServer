package dao


func GetLimsFolderDao() *LimsFolderDao {
	return &LimsFolderDao{
		limsFolders: getContext().DB.C(CollLimsFolder),
		//values:  getContext().DB.C(CollDeviceValues),
	}
}