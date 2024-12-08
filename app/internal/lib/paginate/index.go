package paginate

// func resolveNext[Model_T ModelInterfaceForPagination](
// 	output IPaginationResult,
// 	dataPack *repository.PaginationPack[Model_T],
// 	pageNumber PaginationPage,
// ) error {

// 	if len(dataPack.Data) == 0 {

// 		return nil
// 	}

// 	lastIndex := len(dataPack.Data) - 1

// 	if lastIndex <= 0 {

// 		return nil
// 	}
// 	/*
// 		implementation state is checked at compile time,
// 		no need any assertions at runtime
// 	*/
// 	lastElement := dataPack.Data[lastIndex]

// 	output.GetNavigation().Next = &NavigationQuery{
// 		Cursor: (*lastElement).GetObjectID(),
// 	}

// 	return nil
// }

// func resolvePrev[Model_T ModelInterfaceForPagination](
// 	output IPaginationResult,
// 	dataPack *repository.PaginationPack[Model_T],
// 	pageNumber PaginationPage,
// ) error {

// 	if len(dataPack.Data) == 0 {

// 		return nil
// 	}

// 	/*
// 		implementation state is checked at compile time,
// 		no need any assertions at runtime
// 	*/
// 	firstElement := dataPack.Data[0]

// 	output.GetNavigation().Previous = &NavigationQuery{
// 		Cursor: (*firstElement).GetObjectID(),
// 		IsPrev: true,
// 	}

// 	return nil
// }

// /*
// preparePaginationNavigation expects the input generic type implements
// ModelInterfaceForPagination and whose methods must be implemeted as value
// receiver
// */
// func preparePaginationNavigation[Model_T ModelInterfaceForPagination](
// 	output IPaginationResult,
// 	dataPack *repository.PaginationPack[Model_T],
// 	pageNumber PaginationPage,
// ) error {

// 	output.SetTotalCount(dataPack.Count)

// 	err := resolveNext[Model_T](output, dataPack, pageNumber)

// 	if err != nil {

// 		return err
// 	}

// 	return resolvePrev[Model_T](output, dataPack, pageNumber)
// }
