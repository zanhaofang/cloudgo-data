package entities


//UserInfoAtomicService .
type UserInfoAtomicService struct{}

//UserInfoService .
var UserInfoService = UserInfoAtomicService{}

// Save .
func (*UserInfoAtomicService) Save(u *UserInfo) error {
	session  := mydb.NewSession()
	defer session.Close()
	err := session.Begin()
  _, err = session.Insert(u)
	//tx, err := mydb.Begin()
	//checkErr(err)

	//dao := userInfoDao{sqlt.NewSQLTemplate(tx)}
	//err = dao.Save(u)

	if err == nil {
		session.Commit()
	} else {
		session.Rollback()
		return err
	}
	return nil
}

// FindAll .
func (*UserInfoAtomicService) FindAll() []UserInfo {
	//dao := userInfoDao{sqlt.NewSQLTemplate(mydb)}
	//ulist, err := dao.FindAll()
	ulist := make([]UserInfo, 0)
  err := mydb.Find(&ulist)
	checkErr(err)
	return ulist
}

// FindByID .
func (*UserInfoAtomicService) FindByID(id int) *UserInfo {
	//dao := userInfoDao{sqlt.NewSQLTemplate(mydb)}
	//u, err := dao.FindByID(id)
  var u UserInfo
	_, err := mydb.Id(id).Get(&u)
	checkErr(err)
	return &u
}
