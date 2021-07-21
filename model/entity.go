package model

type (
	ConDescription struct {
		Id int
		Name string
		NameSpace string
		NodeName string
		Status string
		RestartCount int
		StartedAt string
	}
 	ConPage struct{
		Amount int
		CD []ConDescription
    }
)
