package main

import(

  "database/sql"
  "log"
  //github yazılıyor



)


func main()  {
	sql.open("mysql","root:12345678@/demodb")
	if err!=nil{
        panic(err.Error())

	}

	defer db.Close()

	createStatement:="'users'('ID'int(11)NOT NULL KİSİ_KAYİT,'NAME' varchar(45) NOT NULL,'EMAİL'varchar(45) NOT NULL,'MESSAGE' varchar(45) NOT NULL )"
	_,err = db.Ecec("CREATE TABLE IF NOT EXISTS"+ createStatement)

	if err !=nill{
        log.Fatal(err)

	}
	//veri ekleme 
	res,err:= db.Exec("INSERT INTO users(NAME,EMAİL,MESSAGE)VALUES('','','')")

	if err!=nil{
		log.Fatal(err)
	}
   rowCount,err:=res.RowsAffected()
   if err!=nill{
	   log.Fatal(err)
   }

	log.Printf("Inserted %d rows",rowCount)


	lastusers,err1:=res.LasrInsertId()
	if err1 !=nil{
		log.Fatal(err1)
	}
      log.Printf("Inserted ID:%",lastID)



    
var(

     NAME string
	 EMAİL string
	 MESSAGE string
)


// eklenen kayıt getirme 

rows,err:=db.Query("SELECT*FROM users")
if err!=nil{
	log.Fatal(err)
}

for rows.Next(){

	err=rows.Scan(&NAME,&EMAİL,&MESSAGE)
	if err != nil{

		log.Fatal(err)
	}
     // verileri gösterme 
	log.Printf("kisi bilgileri: %q",NAME,EMAİL,MESSAGE)
}

//alternatif

iff err=rows.Err(); err!=nill{
	log.Fatal(err)
}

rows.Close()

}