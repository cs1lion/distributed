package grades

func init() {
	students = []Student{
		{
			ID:        1,
			FirstName: "Nick",
			LastName:  "Carter",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 85,
				},
				{
					Title: "Final Exam",
					Type:  GradeExam,
					Score: 94,
				},
				{
					Title: "Test 1",
					Type:  GradeTest,
					Score: 82,
				},
			},
		},
		{
			ID:        2,
			FirstName: "Jane",
			LastName:  "Doe",
			Grades: []Grade{
				{Title: "Quiz 2", Type: GradeQuiz, Score: 78},
				{Title: "Midterm Exam", Type: GradeExam, Score: 88},
				{Title: "Test 2", Type: GradeTest, Score: 75},
			},
		},
		{
			ID:        3,
			FirstName: "John",
			LastName:  "Smith",
			Grades: []Grade{
				{Title: "Quiz 3", Type: GradeQuiz, Score: 90},
				{Title: "Project Evaluation", Type: GradeExam, Score: 85},
				{Title: "Test 3", Type: GradeTest, Score: 80},
			},
		},
		{
			ID:        4,
			FirstName: "Alice",
			LastName:  "Johnson",
			Grades: []Grade{
				{
					Title: "Group Project",
					Type:  GradeQuiz,
					Score: 89,
				},
				{
					Title: "Class Presentation",
					Type:  GradeTest,
					Score: 91,
				},
				{
					Title: "Midterm Test",
					Type:  GradeExam,
					Score: 75,
				},
			},
		},
	}
}
