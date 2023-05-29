package dao

import (
	bean "oa-review/bean"
	"reflect"
	"testing"
)

func TestReviewerDao_CreateReviewer(t *testing.T) {
	type args struct {
		reviewer *bean.Reviewer
	}
	tests := []struct {
		name    string
		r       *ReviewerDao
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "test0",
			args: args{
				reviewer: &bean.Reviewer{
					Id:           234,
					Name:         "test_reviewer",
					Applications: nil,
					Options:      nil,
					Priority:     234,
				},
			},
			want:    234,
			wantErr: false,
		},
	}
	DropTable()
	ConnDBForUnitTest()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ReviewerDao{}
			got, err := r.CreateReviewer(tt.args.reviewer)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReviewerDao.CreateReviewer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReviewerDao.CreateReviewer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReviewerDao_FindReviewerById(t *testing.T) {
	type args struct {
		reviewerId int64
	}
	tests := []struct {
		name    string
		r       *ReviewerDao
		args    args
		want    *bean.Reviewer
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args: args{reviewerId: 234},
			want: &bean.Reviewer{
				Id:           234,
				Name:         "test_reviewer",
				Applications: nil,
				Options:      nil,
				Priority:     234,
			},
			wantErr: false,
		},
		{
			args:    args{reviewerId: 123},
			want:    nil,
			wantErr: true,
		},
	}
	ConnDBForUnitTest()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ReviewerDao{}
			got, err := r.FindReviewerById(tt.args.reviewerId)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReviewerDao.FindReviewerById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil {
				if got == nil || got.Id != tt.want.Id || got.Name != tt.want.Name || got.Priority != tt.want.Priority {
					t.Errorf("ReviewerDao.FindReviewerById() = %v, want %v", got, tt.want)
					return
				}
			}
		})
	}
}

func TestReviewerDao_AddReviewerOption(t *testing.T) {
	type args struct {
		reviewerId int64
		option     *bean.ReviewOption
	}
	tests := []struct {
		name    string
		r       *ReviewerDao
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				reviewerId: 123,
			},
			wantErr: true,
		}, // not found
		{
			args: args{
				reviewerId: 234,
				option:     &bean.ReviewOption{ApplicationId: 555, ReviewStatus: true},
			},
			wantErr: false,
		}, // ok
	}
	ConnDBForUnitTest()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ReviewerDao{}
			if err := r.AddReviewerOption(tt.args.reviewerId, tt.args.option); (err != nil) != tt.wantErr {
				t.Errorf("ReviewerDao.AddReviewerOption() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestReviewerDao_DeleteReviewerOption(t *testing.T) {
	type args struct {
		reviewerId int64
	}
	tests := []struct {
		name    string
		r       *ReviewerDao
		args    args
		want    *bean.ReviewOption
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args:    args{reviewerId: 123},
			wantErr: true,
		}, // not found
		{
			args:    args{reviewerId: 234},
			wantErr: true,
		}, // empty
		// {
		// 	args:    args{reviewerId: 234},
		// 	want:    &bean.ReviewOption{ApplicationId: 555, ReviewStatus: true},
		// 	wantErr: false,
		// }, // not empty
	}
	ConnDBForUnitTest()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ReviewerDao{}
			got, err := r.DeleteReviewerOption(tt.args.reviewerId)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReviewerDao.DeleteReviewerOption() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReviewerDao.DeleteReviewerOption() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReviewerDao_CheckReviewerExist(t *testing.T) {
	type args struct {
		Id int64
	}
	tests := []struct {
		name    string
		r       *ReviewerDao
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args: args{Id: 123},
			want: false,
		},
		{
			args: args{Id: 234},
			want: true,
		},
	}
	ConnDBForUnitTest()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ReviewerDao{}
			got, err := r.CheckReviewerExist(tt.args.Id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReviewerDao.CheckReviewerExist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReviewerDao.CheckReviewerExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
