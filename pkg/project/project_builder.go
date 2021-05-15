package project

type Builder interface {
	BuildProject(basePath, moduleName, parentVersion string)
}

type BuildMode uint8

const (
	BuildModeOnline  BuildMode = 1
	BuildModeOffline BuildMode = 2
)

func NewProjectBuilder(mode BuildMode) Builder {
	var builder Builder
	switch mode {
	case BuildModeOffline:
		builder = new(OfflineProjectBuilder)
		break
	case BuildModeOnline:
		break
	}
	return builder
}
