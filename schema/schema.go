package schema

type locationSchema struct {
	Lat  float64 `json:"lat" jsonschema:"lat"`
	Long float64 `json:"long" jsonschema:"long"`
}

type scoreSchema struct {
	Type  string `json:"type" jsonschema:"type;description:eg. GPA or PERCENTAGE - [Add score type]"`
	Value string `json:"value" jsonschema:"value;description:eg. 3.4/4.0 - [Add obtained score/total score]"`
}

type resourceSchema struct {
	URL   string `json:"url" jsonschema:"url;description:e.g. http://www.example.com/my-example-slides/;format:uri"`
	Label string `json:"label" jsonschema:"label;description:e.g Slides"`
}

type coreSchema struct {
	Name              string         `json:"name" jsonschema:"name;description:e.g. John Doe"`
	Title             string         `json:"title" jsonschema:"title;description:e.g. Software Engineer"`
	Image             string         `json:"image" jsonschema:"image;description:e.g. example.com/Abcxyz - [URL (as per RFC 3986) to a image in JPEG or PNG format]"`
	Email             string         `json:"email" jsonschema:"email;description:e.g. lucas@example.com;format:email"`
	Phone             string         `json:"phone" jsonschema:"phone;description:e.g. 912-217-7923 - [Phone number]"`
	URL               string         `json:"url" jsonschema:"url;description:e.g. http://www.example.com/my-slides/;format:uri"`
	Summary           string         `json:"summary" jsonschema:"summary;description:Write a short 2-3 sentence biography about yourself"`
	CurrentLocation   locationSchema `json:"currentLocation" jsonschema:"currentLocation;description:Select the location where you currently live.;format:location"`
	PermanentLocation locationSchema `json:"permanentLocation" jsonschema:"permanentLocation;description:Select the location where you permanently live.;format:location"`
}

type workSchema struct {
	Name        string         `json:"name" jsonschema:"name;description:e.g. XYZ Inc. - [Company name]"`
	Description string         `json:"description" jsonschema:"description;description:e.g. A social media company - [Description of the companies primary activity]"`
	Position    string         `json:"position" jsonschema:"position;description:e.g. Software Engineer - [Position at the company]"`
	Location    locationSchema `json:"location" jsonschema:"location;description:e.g. Germany - [Location for this activity];format:location"`
	URL         string         `json:"url" jsonschema:"url;description:e.g. http://xyz.example.com - [Related link to the company website];format:uri"`
	StartDate   string         `json:"startDate" jsonschema:"startDate;description:e.g. 2017-06-28 - [resume.json uses the ISO 8601 date standard];format:date"`
	EndDate     string         `json:"endDate" jsonschema:"endDate;description:e.g. 2018-12-29 - [resume.json uses the ISO 8601 date standard];format:date"`
	Summary     string         `json:"summary" jsonschema:"summary;description:Give an overview of your responsibilities at the company"`
	Highlights  []string       `json:"highlights" jsonschema:"highlights;description:Specify multiple accomplishments;items_description:e.g. Worked with mobile team at Twitter to develop remote debugging tools for mobile browsers"`
}

type educationSchema struct {
	Institution string         `json:"institution" jsonschema:"institution;description:e.g. XYZ Institute of Technology - [Add institute name]"`
	Location    locationSchema `json:"location" jsonschema:"location;description:e.g. Germany - [Location for this institution];format:location"`
	Area        string         `json:"area" jsonschema:"area;description:e.g. Engineering"`
	StudyType   string         `json:"studyType" jsonschema:"studyType;description:e.g. Bachelor"`
	StartDate   string         `json:"startDate" jsonschema:"startDate;description:e.g. 2017-06-29 - [resume.json uses the ISO 8601 date standard];format:date"`
	EndDate     string         `json:"endDate" jsonschema:"endDate;description:e.g. 2013-06-29 - [resume.json uses the ISO 8601 date standard];format:date"`
	Score       scoreSchema    `json:"score" jsonschema:"score;additionalProperties"`
	Courses     []string       `json:"courses" jsonschema:"courses;description:List notable courses/subjects;items_description:e.g. CS302 - Introduction to Algorithms - [Add course name]"`
	Honors      []string       `json:"honors" jsonschema:"honors;description:List education honours;items_description:e.g. Magna Cum Laude"`
}

type volunteerSchema struct {
	Organization string         `json:"organization" jsonschema:"organization;description:e.g. Xyz "`
	Position     string         `json:"position" jsonschema:"position;description:e.g. Open Source Contributor - [Contribution type]"`
	Location     locationSchema `json:"location" jsonschema:"location;description:e.g. Germany - [Location for this activity];format:location"`
	URL          string         `json:"url" jsonschema:"url;description:e.g. http://xyz.example.com - [Related link to support volunteer experience];format:uri"`
	StartDate    string         `json:"startDate" jsonschema:"startDate;description:e.g. 2014-06-29 - [resume.json uses the ISO 8601 date standard];format:date"`
	EndDate      string         `json:"endDate" jsonschema:"endDate;description:e.g. 2017-06-29 - [resume.json uses the ISO 8601 date standard] ;format:date"`
	Summary      string         `json:"summary" jsonschema:"summary;description:Give an overview of your responsibilities at the company"`
	Highlights   []string       `json:"highlights" jsonschema:"highlights;description:Specify accomplishments and achievements;items_description:e.g Invited as a speaker in Xyzcon'17"`
}

type publicationSchema struct {
	Name        string           `json:"name" jsonschema:"name;description:e.g. Deep learning and Artificial Intelligence"`
	Publisher   string           `json:"publisher" jsonschema:"publisher;description:e.g. XYZ, Computer Magazine"`
	ReleaseDate string           `json:"releaseDate" jsonschema:"releaseDate;description:e.g. 2015-08-01 - [resume.json uses the ISO 8601 date standard]"`
	Resources   []resourceSchema `json:"resources" jsonschema:"resources;description:Specify multiple resources with label"`
	URL         string           `json:"url" jsonschema:"url;description:e.g. http://www.computer.org.example.com/csdl/mags/co/2015/10/rx069-abs.html;format:uri"`
	Summary     string           `json:"summary" jsonschema:"summary;description:e.g. Discussion of the advent of deep learning and artificial intelligence - [Short summary of publication]"`
}

type legalSchema struct {
	Name            string           `json:"name" jsonschema:"name;description:e.g. XYZ's patent on LZW compression, a fundamental part of the widely used GIF graphics format - [Add document name]"`
	LegalType       string           `json:"legalType" jsonschema:"legalType;description:e.g. Patent, Trademark, Copyright - [Give the type of this document]"`
	Description     string           `json:"description" jsonschema:"description;description:Give a brief description about this document"`
	ApplicationDate string           `json:"applicationDate" jsonschema:"applicationDate;description:e.g. 2015-08-01 - [resume.json uses the ISO 8601 date standard];format:date"`
	GrantDate       string           `json:"grantDate" jsonschema:"grantDate;description:e.g. 2016-09-01 - [resume.json uses the ISO 8601 date standard];format:date"`
	EndDate         string           `json:"endDate" jsonschema:"endDate;description:e.g. 2020-09-03 - [resume.json uses the ISO 8601 date standard];format:date"`
	Resources       []resourceSchema `json:"resources" jsonschema:"resources;description:Specify multiple resources with label"`
	IDNumber        string           `json:"idNumber" jsonschema:"idNumber;description:e.g. JP2004369746A - [Add the application number or Id Number]  "`
}

type skillSchema struct {
	Name    string      `json:"name" jsonschema:"name;description:e.g. Web Development"`
	Score   scoreSchema `json:"score" jsonschema:"score;additionalProperties;description:Score for the skill name"`
	Keyword []string    `json:"keyword" jsonschema:"keyword;description:List some keywords pertaining to the skill"`
}

type awardSchema struct {
	Title   string `json:"title" jsonschema:"title;description:e.g. Awarded Software Process Achievement Award "`
	Date    string `json:"date" jsonschema:"date;description:e.g. 2016-06-12 - [resume.json uses the ISO 8601 date standard];format:date"`
	Awarder string `json:"awarder" jsonschema:"awarder;description:e.g. IEEE"`
	Summary string `json:"summary" jsonschema:"summary;description:e.g. Received for my work in Deep learning and AI"`
}

type projectSchema struct {
	Name        string           `json:"name" jsonschema:"name;description:e.g. File Transfer application - [Name of the project]"`
	Location    locationSchema   `json:"location" jsonschema:"location;description:e.g France - [Location for this activity];format:location"`
	Description string           `json:"description" jsonschema:"description;description:e.g. Developed a client and server based application - [Short summary of project]"`
	Highlights  []string         `json:"highlights" jsonschema:"highlights;description:Specify multiple features;items_description:e.g. used Java AWT and Swing for client side userinterface"`
	Keywords    []string         `json:"keywords" jsonschema:"keywords;description:Specify special elements involved;items_description:e.g. Java"`
	StartDate   string           `json:"startDate" jsonschema:"startDate;description:e.g. 2017-06-29 - [resume.json uses the ISO 8601 date standard];format:date"`
	EndDate     string           `json:"endDate" jsonschema:"endDate;description:e.g. 2017-06-29 - [resume.json uses the ISO 8601 date standard] ;format:date"`
	Resources   []resourceSchema `json:"resources" jsonschema:"resources;description:Specify multiple resources with label"`
	URL         string           `json:"url" jsonschema:"url;description:e.g. http://www.example.org/csdl/mags/co/1996/10/rx069-abs.html;format:uri"`
	Roles       []string         `json:"roles" jsonschema:"roles;description:Specify your role on this project or in company;items_description:e.g. Team Lead, Speaker, Writer"`
	Entity      string           `json:"entity" jsonschema:"entity;description:e.g. 'greenpeace', 'corporationXYZ' - [Relevant company/entity affiliations]"`
	Type        string           `json:"type" jsonschema:"type;description:e.g. 'volunteering', 'presentation', 'talk', 'application', 'conference'"`
}

type certificateSchema struct {
	Code          string      `json:"code" jsonschema:"code;description:e.g. 1Z0-062"`
	Name          string      `json:"name" jsonschema:"name;description:e.g. XYZ Certified Application Specialist (MCAS) - [Add the certificate name]"`
	Website       string      `json:"website" jsonschema:"website;description:Link to issuing authority's description of the certificate;format:uri"`
	Verification  string      `json:"verification" jsonschema:"verification;description:External candidate verification URL;format:uri"`
	GrantDate     string      `json:"grantDate" jsonschema:"grantDate;description:e.g. 2017-06-29 - [resume.json uses the ISO 8601 date standard];format:date"`
	Score         scoreSchema `json:"score" jsonschema:"score;additionalProperties;description:Exam result (PASS/FAIL, 100%, 100)"`
	EndDate       string      `json:"endDate" jsonschema:"endDate;description:e.g. 2017-06-29 - [resume.json uses the ISO 8601 date standard];format:date"`
	DoesNotExpire bool        `json:"doesNotExpire" jsonschema:"doesNotExpire;format:checkbox"`
}

type referenceSchema struct {
	Name      string `json:"name" jsonschema:"name;description:e.g. Stephan Mark"`
	Company   string `json:"company" jsonschema:"company;description:e.g. Xyz"`
	Position  string `json:"position" jsonschema:"position;description:e.g. Senior Software Engineer"`
	Reference string `json:"reference" jsonschema:"reference;description:e.g. Joe blogs was a great employee, who turned up to work at least once a week. He exceeded my expectations when it came to doing nothing."`
}

type languageSchema struct {
	Language string      `json:"language" jsonschema:"language;description:e.g. English - [Name of language]"`
	Score    scoreSchema `json:"score" jsonschema:"score;additionalProperties;description:Score for the language"`
}

type interestSchema struct {
	Name     string   `json:"name" jsonschema:"name;description:e.g. Machine Learning"`
	Keywords []string `json:"keywords" jsonschema:"keywords;description:Specify keywords associated with the particular interest;items_description:e.g. Neural Networks, Convoluted Neural Networks"`
}

type metaSchema struct {
	Canonical    string `json:"canonical" jsonschema:"canonical;description:URL (as per RFC 3986) to latest version of this document"`
	Version      string `json:"version" jsonschema:"version;description:e.g. v1.0.0 - [A version field which follows semver]"`
	LastModified string `json:"lastModified" jsonschema:"lastModified;description:e.g. 2017-06-29T15:53:00 - [resume.json uses the ISO 8601 date standard];format:date-time"`
}

type Schema struct {
	Core         coreSchema          `json:"core" jsonschema:"core;additionalProperties"`
	Work         []workSchema        `json:"work" jsonschema:"work;items_additionalProperties"`
	Education    []educationSchema   `json:"education" jsonschema:"education;items_additionalProperties"`
	Volunteer    []volunteerSchema   `json:"volunteer" jsonschema:"volunteer;items_additionalProperties"`
	Publications []publicationSchema `json:"publications" jsonschema:"publications;description:Specify your publications;items_additionalProperties"`
	Legal        []legalSchema       `json:"legal" jsonschema:"legal;description:Specify your labels;items_additionalProperties"`
	Skills       []skillSchema       `json:"skills" jsonschema:"skills;description:List out your professional skill-set;items_additionalProperties"`
	Awards       []awardSchema       `json:"awards" jsonschema:"awards;description:Specify any awards you have received throughout your professional career;items_additionalProperties"`
	Projects     []projectSchema     `json:"projects" jsonschema:"projects;description:Specify career projects;items_additionalProperties"`
	Certificates []certificateSchema `json:"certificate" jsonschema:"certificate;items_additionalProperties"`
	References   []referenceSchema   `json:"references" jsonschema:"references;description:List references you have received;items_additionalProperties"`
	Languages    []languageSchema    `json:"languages" jsonschema:"languages;description:List any other languages you speak;items_additionalProperties"`
	Interests    []interestSchema    `json:"interests" jsonschema:"interests;items_additionalProperties"`
	Meta         metaSchema          `json:"meta" jsonschema:"meta;description:The schema version and any other tooling configuration lives here;additionalProperties"`
}
