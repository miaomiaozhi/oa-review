package dao

import (
	"oa-review/db"
	"oa-review/logger"
	v1 "oa-review/models/protoreq/v1"
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

func TestUserDao_CreateUser(t *testing.T) {
	type args struct {
		user *v1.User
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
			name: "test0",
			u:    NewUserDaoInstance(),
			args: args{user: &v1.User{
				Id:       11,
				Password: "123",
				Name:     "test",
				Priority: 11,
			}},
			want:    11,
			wantErr: false,
		},
	}
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
		want    *v1.User
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			u:    NewUserDaoInstance(),
			args: args{11},
			want: &v1.User{
				Id:       11,
				Password: "123",
				Name:     "test", // ok
				// Name:     "hello", // fail
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
			if got.Id != tt.want.Id || got.Password != tt.want.Password || got.Name != tt.want.Name {
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
				userId: 123,
				appId:  312,
			},
			wantErr: true, // 不存在 123
		},
		{
			args: args{
				userId: 11,
				appId:  312,
			},
			wantErr: false, // ok
		},
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
		// {
		// 	name: "test0",
		// 	args: args{
		// 		userId: 123,
		// 	},
		// 	want: false,
		// },
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
