package usecase

import (
	"time"

	"github.com/jiro4989/tene3rm/infra"
	"github.com/jiro4989/tene3rm/repo"
)

type StateUsecase struct {
	datastoreRepo repo.DataStoreRepo
	timeGen       infra.TimeGenerator
}

func NewStateUsecase(datastoreRepo repo.DataStoreRepo, timeGen infra.TimeGenerator) StateUsecase {
	return StateUsecase{
		datastoreRepo: datastoreRepo,
		timeGen:       timeGen,
	}
}

func (s StateUsecase) LoadState(filename string) (infra.StateDTO, error) {
	var data infra.StateDTO
	if err := s.datastoreRepo.LoadJSON(filename, &data); err != nil {
		return infra.StateDTO{}, err
	}
	return data, nil
}

// IsActionLocked は失敗回数が3回以上の場合はロック状態と判定する。
// なんらかのエラーが発生した場合はロック中とする。
func (s StateUsecase) IsActionLocked(filename string, data infra.StateDTO) (bool, error) {
	// 一度も失敗していない場合は Created が nil になるので早期リターン
	if data.Created == nil {
		return false, nil
	}

	// 現在時刻が、記録された時刻の1時間後を超えていた場合は失敗回数をリセットする
	if data.Created.Add(1 * time.Hour).Before(s.timeGen.Now()) {
		_, err := s.ResetFailCount(filename, data)
		if err != nil {
			return true, err
		}
		return false, nil
	}

	return 3 <= data.FailCount, nil
}

func (s StateUsecase) IncrementFailCount(filename string, data infra.StateDTO) (infra.StateDTO, error) {
	return s.updateFailCount(filename, data, data.FailCount+1)
}

func (s StateUsecase) ResetFailCount(filename string, data infra.StateDTO) (infra.StateDTO, error) {
	return s.updateFailCount(filename, data, 0)
}

func (s StateUsecase) updateFailCount(filename string, data infra.StateDTO, failCount int) (infra.StateDTO, error) {
	data.FailCount = failCount
	now := s.timeGen.Now()
	data.Created = &now

	if err := s.datastoreRepo.SaveJSON(filename, data); err != nil {
		return infra.StateDTO{}, err
	}
	return data, nil
}
