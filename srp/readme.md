## Single Responsibility Principle
* The single responsibility principle states that a class or module should have only one reason to change.
* In other words, a type should have a single responsibility, making the code easier to understand and maintain.

### Example
* Let's have a struct User where we have two methods, GetFullName() and Save()
```
type User struct {
    FirstName string
    LastName string
}

func (u *User) GetFullName() string {
    return u.FirstName + " " + u.LastName
}

func (u *User) Save() error {
    // have the db connection
    // save the user to the db
}
```

* In this case the struct User have two responsibilities managing the user data and saving the data to database.
* In order to adhere SRP we should seprate these responsibilities
```
type User struct {
    FirstName string
    LastName string
}

func (u *User) GetFullName() string {
    return u.FirstName + " " + u.LastName
}

type UserDB struct {
    // database connection
    db *mysql.DB
}

func (udb *UserDB) Save(u *User) error {
    // save the user to the db
}
```