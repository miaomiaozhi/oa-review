package dao

import (
	bean "oa-review/bean"
	"oa-review/db"
	"oa-review/logger"
	"testing"
)

func ConnDBForUnitTest() {
	mdb, err := db.NewDB(db.NewDBConfig(
		"root", "mozezhao", "127.0.0.1", 3306, "oa_review",
	))
	if err != nil {
		logger.Info("conn db for unit test error")
		return
	}
	db.SetDB(mdb)
	if db.Migration() != nil {
		logger.Info("conn db for unit test error")
		return
	}
	logger.Info("conn db for unit test success")
}

func DropTable() {
	ConnDBForUnitTest()
	db.GetDB().Migrator().DropTable(&bean.User{})
	db.GetDB().Migrator().DropTable(&bean.Reviewer{})
	db.GetDB().Migrator().DropTable(&bean.Application{})
}

func TestUserDao_CreateUser(t *testing.T) {
	type args struct {
		user *bean.User
	}
	tests := []struct {
		name    string
		u       *UserDao
		args    args
		want    int64
		wantErr bool
	}{
		// ADD
		{
			name: "create_user",
			u:    NewUserDaoInstance(),
			args: args{user: &bean.User{
				Id:           11,
				Password:     "123",
				Name:         "test",
				Applications: nil,
				Priority:     11,
			}},
			want:    11,
			wantErr: false,
		},
		// {
		// 	name: "create_user",
		// 	u:    NewUserDaoInstance(),
		// 	args: args{user: &bean.User{
		// 		Id:           22,
		// 		Password:     "234",
		// 		Name:         "test",
		// 		Applications: nil,
		// 		Priority:     11,
		// 	}},
		// 	want:    22,
		// 	wantErr: false,
		// },
		// {
		// 	name: "create_user",
		// 	u:    NewUserDaoInstance(),
		// 	args: args{user: &bean.User{
		// 		Id:           33,
		// 		Password:     "123",
		// 		Name:         "test",
		// 		Applications: nil,
		// 		Priority:     11,
		// 	}},
		// 	want:    33,
		// 	wantErr: false,
		// },
	}
	DropTable()
	ConnDBForUnitTest()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := tt.u
			got, err := u.CreateUser(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserDao.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UserDao.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDao_FindUserByUserId(t *testing.T) {
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		u       *UserDao
		args    args
		want    *bean.User
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			u:    NewUserDaoInstance(),
			args: args{11},
			want: &bean.User{
				Id:       11,
				Password: "123",
				Name:     "test",
				Priority: 11,
			},
			wantErr: false,
		},
	}
	ConnDBForUnitTest()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewUserDaoInstance()
			got, err := u.FindUserByUserId(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserDao.FindUserByUserId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Id != tt.want.Id || got.Password != tt.want.Password || got.Name != tt.want.Name || got.Priority != tt.want.Priority {
				t.Errorf("UserDao.FindUserByUserId() = %v, want %v", got, tt.want)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("UserDao.FindUserByUserId() = %v, want %v", got, tt.want)
			// }
		})
	}
}

func TestUserDao_AddApplicationForUser(t *testing.T) {
	type args struct {
		userId int64
		appId  int64
	}
	tests := []struct {
		name    string
		u       *UserDao
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				userId: 11,
				appId:  123,
			},
			wantErr: false,
		},
		{
			args: args{
				userId: 11,
				appId:  234,
			},
			wantErr: false,
		},
		// {
		// 	args: args{
		// 		userId: 11,
		// 		appId:  312,
		// 	},
		// 	wantErr: false, // ok
		// },
		// {
		// 	args: args{
		// 		userId: 22,
		// 		appId:  123123,
		// 	},
		// 	wantErr: false, // ok
		// },
	}
	ConnDBForUnitTest()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewUserDaoInstance()
			if err := u.AddApplicationForUser(tt.args.userId, tt.args.appId); (err != nil) != tt.wantErr {
				t.Errorf("UserDao.AddApplicationForUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserDao_CheckUserExist(t *testing.T) {
	type args struct {
		userId int64
	}
	tests := []struct {
		name    string
		u       *UserDao
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test0",
			args: args{
				userId: 123,
			},
			want: false,
		},
		{
			name: "test1",
			args: args{
				userId: 11,
			},
			want: true,
		},
	}
	ConnDBForUnitTest()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewUserDaoInstance()
			got, err := u.CheckUserExist(tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserDao.CheckUserExist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UserDao.CheckUserExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDao_TableSize(t *testing.T) {
	tests := []struct {
		name    string
		u       *UserDao
		want    int64
		wantErr bool
	}{
		{
			want:    1,
			wantErr: false,
		},
	}
	ConnDBForUnitTest()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserDao{}
			got, err := u.TableSize()
			if (err != nil) != tt.wantErr {
				t.Errorf("UserDao.TableSize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UserDao.TableSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
