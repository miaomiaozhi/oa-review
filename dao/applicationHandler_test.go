package dao

import (
	bean "oa-review/bean"
	"testing"
)

func TestApplicationDao_CreateApplication(t *testing.T) {
	type args struct {
		app *bean.Application
	}
	tests := []struct {
		name    string
		a       *ApplicationDao
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args: args{app: &bean.Application{
				Id:               123,
				Context:          "test app",
				ReviewStatus:     false,
				UserId:           123,
				ApprovedReviewer: make(bean.ApproverMap, 0),
			}},
			want:    123,
			wantErr: false,
		},
	}
	ConnDBForUnitTest()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ApplicationDao{}
			got, err := a.CreateApplication(tt.args.app)
			if (err != nil) != tt.wantErr {
				t.Errorf("ApplicationDao.CreateApplication() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ApplicationDao.CreateApplication() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApplicationDao_FindApplicationById(t *testing.T) {
	type args struct {
		appId int64
	}
	tests := []struct {
		name    string
		dao     *ApplicationDao
		args    args
		want    *bean.Application
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args:    args{appId: 777},
			want:    nil,
			wantErr: true,
		}, // not found
		{
			args: args{appId: 222},
			want: &bean.Application{
				Id:               222,
				Context:          "test app",
				ReviewStatus:     false, // diff
				UserId:           123,
				ApprovedReviewer: nil,
			},
			wantErr: false,
		}, // ok
		{
			args: args{appId: 222},
			want: &bean.Application{
				Id:               222,
				Context:          "test app",
				ReviewStatus:     true, // diff
				UserId:           123,
				ApprovedReviewer: nil,
			},
			wantErr: false,
		}, // review status 错误
	}
	ConnDBForUnitTest()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dao := &ApplicationDao{}
			got, err := dao.FindApplicationById(tt.args.appId)
			if (err != nil) != tt.wantErr {
				t.Errorf("ApplicationDao.FindApplicationById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil {
				if got == nil || got.Id != tt.want.Id || got.Context != tt.want.Context || got.ReviewStatus != tt.want.ReviewStatus {
					t.Errorf("ApplicationDao.FindApplicationById() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestApplicationDao_UpdateApprovedReviewerForApplication(t *testing.T) {
	type args struct {
		appId        int64
		reviewerId   int64
		reviewStatus bool
	}
	tests := []struct {
		name    string
		a       *ApplicationDao
		args    args
		wantErr bool
	}{
		{
			args:    args{appId: 222, reviewerId: 234, reviewStatus: true},
			wantErr: false,
		},
		{
			args:    args{appId: 999, reviewerId: 234, reviewStatus: true},
			wantErr: true,
		},
	}
	ConnDBForUnitTest()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ApplicationDao{}
			if err := a.UpdateApprovedReviewerForApplication(tt.args.appId, tt.args.reviewerId, tt.args.reviewStatus); (err != nil) != tt.wantErr {
				t.Errorf("ApplicationDao.UpdateApprovedReviewerForApplication() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestApplicationDao_CheckApplicationExist(t *testing.T) {
	type args struct {
		Id int64
	}
	tests := []struct {
		name    string
		a       *ApplicationDao
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args: args{Id: 222},
			want: true,
		},
		{
			args: args{Id: 223},
			want: false,
		},
	}
	ConnDBForUnitTest()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ApplicationDao{}
			got, err := a.CheckApplicationExist(tt.args.Id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ApplicationDao.CheckApplicationExist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ApplicationDao.CheckApplicationExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
