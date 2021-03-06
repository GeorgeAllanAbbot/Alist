package local

import (
	"context"
	"github.com/alist-org/alist/v3/internal/driver"
	"github.com/alist-org/alist/v3/internal/model"
	"github.com/alist-org/alist/v3/pkg/utils"
	"github.com/pkg/errors"
)

type Driver struct {
	model.Account
	Addition
}

func (d Driver) Config() driver.Config {
	return config
}

func (d *Driver) Init(ctx context.Context, account model.Account) error {
	d.Account = account
	addition := d.Account.Addition
	err := utils.Json.UnmarshalFromString(addition, d.Addition)
	if err != nil {
		return errors.Wrap(err, "error while unmarshal addition")
	}
	return nil
}

func (d *Driver) Drop(ctx context.Context) error {
	return nil
}

func (d *Driver) GetAddition() driver.Additional {
	return d.Addition
}

func (d *Driver) List(ctx context.Context, path string) ([]driver.FileInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (d *Driver) Link(ctx context.Context, path string, args driver.LinkArgs) (*driver.Link, error) {
	//TODO implement me
	panic("implement me")
}

func (d Driver) Other(ctx context.Context, data interface{}) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (d *Driver) MakeDir(ctx context.Context, path string) error {
	//TODO implement me
	panic("implement me")
}

func (d *Driver) Move(ctx context.Context, src, dst string) error {
	//TODO implement me
	panic("implement me")
}

func (d *Driver) Rename(ctx context.Context, src, dst string) error {
	//TODO implement me
	panic("implement me")
}

func (d *Driver) Copy(ctx context.Context, src, dst string) error {
	//TODO implement me
	panic("implement me")
}

func (d *Driver) Remove(ctx context.Context, path string) error {
	//TODO implement me
	panic("implement me")
}

func (d *Driver) Put(ctx context.Context, parentPath string, stream driver.FileStream) error {
	//TODO implement me
	panic("implement me")
}

var _ driver.Driver = (*Driver)(nil)
