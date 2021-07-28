package logic

//func TestLoginService_Login(t *testing.T) {
//	repo.GetDefaultTestDb()
//	type args struct {
//		req *LoginRequest
//		rsp *common.Rsp
//	}
//
//	tests := []struct {
//		name string
//		args args
//		want int
//	}{
//		{
//			"base",
//			args{
//				&LoginRequest{"test", "test", "teacher"},
//				common.NewRsp(),},
//			0,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			ul := &LoginService{}
//			if got := ul.Login(tt.args.req, tt.args.rsp); !reflect.DeepEqual(got.Code(), tt.want) {
//				t.Errorf("LoginService.Login() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
