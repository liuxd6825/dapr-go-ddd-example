package user_impl

import (
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/application/internals/user/executor"
	"github.com/liuxd6825/dapr-go-ddd-sdk/utils/singleutils"
)

func GetUserFindByIdExecutor() executor.UserFindByIdExecutor {
	return singleutils.GetObject[*userFindByIdExecutor]()
}

func GetUserFindByIdsExecutor() executor.UserFindByIdsExecutor {
	return singleutils.GetObject[*userFindByIdsExecutor]()
}

func GetUserFindPagingExecutor() executor.UserFindPagingExecutor {
	return singleutils.GetObject[*userFindPagingExecutor]()
}

func GetUserFindAllExecutor() executor.UserFindAllExecutor {
	return singleutils.GetObject[*userFindAllExecutor]()
}

func GetUserCreateExecutor() executor.UserCreateExecutor {
	return singleutils.GetObject[*userCreateExecutor]()
}

func GetUserCreateManyExecutor() executor.UserCreateManyExecutor {
	return singleutils.GetObject[*userCreateManyExecutor]()
}

func GetUserUpdateExecutor() executor.UserUpdateExecutor {
	return singleutils.GetObject[*userUpdateExecutor]()
}

func GetUserUpdateManyExecutor() executor.UserUpdateManyExecutor {
	return singleutils.GetObject[*userUpdateManyExecutor]()
}

func GetUserDeleteByIdExecutor() executor.UserDeleteByIdExecutor {
	return singleutils.GetObject[*userDeleteByIdExecutor]()
}

func GetUserDeleteManyExecutor() executor.UserDeleteManyExecutor {
	return singleutils.GetObject[*userDeleteManyExecutor]()
}

func GetUserDeleteAllExecutor() executor.UserDeleteAllExecutor {
	return singleutils.GetObject[*userDeleteAllExecutor]()
}

func init() {
	if err := singleutils.Set[*userFindByIdExecutor](func() *userFindByIdExecutor { return newUserFindByIdExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*userFindByIdsExecutor](func() *userFindByIdsExecutor { return newUserFindByIdsExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*userFindPagingExecutor](func() *userFindPagingExecutor { return newUserFindPagingExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*userFindAllExecutor](func() *userFindAllExecutor { return newUserFindAllExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*userCreateExecutor](func() *userCreateExecutor { return newUserCreateExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*userCreateManyExecutor](func() *userCreateManyExecutor { return newUserCreateManyExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*userUpdateExecutor](func() *userUpdateExecutor { return newUserUpdateExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*userUpdateManyExecutor](func() *userUpdateManyExecutor { return newUserUpdateManyExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*userDeleteByIdExecutor](func() *userDeleteByIdExecutor { return newUserDeleteByIdExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*userDeleteManyExecutor](func() *userDeleteManyExecutor { return newUserDeleteManyExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*userDeleteAllExecutor](func() *userDeleteAllExecutor { return newUserDeleteAllExecutor() }); err != nil {
		panic(err)
	}
}
