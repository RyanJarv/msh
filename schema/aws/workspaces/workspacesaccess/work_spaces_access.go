package workspacesaccess

import (
    "time"
)


type WorkSpacesAccess struct {
    ClientIpAddress string `json:"clientIpAddress"`
    ActionType string `json:"actionType"`
    WorkspacesClientProductName string `json:"workspacesClientProductName"`
    LoginTime time.Time `json:"loginTime"`
    ClientPlatform string `json:"clientPlatform"`
    DirectoryId string `json:"directoryId"`
    WorkspaceId string `json:"workspaceId"`
}

func (w *WorkSpacesAccess) SetClientIpAddress(clientIpAddress string) {
    w.ClientIpAddress = clientIpAddress
}

func (w *WorkSpacesAccess) SetActionType(actionType string) {
    w.ActionType = actionType
}

func (w *WorkSpacesAccess) SetWorkspacesClientProductName(workspacesClientProductName string) {
    w.WorkspacesClientProductName = workspacesClientProductName
}

func (w *WorkSpacesAccess) SetLoginTime(loginTime time.Time) {
    w.LoginTime = loginTime
}

func (w *WorkSpacesAccess) SetClientPlatform(clientPlatform string) {
    w.ClientPlatform = clientPlatform
}

func (w *WorkSpacesAccess) SetDirectoryId(directoryId string) {
    w.DirectoryId = directoryId
}

func (w *WorkSpacesAccess) SetWorkspaceId(workspaceId string) {
    w.WorkspaceId = workspaceId
}
