package bl

func newBaseBl(userId string) BaseBL {
	//_permit := rights.NewCheckPermit(userId, dao.GetCrudDispatcher())
	//_permit.SetSharedIDs(NewPermissionBL(userId).GetAllOwners())
	return BaseBL{
		userId: userId,
		//permit: _permit,
	}
}

func NewLimsFolder(userId string) *LimsFolderBL {
	return &LimsFolderBL{
		BaseBL: newBaseBl(userId),
		//dao:    dao.GetDeviceDao(),
	}
}