package usecase_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theluckiesthuman/lexiquery/internal/usecase"
)

func TestDictionaryQuery_Query(t *testing.T) {
	type fields struct {
		url    string
		apiKey string
	}
	type args struct {
		word string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr error
	}{
		{
			name: "success",
			fields: fields{
				url:    "http://test.com",
				apiKey: "test",
			},
			args: args{
				word: "exercise",
			},
			want:    "ek-sər-sīz (noun): the act of bringing into play or realizing in action",
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`[{"hwi":{"prs":[{"mw":"ek-sər-sīz"}]},"fl":"noun","shortdef":["the act of bringing into play or realizing in action"]}]`))
			})
			server := httptest.NewServer(handler)
			defer server.Close()

			d := usecase.NewDictionaryQuery(server.URL, tt.fields.apiKey)
			got, err := d.Query(tt.args.word)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
