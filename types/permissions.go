package types

import "errors"

type PermissionItem struct {
	Op    string `json:"op"`
	Path  string `json:"path"`
	Value string `json:"value"`
}

type PermissionBlock struct {
	PermissionItems []PermissionItem
}

func (p *PermissionBlock) Add(i PermissionItem) {
	p.PermissionItems = append(p.PermissionItems, i)
}
func (p *PermissionBlock) List() ([]PermissionItem, error) {
	EmptyPermissionBlock := errors.New("No Permissions in the block")

	if len(p.PermissionItems) > 0 {
		return p.PermissionItems, nil
	} else {
		return nil, EmptyPermissionBlock
	}

}
