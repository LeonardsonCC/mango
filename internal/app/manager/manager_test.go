package manager_test

import (
	"testing"

	"github.com/LeonardsonCC/mango/internal/app/manager"
)

func TestManager(t *testing.T) {
	t.Run("should get some results from scrappers", func(t *testing.T) {
		m := manager.NewManager()
		results, err := m.SearchManga("naruto")

		if err != nil {
			t.Errorf("failed to get results %s", err.Error())
		}

		if len(results) == 0 {
			t.Error("failed to get any results")
		}
	})
}
