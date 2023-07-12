package scanner

import (
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/mmcdole/gofeed"

	"github.com/temathc/news-aggregator/models"
)

func TestScanner_ScanRss(t *testing.T) {
	time := time.Now()
	guid := "guid"
	wg := sync.WaitGroup{}
	wg.Add(3)
	type fields struct {
		parser           *Mockparser
		publicationsRepo *MockpublicationsRepository
	}
	type args struct {
		link string
		wg   *sync.WaitGroup
	}
	tests := []struct {
		name    string
		setUp   func(f *fields)
		args    args
		wantErr bool
	}{
		{
			name: "success",
			setUp: func(f *fields) {
				f.parser.EXPECT().ParseURL("test_link").Return(&gofeed.Feed{
					Items: []*gofeed.Item{{
						Title:           "test_title",
						Description:     "test_description",
						Link:            "test_link",
						PublishedParsed: &time,
						GUID:            guid,
					}},
				}, nil).Times(1)
				f.publicationsRepo.EXPECT().AddPublications([]models.Publications{
					{
						GUID:        &guid,
						Title:       "test_title",
						Description: "test_description",
						PubTime:     time,
						Link:        "test_link",
					},
				}).Return(nil).Times(1)

			},
			args: args{
				link: "test_link",
				wg:   &wg,
			},
			wantErr: false,
		},
		{
			name: "parse_url_error",
			setUp: func(f *fields) {
				f.parser.EXPECT().ParseURL("test_link").Return(&gofeed.Feed{}, errors.New("Error")).Times(1)
			},
			args: args{
				link: "test_link",
				wg:   &wg,
			},
			wantErr: true,
		},
		{
			name: "add_publication_error",
			setUp: func(f *fields) {
				f.parser.EXPECT().ParseURL("test_link").Return(&gofeed.Feed{
					Items: []*gofeed.Item{{
						Title:           "test_title",
						Description:     "test_description",
						Link:            "test_link",
						PublishedParsed: &time,
						GUID:            guid,
					}},
				}, nil).Times(1)
				f.publicationsRepo.EXPECT().AddPublications([]models.Publications{
					{
						GUID:        &guid,
						Title:       "test_title",
						Description: "test_description",
						PubTime:     time,
						Link:        "test_link",
					},
				}).Return(errors.New("Error")).Times(1)

			},
			args: args{
				link: "test_link",
				wg:   &wg,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				parser:           NewMockparser(ctrl),
				publicationsRepo: NewMockpublicationsRepository(ctrl),
			}
			if tt.setUp != nil {
				tt.setUp(&f)
			}

			h := NewScanner(f.publicationsRepo, f.parser)

			if err := h.ScanRss(tt.args.link, tt.args.wg); (err != nil) != tt.wantErr {
				t.Errorf("ScanRss() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
