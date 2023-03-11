package inMemo

func (in InMemoShopRepository) Delete(id string) (error){
	delete(in.uMap, id)
	return nil
}

