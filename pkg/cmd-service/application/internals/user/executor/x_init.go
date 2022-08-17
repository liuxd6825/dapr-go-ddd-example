package executor

import "github.com/liuxd6825/dapr-go-ddd-sdk/utils/singleutils"

func GetUserCreateCommandExecutor() UserCreateCommandExecutor {
	return singleutils.GetObject[*userCreateCommandExecutor]()
}

func GetUserDeleteCommandExecutor() UserDeleteCommandExecutor {
	return singleutils.GetObject[*userDeleteCommandExecutor]()
}

func GetUserUpdateCommandExecutor() UserUpdateCommandExecutor {
	return singleutils.GetObject[*userUpdateCommandExecutor]()
}

func GetFindAggregateByIdExecutor() FindAggregateByIdExecutor {
	return singleutils.GetObject[*findAggregateByIdExecutor]()
}

func init() {
	if err := singleutils.Set[*userCreateCommandExecutor](func() *userCreateCommandExecutor { return newUserCreateCommandExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*userDeleteCommandExecutor](func() *userDeleteCommandExecutor { return newUserDeleteCommandExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*userUpdateCommandExecutor](func() *userUpdateCommandExecutor { return newUserUpdateCommandExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*findAggregateByIdExecutor](func() *findAggregateByIdExecutor { return newFindAggregateByIdExecutor() }); err != nil {
		panic(err)
	}
}
