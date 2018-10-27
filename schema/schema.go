package schema

type locationSchema struct {
	Lat  float64 `jsonschema:"lat"`
	Long float64 `jsonschema:"long"`
}

type scoreSchema struct {
	Type  string `jsonschema:"type;description:eg. GPA or PERCENTAGE - [Add score type]"`
	Value string `jsonschema:"value;description:eg. 3.4/4.0 - [Add obtained score/total score]"`
}

type resourceSchema struct {
	URL   string `jsonschema:"url"`
	Label string `jsonschema:"label"`
}

type coreSchema struct {
	Name              string         `jsonschema:"name"`
	Title             string         `jsonschema:"title"`
	Image             string         `jsonschema:"image;description:URL (as per RFC 3986) to a image in JPEG or PNG format"`
	Email             string         `jsonschema:"email"`
	Phone             string         `jsonschema:"phone;description:hone numbers are stored as strings so use any format you like"`
	URL               string         `jsonschema:"url;format:uri"`
	Summary           string         `jsonschema:"summary;description:Write a short 2-3 sentence biography about yourself"`
	CurrentLocation   locationSchema `jsonschema:"currentLocation;description:Select the location where you currently live.;format:location"`
	PermanentLocation locationSchema `jsonschema:"permanentLocation;description:Select the location where you permanently live.;format:location"`
}

type workSchema struct {
	Name        string         `jsonschema:"name;description:Company name"`
	Description string         `jsonschema:"description;description:Description of the companies primary activity"`
	Position    string         `jsonschema:"position;description:Position at the company"`
	Location    locationSchema `jsonschema:"location;description:Location for this activity;format:location"`
	URL         string         `jsonschema:"url;description:Related link to the company website;format:uri"`
	StartDate   string         `jsonschema:"startDate;description:resume.json uses the ISO 8601 date standard;format:date"`
	EndDate     string         `jsonschema:"endDate;description:resume.json uses the ISO 8601 date standard;format:date"`
	Summary     string         `jsonschema:"summary;description:Give an overview of your responsibilities at the company"`
	Highlights  []string       `jsonschema:"highlights;description:Specify multiple accomplishments;items_description:"`
}

type educationSchema struct {
	Institution string         `jsonschema:"institution;description:Add institute name"`
	Location    locationSchema `jsonschema:"location;description:Location for this institution;format:location"`
	Area        string         `jsonschema:"area"`
	StudyType   string         `jsonschema:"studyType"`
	StartDate   string         `jsonschema:"startDate;description:resume.json uses the ISO 8601 date standard;format:date"`
	EndDate     string         `jsonschema:"endDate;description:resume.json uses the ISO 8601 date standard;format:date"`
	Score       scoreSchema    `jsonschema:"score;additionalProperties"`
	Courses     []string       `jsonschema:"courses;description:List notable courses/subjects;items_description:Add course name"`
	Honors      []string       `jsonschema:"honors;description:List education honours"`
}

type volunteerSchema struct {
	Organization string         `jsonschema:"organization"`
	Position     string         `jsonschema:"position;description:Contribution type"`
	Location     locationSchema `jsonschema:"location;description:Location for this activity;format:location"`
	URL          string         `jsonschema:"url;description:Related link to support volunteer experience;format:uri"`
	StartDate    string         `jsonschema:"startDate;description:resume.json uses the ISO 8601 date standard;format:date"`
	EndDate      string         `jsonschema:"endDate;description:resume.json uses the ISO 8601 date standard;format:date"`
	Summary      string         `jsonschema:"summary;description:Give an overview of your responsibilities at the company"`
	Highlights   []string       `jsonschema:"highlights;description:Specify accomplishments and achievements"`
}

type publicationSchema struct {
	Name        string           `jsonschema:"name`
	Publisher   string           `jsonschema:"publisher`
	ReleaseDate string           `jsonschema:"releaseDate;description:resume.json uses the ISO 8601 date standard"`
	Resources   []resourceSchema `jsonschema:"resources;description:Specify multiple resources with label"`
	URL         string           `jsonschema:"url;description:format:uri"`
	Summary     string           `jsonschema:"summary;description:Short summary of publication"`
}

type legalSchema struct {
	Name            string           `jsonschema:"name;description:Add document name"`
	LegalType       string           `jsonschema:"legalType;description:Give the type of this document"`
	Description     string           `jsonschema:"description;description:Give a brief description about this document"`
	ApplicationDate string           `jsonschema:"applicationDate;description:resume.json uses the ISO 8601 date standard;format:date"`
	GrantDate       string           `jsonschema:"grantDate;description:resume.json uses the ISO 8601 date standard;format:date"`
	EndDate         string           `jsonschema:"endDate;description:resume.json uses the ISO 8601 date standard;format:date"`
	Resources       []resourceSchema `jsonschema:"resources;description:Specify multiple resources with label"`
	IDNumber        string           `jsonschema:"idNumber;description:Add the application number or Id Number"`
}

type skillSchema struct {
	Name    string `jsonschema:"name`
	Keyword []struct {
		Name  string      `jsonschema:"name;description:Add the skill name"`
		Score scoreSchema `jsonschema:"score;additionalProperties;description:Score for the skill name"`
	} `jsonschema:"keyword;description:List some keywords pertaining to the skill;items_additionalProperties"`
}

type awardSchema struct {
	Title   string `jsonschema:"title`
	Date    string `jsonschema:"date;description:resume.json uses the ISO 8601 date standard;format:date"`
	Awarder string `jsonschema:"awarder`
	Summary string `jsonschema:"summary`
}

type projectSchema struct {
	Name        string           `jsonschema:"name;description:Name of the project"`
	Location    locationSchema   `jsonschema:"location;description:Location for this activity;format:location"`
	Description string           `jsonschema:"description;description:Short summary of project"`
	Highlights  []string         `jsonschema:"highlights;description:Specify multiple features"`
	Keywords    []string         `jsonschema:"keywords;description:Specify special elements involved"`
	StartDate   string           `jsonschema:"startDate;description:resume.json uses the ISO 8601 date standard;format:date"`
	EndDate     string           `jsonschema:"endDate;description:resume.json uses the ISO 8601 date standard;format:date"`
	Resources   []resourceSchema `jsonschema:"resources;description:Specify multiple resources with label"`
	URL         string           `jsonschema:"url;format:uri"`
	Roles       []string         `jsonschema:"roles;description:Specify your role on this project or in company"`
	Entity      string           `jsonschema:"entity;description:Relevant company/entity affiliations"`
	Type        string           `jsonschema:"type"`
}

type certificateSchema struct {
	Code          string      `jsonschema:"code"`
	Name          string      `jsonschema:"name;description:Add the certificate name"`
	Website       string      `jsonschema:"website;description:Link to issuing authority's description of the certificate;format:uri"`
	Verification  string      `jsonschema:"verification;description:External candidate verification URL;format:uri"`
	GrantDate     string      `jsonschema:"grantDate;description:resume.json uses the ISO 8601 date standard;format:date"`
	Score         scoreSchema `jsonschema:"score;additionalProperties;description:Exam result (PASS/FAIL, 100%, 100)"`
	EndDate       string      `jsonschema:"endDate;description:resume.json uses the ISO 8601 date standard;format:date"`
	DoesNotExpire bool        `jsonschema:"doesNotExpire;format:checkbox"`
}

type referenceSchema struct {
	Name      string `jsonschema:"name"`
	Company   string `jsonschema:"company"`
	Position  string `jsonschema:"position"`
	Reference string `jsonschema:"reference"`
}

type languageSchema struct {
	Language string      `jsonschema:"language;description:Name of language"`
	Score    scoreSchema `jsonschema:"score;additionalProperties;description:Score for the language"`
}

type interestSchema struct {
	Name     string   `jsonschema:"name"`
	Keywords []string `jsonschema:"keywords;description:Specify keywords associated with the particular interest"`
}

type metaSchema struct {
	Canonical    string `jsonschema:"canonical;description:URL (as per RFC 3986) to latest version of this document"`
	Version      string `jsonschema:"version"`
	LastModified string `jsonschema:"lastModified;description:resume.json uses the ISO 8601 date standard;format:date"`
}

type Schema struct {
	Core         coreSchema          `jsonschema:"core;additionalProperties"`
	Work         []workSchema        `jsonschema:"work;items_additionalProperties"`
	Education    []educationSchema   `jsonschema:"education;items_additionalProperties"`
	Volunteer    []volunteerSchema   `jsonschema:"volunteer;items_additionalProperties"`
	Publications []publicationSchema `jsonschema:"publications;description:Specify your publications;items_additionalProperties"`
	Legal        []legalSchema       `jsonschema:"legal;description:Specify your labels;items_additionalProperties"`
	Skills       []skillSchema       `jsonschema:"skills;description:List out your professional skill-set;items_additionalProperties"`
	Awards       []awardSchema       `jsonschema:"awards;description:Specify any awards you have received throughout your professional career;items_additionalProperties"`
	Projects     []projectSchema     `jsonschema:"projects;description:Specify career projects;items_additionalProperties"`
	Certificates []certificateSchema `jsonschema:"certificate;items_additionalProperties"`
	References   []referenceSchema   `jsonschema:"references;description:List references you have received;items_additionalProperties"`
	Languages    []languageSchema    `jsonschema:"languages;description:List any other languages you speak;items_additionalProperties"`
	Interests    []interestSchema    `jsonschema:"interests;items_additionalProperties"`
	Meta         metaSchema          `jsonschema:"meta;description:The schema version and any other tooling configuration lives here;additionalProperties"`
}
