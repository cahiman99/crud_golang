package customers

type Customers struct {
	Id            int64
	Name          string `validate:"required" label:"Nama Lengkap"`
	Email         string `validate:"required" label:"Email"`
	NIK           int32  `validate:"required"`
	Gender        string `validate:"required" label:"Jenis Kelamin"`
	Tempat_lahir  string `validate:"required" label:"Tempat Lahir"`
	Tanggal_lahir string `validate:"required" label:"Tanggal Lahir"`
	Alamat        string `validate:"required"`
	No_hp         string `validate:"required" label:"No.HP"`
}
