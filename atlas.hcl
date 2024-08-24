data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "ariga.io/atlas-provider-gorm",
    "load",
    "--path", "./models",
    "--dialect", "mysql", 
  ]
}

env "develop" {
  src = data.external_schema.gorm.url
  dev = "mysql://root:mysql@127.0.0.1:3306/mysqldb"
  url = "mysql://root:mysql@127.0.0.1:3306/mysqldb"
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}
