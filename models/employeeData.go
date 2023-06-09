package models


type Employee struct {        
    ID          int         `json:"id" pg:",pk"`
    Name        string      `json:"name"`
    Leavetype   string      `json:"leavetype"`
    Fromdate    string      `json:"fromdate"`
    Todate      string      `json:"todate"`
    Teamname    string      `json:"teamname"`
    File        string      `json:"file"`
    Reporter    string      `json:"reporter"`

}

