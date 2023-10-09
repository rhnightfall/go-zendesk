
// Code generated by Script. DO NOT EDIT.
// Source: script/codegen/main.go
//
// Generated by this command:
//
//	go run script/codegen/main.go

package zendesk

import "context"

func (z *Client) GetGroupsIterator(ctx context.Context, opts *PaginationOptions) *Iterator[Group] {
	return &Iterator[Group]{
		CommonOptions: opts.CommonOptions,
		pageSize:      opts.PageSize,
		hasMore:       true,
		isCBP:         opts.IsCBP,
		pageAfter:     "",
		pageIndex:     1,
		ctx:           ctx,
		obpFunc:       z.GetGroupsOBP,
		cbpFunc:       z.GetGroupsCBP,
	}
}

func (z *Client) GetGroupsOBP(ctx context.Context, opts *OBPOptions) ([]Group, Page, error) {
	var data struct {
		Groups []Group `json:"groups"`
		Page
	}

	tmp := opts
	if tmp == nil {
		tmp = &OBPOptions{}
	}
	
	u, err := addOptions("/groups.json", tmp)
	
	if err != nil {
		return nil, Page{}, err
	}

	err = getData(z, ctx, u, &data)
	if err != nil {
		return nil, Page{}, err
	}
	return data.Groups, data.Page, nil
}

func (z *Client) GetGroupsCBP(ctx context.Context, opts *CBPOptions) ([]Group, CursorPaginationMeta, error) {
	var data struct {
		Groups []Group `json:"groups"`
		Meta    CursorPaginationMeta `json:"meta"`
	}

	tmp := opts
	if tmp == nil {
		tmp = &CBPOptions{}
	}
	
	u, err := addOptions("/groups.json", tmp)
	
	if err != nil {
		return nil, data.Meta, err
	}

	err = getData(z, ctx, u, &data)
	if err != nil {
		return nil, data.Meta, err
	}
	return data.Groups, data.Meta, nil
}

