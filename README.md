# Owl
Interact with a data persistent layer in a consistent manor.

```
type User struct {
  Name  string
  Email string
}
var users &User

// select
db := owl.Connect("postgres", "db=database,username=username,password=password")
db.Table("users").Select().Where("email = ?", "user@domain.com")
err := db.Execute()

// select specific columns
db := owl.Connect("postgres", "db=database,username=username,password=password")
db.Table("users").Select("name", "email").Where("email = ?", "user@domain.com")
err := db.Execute(&users)

// update
db := owl.Connect("postgres", "db=database,username=username,password=password")
db.Table("users").Update("name", "John Doe").Update("email", "jdoe@company.com")
err := db.Execute()

// delete
db := owl.Connect("postgres", "db=database,username=username,password=password")
db.Table("users").Delete().Where("email = ?", "user@domain.com")
err := db.Execute()

// insert
db := owl.Connect("postgres", "db=database,username=username,password=password")
db.Table("users").Insert("name", "email").Values("Jane Doe", "jdoe@domain.com").Value("Susan Doe", "sdoe@domain.com")
err := db.Execute()
```
