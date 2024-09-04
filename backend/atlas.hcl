data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "ariga.io/atlas-provider-gorm",
    "load",
    "--path", "./models",
    "--dialect", "postgres", 
  ]
}

env "develop" {
  src = data.external_schema.gorm.url
  dev = "postgresql://jureg:password@127.0.0.1:5432/cideclasse?sslmode=disable"
  url = "postgresql://jureg:password@127.0.0.1:5432/cideclasse?sslmode=disable"
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}
