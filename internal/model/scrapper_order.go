package model

type ScrapperOrderSelection struct {
    ID string
    Element string
    Class *string
    Attribute *string
    ElementId *string
}



type ScrapperOrder struct {
    Id  int    
    Url string 
}
