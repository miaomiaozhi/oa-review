package middleware

import (
	"oa-review/logger"
	"testing"
)

func prepareForUnitTest() {
	// path := "./conf/config.json"
	// config, _ := conf.Read(path)
	// conf.InitGlobalConfig(config)
	// logger.Debug("to here")
	// logger.Debug(conf.GetConfig().Conf.MustGetString("web.jwt_token"))
}

func TestGenJwtToken(t *testing.T) {
	type args struct {
		userId   int64
		userName string
		priority int64
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				userId:   11,
				userName: "mzz",
				priority: 123,
			},
		},
	}
	prepareForUnitTest()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenJwtToken(tt.args.userId, tt.args.userName, tt.args.priority)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenJwtToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GenJwtToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseJwtToken(t *testing.T) {
	type args struct {
		jwtToken string
	}
	tests := []struct {
		name string
		args args
		want struct {
			UserId   int64
			UserName string
			Priority int64
		}
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				jwtToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJQcmlvcml0eSI6MTIzLCJVc2VySWQiOjExLCJVc2VyTmFtZSI6Im16eiIsImlhdCI6MTY4NTQzMDI0Mn0.J1rfD7520hAchUnoJkdONwy4KxSoXi-Hxz6mEo0Kytw",
			},
			want: struct {
				UserId   int64
				UserName string
				Priority int64
			}{
				UserId:   11,
				UserName: "mzz",
				Priority: 123,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseJwtToken(tt.args.jwtToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseJwtToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			id := int64(got["UserId"].(float64))
			name := got["UserName"].(string)
			pri := int64(got["Priority"].(float64))
			logger.Info(id, name, pri)
			if !(id == tt.want.UserId && name == tt.want.UserName && pri == tt.want.Priority) {
				t.Errorf("ParseJwtToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
