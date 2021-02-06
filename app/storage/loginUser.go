package storage

import (
	"github.com/luis16121013/courseSoft/app/models"
)

func LoginUser(user,pass string,r *MysqlReponsitory) (models.User,error) {
	u := models.User{}
	//sql := "SELECT ut.id_user, ut.nombre, ut.apellido, ut.dni, rut.descripcion, lt.usuario,lt.pwd FROM usersTable ut, loginTable lt, roleUserTable rut WHERE lt.id_user = ut.id_user AND ut.tipo_user = rut.id_role AND lt.usuario = ? AND lt.pwd =?"
	
	sql := "SELECT * FROM loginTable WHERE usuario = ? AND pwd = ?"
	rows,err := r.Query(sql,user,pass)
	
	if err != nil {
		return models.User{},err
	}

	if rows.Next() {
		//rows.Scan(&u.Id,&u.Nombre,&u.Apellido,&u.Dni,&u.TypeUser)
		rows.Scan(&u.Id,&u.Nombre,&u.TypeUser)
	}

	return u,nil
}
