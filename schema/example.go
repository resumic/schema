package schema

func NewExample() Schema {
	location := locationSchema{
		Lat:  35.712758,
		Long: 51.392114,
	}

	score := scoreSchema{
		Type:  "GPA",
		Value: "3.4",
	}

	resource := resourceSchema{
		URL:   "http://www.example.com/my-example-slides/",
		Label: "Slides",
	}

	core := coreSchema{
		Name:              "John Doe",
		Title:             "Software Engineer",
		Image:             "example.com/Abcxyz",
		Email:             "lucas@example.com",
		Phone:             "912-217-7923",
		URL:               "http://www.example.com/",
		Summary:           "The man who sold the world!",
		CurrentLocation:   location,
		PermanentLocation: location,
	}

	work := workSchema{
		Name:        "XYZ Inc",
		Description: "A social media company",
		Position:    "Software Engineer",
		Location:    location,
		URL:         "http://xyz.example.com",
		StartDate:   "2017-12-29",
		EndDate:     "2018-12-29",
		Summary:     "Developing and maintaining the company website using syna",
		Highlights:  []string{"Worked with mobile team at Twitter to develop remote debugging tools for mobile browsers"},
	}

	education := educationSchema{
		Institution: "XYZ Institute of Technology",
		Location:    location,
		Area:        "Engineering",
		StudyType:   "Bachelor",
		StartDate:   "2017-06-28",
		EndDate:     "2013-06-28",
		Score:       score,
		Courses:     []string{"CS302 - Introduction to Algorithms "},
		Honors:      []string{"Magna Cum Laude"},
	}

	volunteer := volunteerSchema{
		Organization: "Xyz",
		Position:     "Open Source Contributor",
		Location:     location,
		URL:          "http://xyz.example.com",
		StartDate:    "2014-06-29",
		EndDate:      "2017-06-29",
		Summary:      "Frontend developer",
		Highlights:   []string{"Invited as a speaker in Xyzcon'17"},
	}

	publication := publicationSchema{
		Name:        "Deep learning and Artificial Intelligence",
		Publisher:   "XYZ, Computer Magazine",
		ReleaseDate: "2015-08-01",
		Resources:   []resourceSchema{resource},
		URL:         "http://www.computer.org.example.com/csdl/mags/co/2015/10/rx069-abs.html;format:uri",
		Summary:     "Discussion of the advent of deep learning and artificial intelligence",
	}

	legal := legalSchema{
		Name:            "XYZ's patent on LZW compression, a fundamental part of the widely used GIF graphics format",
		LegalType:       "Patent, Trademark, Copyright",
		Description:     "Some legal document!",
		ApplicationDate: "2015-08-01",
		GrantDate:       "2016-09-01",
		EndDate:         "2020-09-03",
		Resources:       []resourceSchema{resource},
		IDNumber:        "JP2004369746A",
	}

	skill := skillSchema{
		Name: "Web Development",
		Score: scoreSchema{
			Type:  "Percentage",
			Value: "80",
		},
		Keyword: []string{"HTML"},
	}

	award := awardSchema{
		Title:   "Awarded Software Process Achievement Award",
		Date:    "2016-06-12",
		Awarder: "IEEE",
		Summary: "Received for my work in Deep learning and AI",
	}

	project := projectSchema{
		Name:        "File Transfer application",
		Location:    location,
		Description: "Developed a client and server based application",
		Highlights:  []string{"used Java AWT and Swing for client side userinterface"},
		Keywords:    []string{"Java"},
		StartDate:   "2016-06-29",
		EndDate:     "2017-03-02",
		Resources:   []resourceSchema{resource},
		URL:         "http://www.example.org/csdl/mags/co/1996/10/rx069-abs.html",
		Roles:       []string{"Team Lead", "Speaker", "Writer"},
		Entity:      "greenpeace",
		Type:        "volunteering",
	}

	certificate := certificateSchema{
		Code:          "1Z0-062",
		Name:          "XYZ Certified Application Specialist (MCAS)",
		Website:       "http://www.example.org",
		Verification:  "http://www.example.org",
		GrantDate:     "2017-06-29",
		Score:         score,
		EndDate:       "2017-06-29",
		DoesNotExpire: true,
	}

	reference := referenceSchema{
		Name:      "Stephan Mark",
		Company:   "Xyz",
		Position:  "Senior Software Engineer",
		Reference: "Joe blogs was a great employee, who turned up to work at least once a week. He exceeded my expectations when it came to doing nothing.",
	}

	language := languageSchema{
		Language: "English",
		Score: scoreSchema{
			Type:  "objective",
			Value: "fluent",
		},
	}

	interest := interestSchema{
		Name:     "Machine Learning",
		Keywords: []string{"Neural Networks"},
	}

	meta := metaSchema{
		Canonical:    "",
		Version:      "v1.0.0",
		LastModified: "2017-06-29T15:53:01",
	}

	schema := Schema{
		Core:         core,
		Work:         []workSchema{work},
		Education:    []educationSchema{education},
		Volunteer:    []volunteerSchema{volunteer},
		Publications: []publicationSchema{publication},
		Legal:        []legalSchema{legal},
		Skills:       []skillSchema{skill},
		Awards:       []awardSchema{award},
		Projects:     []projectSchema{project},
		Certificates: []certificateSchema{certificate},
		References:   []referenceSchema{reference},
		Languages:    []languageSchema{language},
		Interests:    []interestSchema{interest},
		Meta:         meta,
	}

	return schema
}
