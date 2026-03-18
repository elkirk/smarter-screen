package dispatcher

import "testing"

// isBulky tests

func TestIsBulky_ByVolume(t *testing.T) {
	if !isBulky(100, 100, 100) {
		t.Error("expected bulky for volume = 1,000,000")
	}
}

func TestIsBulky_JustUnderVolume(t *testing.T) {
	if isBulky(99, 100, 100) {
		t.Error("expected not bulky for volume = 990,000")
	}
}

func TestIsBulky_WidthAtThreshold(t *testing.T) {
	if !isBulky(150, 1, 1) {
		t.Error("expected bulky for width = 150")
	}
}

func TestIsBulky_HeightAtThreshold(t *testing.T) {
	if !isBulky(1, 150, 1) {
		t.Error("expected bulky for height = 150")
	}
}

func TestIsBulky_LengthAtThreshold(t *testing.T) {
	if !isBulky(1, 1, 150) {
		t.Error("expected bulky for length = 150")
	}
}

func TestIsBulky_AllDimensionsJustUnder(t *testing.T) {
	// 99 * 99 * 99 = 970,299 -> not bulky
	if isBulky(99, 99, 99) {
		t.Error("expected not bulky for 99x99x99")
	}
}

func TestIsBulky_DimensionsUnderButVolumeOver(t *testing.T) {
	// 149 * 149 * 149 = 3,307,949 -> bulky by volume
	if !isBulky(149, 149, 149) {
		t.Error("expected bulky for 149x149x149 (volume exceeds threshold)")
	}
}

// isHeavy tests

func TestIsHeavy_AtThreshold(t *testing.T) {
	if !isHeavy(20) {
		t.Error("expected heavy for mass = 20")
	}
}

func TestIsHeavy_AboveThreshold(t *testing.T) {
	if !isHeavy(25) {
		t.Error("expected heavy for mass = 25")
	}
}

func TestIsHeavy_JustUnder(t *testing.T) {
	if isHeavy(19.99) {
		t.Error("expected not heavy for mass = 19.99")
	}
}

func TestIsHeavy_Light(t *testing.T) {
	if isHeavy(1) {
		t.Error("expected not heavy for mass = 1")
	}
}

// Sort tests

func TestStandardPackage(t *testing.T) {
	result := Sort(50, 50, 50, 10)
	if result != Standard {
		t.Errorf("expected STANDARD, got %s", result)
	}
}

func TestHeavyOnly(t *testing.T) {
	result := Sort(50, 50, 50, 20)
	if result != Special {
		t.Errorf("expected SPECIAL, got %s", result)
	}
}

func TestBulkyByVolume(t *testing.T) {
	result := Sort(100, 100, 100, 10)
	if result != Special {
		t.Errorf("expected SPECIAL, got %s", result)
	}
}

func TestBulkyByDimensionWidth(t *testing.T) {
	result := Sort(150, 1, 1, 10)
	if result != Special {
		t.Errorf("expected SPECIAL, got %s", result)
	}
}

func TestBulkyByDimensionHeight(t *testing.T) {
	result := Sort(1, 150, 1, 10)
	if result != Special {
		t.Errorf("expected SPECIAL, got %s", result)
	}
}

func TestBulkyByDimensionLength(t *testing.T) {
	result := Sort(1, 1, 150, 10)
	if result != Special {
		t.Errorf("expected SPECIAL, got %s", result)
	}
}

func TestRejectedBulkyAndHeavy(t *testing.T) {
	result := Sort(150, 150, 150, 25)
	if result != Rejected {
		t.Errorf("expected REJECTED, got %s", result)
	}
}

func TestRejectedBulkyByVolumeAndHeavy(t *testing.T) {
	result := Sort(100, 100, 100, 20)
	if result != Rejected {
		t.Errorf("expected REJECTED, got %s", result)
	}
}

func TestExactVolumeThreshold(t *testing.T) {
	// 100 * 100 * 100 = 1,000,000 -> bulky
	result := Sort(100, 100, 100, 19)
	if result != Special {
		t.Errorf("expected SPECIAL, got %s", result)
	}
}

func TestJustUnderVolumeThreshold(t *testing.T) {
	// 99 * 100 * 100 = 990,000 -> not bulky
	result := Sort(99, 100, 100, 19)
	if result != Standard {
		t.Errorf("expected STANDARD, got %s", result)
	}
}

func TestExactMassThreshold(t *testing.T) {
	result := Sort(10, 10, 10, 20)
	if result != Special {
		t.Errorf("expected SPECIAL, got %s", result)
	}
}

func TestJustUnderMassThreshold(t *testing.T) {
	result := Sort(10, 10, 10, 19)
	if result != Standard {
		t.Errorf("expected STANDARD, got %s", result)
	}
}

func TestDimensionJustUnder150(t *testing.T) {
	// 149 * 149 * 149 = 3,307,949 -> still bulky by volume
	result := Sort(149, 149, 149, 19)
	if result != Special {
		t.Errorf("expected SPECIAL, got %s", result)
	}
}

func TestSmallDimensionsNotBulky(t *testing.T) {
	// 99 * 99 * 99 = 970,299 -> not bulky
	result := Sort(99, 99, 99, 19)
	if result != Standard {
		t.Errorf("expected STANDARD, got %s", result)
	}
}
