package enum

type CardTypeStatus int

const (
	_CardTypeStatus = iota
	SINGLE          //单根
	DOUBLE          //对子
	THREE           //三不带
	THREE_AND_ONE   //三带一
	BOMB            //炸弹
	FOUR_TWO        //四带二
	PLANE           //飞机
	PLANE_EMPTY     //三不带飞机
	DOUBLE_ALONE    //连对
	SINGLE_ALONE    //顺子
	KING_BOMB       //王炸
	ERROR_TYPE      //非法类型
)

func GetStatus(value int) CardTypeStatus {
	switch value {
	case 0:
		return SINGLE
	case 1:
		return DOUBLE
	case 2:
		return THREE
	case 3:
		return THREE_AND_ONE
	case 4:
		return BOMB
	case 5:
		return FOUR_TWO
	case 6:
		return PLANE
	case 7:
		return PLANE_EMPTY
	case 8:
		return DOUBLE_ALONE
	case 9:
		return SINGLE_ALONE
	case 10:
		return KING_BOMB
	case 11:
		return ERROR_TYPE
	}
	return ERROR_TYPE
}

func GetIntValue(value CardTypeStatus) int {
	switch value {
	case SINGLE:
		return 0
	case DOUBLE:
		return 1
	case THREE:
		return 2
	case THREE_AND_ONE:
		return 3
	case BOMB:
		return 4
	case FOUR_TWO:
		return 5
	case PLANE:
		return 6
	case PLANE_EMPTY:
		return 7
	case DOUBLE_ALONE:
		return 8
	case SINGLE_ALONE:
		return 9
	case KING_BOMB:
		return 10
	case ERROR_TYPE:
		return 11
	}
	return 0
}
