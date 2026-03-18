package dispatcher

import "testing"

func TestStandardPackage(t *testing.T) {
	result := Sort(50, 50, 50, 10)
	if result != "STANDARD" {
		t.Errorf("expected STANDARD, got %s", result)
	}
}

func TestHeavyOnly(t *testing.T) {
	result := Sort(50, 50, 50, 20)
	if result != "SPECIAL" {
		t.Errorf("expected SPECIAL, got %s", result)
	}
}

func TestBulkyByVolume(t *testing.T) {
	result := Sort(100, 100, 100, 10)
	if result != "SPECIAL" {
		t.Errorf("expected SPECIAL, got %s", result)
	}
}

func TestBulkyByDimensionWidth(t *testing.T) {
	result := Sort(150, 1, 1, 10)
	if result != "SPECIAL" {
		t.Errorf("expected SPECIAL, got %s", result)
	}
}

func TestBulkyByDimensionHeight(t *testing.T) {
	result := Sort(1, 150, 1, 10)
	if result != "SPECIAL" {
		t.Errorf("expected SPECIAL, got %s", result)
	}
}

func TestBulkyByDimensionLength(t *testing.T) {
	result := Sort(1, 1, 150, 10)
	if result != "SPECIAL" {
		t.Errorf("expected SPECIAL, got %s", result)
	}
}

func TestRejectedBulkyAndHeavy(t *testing.T) {
	result := Sort(150, 150, 150, 25)
	if result != "REJECTED" {
		t.Errorf("expected REJECTED, got %s", result)
	}
}

func TestRejectedBulkyByVolumeAndHeavy(t *testing.T) {
	result := Sort(100, 100, 100, 20)
	if result != "REJECTED" {
		t.Errorf("expected REJECTED, got %s", result)
	}
}

func TestExactVolumeThreshold(t *testing.T) {
	// 100 * 100 * 100 = 1,000,000 -> bulky
	result := Sort(100, 100, 100, 19)
	if result != "SPECIAL" {
		t.Errorf("expected SPECIAL, got %s", result)
	}
}

func TestJustUnderVolumeThreshold(t *testing.T) {
	// 99 * 100 * 100 = 990,000 -> not bulky
	result := Sort(99, 100, 100, 19)
	if result != "STANDARD" {
		t.Errorf("expected STANDARD, got %s", result)
	}
}

func TestExactMassThreshold(t *testing.T) {
	result := Sort(10, 10, 10, 20)
	if result != "SPECIAL" {
		t.Errorf("expected SPECIAL, got %s", result)
	}
}

func TestJustUnderMassThreshold(t *testing.T) {
	result := Sort(10, 10, 10, 19)
	if result != "STANDARD" {
		t.Errorf("expected STANDARD, got %s", result)
	}
}

func TestDimensionJustUnder150(t *testing.T) {
	// 149 * 149 * 149 = 3,307,949 -> still bulky by volume
	result := Sort(149, 149, 149, 19)
	if result != "SPECIAL" {
		t.Errorf("expected SPECIAL, got %s", result)
	}
}

func TestSmallDimensionsNotBulky(t *testing.T) {
	// 99 * 99 * 99 = 970,299 -> not bulky
	result := Sort(99, 99, 99, 19)
	if result != "STANDARD" {
		t.Errorf("expected STANDARD, got %s", result)
	}
}
