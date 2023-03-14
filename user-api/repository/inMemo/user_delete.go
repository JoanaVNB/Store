package inMemo

func (in InMemoRepository) Delete(id string) (error){
	delete(in.uMap, id)
	return nil
}

