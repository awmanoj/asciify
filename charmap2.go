package main

// Based on an earlier article (reference not available anymore)
var charmap2 = map[int]string{
	0:   "Q", // "Q"	char
	1:   "Q", //"Q"	char
	2:   "Q", //"Q"	char
	3:   "Q", //"Q"	char
	4:   "Q", //"Q"	char
	5:   "Q", //"Q"	char
	6:   "Q", //"Q"	char
	7:   "Q", //"Q"	char
	8:   "Q", //"Q"	char
	9:   "Q", //"Q"	char
	10:  "Q", //"Q"	char
	11:  "R", //"R"	char
	12:  "R", //"R"	char
	13:  "R", //"R"	char
	14:  "R", //"R"	char
	15:  "R", //"R"	char
	16:  "R", //"R"	char
	17:  "R", //"R"	char
	18:  "R", //"R"	char
	19:  "R", //"R"	char
	20:  "R", //"R"	char
	21:  "R", //"R"	char
	22:  "R", //"R"	char
	23:  "M", //"M"	char
	24:  "M", //"M"	char
	25:  "M", //"M"	char
	26:  "M", //"M"	char
	27:  "M", //"M"	char
	28:  "Z", //"Z"	char
	29:  "Z", //"Z"	char
	30:  "Z", //"Z"	char
	31:  "Z", //"Z"	char
	32:  "D", //"D"	char
	33:  "D", //"D"	char
	34:  "D", //"D"	char
	35:  "D", //"D"	char
	36:  "E", //"E"	char
	37:  "E", //"E"	char
	38:  "E", //"E"	char
	39:  "E", //"E"	char
	40:  "b", //"b"	char
	41:  "b", //"b"	char
	42:  "b", //"b"	char
	43:  "b", //"b"	char
	44:  "0", //"0"	char
	45:  "0", //"0"	char
	46:  "0", //"0"	char
	47:  "0", //"0"	char
	48:  "0", //"0"	char
	49:  "9", //"9"	char
	50:  "9", //"9"	char
	51:  "9", //"9"	char
	52:  "9", //"9"	char
	53:  "p", //"p"	char
	54:  "p", //"p"	char
	55:  "p", //"p"	char
	56:  "p", //"p"	char
	57:  "P", //"P"	char
	58:  "P", //"P"	char
	59:  "P", //"P"	char
	60:  "P", //"P"	char
	61:  "X", //"X"	char
	62:  "X", //"X"	char
	63:  "X", //"X"	char
	64:  "X", //"X"	char
	65:  "S", //"S"	char
	66:  "S", //"S"	char
	67:  "S", //"S"	char
	68:  "S", //"S"	char
	69:  "h", //"h"	char
	70:  "h", //"h"	char
	71:  "h", //"h"	char
	72:  "h", //"h"	char
	73:  "h", //"h"	char
	74:  "F", //"F"	char
	75:  "F", //"F"	char
	76:  "F", //"F"	char
	77:  "F", //"F"	char
	78:  "1", //"1"	char
	79:  "1", //"1"	char
	80:  "1", //"1"	char
	81:  "1", //"1"	char
	82:  "2", //"2"	char
	83:  "2", //"2"	char
	84:  "2", //"2"	char
	85:  "2", //"2"	char
	86:  "f", //"f"	char
	87:  "f", //"f"	char
	88:  "f", //"f"	char
	89:  "f", //"f"	char
	90:  "t", //"t"	char
	91:  "t", //"t"	char
	92:  "t", //; //"t"	char
	93:  "t", //"t"	char
	94:  "t", //"t"	char
	95:  "U", //"U"	char
	96:  "U", //"U"	char
	97:  "U", //"U"	char
	98:  "U", //"U"	char
	99:  "j", //"j"	char
	100: "j", //"j"	char
	101: "j", //"j"	char
	102: "j", //"j"	char
	103: "J", //"J"	char
	104: "J", //"J"	char
	105: "J", //"J"	char
	106: "J", //"J"	char
	107: "J", //"J"	char
	108: "J", //"J"	char
	109: "J", //"J"	char
	110: "J", //"J"	char
	111: "Y", //"Y"	char
	112: "Y", //"Y"	char
	113: "Y", //"Y"	char
	114: "Y", //"Y"	char
	115: "c", //"c"	char
	116: "c", //"c"	char
	117: "c", //"c"	char
	118: "c", //"c"	char
	119: "c", //"c"	char
	120: "L", //"L"	char
	121: "L", //"L"	char
	122: "L", //"L"	char
	123: "L", //"L"	char
	124: "L", //"L"	char
	125: "L", //"L"	char
	126: "L", // ; //"L"	char
	127: "L", //"L"	char
	128: "7", //"7"	char
	129: "7", //"7"	char
	130: "7", //"7"	char
	131: "7", //"7"	char
	132: "7", //"7"	char
	133: "7", //"7"	char
	134: "7", //"7"	char
	135: "7", //"7"	char
	136: "7", //"7"	char
	137: "7", //"7"	char
	138: "7", //"7"	char
	139: "7", // "7"	char
	140: "7", //"7"	char
	141: "r", //"r"	char
	142: "r", //"r"	char
	143: "r", //"r"	char
	144: "r", //"r"	char
	145: "r", //"r"	char
	146: "r", //"r"	char
	147: "r", //"r"	char
	148: "r", //"r"	char
	149: "r", //"r"	char
	150: "r", //"r"	char
	151: "r", //"r"	char
	152: "r", //"r"	char
	153: "r", //"r"	char
	154: "r", //"r"	char
	155: "r", //"r"	char
	156: "r", //"r"	char
	157: "i", //"i"	char
	158: "i", //"i"	char
	159: "i", //"i"	char
	160: "i", //"i"	char
	161: "i", //"i"	char
	162: "i", //"i"	char
	163: "i", // "i"	char
	164: "i", //"i"	char
	165: "i", //"i"	char
	166: "i", //"i"	char
	167: "i", //"i"	char
	168: "i", // "i"	char
	169: "i", //"i"	char
	170: ":", //":"	char
	171: ":", //":"	char
	172: ":", // ":"	char
	173: ":", //":"	char
	174: ":", //":"	char
	175: ":", //":"	char
	176: ":", //":"	char
	177: ":", //":"	char
	178: ":", //":"	char
	179: ":", //":"	char
	180: ":", //":"	char
	181: ":", //":"	char
	182: ":", //":"	char
	183: ":", //":"	char
	184: ":", //":"	char
	185: ":", //":"	char
	186: ":", //":"	char
	187: ":", //":"	char
	188: ":", //":"	char
	189: ":", //":"	char
	190: ":", //":"	char
	191: ":", //":"	char
	192: ":", //":"	char
	193: ":", //":"	char
	194: ":", // ":"	char
	195: ".", //"."	char
	196: ".", //"."	char
	197: ".", // "."	char
	198: ".", //"."	char
	199: ".", //"."	char
	200: ".", //"."	char
	201: ".", //"."	char
	202: ".", //"."	char
	203: ".", //"."	char
	204: ".", // "."	char
	205: ".", //"."	char
	206: ".", //"."	char
	207: ".", //"."	char
	208: ".", //"."	char
	209: ".", //"."	char
	210: ".", //"."	char
	211: ".", //"."	char
	212: ".", //"."	char
	213: ".", //"."	char
	214: ".", // "."	char
	215: ".", //"."	char
	216: ".", //"."	char
	217: ".", //"."	char
	218: ".", //"."	char
	219: ".", //"."	char
	220: ".", //"."	char
	221: ".", //"."	char
	222: ".", //"."	char
	223: ".", //"."	char
	224: ",", //","	char
	225: ",", //","	char
	226: ",", //","	char
	227: ",", //","	char
	228: ",", //","	char
	229: ",", //","	char
	230: ",", //","	char
	231: ",", //","	char
	232: ",", //","	char
	233: ",", //","	char
	234: ",", //","	char
	235: ",", //","	char
	236: ",", //","	char
	237: ",", //","	char
	238: ",", //","	char
	239: ",", //","	char
	240: ",", //","	char
	241: ",", //","	char
	242: ",", // ","	char
	243: ",", //","	char
	244: " ", //","	char
	245: " ", //" "	char
	246: " ", //" "	char
	247: " ", //" "	char
	248: " ", //" "	char
	249: " ", //" "	char
	250: " ", //" "	char
	251: " ", //" "	char
	252: " ", //" "	char
	253: " ", //" "	char
	254: " ", //" "	char
	255: " ", //" "	char
}
