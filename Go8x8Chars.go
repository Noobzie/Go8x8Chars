package Go8x8Chars

// Functions starting with c are Capital letters, s Small letters

// GetCharacter returns teh character corresponding to the char string.
func GetCharacter(char string) [][]int {
	switch char {
	case "A":
		return cA()
	case "B":
		return cB()
	case "C":
		return cC()
	case "D":
		return cD()
	case "E":
		return cE()
	case "F":
		return cF()
	case "G":
		return cG()
	case "H":
		return cH()
	case "I":
		return cI()
	case "J":
		return cJ()
	case "K":
		return cK()
	case "L":
		return cL()
	case "M":
		return cM()
	case "N":
		return cN()
	case "O":
		return cO()
	case "P":
		return cP()
	case "Q":
		return cQ()
	case "R":
		return cR()
	case "S":
		return cS()
	case "T":
		return cT()
	case "U":
		return cU()
	case "V":
		return cV()
	case "W":
		return cW()
	case "X":
		return cX()
	case "Y":
		return cY()
	case "Z":
		return cZ()
	case "1":
		return one()
	case "2":
		return two()
	case "3":
		return three()
	case "4":
		return four()
	case "5":
		return five()
	case "6":
		return six()
	case "7":
		return seven()
	case "8":
		return eight()
	case "9":
		return nine()
	case "0":
		return zero()
	default:
		return space()
	}
}

// GetCharacterWithBorder return a character with on the right and bottom a line of 0's
func GetCharacterWithBorder(char string) [][]int {
	character := GetCharacter(char)
	for i := 0; i < 8; i++ {
		currentChar := character[i]
		currentChar = append(currentChar, 0)
		character[i] = currentChar
	}
	bottomRow := [][]int{{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	character = append(character, bottomRow...)
	return character
}

func cA() [][]int {
	A := [][]int{
		{0, 0, 0, 1, 1, 0, 0, 0},
		{0, 0, 1, 0, 0, 1, 0, 0},
		{0, 1, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 1, 0},
		{0, 1, 1, 1, 1, 1, 1, 0},
		{0, 1, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 1, 0},
		{1, 1, 1, 0, 0, 1, 1, 1},
	}
	return A
}

func cB() [][]int {
	B := [][]int{
		{1, 1, 1, 1, 1, 1, 0, 0},
		{0, 1, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 1, 0, 0},
		{0, 1, 1, 1, 1, 1, 1, 0},
		{0, 1, 0, 0, 0, 0, 0, 1},
		{0, 1, 0, 0, 0, 0, 0, 1},
		{0, 1, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 0},
	}
	return B
}

func cC() [][]int {
	C := [][]int{
		{0, 0, 1, 1, 1, 1, 0, 1},
		{0, 1, 0, 0, 0, 0, 1, 1},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{0, 1, 0, 0, 0, 0, 1, 1},
		{0, 0, 1, 1, 1, 1, 0, 1},
	}
	return C
}

func cD() [][]int {
	D := [][]int{
		{1, 1, 1, 1, 1, 1, 0, 0},
		{0, 1, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 0, 1},
		{0, 1, 0, 0, 0, 0, 0, 1},
		{0, 1, 0, 0, 0, 0, 0, 1},
		{0, 1, 0, 0, 0, 0, 0, 1},
		{0, 1, 0, 0, 0, 0, 1, 0},
		{1, 1, 1, 1, 1, 1, 0, 0},
	}
	return D
}

func cE() [][]int {
	E := [][]int{
		{1, 1, 1, 1, 1, 1, 1, 1},
		{0, 1, 0, 0, 0, 0, 0, 1},
		{0, 1, 0, 0, 0, 0, 0, 0},
		{0, 1, 1, 1, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 1},
	}
	return E
}

func cF() [][]int {
	F := [][]int{
		{1, 1, 1, 1, 1, 1, 1, 1},
		{0, 1, 0, 0, 0, 0, 0, 1},
		{0, 1, 0, 0, 0, 0, 0, 1},
		{0, 1, 0, 0, 1, 0, 0, 0},
		{0, 1, 1, 1, 1, 0, 0, 0},
		{0, 1, 0, 0, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0, 0, 0},
	}
	return F
}

func cG() [][]int {
	G := [][]int{
		{0, 0, 1, 1, 1, 1, 1, 1},
		{0, 1, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 1, 1, 1, 1, 1},
		{1, 0, 0, 1, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{0, 1, 0, 0, 0, 0, 0, 1},
		{0, 0, 1, 1, 1, 1, 1, 0},
	}
	return G
}

func cH() [][]int {
	H := [][]int{
		{1, 1, 1, 0, 0, 1, 1, 1},
		{0, 1, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 1, 0},
		{0, 1, 1, 1, 1, 1, 1, 0},
		{0, 1, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 1, 0},
		{1, 1, 1, 0, 0, 1, 1, 1},
	}
	return H
}

func cI() [][]int {
	I := [][]int{
		{1, 1, 1, 1, 1, 1, 1, 0},
		{0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 0},
	}
	return I
}

func cJ() [][]int {
	J := [][]int{
		{0, 0, 1, 1, 1, 1, 1, 0},
		{0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 1, 0, 0},
		{1, 1, 0, 0, 0, 1, 0, 0},
		{1, 0, 0, 0, 0, 1, 0, 0},
		{1, 0, 0, 0, 0, 1, 0, 0},
		{0, 1, 1, 1, 1, 0, 0, 0},
	}
	return J
}

func cK() [][]int {
	K := [][]int{
		{1, 1, 1, 0, 1, 1, 1, 1},
		{0, 1, 0, 0, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 0, 0, 0},
		{0, 1, 1, 1, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 1, 0, 0},
		{0, 1, 0, 0, 0, 0, 1, 0},
		{1, 1, 1, 0, 1, 1, 1, 1},
	}
	return K
}

func cL() [][]int {
	L := [][]int{
		{1, 1, 1, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 1},
		{0, 1, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 1},
	}
	return L
}

func cM() [][]int {
	M := [][]int{
		{1, 1, 0, 0, 0, 1, 1, 1},
		{0, 1, 1, 0, 1, 0, 1, 0},
		{0, 1, 0, 1, 0, 0, 1, 0},
		{0, 1, 0, 1, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 1, 0},
		{1, 1, 1, 0, 0, 1, 1, 1},
	}
	return M
}

func cN() [][]int {
	N := [][]int{
		{1, 1, 0, 0, 0, 1, 1, 1},
		{0, 1, 1, 0, 0, 0, 1, 0},
		{0, 1, 0, 1, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 1, 1, 0},
		{0, 1, 0, 0, 0, 0, 1, 0},
		{1, 1, 1, 0, 0, 1, 1, 1},
	}
	return N
}

func cO() [][]int {
	O := [][]int{
		{0, 0, 1, 1, 1, 1, 0, 0},
		{0, 1, 0, 0, 0, 0, 1, 0},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{0, 1, 0, 0, 0, 0, 1, 0},
		{0, 0, 1, 1, 1, 1, 0, 0},
	}
	return O
}

func cP() [][]int {
	P := [][]int{
		{1, 1, 1, 1, 1, 1, 1, 0},
		{0, 1, 0, 0, 0, 0, 0, 1},
		{0, 1, 0, 0, 0, 0, 0, 1},
		{0, 1, 0, 0, 0, 0, 0, 1},
		{0, 1, 1, 1, 1, 1, 1, 0},
		{0, 1, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0, 0, 0},
	}
	return P
}

func cQ() [][]int {
	Q := [][]int{
		{0, 1, 1, 1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 1, 0, 0, 0, 1},
		{0, 1, 1, 1, 1, 1, 1, 0},
		{0, 0, 0, 1, 0, 0, 0, 0},
	}
	return Q
}

func cR() [][]int {
	R := [][]int{
		{1, 1, 1, 1, 1, 1, 1, 0},
		{0, 1, 0, 0, 0, 0, 0, 1},
		{0, 1, 0, 0, 0, 0, 0, 1},
		{0, 1, 0, 0, 0, 0, 0, 1},
		{0, 1, 1, 1, 1, 1, 1, 0},
		{0, 1, 0, 0, 0, 1, 0, 0},
		{0, 1, 0, 0, 0, 0, 1, 0},
		{1, 1, 1, 0, 0, 1, 1, 1},
	}
	return R
}

func cS() [][]int {
	S := [][]int{
		{0, 1, 1, 1, 1, 1, 0, 1},
		{1, 0, 0, 0, 0, 0, 1, 1},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{0, 1, 1, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 0},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{1, 1, 0, 0, 0, 0, 0, 1},
		{1, 0, 1, 1, 1, 1, 1, 0},
	}
	return S
}

func cT() [][]int {
	T := [][]int{
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 1, 0, 0, 1},
		{0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 1, 1, 1, 0, 0},
	}
	return T
}

func cU() [][]int {
	U := [][]int{
		{1, 1, 1, 0, 0, 1, 1, 1},
		{0, 1, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 1, 0},
		{0, 0, 1, 1, 1, 1, 0, 0},
	}
	return U
}

func cV() [][]int {
	V := [][]int{
		{1, 1, 1, 0, 0, 1, 1, 1},
		{0, 1, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 1, 0},
		{0, 0, 1, 0, 0, 0, 1, 0},
		{0, 0, 1, 0, 0, 1, 0, 0},
		{0, 0, 0, 1, 0, 1, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0},
	}
	return V
}

func cW() [][]int {
	W := [][]int{
		{1, 1, 1, 0, 0, 1, 1, 1},
		{0, 1, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 1, 0, 0, 1, 0},
		{0, 1, 0, 1, 0, 0, 1, 0},
		{0, 1, 0, 1, 0, 0, 1, 0},
		{0, 0, 1, 0, 1, 1, 0, 0},
	}
	return W
}

func cX() [][]int {
	X := [][]int{
		{1, 1, 1, 0, 0, 1, 1, 1},
		{0, 1, 0, 0, 0, 0, 1, 0},
		{0, 0, 1, 0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1, 0, 0, 0},
		{0, 0, 1, 0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0, 1, 0, 0},
		{0, 1, 0, 0, 0, 0, 1, 0},
		{1, 1, 1, 0, 0, 1, 1, 1},
	}
	return X
}

func cY() [][]int {
	Y := [][]int{
		{1, 1, 1, 0, 0, 1, 1, 1},
		{0, 1, 0, 0, 0, 0, 1, 0},
		{0, 0, 1, 0, 0, 1, 0, 0},
		{0, 0, 0, 1, 0, 1, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 1, 1, 1, 0, 0},
	}
	return Y
}

func cZ() [][]int {
	Z := [][]int{
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 1, 0},
		{1, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 1},
		{0, 1, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 1},
	}
	return Z
}

func zero() [][]int {
	zero := [][]int{
		{0, 0, 1, 1, 1, 1, 0, 0},
		{0, 1, 0, 0, 0, 0, 1, 0},
		{1, 0, 0, 0, 0, 1, 0, 1},
		{1, 0, 0, 0, 1, 0, 0, 1},
		{1, 0, 0, 1, 0, 0, 0, 1},
		{1, 0, 1, 0, 0, 0, 0, 1},
		{0, 1, 0, 0, 0, 0, 1, 0},
		{0, 0, 1, 1, 1, 1, 0, 0},
	}
	return zero
}

func one() [][]int {
	one := [][]int{
		{0, 0, 1, 1, 0, 0, 0, 0},
		{0, 1, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1},
	}
	return one
}

func two() [][]int {
	two := [][]int{
		{0, 1, 1, 1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 1, 1, 0},
		{0, 0, 0, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 1},
	}
	return two
}

func three() [][]int {
	three := [][]int{
		{0, 1, 1, 1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 1, 1, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{0, 1, 1, 1, 1, 1, 1, 0},
	}
	return three
}

func four() [][]int {
	four := [][]int{
		{0, 0, 0, 0, 0, 1, 1, 0},
		{0, 0, 0, 0, 1, 0, 1, 0},
		{0, 0, 0, 1, 0, 0, 1, 0},
		{0, 0, 1, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 1, 0},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 1, 1, 1},
	}
	return four
}

func five() [][]int {
	five := [][]int{
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{0, 1, 1, 1, 1, 1, 1, 0},
	}
	return five
}

func six() [][]int {
	six := [][]int{
		{0, 1, 1, 1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{0, 1, 1, 1, 1, 1, 1, 0},
	}
	return six
}

func seven() [][]int {
	seven := [][]int{
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 1, 1, 1, 0, 0, 0},
	}
	return seven
}

func eight() [][]int {
	eight := [][]int{
		{0, 1, 1, 1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{0, 1, 1, 1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{0, 1, 1, 1, 1, 1, 1, 0},
	}
	return eight
}

func nine() [][]int {
	nine := [][]int{
		{0, 1, 1, 1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{0, 1, 1, 1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{0, 1, 1, 1, 1, 1, 1, 0},
	}
	return nine
}

func exclamation() [][]int {
	space := [][]int{
		{0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0},
	}
	return space
}

func space() [][]int {
	space := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}
	return space
}
