package model

type Student struct {
	ID          uint
	StudentName string
	ClassID     uint
	ClassInfo   Class `gorm:"foreignKey:ClassID;->"`
}

type Class struct {
	ID        uint
	ClassName string
}

type Subject struct {
	ID          uint
	SubjectName string
}

type StudentSubject struct {
	StudentID uint
	SubjectID uint
}