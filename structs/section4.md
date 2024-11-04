# 4-42. Defining Structs

# 4-43. Declaring Structs

# 4-44. Updating Struct Values

    var alex person
    fmt.Println(alex)
    // output: { }

    fmt.Printf("%+v", alex)
    // output: {firstName: lastName:}
     ---

    var alex person
    alex.firstName = "Alex"
    alex.lastName = "Anderson"
    fmt.Println(alex) or fmt.Printf("%v", alex)
    //{Alex Anderson}
    fmt.Printf("%+v", alex)
    //{firstName:Alex lastName:Anderson}
