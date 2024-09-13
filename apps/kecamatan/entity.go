package kecamatan

type Kecamatan struct {
	KecamatanName string `db:"kecamatan_name"`
	KelurahanName string `db:"kelurahan_name"`
	Paslon1       int32  `db:"paslon1"`
	Paslon2       int32  `db:"paslon2"`
	Paslon3       int32  `db:"paslon3"`
	Paslon4       int32  `db:"paslon4"`
	SuaraSah      int32  `db:"suara_sah"`
	SuaraTidakSah int32  `db:"suara_tidak_sah"`
	TotalSuara    int32  `db:"total_suara"`
	Persentase    int32  `db:"persentase"`
	TotalVoters   int32  `db:"total_voters"`
	TotalTPS      int32  `db:"total_tps"`
	Sudah         int32  `db:"sudah"`
	Belum         int32  `db:"belum"`
	PP            int32  `db:"pp"`
	Code          string `db:"code"`
}

type Kelurahan struct {
	KecamatanName string `db:"kecamatan_name"`
	KelurahanName string `db:"kelurahan_name"`
	Paslon1       int32  `db:"paslon1"`
	Paslon2       int32  `db:"paslon2"`
	Paslon3       int32  `db:"paslon3"`
	Paslon4       int32  `db:"paslon4"`
	SuaraSah      int32  `db:"suara_sah"`
	SuaraTidakSah int32  `db:"suara_tidak_sah"`
	TotalSuara    int32  `db:"total_suara"`
	Persentase    int32  `db:"persentase"`
	TotalVoters   int32  `db:"total_voters"`
	TotalTPS      int32  `db:"total_tps"`
	Sudah         int32  `db:"sudah"`
	Belum         int32  `db:"belum"`
	PP            int32  `db:"pp"`
	Code          string `db:"code"`
}

func (k Kecamatan) ToKecamatanListResponse() KecamatanListResponse {
	return KecamatanListResponse{
		KecamatanName: k.KecamatanName,
		Paslon1:       k.Paslon1,
		Paslon2:       k.Paslon2,
		Paslon3:       k.Paslon3,
		Paslon4:       k.Paslon4,
		SuaraSah:      k.SuaraSah,
		SuaraTidakSah: k.SuaraTidakSah,
		TotalVoters:   k.TotalVoters,
		TotalTPS:      k.TotalTPS,
		Sudah:         k.Sudah,
		Belum:         k.Belum,
		PP:            k.PP,
		Code:          k.Code,
	}
}

func (k Kelurahan) ToKelurahanListResponse() KelurahanListResponse {
	return KelurahanListResponse{
		KecamatanName: k.KecamatanName,
		KelurahanName: k.KelurahanName,
		Paslon1:       k.Paslon1,
		Paslon2:       k.Paslon2,
		Paslon3:       k.Paslon3,
		Paslon4:       k.Paslon4,
		SuaraSah:      k.SuaraSah,
		SuaraTidakSah: k.SuaraTidakSah,
		TotalVoters:   k.TotalVoters,
		TotalTPS:      k.TotalTPS,
		Sudah:         k.Sudah,
		Belum:         k.Belum,
		PP:            k.PP,
		Code:          k.Code,
	}
}

func (k Kecamatan) ToGetVoterKecamatanResponse() GetVoterKecamatanResponse {
	return GetVoterKecamatanResponse{
		KecamatanName: k.KecamatanName,
		Paslon1:       k.Paslon1,
		Paslon2:       k.Paslon2,
		Paslon3:       k.Paslon3,
		Paslon4:       k.Paslon4,
		SuaraTidakSah: k.SuaraTidakSah,
	}
}

func (k Kecamatan) ToAllVoterResponse() AllVoterResponse {
	return AllVoterResponse{
		TotalSuara: k.TotalSuara,
		Persentase: k.Persentase,
	}
}
