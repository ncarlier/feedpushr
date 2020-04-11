package output

import (
	"sync"

	"github.com/ncarlier/feedpushr/v3/pkg/cache"
	"github.com/ncarlier/feedpushr/v3/pkg/model"
	"github.com/ncarlier/feedpushr/v3/pkg/output/plugins"
	"github.com/ncarlier/feedpushr/v3/pkg/plugin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Manager is the object that manage outputs.
type Manager struct {
	lock       sync.RWMutex
	plugins    map[string]model.OutputPlugin
	processors map[string]*Processor
	cache      *cache.Manager
	logger     zerolog.Logger
}

// NewOutputManager creates a new output manager
func NewOutputManager(cache *cache.Manager) (*Manager, error) {
	manager := &Manager{
		plugins:    plugins.GetBuiltinOutputPlugins(),
		processors: make(map[string]*Processor),
		cache:      cache,
		logger:     log.With().Str("component", "output-manager").Logger(),
	}

	// Register external output plugins...
	err := plugin.GetRegistry().ForEachOutputPlugin(func(plug model.OutputPlugin) error {
		manager.plugins[plug.Spec().Name] = plug
		return nil
	})
	if err != nil {
		return nil, err
	}
	return manager, err
}

// GetAvailableOutputs get all available outputs
func (m *Manager) GetAvailableOutputs() []model.Spec {
	result := []model.Spec{}
	for _, plugin := range m.plugins {
		result = append(result, plugin.Spec())
	}
	return result
}

// Push articles to output processors
func (m *Manager) Push(articles []*model.Article) {
	for _, processor := range m.processors {
		processor.Process(articles)
	}
}

// Shutdown stop the manager (aka. stop all processors)
func (m *Manager) Shutdown() {
	m.logger.Debug().Msg("shutting down all processors")
	for _, processor := range m.processors {
		processor.Shutdown()
	}
	m.logger.Debug().Msg("all processors stopped")
}
