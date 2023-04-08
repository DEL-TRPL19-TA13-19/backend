package ahp

func GetRatioIndex() [15]float64 {
	return [15]float64{0, 0, 0.58, 0.9, 1.12, 1.24, 1.32, 1.41, 1.46, 1.49, 1.51, 1.48, 1.56, 1.57, 1.59}
}

func TimbulanSampahSubCriteria() map[string]float64 {
	return map[string]float64{
		"Jaringan Jalan":      0.282,
		"Perumahan":           0.438,
		"Fasilitas Komersial": 0.222,
		"Fasilitas Umum":      0.174,
		"Fasilitas Sosial":    0.103,
		"Ruang Terbuka":       0.038}
}

func JarakTPASubCriteria() map[string]float64 {
	return map[string]float64{
		"Alternatif berada di jangkauan layanan TPA":               0.669,
		"Alternatif berada di batas terjauh jangkauan layanan TPA": 0.267,
		"Alternatif tidak berada di jangkauan TPA":                 0.064}
}

func KondisiTanahSubCriteria() map[string]float64 {
	return map[string]float64{
		"Tanah keras tidak memiliki unsur organik dan unsur hara dan kedap air": 0.503,
		"Tanah keras tidak memiliki unsur organik dan unsur hara":               0.267,
		"Tanah keras tidak memiliki salah satu unsur hara atau unsur organik":   0.260,
		"Tanah keras memiliki unsur organik dan unsur hara":                     0.134,
		"Bukan tanah keras": 0.068,
	}
}

func JarakPemukimanSubCriteria() map[string]float64 {
	return map[string]float64{
		"0m-100m":   0.503,
		"101m-200m": 0.260,
		"201m-300m": 0.134,
		"301m-400m": 0.068,
		"401m-500m": 0.035,
	}
}

func JarakSungaiSubCriteria() map[string]float64 {
	return map[string]float64{
		"Lokasi memenuhi peli banjir":          0.669,
		"Lokasi memenuhi sebagian peli banjir": 0.267,
		"Lokasi tidak memenuhi peli banjir":    0.064,
	}
}

func PartisipasiMasyarakatSubCriteria() map[string]float64 {
	return map[string]float64{
		"<20% Masyarakat Setuju":    0.503,
		"21%-40% Masyarakat Setuju": 0.260,
		"41%-60% Masyarakat Setuju": 0.134,
		"61%-81% Masyarakat Setuju": 0.068,
		">80% Masyarakat Setuju":    0.035,
	}
}

func CakupanRumahSubCriteria() map[string]float64 {
	return map[string]float64{
		"<40 Rumah":     0.503,
		"41-80 Rumah":   0.260,
		"81-120 Rumah":  0.134,
		"121-160 Rumah": 0.068,
		">160 Rumah":    0.035,
	}
}

func AksesibilitasSubCriteria() map[string]float64 {
	return map[string]float64{
		"Kondisi jalan bagus dan bisa dilewati kendaraan pengangkut sampah":                                                                                    0.669,
		"Kondisi jalan bagus, tetapi tidak bisa dilewati kendaraan pengangkut sampah atau jalan tidak bagus, tetapi bisa dilewati kendaraan pengangkut sampah": 0.267,
		"Kondisi jalan tidak bagus dan tidak bisa dilewati kendaraan pengangkut sampah":                                                                        0.064,
	}
}
