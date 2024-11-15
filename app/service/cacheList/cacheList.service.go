package cacheListService

import (
	"app/internal/common"
	libCommon "app/internal/lib/common"
	libError "app/internal/lib/error"
	"app/internal/memoryCache"
	"context"
	"errors"
	"time"
)

var (
	ERR_REFRESH_TOKEN_ID_NOT_EXISTS = errors.New("CacheListManipulator[Key_T, Value_T]] error: token id not exists")
	ERR_UPDATE_NOT_APPROVED         = errors.New("CacheListManipulator[Key_T, Value_T]] error: update not approved")
)

type (
	CacheListManipulator[Key_T comparable, Value_T comparable] struct {
		CacheListLogger
		name        string
		CacheClient *memoryCache.CacheClient[Key_T, Value_T]
	}
)

func NewCacheListManipulator[Key_T, Value_T comparable](name string) *CacheListManipulator[Key_T, Value_T] {

	return &CacheListManipulator[Key_T, Value_T]{
		name: name,
	}
}

func (this *CacheListManipulator[Key_T, Value_T]) Has(tokenId Key_T, ctx context.Context) (res bool, err error) {

	endLogFn := this.PushTraceCondWithMessurement("check_"+this.name+"_cache_key_exist", "true", ctx)

	defer func() {

		endLogFn(err, "false")
	}()

	_, exists, err := this.CacheClient.Read(ctx, tokenId)

	if err != nil {

		return false, libError.NewInternal(err)
	}

	return exists, nil
}

func (this *CacheListManipulator[Key_T, Value_T]) Get(tokenID Key_T, ctx context.Context) (retVal Value_T, exists bool, err error) {

	endLogFn := this.PushTraceCondWithMessurement("get_"+this.name+"_cache_key", libCommon.Ternary(exists, "exists", "absent"), ctx)

	defer func() {

		if err != nil {

			endLogFn(err, err.Error())
			return
		}

		endLogFn(err, "")
	}()

	val, exists, err := this.CacheClient.Read(ctx, tokenID)

	if err != nil {

		return val, false, libError.NewInternal(err)
	}

	return val, exists, nil
}

func (this *CacheListManipulator[Key_T, Value_T]) Read(
	tokenID Key_T,
	readFunc func(ctx memoryCache.IHoldContext[Key_T, Value_T], value Value_T) error,
	ctx context.Context,
) (err error) {

	endLogFn := this.PushTraceCondWithMessurement("read_"+this.name+"_cache_key", "success", ctx)

	defer func() {

		if err != nil {

			endLogFn(err, err.Error())
			return
		}

		endLogFn(err, "")
	}()

	err = this.CacheClient.Hold(ctx, tokenID, readFunc)

	if !errors.Is(err, common.ERR_INTERNAL) {

		return libError.NewInternal(err)
	}

	if err != nil {

		return err
	}

	return nil
}

func (this *CacheListManipulator[Key_T, Value_T]) Update(
	tokenId Key_T,
	updatefunc func(ctx memoryCache.IUpdateContext[Key_T, Value_T], val Value_T) (Value_T, error),
	ctx context.Context,
) (err error) {

	endLogFn := this.PushTraceCondWithMessurement("update_"+this.name+"_cache_key", "success", ctx)

	defer func() {

		if err != nil {

			endLogFn(err, err.Error())
			return
		}

		endLogFn(err, "")
	}()

	err = this.CacheClient.Update(ctx, tokenId, updatefunc)

	if !errors.Is(err, common.ERR_INTERNAL) {

		return libError.NewInternal(err)
	}

	if err != nil {

		return err
	}

	return nil
}

func (this *CacheListManipulator[Key_T, Value_T]) Set(
	tokenId Key_T,
	value Value_T,
	ctx context.Context,
) (success bool, err error) {

	endLogFn := this.PushTraceCondWithMessurement("set_"+this.name+"_cache_key", libCommon.Ternary(success, "success", "failed"), ctx)

	defer func() {

		if err != nil {

			endLogFn(err, err.Error())
			return
		}

		endLogFn(err, "")
	}()

	err = this.CacheClient.Set(ctx, tokenId, value)

	if err != nil {

		return false, libError.NewInternal(err)
	}

	return true, nil
}

func (this *CacheListManipulator[Key_T, Value_T]) SetWithExpire(
	tokenID Key_T, value Value_T, expire time.Time, ctx context.Context,
) (err error) {

	endLogFn := this.PushTraceCondWithMessurement("set_"+this.name+"_cache_key_with_expire_time"+expire.String(), "success", ctx)

	defer func() {

		if err != nil {

			endLogFn(err, err.Error())
			return
		}

		endLogFn(err, "")
	}()

	err = this.CacheClient.SetWithExpire(ctx, tokenID, value, expire)

	if err != nil {

		return libError.NewInternal(err)
	}

	return nil
}

func (this *CacheListManipulator[Key_T, Value_T]) Delete(key Key_T, ctx context.Context) (isDeleted bool, err error) {

	endLogFn := this.PushTraceCondWithMessurement("delete_"+this.name+"_cache_key", libCommon.Ternary(isDeleted, "success", "failed"), ctx)

	defer func() {

		if err != nil {

			endLogFn(err, err.Error())
			return
		}

		endLogFn(err, "")
	}()

	isDeleted, err = this.CacheClient.Delete(ctx, key)

	if err != nil {

		return false, libError.NewInternal(err)
	}

	return isDeleted, nil
}
