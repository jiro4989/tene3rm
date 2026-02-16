package usecase

import (
	"errors"
	"path/filepath"
	"testing"
	"time"

	"github.com/jiro4989/tene3rm/infra"
	"github.com/stretchr/testify/assert"
)

func TestStateUsecase_LoadState(t *testing.T) {
	ti := time.Date(2026, time.February, 16, 0, 0, 0, 1, time.Local)
	data := make(map[string][]byte)

	tests := []struct {
		name     string
		uc       StateUsecase
		filename string
		want     infra.StateDTO
		wantErr  bool
	}{
		{
			name: "正常系: ファイルを正常にロードできる",
			uc: NewStateUsecase(
				infra.NewFileRepo(filepath.Join("..", "testdata")),
				infra.NewMockTimeGenerator(time.Now()),
			),
			filename: "state1.json",
			want: infra.StateDTO{
				FailCount: 1,
				Created:   &ti,
			},
			wantErr: false,
		},
		{
			name: "正常系: ファイルが存在しない場合はゼロ値の構造体が返される",
			uc: NewStateUsecase(
				infra.NewFileRepo(filepath.Join("..", "testdata")),
				infra.NewMockTimeGenerator(time.Now()),
			),
			filename: "not_exists.json",
			want: infra.StateDTO{
				FailCount: 0,
				Created:   nil,
			},
			wantErr: false,
		},
		{
			name: "異常系: エラーが発生する場合はエラーを返す",
			uc: NewStateUsecase(
				infra.NewInMemoryRepo(data, errors.New("error")),
				infra.NewMockTimeGenerator(time.Now()),
			),
			filename: "not_exists.json",
			want:     infra.StateDTO{},
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := assert.New(t)

			got, err := tt.uc.LoadState(tt.filename)
			if tt.wantErr {
				a.Error(err)
				return
			}

			a.NoError(err)
			a.Equal(tt.want.FailCount, got.FailCount)
			if tt.want.Created != nil {
				// Location がローカルと CI で異なることでテストがコケることがあ
				// るので、Location 以外をチェック
				a.Equal(tt.want.Created.Year(), got.Created.Year())
				a.Equal(tt.want.Created.Month(), got.Created.Month())
				a.Equal(tt.want.Created.Day(), got.Created.Day())
				a.Equal(tt.want.Created.Hour(), got.Created.Hour())
				a.Equal(tt.want.Created.Minute(), got.Created.Minute())
				a.Equal(tt.want.Created.Second(), got.Created.Second())
				a.Equal(tt.want.Created.Nanosecond(), got.Created.Nanosecond())
			}
		})
	}
}

func TestStateUsecase_IsActionLocked(t *testing.T) {
	ti := time.Date(2026, time.February, 16, 0, 0, 0, 1, time.Local)
	data := make(map[string][]byte)

	tests := []struct {
		name     string
		uc       StateUsecase
		filename string
		data     infra.StateDTO
		want     bool
		wantErr  bool
	}{
		{
			name: "正常系: 時間経過が1時間以内の場合は FailCount で判定する",
			uc: NewStateUsecase(
				infra.NewFileRepo(filepath.Join("..", "testdata")),
				infra.NewMockTimeGenerator(time.Date(2026, time.February, 16, 0, 10, 0, 1, time.Local)),
			),
			filename: "state2.json",
			data: infra.StateDTO{
				FailCount: 0,
				Created:   &ti,
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "正常系: FailCount = 2 の場合は false",
			uc: NewStateUsecase(
				infra.NewFileRepo(filepath.Join("..", "testdata")),
				infra.NewMockTimeGenerator(time.Date(2026, time.February, 16, 0, 10, 0, 1, time.Local)),
			),
			filename: "state2.json",
			data: infra.StateDTO{
				FailCount: 2,
				Created:   &ti,
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "正常系: FailCount = 3 の場合は true",
			uc: NewStateUsecase(
				infra.NewFileRepo(filepath.Join("..", "testdata")),
				infra.NewMockTimeGenerator(time.Date(2026, time.February, 16, 0, 10, 0, 1, time.Local)),
			),
			filename: "state2.json",
			data: infra.StateDTO{
				FailCount: 3,
				Created:   &ti,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "正常系: Created が nil のときは false",
			uc: NewStateUsecase(
				infra.NewFileRepo(filepath.Join("..", "testdata")),
				infra.NewMockTimeGenerator(time.Now()),
			),
			filename: "state2.json",
			data: infra.StateDTO{
				FailCount: 0,
				Created:   nil,
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "正常系: 1時間以上経過している場合は Reset される",
			uc: NewStateUsecase(
				infra.NewFileRepo(filepath.Join("..", "testdata")),
				infra.NewMockTimeGenerator(time.Date(2026, time.February, 16, 12, 0, 0, 1, time.Local)),
			),
			filename: "state2.json",
			data: infra.StateDTO{
				FailCount: 0,
				Created:   &ti,
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "異常系: エラーが発生した場合、エラーが返る",
			uc: NewStateUsecase(
				infra.NewInMemoryRepo(data, errors.New("error")),
				infra.NewMockTimeGenerator(time.Date(2026, time.February, 16, 12, 0, 0, 1, time.Local)),
			),
			filename: "state3.json",
			data: infra.StateDTO{
				FailCount: 0,
				Created:   &ti,
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := assert.New(t)

			got, err := tt.uc.IsActionLocked(tt.filename, tt.data)
			if tt.wantErr {
				a.Error(err)
				return
			}

			a.NoError(err)
			a.Equal(tt.want, got)
		})
	}
}

func TestStateUsecase_IncrementFailCount(t *testing.T) {
	t1 := time.Date(2026, time.February, 16, 0, 0, 0, 1, time.Local)
	t2 := time.Date(2026, time.February, 16, 0, 10, 0, 1, time.Local)
	data := make(map[string][]byte)

	tests := []struct {
		name     string
		uc       StateUsecase
		filename string
		data     infra.StateDTO
		want     infra.StateDTO
		wantErr  bool
	}{
		{
			name: "正常系: FailCount を加算して Created を更新する",
			uc: NewStateUsecase(
				infra.NewDefaultInMemoryRepo(),
				infra.NewMockTimeGenerator(t2),
			),
			filename: "key1.json",
			data: infra.StateDTO{
				FailCount: 0,
				Created:   &t1,
			},
			want: infra.StateDTO{
				FailCount: 1,
				Created:   &t2,
			},
			wantErr: false,
		},
		{
			name: "異常系: エラーが発生したらエラーが返る",
			uc: NewStateUsecase(
				infra.NewInMemoryRepo(data, errors.New("error")),
				infra.NewMockTimeGenerator(t2),
			),
			filename: "key1.json",
			data: infra.StateDTO{
				FailCount: 0,
				Created:   &t1,
			},
			want:    infra.StateDTO{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := assert.New(t)

			got, err := tt.uc.IncrementFailCount(tt.filename, tt.data)
			if tt.wantErr {
				a.Error(err)
				return
			}

			a.NoError(err)
			a.Equal(tt.want, got)
		})
	}
}

func TestStateUsecase_ResetFailCount(t *testing.T) {
	t1 := time.Date(2026, time.February, 16, 0, 0, 0, 1, time.Local)
	t2 := time.Date(2026, time.February, 16, 0, 10, 0, 1, time.Local)
	data := make(map[string][]byte)

	tests := []struct {
		name     string
		uc       StateUsecase
		filename string
		data     infra.StateDTO
		want     infra.StateDTO
		wantErr  bool
	}{
		{
			name: "正常系: FailCount を0 にして Created を更新する",
			uc: NewStateUsecase(
				infra.NewDefaultInMemoryRepo(),
				infra.NewMockTimeGenerator(t2),
			),
			filename: "key1.json",
			data: infra.StateDTO{
				FailCount: 3,
				Created:   &t1,
			},
			want: infra.StateDTO{
				FailCount: 0,
				Created:   &t2,
			},
			wantErr: false,
		},
		{
			name: "異常系: エラーが発生したらエラーが返る",
			uc: NewStateUsecase(
				infra.NewInMemoryRepo(data, errors.New("error")),
				infra.NewMockTimeGenerator(t2),
			),
			filename: "key1.json",
			data: infra.StateDTO{
				FailCount: 3,
				Created:   &t1,
			},
			want:    infra.StateDTO{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := assert.New(t)

			got, err := tt.uc.ResetFailCount(tt.filename, tt.data)
			if tt.wantErr {
				a.Error(err)
				return
			}

			a.NoError(err)
			a.Equal(tt.want, got)
		})
	}
}
