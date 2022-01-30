package tokenengine

import (
	"log"
	"testing"
)

func TestTokenGeneratorEngine_VerifyToken(t *testing.T) {
	type fields struct {
		Log log.Logger
	}
	type args struct {
		tokenStr   string
		privateKey string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:   "negative",
			fields: fields{log.Logger{}},
			args: args{
				tokenStr:   "",
				privateKey: "",
			},
			want:    false,
			wantErr: true,
		},
		{
			name:   "negative",
			fields: fields{log.Logger{}},
			args: args{
				tokenStr:   "",
				privateKey: "j0hn",
			},
			want:    false,
			wantErr: true,
		},
		{
			name:   "negative",
			fields: fields{log.Logger{}},
			args: args{
				tokenStr:   "SFMyNTY=.eyJleHAiOiIxNjQzNTcyMjI2IiwidXNlciI6Impvc2hpbi5qb2huc29uIn0=.rp1MtE89fPwGmRTVS1bpe1V6FbMOcq4Ug6XpUz33mYQ=",
				privateKey: "j0hn",
			},
			want:    false,
			wantErr: false,
		},
		{
			name:   "positive",
			fields: fields{log.Logger{}},
			args: args{
				tokenStr:   "SFMyNTY=.eyJleHAiOiIxNjQzNTcyMjI2IiwidXNlciI6Impvc2hpbi5qb2huc29uIn0=.rp1MtE89fPwGmRTVS1bpe1V6FbMOcq4Ug6XpUz33mYQ=",
				privateKey: "j0sh19",
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tm := TokenGeneratorEngine{
				Log: tt.fields.Log,
			}
			got, err := tm.VerifyToken(tt.args.tokenStr, tt.args.privateKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("VerifyToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("VerifyToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}
