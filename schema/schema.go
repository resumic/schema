package schema

type locationSchema struct {
	Lat  float64 `jsonschema:"lat"`
	Long float64 `jsonschema:"long"`
}

type coreSchema struct {
	Name              string         `jsonschema:"name,description:e.g. John Doe"`
	Title             string         `jsonschema:"title,description:e.g. Software Engineer"`
	Image             string         `jsonschema:"image,description:e.g. example.com/Abcxyz - [URL (as per RFC 3986) to a image in JPEG or PNG format]"`
	Email             string         `jsonschema:"email,description:e.g. lucas@example.com,format:email"`
	Phone             string         `jsonschema:"phone,description:e.g. 912-217-7923 - [Phone numbers are stored as strings so use any format you like]"`
	URL               string         `jsonschema:"url,description:e.g. http://www.example.com/my-slides/,format:uri"`
	Summary           string         `jsonschema:"summary,description:Write a short 2-3 sentence biography about yourself"`
	CurrentLocation   locationSchema `jsonschema:"currentLocation,description:Select the location where you currently live.,format:location"`
	PermanentLocation locationSchema `jsonschema:"permanentLocation,description:Select the location where you permanently live.,format:location"`
}

type workSchema struct {
	Name        string   `jsonschema:"name,description:e.g. XYZ Inc. - [Company name]"`
	Description string   `jsonschema:"description,description:e.g. A social media company - [Description of the companies primary activity]"`
	Position    string   `jsonschema:"position,description:e.g. Software Engineer - [Position at the company]"`
	Location    string   `jsonschema:"location,description:e.g. Germany - [Location for this activity],format:location"`
	URL         string   `jsonschema:"url,description:e.g. http://xyz.example.com - [Related link to the company website],format:uri"`
	StartDate   string   `jsonschema:"startDate,description:e.g. 2017-06-28 - [resume.json uses the ISO 8601 date standard],format:date"`
	EndDate     string   `jsonschema:"endDate,description:e.g. 2018-12-29 - [resume.json uses the ISO 8601 date standard],format:date"`
	Summary     string   `jsonschema:"summary,description:Give an overview of your responsibilities at the company"`
	Highlights  []string `jsonschema:"highlights,description:Specify multiple accomplishments,items_description:e.g. Worked with mobile team at Twitter to develop remote debugging tools for mobile browsers"`
}

type educationSchema struct {
	Institution string         `jsonschema:"institution,description:e.g. XYZ Institute of Technology - [Add institute name]"`
	Location    locationSchema `jsonschema:"location,description:e.g. Germany - [Location for this institution],format:location"`
	Area        string         `jsonschema:"area,description:e.g. Engineering"`
	StudyType   string         `jsonschema:"studyType,description:e.g. Bachelor"`
	StartDate   string         `jsonschema:"startDate,description:e.g. 2017-06-29 - [resume.json uses the ISO 8601 date standard],format:date"`
	EndDate     string         `jsonschema:"endDate,description:e.g. 2013-06-29 - [resume.json uses the ISO 8601 date standard],format:date"`
	Score       struct {
		Scoretype  string `jsonschema:"scoretype,description:eg. GPA or PERCENTAGE - [Add score type]"`
		Scorevalue string `jsonschema:"scorevalue,description:eg. 3.4/4.0 - [Add obtained score/total score]"`
	} `jsonschema:"score,additionalProperties"`
	Courses []string `jsonschema:"courses,description:List notable courses/subjects,items_description:e.g. CS302 - Introduction to Algorithms - [Add course name]"`
	Honors  []string `jsonschema:"honors,description:List education honours,items_description:e.g. Magna Cum Laude"`
}

type volunteerSchema struct {
	Organization string         `jsonschema:"organization,description:e.g. Xyz "`
	Position     string         `jsonschema:"position,description:e.g. Open Source Contributor - [Contribution type]"`
	Location     locationSchema `jsonschema:"location,description:e.g. Germany - [Location for this activity],format:location"`
	URL          string         `jsonschema:",description:e.g. http://xyz.example.com - [Related link to support volunteer experience],format:uri"`
	StartDate    string         `jsonschema:"startDate,description:e.g. 2014-06-29 - [resume.json uses the ISO 8601 date standard],format:date"`
	EndDate      string         `jsonschema:"endDate,description:e.g. 2017-06-29 - [resume.json uses the ISO 8601 date standard] ,format:date"`
	Summary      string         `jsonschema:"summary,description:Give an overview of your responsibilities at the company"`
	Highlights   []string       `jsonschema:"highlights,description:Specify accomplishments and achievements,items_description:e.g Invited as a speaker in Xyzcon'17"`
}

type publicationSchema struct {
	Name        string `jsonschema:"name,description:e.g. Deep learning and Artificial Intelligence"`
	publisher   string `jsonschema:"publisher,description:e.g. XYZ Computer Magazine"`
	ReleaseDate string `jsonschema:"releaseDate,description:e.g. 2015-08-01 - [resume.json uses the ISO 8601 date standard]"`
	Resources   []struct {
		URL   string `jsonschema:"url,description:e.g. http://www.example.com/my-example-slides/,format:uri"`
		Label string `jsonschema:"label,description:e.g Slides"`
	} `jsonschema:"resources,description:Specify multiple resources with label"`
	URL     string `jsonschema:"url,description:e.g. http://www.computer.org.example.com/csdl/mags/co/2015/10/rx069-abs.html,format:uri"`
	Summary string `jsonschema:"summary,description:e.g. Discussion of the advent of deep learning and artificial intelligence - [Short summary of publication]"`
}

type legalSchema struct {
	Name            string `jsonschema:"name,description:e.g. XYZ's patent on LZW compression a fundamental part of the widely used GIF graphics format - [Add document name]"`
	legalType       string `jsonschema:"legalType,description:e.g. Patent Trademark Copyright - [Give the type of this document]"`
	Description     string `jsonschema:"description,description:Give a brief description about this document"`
	ApplicationDate string `jsonschema:"applicationDate,description:e.g. 2015-08-01 - [resume.json uses the ISO 8601 date standard],format:date"`
	GrantDate       string `jsonschema:"grantDate,description:e.g. 2016-09-01 - [resume.json uses the ISO 8601 date standard],format:date"`
	EndDate         string `jsonschema:"endDate,description:e.g. 2020-09-03 - [resume.json uses the ISO 8601 date standard],format:date"`
	Resources       []struct {
		URL   string `jsonschema:"url,description:e.g. http://www.example.com/my-patent-slides/,format:uri"`
		Label string `jsonschema:"label,description:e.g Slides"`
	} `jsonschema:"resources,description:Specify multiple resources with label"`
	IDNumber string `jsonschema:"idNumber,description:e.g. JP2004369746A - [Add the application number or Id Number]  "`
}

type skillSchema struct {
	Name    string `jsonschema:"name,description:e.g. Web Development"`
	Keyword []struct {
		Name  string  `jsonschema:"name,description:e.g. HTML - [Add the skill name]"`
		Score float64 `jsonschema:"score,description:e.g. 20 - [Score for the skill name]"`
	} `jsonschema:"keyword,description:List some keywords pertaining to the skill,items_additionalProperties"`
}

type awardSchema struct {
	Title   string `jsonschema:"title,description:e.g. Awarded Software Process Achievement Award "`
	Date    string `jsonschema:"date,description:e.g. 2016-06-12 - [resume.json uses the ISO 8601 date standard],format:date"`
	Awarder string `jsonschema:"awarder,description:e.g. IEEE"`
	Summary string `jsonschema:"summary,description:e.g. Received for my work in Deep learning and AI"`
}

type projectSchema struct {
	Name        string         `jsonschema:"name,description:e.g. File Transfer application - [Name of the project]"`
	Location    locationSchema `jsonschema:"location,description:e.g France - [Location for this activity],format:location"`
	Description string         `jsonschema:"description,description:e.g. Developed a client and server based application - [Short summary of project]"`
	Highlights  []string       `jsonschema:"highlights,description:Specify multiple features,items_description:e.g. used Java AWT and Swing for client side userinterface"`
	Keywords    []string       `jsonschema:"keywords,description:Specify special elements involved,items_description:e.g. Java"`
	StartDate   string         `jsonschema:"startDate,description:e.g. 2017-06-29 - [resume.json uses the ISO 8601 date standard],format:date"`
	EndDate     string         `jsonschema:"endDate,description:e.g. 2017-06-29 - [resume.json uses the ISO 8601 date standard] ,format:date"`
	Resources   []struct {
		URL   string `jsonschema:"url,description:e.g. http://www.example.com/my-awesome-slides/"`
		Label string `jsonschema:"label,description:e.g Slides"`
	} `jsonschema:"resources,description:Specify multiple resources with label"`
	URL    string   `jsonschema:"url,description:e.g. http://www.example.org/csdl/mags/co/1996/10/rx069-abs.html,format:uri"`
	Roles  []string `jsonschema:"roles,description:e.g. http://www.example.org/csdl/mags/co/1996/10/rx069-abs.html,items_description:e.g. Team Lead Speaker Writer"`
	Entity string   `jsonschema:"entity,description:e.g. 'greenpeace' 'corporationXYZ' - [Relevant company/entity affiliations]"`
	Type   string   `jsonschema:"type,description:e.g. 'volunteering' 'presentation' 'talk' 'application' 'conference'"`
}

type certificateSchema struct {
	Code          string `jsonschema:"code,description:e.g. 1Z0-062"`
	Name          string `jsonschema:"name,description:e.g. XYZ Certified Application Specialist (MCAS) - [Add the certificate name]"`
	Website       string `jsonschema:"website,description:Link to issuing authority's description of the certificate,format:uri"`
	Verification  string `jsonschema:"verification,description:External candidate verification URL,format:uri"`
	GrantDate     string `jsonschema:"grantDate,description:e.g. 2017-06-29 - [resume.json uses the ISO 8601 date standard],format:date"`
	Score         string `jsonschema:"score,description:e.g. 95% - [Exam result (PASS/FAIL 100% 100)]"`
	EndDate       string `jsonschema:"endDate,description:e.g. 2017-06-29 - [resume.json uses the ISO 8601 date standard],format:date"`
	DoesNotExpire bool   `jsonschema:"doesNotExpire,format:checkbox"`
}

type referenceSchema struct {
	Name      string `jsonschema:"name,description:e.g. Stephan Mark"`
	Company   string `jsonschema:"company,description:e.g. Xyz"`
	Position  string `jsonschema:"position,description:e.g. Senior Software Engineer"`
	Reference string `jsonschema:"reference,description:e.g. Joe blogs was a great employee who turned up to work at least once a week. He exceeded my expectations when it came to doing nothing."`
}

type languageSchema struct {
	Language string  `jsonschema:"language,description:e.g. English - [Name of language]"`
	Score    float64 `jsonschema:"score,description:e.g. 20 - [Score for the language]"`
}

type interestSchema struct {
	Name     string   `jsonschema:"name,description:e.g. Machine Learning"`
	Keywords []string `jsonschema:"keywords,description:Specify keywords associated with the particular interest,items_description:e.g. Neural Networks Convoluted Neural Networks"`
}

type metaSchema struct {
	Canonical    string `jsonschema:"canonical,description:URL (as per RFC 3986) to latest version of this document"`
	Version      string `jsonschema:"version,description:e.g. v1.0.0 - [A version field which follows semver]"`
	LastModified string `jsonschema:"lastModified,description:e.g. 2017-06-29T15:53:00 - [resume.json uses the ISO 8601 date standard],format:date"`
}

type Schema struct {
	Core         coreSchema          `jsonschema:"core,additionalProperties"`
	Work         []workSchema        `jsonschema:"work,items_additionalProperties"`
	Education    []educationSchema   `jsonschema:"education,items_additionalProperties"`
	Volunteer    []volunteerSchema   `jsonschema:"volunteer,items_additionalProperties"`
	Publications []publicationSchema `jsonschema:"publications,description:Specify your publications,items_additionalProperties"`
	Legal        []legalSchema       `jsonschema:"legal,description:Specify your labels,items_additionalProperties"`
	Skills       []skillSchema       `jsonschema:"skills,description:List out your professional skill-set,items_additionalProperties"`
	Awards       []awardSchema       `jsonschema:"awards,description:Specify any awards you have received throughout your professional career,items_additionalProperties"`
	Projects     []projectSchema     `jsonschema:"projects,description:Specify career projects,items_additionalProperties"`
	Certificates []certificateSchema `jsonschema:"certificate,items_additionalProperties"`
	References   []referenceSchema   `jsonschema:"references,description:List references you have received,items_additionalProperties"`
	Languages    []languageSchema    `jsonschema:"languages,description:List any other languages you speak,items_additionalProperties"`
	Interests    []interestSchema    `jsonschema:"interests,items_additionalProperties"`
	Meta         metaSchema          `jsonschema:"meta,description:The schema version and any other tooling configuration lives here,additionalProperties"`
}
