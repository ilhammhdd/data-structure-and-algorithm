package binarySearchTree

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeftReplacement_DirectLeaf(t *testing.T) {
	//         8
	//    4
	// 2     5
	givenToInsert := []int{8, 4, 2, 5}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	//    8
	// 2
	//   5
	actualReplacementParent, actualReplacement := LeftReplacement(root.left)
	assert.Equal(t, 4, actualReplacementParent.Val)
	assert.Equal(t, 2, actualReplacement.Val)
}

func TestLeftReplacement_DirectWithLeftChild(t *testing.T) {
	//	        8
	//	   4
	//	2     5
	// 1
	givenToInsert := []int{8, 4, 2, 5, 1}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	//	        8
	//	   2
	//	1     5
	actualReplacementParent, actualReplacement := LeftReplacement(root.left)
	assert.Equal(t, 4, actualReplacementParent.Val)
	assert.Equal(t, 2, actualReplacement.Val)
}

func TestLeftReplacement_RightLeaf(t *testing.T) {
	//	        8
	//	    4
	//   1
	// 0   2
	givenToInsert := []int{8, 4, 1, 0, 2}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	//	        8
	//	    2
	//   1
	// 0
	actualReplacementParent, actualReplacement := LeftReplacement(root.left)
	assert.Equal(t, 1, actualReplacementParent.Val)
	assert.Equal(t, 2, actualReplacement.Val)
}

func TestLeftReplacement_RightLeafWithLeftChild(t *testing.T) {
	//	        8
	//	    4
	//   1
	// 0   3
	//    2
	givenToInsert := []int{8, 4, 1, 0, 3, 2}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	//	        8
	//	    3
	//   1
	// 0   2
	actualReplacementParent, actualReplacement := LeftReplacement(root.left)
	assert.Equal(t, 1, actualReplacementParent.Val)
	assert.Equal(t, 3, actualReplacement.Val)
}

func TestRightReplacement_DirectLeaf(t *testing.T) {
	// 2
	//   5
	//    8
	givenToInsert := []int{2, 5, 8}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	// 2
	//   8
	actualReplacementParent, actualReplacement := RightReplacement(root.right)
	assert.Equal(t, 5, actualReplacementParent.Val)
	assert.Equal(t, 8, actualReplacement.Val)
}

func TestRightReplacement_DirectWithRightChild(t *testing.T) {
	// 2
	//    5
	//      8
	//       9
	givenToInsert := []int{2, 5, 8, 9}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	// 2
	//   8
	//    9
	actualReplacementParent, actualReplacement := RightReplacement(root.right)
	assert.Equal(t, 5, actualReplacementParent.Val)
	assert.Equal(t, 8, actualReplacement.Val)
}

func TestRightReplacement_Leaf(t *testing.T) {
	// 2
	//    5
	//      8
	//     7
	givenToInsert := []int{2, 5, 8, 7}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	// 2
	//    7
	//      8
	actualReplacementParent, actualReplacement := RightReplacement(root.right)
	assert.Equal(t, 8, actualReplacementParent.Val)
	assert.Equal(t, 7, actualReplacement.Val)
}

func TestRightReplacement_WithRightChild(t *testing.T) {
	// 2
	//    5
	//       9
	//     7
	//      8
	givenToInsert := []int{2, 5, 9, 7, 8}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	// 2
	//    7
	//       9
	//     8
	actualReplacementParent, actualReplacement := RightReplacement(root.right)
	assert.Equal(t, 9, actualReplacementParent.Val)
	assert.Equal(t, 7, actualReplacement.Val)
}

func TestRemove_NoExistsLess(t *testing.T) {
	//          8
	//     4            12
	//  2     6     10      14
	// 1 3   5 7   9 11   13  15
	givenToInsert := []int{8, 4, 12, 2, 6, 10, 14, 1, 3, 5, 7, 9, 11, 13, 15}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	root = Remove(root, 0)
	expectedInOrder := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	assert.Equal(t, expectedInOrder, InOrder(root))
}

func TestRemove_NoExistsGreater(t *testing.T) {
	//          8
	//     4            12
	//  2     6     10      14
	// 1 3   5 7   9 11   13  15
	givenToInsert := []int{8, 4, 12, 2, 6, 10, 14, 1, 3, 5, 7, 9, 11, 13, 15}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	root = Remove(root, 16)
	expectedInOrder := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	assert.Equal(t, expectedInOrder, InOrder(root))
}

func TestRemove_Root(t *testing.T) {
	// 8
	givenToInsert := []int{8}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	//
	root = Remove(root, 8)
	expectedInOrder := []int{}
	assert.Equal(t, expectedInOrder, InOrder(root))
}

func TestRemove_Root_LeftReplacementDirect(t *testing.T) {
	//  8
	// 7
	givenToInsert := []int{8, 7}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	// 7
	root = Remove(root, 8)
	expectedInOrder := []int{7}
	assert.Equal(t, expectedInOrder, InOrder(root))
}

func TestRemove_Root_RightReplacementDirect(t *testing.T) {
	// 8
	//  9
	givenToInsert := []int{8, 9}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	// 7
	root = Remove(root, 8)
	expectedInOrder := []int{9}
	assert.Equal(t, expectedInOrder, InOrder(root))
}

func TestRemove_RootLeftLeafReplacement(t *testing.T) {
	//          8
	//     4            12
	//  2     6     10      14
	// 1 3   5 7   9 11   13  15
	givenToInsert := []int{8, 4, 12, 2, 6, 10, 14, 1, 3, 5, 7, 9, 11, 13, 15}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	//          7
	//     4            12
	//  2     6     10      14
	// 1 3   5     9 11   13  15
	expectedInOrder := []int{1, 2, 3, 4, 5, 6, 7, 9, 10, 11, 12, 13, 14, 15}
	expectedPreOrder := []int{7, 4, 2, 1, 3, 6, 5, 12, 10, 9, 11, 14, 13, 15}
	expectedPostOrder := []int{1, 3, 2, 5, 6, 4, 9, 11, 10, 13, 15, 14, 12, 7}
	root = Remove(root, 8)
	assert.Equal(t, expectedInOrder, InOrder(root))
	assert.Equal(t, expectedPreOrder, PreOrder(root))
	assert.Equal(t, expectedPostOrder, PostOrder(root))
}

func TestRemove_RootLeftReplacement(t *testing.T) {
	//          8
	//     4            12
	//  2     6     10      14
	// 1 3   5     9 11   13  15
	givenToInsert := []int{8, 4, 12, 2, 6, 10, 14, 1, 3, 5, 9, 11, 13, 15}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	//          6
	//     4            12
	//  2     5     10      14
	// 1 3         9 11   13  15
	expectedInorder := []int{1, 2, 3, 4, 5, 6, 9, 10, 11, 12, 13, 14, 15}
	expectedPreOrder := []int{6, 4, 2, 1, 3, 5, 12, 10, 9, 11, 14, 13, 15}
	expectedPostOrder := []int{1, 3, 2, 5, 4, 9, 11, 10, 13, 15, 14, 12, 6}
	root = Remove(root, 8)
	assert.Equal(t, expectedInorder, InOrder(root))
	assert.Equal(t, expectedPreOrder, PreOrder(root))
	assert.Equal(t, expectedPostOrder, PostOrder(root))
}

func TestRemove_RootRightLeafReplacement(t *testing.T) {
	//   8
	//      12
	//  10      14
	// 9 11   13  15
	givenToInsert := []int{8, 12, 10, 14, 9, 11, 13, 15}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	//   9
	//      12
	//  10      14
	//   11   13  15
	expectedInOrder := []int{9, 10, 11, 12, 13, 14, 15}
	expectedPreOrder := []int{9, 12, 10, 11, 14, 13, 15}
	expectedPostOrder := []int{11, 10, 13, 15, 14, 12, 9}
	root = Remove(root, 8)
	assert.Equal(t, expectedInOrder, InOrder(root))
	assert.Equal(t, expectedPreOrder, PreOrder(root))
	assert.Equal(t, expectedPostOrder, PostOrder(root))
}

func TestRemove_RootRightReplacement(t *testing.T) {
	//   8
	//      12
	//  10      14
	//   11   13  15
	givenToInsert := []int{8, 12, 10, 14, 11, 13, 15}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	//   10
	//      12
	//  11      14
	//        13  15
	expectedInOrder := []int{10, 11, 12, 13, 14, 15}
	expectedPreOrder := []int{10, 12, 11, 14, 13, 15}
	expectedPostOrder := []int{11, 13, 15, 14, 12, 10}
	root = Remove(root, 8)
	assert.Equal(t, expectedInOrder, InOrder(root))
	assert.Equal(t, expectedPreOrder, PreOrder(root))
	assert.Equal(t, expectedPostOrder, PostOrder(root))
}

func TestRemove_LeftReplacement_DirectLeaf(t *testing.T) {
	//         8
	//    4
	// 2     5
	givenToInsert := []int{8, 4, 2, 5}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	//    8
	// 2
	//   5
	expectedInOrder := []int{2, 5, 8}
	expectedPreOrder := []int{8, 2, 5}
	expectedPostOrder := []int{5, 2, 8}
	root = Remove(root, 4)
	assert.Equal(t, expectedInOrder, InOrder(root))
	assert.Equal(t, expectedPreOrder, PreOrder(root))
	assert.Equal(t, expectedPostOrder, PostOrder(root))
}

func TestRemove_LeftReplacement_DirectWithLeftChild(t *testing.T) {
	//	        8
	//	   4
	//	2     5
	// 1
	givenToInsert := []int{8, 4, 2, 5, 1}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	//	        8
	//	   2
	//	1     5
	root = Remove(root, 4)
	expectedInOrder := []int{1, 2, 5, 8}
	expectedPreOrder := []int{8, 2, 1, 5}
	expectedPostOrder := []int{1, 5, 2, 8}
	assert.Equal(t, expectedInOrder, InOrder(root))
	assert.Equal(t, expectedPreOrder, PreOrder(root))
	assert.Equal(t, expectedPostOrder, PostOrder(root))
}

func TestRemove_LeftReplacement_RightLeaf(t *testing.T) {
	//	        8
	//	    4
	//   1
	// 0   2
	givenToInsert := []int{8, 4, 1, 0, 2}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	//	        8
	//	    2
	//   1
	// 0
	root = Remove(root, 4)
	expectedInOrder := []int{0, 1, 2, 8}
	expectedPreOrder := []int{8, 2, 1, 0}
	expectedPostOrder := []int{0, 1, 2, 8}
	assert.Equal(t, expectedInOrder, InOrder(root))
	assert.Equal(t, expectedPreOrder, PreOrder(root))
	assert.Equal(t, expectedPostOrder, PostOrder(root))
}

func TestRemove_LeftReplacement_RightLeafWithLeftChild(t *testing.T) {
	//	        8
	//	    4
	//   1
	// 0   3
	//    2
	givenToInsert := []int{8, 4, 1, 0, 3, 2}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	//	        8
	//	    3
	//   1
	// 0   2
	root = Remove(root, 4)
	expectedInOrder := []int{0, 1, 2, 3, 8}
	expectedPreOrder := []int{8, 3, 1, 0, 2}
	expectedPostOrder := []int{0, 2, 1, 3, 8}
	assert.Equal(t, expectedInOrder, InOrder(root))
	assert.Equal(t, expectedPreOrder, PreOrder(root))
	assert.Equal(t, expectedPostOrder, PostOrder(root))
}

func TestRemove_RightReplacement_DirectLeaf(t *testing.T) {
	// 2
	//   5
	//    8
	givenToInsert := []int{2, 5, 8}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	// 2
	//   8
	root = Remove(root, 5)
	expectedInOrder := []int{2, 8}
	expectedPreOrder := []int{2, 8}
	expectedPostOrder := []int{8, 2}
	assert.Equal(t, expectedInOrder, InOrder(root))
	assert.Equal(t, expectedPreOrder, PreOrder(root))
	assert.Equal(t, expectedPostOrder, PostOrder(root))
}

func TestRemove_RightReplacement_DirectWithRightChild(t *testing.T) {
	// 2
	//    5
	//      8
	//       9
	givenToInsert := []int{2, 5, 8, 9}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	// 2
	//   8
	//    9
	root = Remove(root, 5)
	expectedInOrder := []int{2, 8, 9}
	expectedPreOrder := []int{2, 8, 9}
	expectedPostOrder := []int{9, 8, 2}
	assert.Equal(t, expectedInOrder, InOrder(root))
	assert.Equal(t, expectedPreOrder, PreOrder(root))
	assert.Equal(t, expectedPostOrder, PostOrder(root))
}

func TestRemove_RightReplacement_Leaf(t *testing.T) {
	// 2
	//    5
	//      8
	//     7
	givenToInsert := []int{2, 5, 8, 7}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	// 2
	//    7
	//      8
	root = Remove(root, 5)
	expectedInOrder := []int{2, 7, 8}
	expectedPreOrder := []int{2, 7, 8}
	expectedPostOrder := []int{8, 7, 2}
	assert.Equal(t, expectedInOrder, InOrder(root))
	assert.Equal(t, expectedPreOrder, PreOrder(root))
	assert.Equal(t, expectedPostOrder, PostOrder(root))
}

func TestRemove_RightReplacement_WithRightChild(t *testing.T) {
	// 2
	//    5
	//       9
	//     7
	//      8
	givenToInsert := []int{2, 5, 9, 7, 8}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	// 2
	//    7
	//       9
	//     8
	root = Remove(root, 5)
	expectedInOrder := []int{2, 7, 8, 9}
	expectedPreOrder := []int{2, 7, 9, 8}
	expectedPostOrder := []int{8, 9, 7, 2}
	assert.Equal(t, expectedInOrder, InOrder(root))
	assert.Equal(t, expectedPreOrder, PreOrder(root))
	assert.Equal(t, expectedPostOrder, PostOrder(root))
}

func TestRemove_BigTree(t *testing.T) {
	givenToInsert := []int{214103, 277791, 32618, 196194, 175494, 786510, 762341, 15292, 348133, 434701, 242996, 716769, 147251, 966489, 152487, 962735, 242016, 431853, 734739, 33135, 232627, 62058, 106905, 370494, 374687, 192143, 65398, 983597, 773537, 124615, 454433, 746607, 497108, 936601, 380245, 215590, 86775, 558615, 14862, 146551, 394259, 393222, 726954, 624114, 516423, 377541, 299198, 754144, 959206, 475227, 599726, 348460, 609710, 2286, 723993, 532039, 515484, 298290, 543561, 31804, 604233, 114377, 656654, 472469, 425773, 839342, 304067, 906501, 448751, 242339, 623804, 369901, 755855, 744683, 524603, 544011, 662969, 654664, 778640, 221201, 254046, 630470, 545375, 929724, 587291, 186061, 138376, 45972, 43473, 885358, 877574, 835993, 700238, 927649, 862305, 522274, 642842, 323806, 546460, 619492, 819572, 373089, 74976, 983929, 501649, 366925, 278711, 652318, 930057, 682990, 522956, 830751, 766854, 784776, 263712, 581000, 32227, 898833, 485163, 499664, 655322, 370016, 114678, 248711, 324673, 431198, 400260, 709368, 634951, 539616, 537605, 763574, 981496, 310964, 347125, 700023, 2172, 253839, 805181, 561432, 943973, 592903, 813103, 161479, 261429, 297651, 839748, 699066, 564279, 357334, 804573, 606644, 445073, 626044, 109971, 482756, 232427, 9568, 550077, 6849, 725305, 743205, 332316, 911117, 770983, 619330, 594820, 254974, 675148, 389625, 609851, 714110, 195395, 739922, 699351, 523240, 922335, 604636, 852813, 719522, 249886, 452821, 671537, 139986, 522060, 467971, 445446, 177657, 141093, 824116, 859584, 805913, 487842, 919018, 404917, 753962, 872994, 864305, 154526, 193376, 773175, 489167, 852730, 393437, 42222, 563876, 343152, 108499, 442741, 91051, 663164, 540108, 156479, 611469, 696950, 377759, 766249, 29585, 639796, 406563, 181850, 648330, 507177, 730389, 486496, 880852, 916157, 206617, 459221, 963892, 944504, 939445, 378548, 770015, 640051, 341622, 610556, 435101, 378967, 540588, 341962, 895604, 371450, 273826, 468123, 821758, 770529, 913260, 503645, 777797, 440541, 533485, 489407, 572395, 183543, 41373, 886064, 818454, 959148, 96999, 258982, 781128, 915312, 639725, 729623, 523128, 882763, 405692, 499286, 570022, 862835, 684376, 649799, 128385, 935341, 186425, 482719, 738443, 505020, 70256, 664698, 923240, 397841, 227119, 930563, 119147, 105013, 534891, 483925, 980729, 811969, 66838, 75801, 141732, 675095, 776354, 510380, 825370, 271596, 397113, 150592, 739336, 846068, 358444, 162563, 178153, 290921, 965514, 940412, 195746, 115986, 589471, 25086, 523619, 88216, 807265, 763777, 834524, 513766, 206432, 742055, 722040, 408645, 773882, 457267, 44248, 534348, 787505, 304301, 610707, 476391, 9480, 391503, 881962, 204974, 189380, 292552, 951953, 59427, 636546, 432279, 346979, 686178, 589135, 700858, 988250, 971406, 819051, 899753, 616552, 545497, 898971, 881297, 492661, 907127, 163572, 286870, 823045, 70688, 893031, 363327, 945375, 467185, 156435, 709635, 454497, 658249, 827053, 228177, 674884, 703487, 183505, 781929, 211843, 352109, 558864, 345796, 720050, 104817, 780679, 117643, 134505, 937931, 481865, 71608, 4504, 848037, 808643, 112162, 923695, 193831, 477149, 841115, 258796, 163593, 606124, 552983, 84207, 258044, 106774, 282934, 71223, 647792, 703912, 655140, 716187, 651585, 753273, 857984, 471032, 379180, 777439, 529869, 503634, 830423, 127704, 302506, 397321, 206196, 483484, 302435, 145818, 580394, 422650, 126152, 415188, 782657, 400299, 60087, 451171, 214954, 530783, 186439, 296424, 991699, 499897, 387639, 822710, 102213, 688173, 208890, 388669, 818517, 963843, 802115, 531367, 180261, 371456, 462921, 230209, 322668, 688478, 854326, 462136, 733938, 467128, 736537, 865527, 132675, 835497, 684793, 715945, 68439, 598216, 701855, 340826, 646639, 373529, 975829, 356141, 491764, 695432, 909574, 912042, 528321, 691101, 788363, 102169, 763131, 945684, 936382, 431689, 200386, 345792, 964809, 117252, 503175, 557673, 738890, 683296, 157288, 506661, 957483, 771278, 168559, 275792, 475838, 556881, 436812, 344359, 767293, 217690, 834922, 668708, 610962, 903090, 261559, 255333, 345069, 848826, 455758, 354141, 354962, 228644, 37657, 76739, 901304, 576700, 151996, 397504, 11084, 453115, 441362, 260997, 23657, 933338, 245362, 401158, 308908, 368679, 148232, 923765, 705635, 820188, 96600, 940837, 689963, 147918, 266364, 962820, 490362, 48749, 381163, 86496, 567021, 403265, 198788, 132352, 757134, 378283, 570523, 205044, 375643, 208634, 754975, 515909, 725109, 181390, 438149, 959376, 743216, 357144, 102842, 920827, 824677, 304738, 853501, 646649, 856213, 208327, 43376, 553453, 281389, 691366, 519990, 605677, 263004, 130832, 446469, 539448, 822664, 897513, 832730, 711808, 870158, 418458, 769021, 101747, 307016, 990445, 752819, 46170, 316725, 727982, 443414, 644015, 26447, 532096, 38281, 448720, 706781, 948329, 402507, 610669, 735419, 535847, 332899, 208554, 391492, 682746, 817070, 213619, 155651, 358895, 660283, 919810, 504831, 20207, 590021, 295834, 739624, 196817, 52649, 612164, 976468, 214769, 439748, 604340, 441404, 25573, 704355, 608329, 480439, 381293, 169814, 624225, 407262, 888657, 47742, 323386, 895189, 559761, 231173, 616050, 883140, 204989, 894378, 321864, 152191, 35051, 478639, 197108, 176976, 343847, 226419, 347978, 465417, 580465, 182363, 571645, 91113, 603744, 883339, 95169, 976387, 348436, 805775, 5126, 76133, 385753, 300200, 612854, 960958, 711105, 598529, 694602, 709500, 511607, 212815, 179414, 108245, 135804, 476071, 863321, 56712, 59590, 976302, 221135, 456532, 865857, 435004, 393254, 9558, 850962, 341882, 674737, 981426, 660826, 784307, 69104, 631367, 198367, 841577, 741004, 893348, 293976, 496423, 262361, 304387, 552866, 870528, 735499, 535425, 532641, 374247, 797236, 980560, 814783, 651892, 717857, 892600, 780291, 234121, 192380, 758737, 520866, 511275, 781160, 187567, 493857, 907805, 617394, 482970, 886222, 624627, 836311, 132972, 811621, 352661, 733167, 224327, 979172, 396094, 12420, 783501, 361020, 28459, 935496, 127653, 165041, 858098, 673098, 988306, 283714, 563698, 390141, 620364, 80082, 540052, 699330, 958848, 154050, 186018, 320447, 850177, 128389, 899781, 878624, 389794, 93100, 447924, 416813, 563011, 484401, 371406, 201366, 743855, 412643, 33432, 233821, 739484, 754146, 941601, 503306, 848522, 615851, 430141, 801392, 836136, 945411, 1753, 144519, 350604, 891147, 815774, 859712, 932983, 105682, 227661, 594997, 960259, 889320, 579985, 151334, 196252, 944086, 311721, 528707, 768869, 599013, 462774, 877243, 610104, 859259, 994086, 404328, 229249, 74546, 429376, 229138, 454650, 936538, 106109, 757089, 319346, 937184, 803031, 282869, 419207, 735677, 109701, 790863, 394052, 984562, 867859, 593205, 646693, 322314, 533181, 25026, 97206, 896877, 311468, 980580, 56182, 687268, 898488, 402651, 841685, 330348, 412016, 91002, 742581, 990884, 517871, 347965, 662155, 921080, 858600, 591029, 765590, 68624, 478833, 106489, 162409, 287189, 272826, 403586, 839669, 356248, 224551, 195255, 131908, 574604, 110941, 580451, 910094, 311844, 100148, 65637, 813292, 731608, 281062, 265028, 680534, 145360, 24713, 941829, 553517, 766998, 973094, 42765, 43178, 709075, 107213, 880430, 161161, 370665, 815936, 767461, 486503, 62064, 880952, 900297, 942172, 805140, 810108, 317980, 735617, 400418, 812798, 46538, 363350, 983663, 356880, 190807, 872360, 328993, 129769, 799876, 19445, 124999, 878669, 647718, 276296, 251888, 620944, 762383, 837040, 765508, 431219, 212240, 912467, 66359, 370661, 252553, 232696, 10869, 840700, 960691, 661330, 779863, 362145, 930856, 247099, 111658, 156380, 87978, 911957, 286689, 691757, 766454, 723992, 631845, 946088, 325699, 605783, 207751, 845924, 692957, 341083, 229566, 417572, 103719, 720993, 379873, 985527, 177319, 966069, 33273, 993989, 328988, 602332, 450744, 33607, 969877, 779206, 432634, 441846, 521822, 259204, 553576, 183328, 765323, 67706, 81201, 628595, 374780, 373939, 876017, 410953, 212850, 632757, 973166, 569560, 118484, 341308, 221851, 482255, 312363}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	const givenToRemove int = 631845
	root = Remove(root, givenToRemove)

	slices.Sort(givenToInsert)
	givenToDeleteIdx, _ := slices.BinarySearch(givenToInsert, givenToRemove)
	expectedInOrder := append(givenToInsert[:givenToDeleteIdx], givenToInsert[givenToDeleteIdx+1:]...)
	assert.Equal(t, expectedInOrder, InOrder(root))
}
