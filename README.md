# Data Object <a href="https://github.com/gouniverse/dataobject" style="float:right:"><img src="https://gitpod.io/button/open-in-gitpod.svg" alt="Open in Gitpod" loading="lazy"></a>

A data object is a special purpose structure that is designed
to hold data and track the changes to allow efficient 
serialization to a data store.

It follows the following principles:

1. Has a default, non-argument constructor
2. Has a unique identifier (ID) allowing to safely find the object among any other
3. All the fields are private, thus non-modifiable from outside
4. The fields are only accessed via public setter (mutator) and getter (accessor) methods
5. All changes are tracked, and returned on request as a map of type map[string]string
6. All data can be returned on request as a map of type map[string]string

Any object can be considered a data object as long as it adheres to the above principles, 
regardless of its specific implementation.

The implemetation in this repo is just one way to implement the above principles. Other variations are possible to suit specific needs.

The concept is a bit similar to a POJO and a Java Bean.

## Usage

For an object using the above specifications

```
// Create new user with autogenerated default ID
user := NewUser()

// Create new user with already existing data (i.e. from database)
user := NewUserFromData(data)

// returns the ID of the object
id := user.ID()

// example setter method
user.SetFirstName("John")

// example getter method
firstName := user.FirstName()

// find if the object has been modified
isDirty := user.IsDirty()

// reurns the changed data
dataChanged := user.DataChanged()

// reurns all the data
data := user.Data()
```

## Saving data

Saving the data is left to the end user, as it is specific for each data store.

Some stores (i.e. relational databases) allow to only change specific fields,
which is why the DataChanged() getter method should be used to identify
the changed fields, then only save the changes efficiently.

```
func SaveUserChanges(user User) bool {
    if !user.IsDirty() {
        return true
    }

    changedData := user.DataChanged()

    return save(changedData)
}
```

Other stores (i.e. document stores, file stores) save all the fields each time, 
which is where the Data() getter method should be used

```
func SaveFullUser(user User) bool {
    if !user.IsDirty() {
        return true
    }

    allData := user.Data()

    return save(allData)
}
```