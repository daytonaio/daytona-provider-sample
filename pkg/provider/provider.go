package provider

import (
	"encoding/json"
	"io"

	internal "github.com/daytonaio/daytona-provider-sample/internal"
	log_writers "github.com/daytonaio/daytona-provider-sample/internal/log"
	provider_types "github.com/daytonaio/daytona-provider-sample/pkg/types"

	"github.com/daytonaio/daytona/pkg/logger"
	"github.com/daytonaio/daytona/pkg/provider"
	"github.com/daytonaio/daytona/pkg/types"
)

type SampleProvider struct {
	BasePath          *string
	ServerDownloadUrl *string
	ServerVersion     *string
	ServerUrl         *string
	ServerApiUrl      *string
	LogsDir           *string
	OwnProperty       string
}

func (p *SampleProvider) Initialize(req provider.InitializeProviderRequest) (*types.Empty, error) {
	p.OwnProperty = "my-own-property"

	p.BasePath = &req.BasePath
	p.ServerDownloadUrl = &req.ServerDownloadUrl
	p.ServerVersion = &req.ServerVersion
	p.ServerUrl = &req.ServerUrl
	p.ServerApiUrl = &req.ServerApiUrl
	p.LogsDir = &req.LogsDir

	return new(types.Empty), nil
}

func (p SampleProvider) GetInfo() (provider.ProviderInfo, error) {
	return provider.ProviderInfo{
		Name:    "provider-sample",
		Version: internal.Version,
	}, nil
}

func (p SampleProvider) GetTargetManifest() (*provider.ProviderTargetManifest, error) {
	return provider_types.GetTargetManifest(), nil
}

func (p SampleProvider) GetDefaultTargets() (*[]provider.ProviderTarget, error) {
	info, err := p.GetInfo()
	if err != nil {
		return nil, err
	}

	defaultTargets := []provider.ProviderTarget{
		{
			Name:         "default-target",
			ProviderInfo: info,
			Options:      "{\n\t\"Required STring\": \"default-required-string\"\n}",
		},
	}
	return &defaultTargets, nil
}

func (p SampleProvider) CreateWorkspace(workspaceReq *provider.WorkspaceRequest) (*types.Empty, error) {
	logWriter := io.MultiWriter(&log_writers.InfoLogWriter{})
	if p.LogsDir != nil {
		wsLogWriter := logger.GetWorkspaceLogger(*p.LogsDir, workspaceReq.Workspace.Id)
		logWriter = io.MultiWriter(&log_writers.InfoLogWriter{}, wsLogWriter)
		defer wsLogWriter.Close()
	}

	logWriter.Write([]byte("Workspace created\n"))

	return new(types.Empty), nil
}

func (p SampleProvider) StartWorkspace(workspaceReq *provider.WorkspaceRequest) (*types.Empty, error) {
	return new(types.Empty), nil
}

func (p SampleProvider) StopWorkspace(workspaceReq *provider.WorkspaceRequest) (*types.Empty, error) {
	return new(types.Empty), nil
}

func (p SampleProvider) DestroyWorkspace(workspaceReq *provider.WorkspaceRequest) (*types.Empty, error) {
	return new(types.Empty), nil
}

func (p SampleProvider) GetWorkspaceInfo(workspaceReq *provider.WorkspaceRequest) (*types.WorkspaceInfo, error) {
	providerMetadata, err := p.getWorkspaceMetadata(workspaceReq)
	if err != nil {
		return nil, err
	}

	workspaceInfo := &types.WorkspaceInfo{
		Name:             workspaceReq.Workspace.Name,
		ProviderMetadata: providerMetadata,
	}

	projectInfos := []*types.ProjectInfo{}
	for _, project := range workspaceReq.Workspace.Projects {
		projectInfo, err := p.GetProjectInfo(&provider.ProjectRequest{
			TargetOptions: workspaceReq.TargetOptions,
			Project:       project,
		})
		if err != nil {
			return nil, err
		}
		projectInfos = append(projectInfos, projectInfo)
	}
	workspaceInfo.Projects = projectInfos

	return workspaceInfo, nil
}

func (p SampleProvider) CreateProject(projectReq *provider.ProjectRequest) (*types.Empty, error) {
	logWriter := io.MultiWriter(&log_writers.InfoLogWriter{})
	if p.LogsDir != nil {
		wsLogWriter := logger.GetWorkspaceLogger(*p.LogsDir, projectReq.Project.WorkspaceId)
		projectLogWriter := logger.GetProjectLogger(*p.LogsDir, projectReq.Project.WorkspaceId, projectReq.Project.Name)
		logWriter = io.MultiWriter(&log_writers.InfoLogWriter{}, wsLogWriter, projectLogWriter)
		defer wsLogWriter.Close()
		defer projectLogWriter.Close()
	}

	logWriter.Write([]byte("Project created\n"))

	return new(types.Empty), nil
}

func (p SampleProvider) StartProject(projectReq *provider.ProjectRequest) (*types.Empty, error) {
	return new(types.Empty), nil
}

func (p SampleProvider) StopProject(projectReq *provider.ProjectRequest) (*types.Empty, error) {
	return new(types.Empty), nil
}

func (p SampleProvider) DestroyProject(projectReq *provider.ProjectRequest) (*types.Empty, error) {
	return new(types.Empty), nil
}

func (p SampleProvider) GetProjectInfo(projectReq *provider.ProjectRequest) (*types.ProjectInfo, error) {
	providerMetadata := provider_types.ProjectMetadata{
		Property: projectReq.Project.Name,
	}

	metadataString, err := json.Marshal(providerMetadata)
	if err != nil {
		return nil, err
	}

	projectInfo := &types.ProjectInfo{
		Name:             projectReq.Project.Name,
		IsRunning:        true,
		Created:          "Created at ...",
		Started:          "Started at ...",
		Finished:         "Finished at ...",
		ProviderMetadata: string(metadataString),
	}

	return projectInfo, nil
}

func (p SampleProvider) getWorkspaceMetadata(workspaceReq *provider.WorkspaceRequest) (string, error) {
	metadata := provider_types.WorkspaceMetadata{
		Property: workspaceReq.Workspace.Id,
	}

	jsonContent, err := json.Marshal(metadata)
	if err != nil {
		return "", err
	}

	return string(jsonContent), nil
}
