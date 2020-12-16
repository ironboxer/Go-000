//+build wireinject 忽略编译

package main

import (
	"Week04/internal/biz"
	"Week04/internal/data"
	"github.com/google/wire"
)

func InitTagUsecase() *biz.TagUsecase {
	wire.Build(biz.NewTagUsecase, data.NewTagRepo)
	return &biz.TagUsecase{}
}
